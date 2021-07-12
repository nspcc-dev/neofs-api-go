package token

import (
	"crypto/ecdsa"
	"errors"

	neofsecdsa "github.com/nspcc-dev/neofs-api-go/crypto/ecdsa"
	"github.com/nspcc-dev/neofs-api-go/pkg/acl/eacl"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	"github.com/nspcc-dev/neofs-api-go/v2/acl"
	apicrypto "github.com/nspcc-dev/neofs-api-go/v2/crypto"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	v2signature "github.com/nspcc-dev/neofs-api-go/v2/signature"
)

var (
	errNilBearerToken     = errors.New("bearer token is not set")
	errNilBearerTokenBody = errors.New("bearer token body is not set")
	errNilBearerTokenEACL = errors.New("bearer token EACL table is not set")
)

type BearerToken struct {
	token acl.BearerToken
}

// ToV2 converts BearerToken to v2 BearerToken message.
//
// Nil BearerToken converts to nil.
func (b *BearerToken) ToV2() *acl.BearerToken {
	if b == nil {
		return nil
	}

	return &b.token
}

func (b *BearerToken) SetLifetime(exp, nbf, iat uint64) {
	body := b.token.GetBody()
	if body == nil {
		body = new(acl.BearerTokenBody)
	}

	lt := new(acl.TokenLifetime)
	lt.SetExp(exp)
	lt.SetNbf(nbf)
	lt.SetIat(iat)

	body.SetLifetime(lt)
	b.token.SetBody(body)
}

func (b *BearerToken) SetEACLTable(table *eacl.Table) {
	body := b.token.GetBody()
	if body == nil {
		body = new(acl.BearerTokenBody)
	}

	body.SetEACL(table.ToV2())
	b.token.SetBody(body)
}

func (b *BearerToken) SetOwner(id *owner.ID) {
	body := b.token.GetBody()
	if body == nil {
		body = new(acl.BearerTokenBody)
	}

	body.SetOwnerID(id.ToV2())
	b.token.SetBody(body)
}

// SignTokenECDSA signs BearerToken with ecdsa.PrivateKey.
//
// Key must not be nil.
func (b *BearerToken) SignTokenECDSA(key ecdsa.PrivateKey) error {
	err := sanityCheck(b)
	if err != nil {
		return err
	}

	sig := b.token.GetSignature()
	if sig == nil {
		sig = new(refs.Signature)
		b.token.SetSignature(sig)
	}

	var p apicrypto.SignPrm

	p.SetProtoMarshaler(v2signature.StableMarshalerCrypto(b.token.GetBody()))
	p.SetTargetSignature(sig)

	return apicrypto.Sign(neofsecdsa.Signer(key), p)
}

// Issuer returns owner.ID associated with the key that signed bearer token.
// To pass node validation it should be owner of requested container. Returns
// nil if token is not signed.
func (b *BearerToken) Issuer() *owner.ID {
	var key ecdsa.PublicKey

	err := neofsecdsa.UnmarshalPublicKey(&key, b.token.GetSignature().GetKey())
	if err != nil {
		return nil
	}

	wallet, err := owner.NEO3WalletFromECDSAPublicKey(key)
	if err != nil {
		return nil
	}

	return owner.NewIDFromNeo3Wallet(wallet)
}

// NewBearerToken creates and initializes blank BearerToken.
//
// Defaults:
//  - signature: nil;
//  - eacl: nil;
//  - ownerID: nil;
//  - exp: 0;
//  - nbf: 0;
//  - iat: 0.
func NewBearerToken() *BearerToken {
	b := new(BearerToken)
	b.token = acl.BearerToken{}
	b.token.SetBody(new(acl.BearerTokenBody))

	return b
}

// ToV2 converts BearerToken to v2 BearerToken message.
func NewBearerTokenFromV2(v2 *acl.BearerToken) *BearerToken {
	if v2 == nil {
		v2 = new(acl.BearerToken)
	}

	return &BearerToken{
		token: *v2,
	}
}

// sanityCheck if bearer token is ready to be issued.
func sanityCheck(b *BearerToken) error {
	switch {
	case b == nil:
		return errNilBearerToken
	case b.token.GetBody() == nil:
		return errNilBearerTokenBody
	case b.token.GetBody().GetEACL() == nil:
		return errNilBearerTokenEACL
	}

	// consider checking EACL sanity there, lifetime correctness, etc.

	return nil
}

// Marshal marshals BearerToken into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (b *BearerToken) Marshal(bs ...[]byte) ([]byte, error) {
	var buf []byte
	if len(bs) > 0 {
		buf = bs[0]
	}

	return b.ToV2().
		StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of BearerToken.
func (b *BearerToken) Unmarshal(data []byte) error {
	fV2 := new(acl.BearerToken)
	if err := fV2.Unmarshal(data); err != nil {
		return err
	}

	*b = *NewBearerTokenFromV2(fV2)

	return nil
}

// MarshalJSON encodes BearerToken to protobuf JSON format.
func (b *BearerToken) MarshalJSON() ([]byte, error) {
	return b.ToV2().
		MarshalJSON()
}

// UnmarshalJSON decodes BearerToken from protobuf JSON format.
func (b *BearerToken) UnmarshalJSON(data []byte) error {
	fV2 := new(acl.BearerToken)
	if err := fV2.UnmarshalJSON(data); err != nil {
		return err
	}

	*b = *NewBearerTokenFromV2(fV2)

	return nil
}
