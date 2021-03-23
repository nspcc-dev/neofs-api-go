package signature

import (
	"crypto/ecdsa"
	"fmt"
	"io"

	"github.com/nspcc-dev/neofs-api-go/util/proto"
	"github.com/nspcc-dev/neofs-api-go/util/signature"
	"github.com/nspcc-dev/neofs-api-go/v2/accounting"
	"github.com/nspcc-dev/neofs-api-go/v2/container"
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	"github.com/nspcc-dev/neofs-api-go/v2/object"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/session"
	"github.com/pkg/errors"
)

type serviceRequest interface {
	GetMetaHeader() *session.RequestMetaHeader
	GetVerificationHeader() *session.RequestVerificationHeader
	SetVerificationHeader(*session.RequestVerificationHeader)
}

type serviceResponse interface {
	GetMetaHeader() *session.ResponseMetaHeader
	GetVerificationHeader() *session.ResponseVerificationHeader
	SetVerificationHeader(*session.ResponseVerificationHeader)
}

type stableMarshaler interface {
	MarshalStream(s proto.Stream) (int, error)
	StableSize() int
}

type StableMarshalerWrapper struct {
	SM stableMarshaler
}

type metaHeader interface {
	stableMarshaler
	getOrigin() metaHeader
}

type verificationHeader interface {
	stableMarshaler

	GetBodySignature() *refs.Signature
	SetBodySignature(*refs.Signature)
	GetMetaSignature() *refs.Signature
	SetMetaSignature(*refs.Signature)
	GetOriginSignature() *refs.Signature
	SetOriginSignature(*refs.Signature)

	setOrigin(stableMarshaler)
	getOrigin() verificationHeader
}

type requestMetaHeader struct {
	*session.RequestMetaHeader
}

type responseMetaHeader struct {
	*session.ResponseMetaHeader
}

type requestVerificationHeader struct {
	*session.RequestVerificationHeader
}

type responseVerificationHeader struct {
	*session.ResponseVerificationHeader
}

func (h *requestMetaHeader) getOrigin() metaHeader {
	return &requestMetaHeader{
		RequestMetaHeader: h.GetOrigin(),
	}
}

func (h *responseMetaHeader) getOrigin() metaHeader {
	return &responseMetaHeader{
		ResponseMetaHeader: h.GetOrigin(),
	}
}

func (h *requestVerificationHeader) getOrigin() verificationHeader {
	if origin := h.GetOrigin(); origin != nil {
		return &requestVerificationHeader{
			RequestVerificationHeader: origin,
		}
	}

	return nil
}

func (h *requestVerificationHeader) setOrigin(m stableMarshaler) {
	if m != nil {
		h.SetOrigin(m.(*session.RequestVerificationHeader))
	}
}

func (r *responseVerificationHeader) getOrigin() verificationHeader {
	if origin := r.GetOrigin(); origin != nil {
		return &responseVerificationHeader{
			ResponseVerificationHeader: origin,
		}
	}

	return nil
}

func (r *responseVerificationHeader) setOrigin(m stableMarshaler) {
	if m != nil {
		r.SetOrigin(m.(*session.ResponseVerificationHeader))
	}
}

func (s StableMarshalerWrapper) WriteSignedDataTo(w io.Writer) (int, error) {
	if s.SM != nil {
		ss := proto.NewStream(w)
		return s.SM.MarshalStream(ss)
	}
	return 0, nil
}

func keySignatureHandler(s *refs.Signature) signature.KeySignatureHandler {
	return func(key []byte, sig []byte) {
		s.SetKey(key)
		s.SetSign(sig)
	}
}

func SignServiceMessage(key *ecdsa.PrivateKey, msg interface{}) error {
	var (
		body, meta, verifyOrigin stableMarshaler
		verifyHdr                verificationHeader
		verifyHdrSetter          func(verificationHeader)
	)

	switch v := msg.(type) {
	case nil:
		return nil
	case serviceRequest:
		body = serviceMessageBody(v)
		meta = v.GetMetaHeader()
		verifyHdr = &requestVerificationHeader{new(session.RequestVerificationHeader)}
		verifyHdrSetter = func(h verificationHeader) {
			v.SetVerificationHeader(h.(*requestVerificationHeader).RequestVerificationHeader)
		}

		if h := v.GetVerificationHeader(); h != nil {
			verifyOrigin = h
		}
	case serviceResponse:
		body = serviceMessageBody(v)
		meta = v.GetMetaHeader()
		verifyHdr = &responseVerificationHeader{new(session.ResponseVerificationHeader)}
		verifyHdrSetter = func(h verificationHeader) {
			v.SetVerificationHeader(h.(*responseVerificationHeader).ResponseVerificationHeader)
		}

		if h := v.GetVerificationHeader(); h != nil {
			verifyOrigin = h
		}
	default:
		panic(fmt.Sprintf("unsupported session message %T", v))
	}

	if verifyOrigin == nil {
		// sign session message body
		if err := signServiceMessagePart(key, body, verifyHdr.SetBodySignature); err != nil {
			return errors.Wrap(err, "could not sign body")
		}
	}

	// sign meta header
	if err := signServiceMessagePart(key, meta, verifyHdr.SetMetaSignature); err != nil {
		return errors.Wrap(err, "could not sign meta header")
	}

	// sign verification header origin
	if err := signServiceMessagePart(key, verifyOrigin, verifyHdr.SetOriginSignature); err != nil {
		return errors.Wrap(err, "could not sign origin of verification header")
	}

	// wrap origin verification header
	verifyHdr.setOrigin(verifyOrigin)

	// update matryoshka verification header
	verifyHdrSetter(verifyHdr)

	return nil
}

