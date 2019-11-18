package service

import (
	"crypto/ecdsa"

	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/nspcc-dev/neofs-proto/internal"
	"github.com/pkg/errors"
)

// ErrWrongSignature should be raised when wrong signature is passed into VerifyRequest.
const ErrWrongSignature = internal.Error("wrong signature")

// SignedRequest interface allows sign and verify requests.
type SignedRequest interface {
	PrepareData() ([]byte, error)
	GetSignature() []byte
	SetSignature([]byte)
}

// SignRequest with passed private key.
func SignRequest(r SignedRequest, key *ecdsa.PrivateKey) error {
	var signature []byte
	if data, err := r.PrepareData(); err != nil {
		return err
	} else if signature, err = crypto.Sign(key, data); err != nil {
		return errors.Wrap(err, "could not sign data")
	}

	r.SetSignature(signature)

	return nil
}

// VerifyRequest by passed public keys.
func VerifyRequest(r SignedRequest, keys ...*ecdsa.PublicKey) bool {
	data, err := r.PrepareData()
	if err != nil {
		return false
	}
	for i := range keys {
		if err := crypto.Verify(keys[i], data, r.GetSignature()); err == nil {
			return true
		}
	}
	return false
}
