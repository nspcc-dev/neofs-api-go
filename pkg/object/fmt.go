package object

import (
	"crypto/ecdsa"
	"crypto/sha256"

	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/util/signature"
	signatureV2 "github.com/nspcc-dev/neofs-api-go/v2/signature"
	"github.com/pkg/errors"
)

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
		CalculatePayloadChecksum(obj.GetPayload()),
	)
}

// VerifyPayloadChecksum checks if payload checksum in the object
// corresponds to its payload.
func VerifyPayloadChecksum(obj *Object) error {
	if !pkg.EqualChecksums(
		obj.GetPayloadChecksum(),
		CalculatePayloadChecksum(obj.GetPayload()),
	) {
		return errors.New("payload checksum mismatch")
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

	if !id.Equal(obj.GetID()) {
		return errors.New("incorrect object identifier")
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
	sig, err := CalculateIDSignature(key, obj.GetID())
	if err != nil {
		return err
	}

	obj.SetSignature(sig)

	return nil
}

func VerifyIDSignature(obj *Object) error {
	return signature.VerifyDataWithSource(
		signatureV2.StableMarshalerWrapper{
			SM: obj.GetID().ToV2(),
		},
		func() ([]byte, []byte) {
			sig := obj.GetSignature()

			return sig.GetKey(), sig.GetSign()
		},
	)
}

// SetIDWithSignature sets object identifier and signature.
func SetIDWithSignature(key *ecdsa.PrivateKey, obj *RawObject) error {
	if err := CalculateAndSetID(obj); err != nil {
		return errors.Wrap(err, "could not set identifier")
	}

	if err := CalculateAndSetSignature(key, obj); err != nil {
		return errors.Wrap(err, "could not set signature")
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
		return errors.Wrap(err, "invalid header structure")
	}

	if err := VerifyPayloadChecksum(obj); err != nil {
		return errors.Wrap(err, "invalid payload checksum")
	}

	return nil
}

// CheckHeaderVerificationFields checks all verification fields except payload.
func CheckHeaderVerificationFields(obj *Object) error {
	if err := VerifyIDSignature(obj); err != nil {
		return errors.Wrap(err, "invalid signature")
	}

	if err := VerifyID(obj); err != nil {
		return errors.Wrap(err, "invalid identifier")
	}

	return nil
}
