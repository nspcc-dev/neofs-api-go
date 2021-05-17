package object

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/util/signature"
	signatureV2 "github.com/nspcc-dev/neofs-api-go/v2/signature"
)

var errCheckSumMismatch = errors.New("payload checksum mismatch")

var errIncorrectID = errors.New("incorrect object identifier")

// CalculatePayloadChecksum calculates and returns checksum of
// object payload bytes.
func CalculatePayloadChecksum(payload []byte) *pkg.Checksum {
	res := pkg.NewChecksum()
	res.SetSHA256(sha256.Sum256(payload))

	return res
}

// CalculateAndSetPayloadChecksum calculates checksum of current
// object payload and writes it to the object.
func CalculateAndSetPayloadChecksum(obj *RawObject) {
	obj.SetPayloadChecksum(
		CalculatePayloadChecksum(obj.Payload()),
	)
}

// VerifyPayloadChecksum checks if payload checksum in the object
// corresponds to its payload.
func VerifyPayloadChecksum(obj *Object) error {
	if !pkg.EqualChecksums(
		obj.PayloadChecksum(),
		CalculatePayloadChecksum(obj.Payload()),
	) {
		return errCheckSumMismatch
	}

	return nil
}

// CalculateID calculates identifier for the object.
func CalculateID(obj *Object) (*ID, error) {
	data, err := obj.ToV2().GetHeader().StableMarshal(nil)
	if err != nil {
		return nil, err
	}

	id := NewID()
	id.SetSHA256(sha256.Sum256(data))

	return id, nil
}

// CalculateAndSetID calculates identifier for the object
// and writes the result to it.
func CalculateAndSetID(obj *RawObject) error {
	id, err := CalculateID(obj.Object())
	if err != nil {
		return err
	}

	obj.SetID(id)

	return nil
}

// VerifyID checks if identifier in the object corresponds to
// its structure.
func VerifyID(obj *Object) error {
	id, err := CalculateID(obj)
	if err != nil {
		return err
	}

	if !id.Equal(obj.ID()) {
		return errIncorrectID
	}

	return nil
}

func CalculateIDSignature(key *ecdsa.PrivateKey, id *ID) (*pkg.Signature, error) {
	sig := pkg.NewSignature()

	if err := signature.SignDataWithHandler(
		key,
		signatureV2.StableMarshalerWrapper{
			SM: id.ToV2(),
		},
		func(key, sign []byte) {
			sig.SetKey(key)
			sig.SetSign(sign)
		},
	); err != nil {
		return nil, err
	}

	return sig, nil
}

func CalculateAndSetSignature(key *ecdsa.PrivateKey, obj *RawObject) error {
	sig, err := CalculateIDSignature(key, obj.ID())
	if err != nil {
		return err
	}

	obj.SetSignature(sig)

	return nil
}

func VerifyIDSignature(obj *Object) error {
	return signature.VerifyDataWithSource(
		signatureV2.StableMarshalerWrapper{
			SM: obj.ID().ToV2(),
		},
		func() ([]byte, []byte) {
			sig := obj.Signature()

			return sig.Key(), sig.Sign()
		},
	)
}

// SetIDWithSignature sets object identifier and signature.
func SetIDWithSignature(key *ecdsa.PrivateKey, obj *RawObject) error {
	if err := CalculateAndSetID(obj); err != nil {
		return fmt.Errorf("could not set identifier: %w", err)
	}

	if err := CalculateAndSetSignature(key, obj); err != nil {
		return fmt.Errorf("could not set signature: %w", err)
	}

	return nil
}

// SetVerificationFields calculates and sets all verification fields of the object.
func SetVerificationFields(key *ecdsa.PrivateKey, obj *RawObject) error {
	CalculateAndSetPayloadChecksum(obj)

	return SetIDWithSignature(key, obj)
}

// CheckVerificationFields checks all verification fields of the object.
func CheckVerificationFields(obj *Object) error {
	if err := CheckHeaderVerificationFields(obj); err != nil {
		return fmt.Errorf("invalid header structure: %w", err)
	}

	if err := VerifyPayloadChecksum(obj); err != nil {
		return fmt.Errorf("invalid payload checksum: %w", err)
	}

	return nil
}

// CheckHeaderVerificationFields checks all verification fields except payload.
func CheckHeaderVerificationFields(obj *Object) error {
	if err := VerifyIDSignature(obj); err != nil {
		return fmt.Errorf("invalid signature: %w", err)
	}

	if err := VerifyID(obj); err != nil {
		return fmt.Errorf("invalid identifier: %w", err)
	}

	return nil
}
