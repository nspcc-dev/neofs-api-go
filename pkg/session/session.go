package session

import (
	"crypto/ecdsa"

	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	"github.com/nspcc-dev/neofs-api-go/util/signature"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/session"
	v2signature "github.com/nspcc-dev/neofs-api-go/v2/signature"
)

// Token represents NeoFS API v2-compatible
// session token.
type Token session.SessionToken

// NewTokenFromV2 wraps session.SessionToken message structure
// into Token.
func NewTokenFromV2(tV2 *session.SessionToken) *Token {
	return (*Token)(tV2)
}

// NewToken creates and returns blank Token.
func NewToken() *Token {
	return NewTokenFromV2(new(session.SessionToken))
}

// ToV2 converts Token to session.SessionToken message structure.
func (t *Token) ToV2() *session.SessionToken {
	return (*session.SessionToken)(t)
}

func (t *Token) setBodyField(setter func(*session.SessionTokenBody)) {
	token := (*session.SessionToken)(t)
	body := token.GetBody()

	if body == nil {
		body = new(session.SessionTokenBody)
		token.SetBody(body)
	}

	setter(body)
}

// ID returns Token identifier.
func (t *Token) ID() []byte {
	return (*session.SessionToken)(t).
		GetBody().
		GetID()
}

// SetID sets Token identifier.
func (t *Token) SetID(v []byte) {
	t.setBodyField(func(body *session.SessionTokenBody) {
		body.SetID(v)
	})
}

// OwnerID returns Token's owner identifier.
func (t *Token) OwnerID() *owner.ID {
	return owner.NewIDFromV2(
		(*session.SessionToken)(t).
			GetBody().
			GetOwnerID(),
	)
}

// SetOwnerID sets Token's owner identifier.
func (t *Token) SetOwnerID(v *owner.ID) {
	t.setBodyField(func(body *session.SessionTokenBody) {
		body.SetOwnerID(v.ToV2())
	})
}

// SessionKey returns public key of the session
// in a binary format.
func (t *Token) SessionKey() []byte {
	return (*session.SessionToken)(t).
		GetBody().
		GetSessionKey()
}

// SetSessionKey sets public key of the session
// // in a binary format.
func (t *Token) SetSessionKey(v []byte) {
	t.setBodyField(func(body *session.SessionTokenBody) {
		body.SetSessionKey(v)
	})
}

// Sign calculates and writes signature of the Token data.
//
// Returns signature calculation errors.
func (t *Token) Sign(key *ecdsa.PrivateKey) error {
	tV2 := (*session.SessionToken)(t)

	signedData := v2signature.StableMarshalerWrapper{
		SM: tV2.GetBody(),
	}

	return signature.SignDataWithHandler(key, signedData, func(key, sig []byte) {
		tSig := tV2.GetSignature()
		if tSig == nil {
			tSig = new(refs.Signature)
		}

		tSig.SetKey(key)
		tSig.SetSign(sig)

		tV2.SetSignature(tSig)
	})
}

// VerifySignature checks if token signature is
// presented and valid.
func (t *Token) VerifySignature() bool {
	tV2 := (*session.SessionToken)(t)

	signedData := v2signature.StableMarshalerWrapper{
		SM: tV2.GetBody(),
	}

	return signature.VerifyDataWithSource(signedData, func() (key, sig []byte) {
		tSig := tV2.GetSignature()
		return tSig.GetKey(), tSig.GetSign()
	}) == nil
}

// Signature returns Token signature.
func (t *Token) Signature() *pkg.Signature {
	return pkg.NewSignatureFromV2(
		(*session.SessionToken)(t).
			GetSignature(),
	)
}

// Marshal marshals Token into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (t *Token) Marshal(bs ...[]byte) ([]byte, error) {
	var buf []byte
	if len(bs) > 0 {
		buf = bs[0]
	}

	return (*session.SessionToken)(t).
		StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of Token.
func (t *Token) Unmarshal(data []byte) error {
	return (*session.SessionToken)(t).
		Unmarshal(data)
}

// MarshalJSON encodes Token to protobuf JSON format.
func (t *Token) MarshalJSON() ([]byte, error) {
	return (*session.SessionToken)(t).
		MarshalJSON()
}

// UnmarshalJSON decodes Token from protobuf JSON format.
func (t *Token) UnmarshalJSON(data []byte) error {
	return (*session.SessionToken)(t).
		UnmarshalJSON(data)
}
