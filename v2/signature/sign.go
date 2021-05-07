package signature

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/nspcc-dev/neofs-api-go/util/signature"
	"github.com/nspcc-dev/neofs-api-go/v2/accounting"
	"github.com/nspcc-dev/neofs-api-go/v2/container"
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	"github.com/nspcc-dev/neofs-api-go/v2/object"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/reputation"
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
	StableMarshal([]byte) ([]byte, error)
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

func (s StableMarshalerWrapper) ReadSignedData(buf []byte) ([]byte, error) {
	if s.SM != nil {
		return s.SM.StableMarshal(buf)
	}

	return nil, nil
}

func (s StableMarshalerWrapper) SignedDataSize() int {
	if s.SM != nil {
		return s.SM.StableSize()
	}

	return 0
}

func keySignatureHandler(s *refs.Signature) signature.KeySignatureHandler {
	return func(key []byte, sig []byte) {
		s.SetKey(key)
		s.SetSign(sig)
	}
}

func keySignatureSource(s *refs.Signature) signature.KeySignatureSource {
	return func() ([]byte, []byte) {
		return s.GetKey(), s.GetSign()
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

func VerifyServiceMessage(msg interface{}) error {
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

	return verifyMatryoshkaLevel(serviceMessageBody(msg), meta, verify)
}

func verifyMatryoshkaLevel(body stableMarshaler, meta metaHeader, verify verificationHeader) error {
	if err := verifyServiceMessagePart(meta, verify.GetMetaSignature); err != nil {
		return errors.Wrap(err, "could not verify meta header")
	}

	origin := verify.getOrigin()

	if err := verifyServiceMessagePart(origin, verify.GetOriginSignature); err != nil {
		return errors.Wrap(err, "could not verify origin of verification header")
	}

	if origin == nil {
		if err := verifyServiceMessagePart(body, verify.GetBodySignature); err != nil {
			return errors.Wrap(err, "could not verify body")
		}

		return nil
	}

	if verify.GetBodySignature() != nil {
		return errors.New("body signature at the matryoshka upper level")
	}

	return verifyMatryoshkaLevel(body, meta.getOrigin(), origin)
}

func verifyServiceMessagePart(part stableMarshaler, sigRdr func() *refs.Signature) error {
	return signature.VerifyDataWithSource(
		&StableMarshalerWrapper{part},
		keySignatureSource(sigRdr()),
	)
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

		/* Reputation */
	case *reputation.AnnounceLocalTrustRequest:
		return v.GetBody()
	case *reputation.AnnounceLocalTrustResponse:
		return v.GetBody()
	case *reputation.AnnounceIntermediateResultRequest:
		return v.GetBody()
	case *reputation.AnnounceIntermediateResultResponse:
		return v.GetBody()
	}
}
