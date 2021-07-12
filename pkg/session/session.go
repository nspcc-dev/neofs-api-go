package session

import (
	"crypto/ecdsa"

	cryptoalgo "github.com/nspcc-dev/neofs-api-go/crypto/algo"
	neofsecdsa "github.com/nspcc-dev/neofs-api-go/crypto/ecdsa"
	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	apicrypto "github.com/nspcc-dev/neofs-api-go/v2/crypto"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/session"
	v2signature "github.com/nspcc-dev/neofs-api-go/v2/signature"
)

// Token represents NeoFS API v2-compatible
// session token.
type Token session.SessionToken

// NewTokenFromV2 wraps session.SessionToken message structure
// into Token.
//
// Nil session.SessionToken converts to nil.
func NewTokenFromV2(tV2 *session.SessionToken) *Token {
	return (*Token)(tV2)
}

// NewToken creates and returns blank Token.
//
// Defaults:
//  - body: nil;
//  - id: nil;
//  - ownerId: nil;
//  - sessionKey: nil;
//  - exp: 0;
//  - iat: 0;
//  - nbf: 0;
func NewToken() *Token {
	return NewTokenFromV2(new(session.SessionToken))
}

// ToV2 converts Token to session.SessionToken message structure.
//
// Nil Token converts to nil.
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
// in a binary format.
func (t *Token) SetSessionKey(v []byte) {
	t.setBodyField(func(body *session.SessionTokenBody) {
		body.SetSessionKey(v)
	})
}

func (t *Token) setLifetimeField(f func(*session.TokenLifetime)) {
	t.setBodyField(func(body *session.SessionTokenBody) {
		lt := body.GetLifetime()
		if lt == nil {
			lt = new(session.TokenLifetime)
			body.SetLifetime(lt)
		}

		f(lt)
	})
}

// Exp returns epoch number of the token expiration.
func (t *Token) Exp() uint64 {
	return (*session.SessionToken)(t).
		GetBody().
		GetLifetime().
		GetExp()
}

// SetExp sets epoch number of the token expiration.
func (t *Token) SetExp(exp uint64) {
	t.setLifetimeField(func(lt *session.TokenLifetime) {
		lt.SetExp(exp)
	})
}

// Nbf returns starting epoch number of the token.
func (t *Token) Nbf() uint64 {
	return (*session.SessionToken)(t).
		GetBody().
		GetLifetime().
		GetNbf()
}

// SetNbf sets starting epoch number of the token.
func (t *Token) SetNbf(nbf uint64) {
	t.setLifetimeField(func(lt *session.TokenLifetime) {
		lt.SetNbf(nbf)
	})
}

// Iat returns starting epoch number of the token.
func (t *Token) Iat() uint64 {
	return (*session.SessionToken)(t).
		GetBody().
		GetLifetime().
		GetIat()
}

// SetIat sets the number of the epoch in which the token was issued.
func (t *Token) SetIat(iat uint64) {
	t.setLifetimeField(func(lt *session.TokenLifetime) {
		lt.SetIat(iat)
	})
}

// SignECDSA calculates and writes ECDSA signature of the Token data.
//
// Returns signature calculation errors.
func (t *Token) SignECDSA(key ecdsa.PrivateKey) error {
	tV2 := (*session.SessionToken)(t)

	tSig := tV2.GetSignature()
	if tSig == nil {
		tSig = new(refs.Signature)
		tV2.SetSignature(tSig)
	}

	var p apicrypto.SignPrm

	p.SetProtoMarshaler(v2signature.StableMarshalerCrypto(tV2.GetBody()))
	p.SetTargetSignature(tSig)

	return apicrypto.Sign(neofsecdsa.Signer(key), p)
}

// VerifySignature checks if token signature is
// presented and valid.
func (t *Token) VerifySignature() bool {
	tV2 := (*session.SessionToken)(t)

	sig := tV2.GetSignature()

	key, err := cryptoalgo.UnmarshalKey(cryptoalgo.ECDSA, sig.GetKey())
	if err != nil {
		return false
	}

	var p apicrypto.VerifyPrm

	p.SetProtoMarshaler(v2signature.StableMarshalerCrypto(tV2.GetBody()))
	p.SetSignature(sig.GetSign())

	return apicrypto.Verify(key, p)
}

// Signature returns Token signature.
func (t *Token) Signature() *pkg.Signature {
	return pkg.NewSignatureFromV2(
		(*session.SessionToken)(t).
			GetSignature(),
	)
}

// SetContext sets context of the Token.
//
// Supported contexts:
//  - *ContainerContext.
//
// Resets context if it is not supported.
func (t *Token) SetContext(v interface{}) {
	var cV2 session.SessionTokenContext

	switch c := v.(type) {
	case *ContainerContext:
		cV2 = c.ToV2()
	}

	t.setBodyField(func(body *session.SessionTokenBody) {
		body.SetContext(cV2)
	})
}

// Context returns context of the Token.
//
// Supports same contexts as SetContext.
//
// Returns nil if context is not supported.
func (t *Token) Context() interface{} {
	switch v := (*session.SessionToken)(t).
		GetBody().
		GetContext(); c := v.(type) {
	default:
		return nil
	case *session.ContainerSessionContext:
		return ContainerContextFromV2(c)
	}
}

// GetContainerContext is a helper function that casts
// Token context to ContainerContext.
//
// Returns nil if context is not a ContainerContext.
func GetContainerContext(t *Token) *ContainerContext {
	c, _ := t.Context().(*ContainerContext)
	return c
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
