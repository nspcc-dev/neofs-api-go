package object

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"errors"
	"fmt"

	cryptoalgo "github.com/nspcc-dev/neofs-api-go/crypto/algo"
	neofsecdsa "github.com/nspcc-dev/neofs-api-go/crypto/ecdsa"
	"github.com/nspcc-dev/neofs-api-go/pkg"
	apicrypto "github.com/nspcc-dev/neofs-api-go/v2/crypto"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
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

// CalculateIDSignatureECDSA calculates and returns ECDSA signature of ID.
//
// Key must not be nil.
func CalculateIDSignatureECDSA(key ecdsa.PrivateKey, id *ID) (*pkg.Signature, error) {
	sigV2 := new(refs.Signature)

	var p apicrypto.SignPrm

	p.SetProtoMarshaler(signatureV2.StableMarshalerCrypto(id.ToV2()))
	p.SetTargetSignature(sigV2)

	if err := apicrypto.Sign(neofsecdsa.Signer(key), p); err != nil {
		return nil, err
	}

	return pkg.NewSignatureFromV2(sigV2), nil
}

// CalculateAndSetECDSASignature calculates ECDSA signature of object ID and
// writes it to RawObject.
//
// Key must not be nil.
func CalculateAndSetECDSASignature(key ecdsa.PrivateKey, obj *RawObject) error {
	sig, err := CalculateIDSignatureECDSA(key, obj.ID())
	if err != nil {
		return err
	}

	obj.SetSignature(sig)

	return nil
}

func VerifyIDSignature(obj *Object) error {
	sig := obj.Signature().ToV2()

	key, err := cryptoalgo.UnmarshalKey(cryptoalgo.ECDSA, sig.GetKey())
	if err != nil {
		return err
	}

	var p apicrypto.VerifyPrm

	p.SetProtoMarshaler(signatureV2.StableMarshalerCrypto(obj.ID().ToV2()))
	p.SetSignature(sig.GetSign())

	if !apicrypto.Verify(key, p) {
		return errors.New("invalid signature")
	}

	return nil
}

// SetIDWithSignatureECDSA sets object identifier and ECDSA signature.
//
// Key must not be nil.
func SetIDWithSignatureECDSA(key ecdsa.PrivateKey, obj *RawObject) error {
	if err := CalculateAndSetID(obj); err != nil {
		return fmt.Errorf("could not set identifier: %w", err)
	}

	if err := CalculateAndSetECDSASignature(key, obj); err != nil {
		return fmt.Errorf("could not set signature: %w", err)
	}

	return nil
}

// SetVerificationFields calculates and sets all verification fields of the object.
//
// Key must not be nil.
func SetVerificationFieldsECDSA(key ecdsa.PrivateKey, obj *RawObject) error {
	CalculateAndSetPayloadChecksum(obj)

	return SetIDWithSignatureECDSA(key, obj)
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