func signServiceMessagePart(key *ecdsa.PrivateKey, part stableMarshaler, sigWrite func(*refs.Signature)) error {
	sig := new(refs.Signature)

	// sign part
	if err := signature.SignDataWithHandler(
		key,
		&StableMarshalerWrapper{part},
		keySignatureHandler(sig),
	); err != nil {
		return err
	}

	// write part signature
	sigWrite(sig)

	return nil
}

func VerifyServiceMessage(msg interface{}, opts ...signature.SignOption) error {
	cfg := signature.DefaultOptions()

	for i := range opts {
		opts[i](cfg)
	}

	var (
		meta   metaHeader
		verify verificationHeader
	)

	switch v := msg.(type) {
	case nil:
		return nil
	case serviceRequest:
		meta = &requestMetaHeader{
			RequestMetaHeader: v.GetMetaHeader(),
		}

		verify = &requestVerificationHeader{
			RequestVerificationHeader: v.GetVerificationHeader(),
		}
	case serviceResponse:
		meta = &responseMetaHeader{
			ResponseMetaHeader: v.GetMetaHeader(),
		}

		verify = &responseVerificationHeader{
			ResponseVerificationHeader: v.GetVerificationHeader(),
		}
	default:
		panic(fmt.Sprintf("unsupported session message %T", v))
	}

	return verifyMatryoshkaLevel(serviceMessageBody(msg), meta, verify, cfg)
}

func verifyMatryoshkaLevel(body stableMarshaler, meta metaHeader, verify verificationHeader,
	cfg *signature.Options) error {
	if err := verifyServiceMessagePart(meta, verify.GetMetaSignature(), cfg); err != nil {
		return errors.Wrap(err, "could not verify meta header")
	}

	origin := verify.getOrigin()

	if err := verifyServiceMessagePart(origin, verify.GetOriginSignature(), cfg); err != nil {
		return errors.Wrap(err, "could not verify origin of verification header")
	}

	if origin == nil {
		if err := verifyServiceMessagePart(body, verify.GetBodySignature(), cfg); err != nil {
			return errors.Wrap(err, "could not verify body")
		}

		return nil
	}

	if verify.GetBodySignature() != nil {
		return errors.New("body signature at the matryoshka upper level")
	}

	return verifyMatryoshkaLevel(body, meta.getOrigin(), origin, cfg)
}

func verifyServiceMessagePart(part stableMarshaler, sig *refs.Signature, cfg *signature.Options) error {
	return signature.VerifyData(&StableMarshalerWrapper{part}, sig.GetKey(), sig.GetSign(),
		signature.WithUnmarshalPublicKey(cfg.UnmarshalPublic))
}

func serviceMessageBody(req interface{}) stableMarshaler {
	switch v := req.(type) {
	default:
		panic(fmt.Sprintf("unsupported session message %T", req))

		/* Accounting */
	case *accounting.BalanceRequest:
		return v.GetBody()
	case *accounting.BalanceResponse:
		return v.GetBody()

		/* Session */
	case *session.CreateRequest:
		return v.GetBody()
	case *session.CreateResponse:
		return v.GetBody()

		/* Container */
	case *container.PutRequest:
		return v.GetBody()
	case *container.PutResponse:
		return v.GetBody()
	case *container.DeleteRequest:
		return v.GetBody()
	case *container.DeleteResponse:
		return v.GetBody()
	case *container.GetRequest:
		return v.GetBody()
	case *container.GetResponse:
		return v.GetBody()
	case *container.ListRequest:
		return v.GetBody()
	case *container.ListResponse:
		return v.GetBody()
	case *container.SetExtendedACLRequest:
		return v.GetBody()
	case *container.SetExtendedACLResponse:
		return v.GetBody()
	case *container.GetExtendedACLRequest:
		return v.GetBody()
	case *container.GetExtendedACLResponse:
		return v.GetBody()
	case *container.AnnounceUsedSpaceRequest:
		return v.GetBody()
	case *container.AnnounceUsedSpaceResponse:
		return v.GetBody()

		/* Object */
	case *object.PutRequest:
		return v.GetBody()
	case *object.PutResponse:
		return v.GetBody()
	case *object.GetRequest:
		return v.GetBody()
	case *object.GetResponse:
		return v.GetBody()
	case *object.HeadRequest:
		return v.GetBody()
	case *object.HeadResponse:
		return v.GetBody()
	case *object.SearchRequest:
		return v.GetBody()
	case *object.SearchResponse:
		return v.GetBody()
	case *object.DeleteRequest:
		return v.GetBody()
	case *object.DeleteResponse:
		return v.GetBody()
	case *object.GetRangeRequest:
		return v.GetBody()
	case *object.GetRangeResponse:
		return v.GetBody()
	case *object.GetRangeHashRequest:
		return v.GetBody()
	case *object.GetRangeHashResponse:
		return v.GetBody()

		/* Netmap */
	case *netmap.LocalNodeInfoRequest:
		return v.GetBody()
	case *netmap.LocalNodeInfoResponse:
		return v.GetBody()
	case *netmap.NetworkInfoRequest:
		return v.GetBody()
	case *netmap.NetworkInfoResponse:
		return v.GetBody()
	}
}
