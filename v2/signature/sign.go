package signature

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/nspcc-dev/neofs-api-go/util/signature"
	"github.com/nspcc-dev/neofs-api-go/v2/accounting"
	"github.com/nspcc-dev/neofs-api-go/v2/service"
	"github.com/pkg/errors"
)

type SignedRequest interface {
	GetRequestMetaHeader() *service.RequestMetaHeader
	GetRequestVerificationHeader() *service.RequestVerificationHeader
	SetRequestVerificationHeader(*service.RequestVerificationHeader)
}

type stableMarshaler interface {
	StableMarshal([]byte) ([]byte, error)
	StableSize() int
}

type stableMarshalerWrapper struct {
	sm stableMarshaler
}

func (s stableMarshalerWrapper) ReadSignedData(buf []byte) ([]byte, error) {
	return s.sm.StableMarshal(buf)
}

func (s stableMarshalerWrapper) SignedDataSize() int {
	return s.sm.StableSize()
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

func requestBody(req SignedRequest) stableMarshaler {
	switch v := req.(type) {
	case *accounting.BalanceRequest:
		return v.GetBody()
	default:
		panic(fmt.Sprintf("unknown request %T", req))
	}
}

func SignRequest(key *ecdsa.PrivateKey, req SignedRequest) error {
	if req == nil {
		return nil
	}

	// create new level of matryoshka
	verifyHdr := new(service.RequestVerificationHeader)

	// attach the previous matryoshka
	verifyHdr.SetOrigin(req.GetRequestVerificationHeader())

	// sign request body
	if err := signRequestPart(key, requestBody(req), verifyHdr.SetBodySignature); err != nil {
		return errors.Wrap(err, "could not sign request body")
	}

	// sign meta header
	if err := signRequestPart(key, req.GetRequestMetaHeader(), verifyHdr.SetMetaSignature); err != nil {
		return errors.Wrap(err, "could not sign request meta header")
	}

	// sign verification header origin
	if err := signRequestPart(key, verifyHdr.GetOrigin(), verifyHdr.SetOriginSignature); err != nil {
		return errors.Wrap(err, "could not sign origin of request verification header")
	}

	// make a new top of the matryoshka
	req.SetRequestVerificationHeader(verifyHdr)

	return nil
}

func signRequestPart(key *ecdsa.PrivateKey, part stableMarshaler, sigWrite func(*service.Signature)) error {
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

func VerifyRequest(req SignedRequest) error {
	verifyHdr := req.GetRequestVerificationHeader()

	// verify body signature
	if err := verifyRequestPart(requestBody(req), verifyHdr.GetBodySignature); err != nil {
		return errors.Wrap(err, "could not verify request body")
	}

	// verify meta header
	if err := verifyRequestPart(req.GetRequestMetaHeader(), verifyHdr.GetMetaSignature); err != nil {
		return errors.Wrap(err, "could not verify request meta header")
	}

	// verify verification header origin
	if err := verifyRequestPart(verifyHdr.GetOrigin(), verifyHdr.GetOriginSignature); err != nil {
		return errors.Wrap(err, "could not verify origin of request verification header")
	}

	return nil
}

func verifyRequestPart(part stableMarshaler, sigRdr func() *service.Signature) error {
	if err := signature.VerifyDataWithSource(
		&stableMarshalerWrapper{part},
		keySignatureSource(sigRdr()),
	); err != nil {
		return err
	}

	return nil
}
