package signature

import (
	"crypto/ecdsa"

	"github.com/pkg/errors"
)

type SignedRequest interface {
	RequestBody() DataSource
	RequestMetaHeader() DataSource
	OriginVerificationHeader() DataSource

	SetBodySignatureWithKey(key, sig []byte)
	BodySignatureWithKey() (key, sig []byte)

	SetMetaSignatureWithKey(key, sig []byte)
	MetaSignatureWithKey() (key, sig []byte)

	SetOriginSignatureWithKey(key, sig []byte)
	OriginSignatureWithKey() (key, sig []byte)
}

func SignRequest(key *ecdsa.PrivateKey, src SignedRequest) error {
	if src == nil {
		return errors.New("nil source")
	}

	// sign body
	if err := SignDataWithHandler(key, src.RequestBody(), src.SetBodySignatureWithKey); err != nil {
		return errors.Wrap(err, "could not sign body")
	}

	// sign meta
	if err := SignDataWithHandler(key, src.RequestMetaHeader(), src.SetMetaSignatureWithKey); err != nil {
		return errors.Wrap(err, "could not sign meta header")
	}

	// sign verify origin
	if err := SignDataWithHandler(key, src.OriginVerificationHeader(), src.SetOriginSignatureWithKey); err != nil {
		return errors.Wrap(err, "could not sign verification header origin")
	}

	return nil
}

func VerifyRequest(src SignedRequest) error {
	// verify body signature
	if err := VerifyDataWithSource(src.RequestBody(), src.BodySignatureWithKey); err != nil {
		return errors.Wrap(err, "could not verify body")
	}

	// verify meta header
	if err := VerifyDataWithSource(src.RequestMetaHeader(), src.MetaSignatureWithKey); err != nil {
		return errors.Wrap(err, "could not verify meta header")
	}

	// verify verification header origin
	if err := VerifyDataWithSource(src.OriginVerificationHeader(), src.OriginSignatureWithKey); err != nil {
		return errors.Wrap(err, "could not verify verification header origin")
	}

	return nil
}
