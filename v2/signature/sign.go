package signature

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/nspcc-dev/neofs-api-go/util/signature"
	"github.com/nspcc-dev/neofs-api-go/v2/accounting"
	"github.com/nspcc-dev/neofs-api-go/v2/service"
	"github.com/pkg/errors"
)

type serviceRequest interface {
	GetMetaHeader() *service.RequestMetaHeader
	GetVerificationHeader() *service.RequestVerificationHeader
	SetVerificationHeader(*service.RequestVerificationHeader)
}

type serviceResponse interface {
	GetMetaHeader() *service.ResponseMetaHeader
	GetVerificationHeader() *service.ResponseVerificationHeader
	SetVerificationHeader(*service.ResponseVerificationHeader)
}

type stableMarshaler interface {
	StableMarshal([]byte) ([]byte, error)
	StableSize() int
}

type stableMarshalerWrapper struct {
	sm stableMarshaler
}

type metaHeader interface {
	stableMarshaler
	getOrigin() metaHeader
}

type verificationHeader interface {
	stableMarshaler

	GetBodySignature() *service.Signature
	SetBodySignature(*service.Signature)
	GetMetaSignature() *service.Signature
	SetMetaSignature(*service.Signature)
	GetOriginSignature() *service.Signature
	SetOriginSignature(*service.Signature)

	setOrigin(stableMarshaler)
	getOrigin() verificationHeader
}

type requestMetaHeader struct {
	*service.RequestMetaHeader
}

type responseMetaHeader struct {
	*service.ResponseMetaHeader
}

type requestVerificationHeader struct {
	*service.RequestVerificationHeader
}

type responseVerificationHeader struct {
	*service.ResponseVerificationHeader
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
		h.SetOrigin(m.(*service.RequestVerificationHeader))
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
		r.SetOrigin(m.(*service.ResponseVerificationHeader))
	}
}

func (s stableMarshalerWrapper) ReadSignedData(buf []byte) ([]byte, error) {
	if s.sm != nil {
		return s.sm.StableMarshal(buf)
	}

	return nil, nil
}

func (s stableMarshalerWrapper) SignedDataSize() int {
	if s.sm != nil {
		return s.sm.StableSize()
	}

	return 0
}

func keySignatureHandler(s *service.Signature) signature.KeySignatureHandler {
	return func(key []byte, sig []byte) {
		s.SetKey(key)
		s.SetSign(sig)
	}
}

func keySignatureSource(s *service.Signature) signature.KeySignatureSource {
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
		verifyHdr = &requestVerificationHeader{new(service.RequestVerificationHeader)}
		verifyHdrSetter = func(h verificationHeader) {
			v.SetVerificationHeader(h.(*requestVerificationHeader).RequestVerificationHeader)
		}

		if h := v.GetVerificationHeader(); h != nil {
			verifyOrigin = h
		}
	case serviceResponse:
		body = serviceMessageBody(v)
		meta = v.GetMetaHeader()
		verifyHdr = &responseVerificationHeader{new(service.ResponseVerificationHeader)}
		verifyHdrSetter = func(h verificationHeader) {
			v.SetVerificationHeader(h.(*responseVerificationHeader).ResponseVerificationHeader)
		}

		if h := v.GetVerificationHeader(); h != nil {
			verifyOrigin = h
		}
	default:
		panic(fmt.Sprintf("unsupported service message %T", v))
	}

	if verifyOrigin == nil {
		// sign service message body
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

func signServiceMessagePart(key *ecdsa.PrivateKey, part stableMarshaler, sigWrite func(*service.Signature)) error {
	sig := new(service.Signature)

	// sign part
	if err := signature.SignDataWithHandler(
		key,
		&stableMarshalerWrapper{part},
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
		panic(fmt.Sprintf("unsupported service message %T", v))
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

func verifyServiceMessagePart(part stableMarshaler, sigRdr func() *service.Signature) error {
	return signature.VerifyDataWithSource(
		&stableMarshalerWrapper{part},
		keySignatureSource(sigRdr()),
	)
}

func serviceMessageBody(req interface{}) stableMarshaler {
	switch v := req.(type) {
	case *accounting.BalanceRequest:
		return v.GetBody()
	case *accounting.BalanceResponse:
		return v.GetBody()
	default:
		panic(fmt.Sprintf("unsupported service message %T", req))
	}
}
