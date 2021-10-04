package token

import (
	"crypto/ecdsa"
	"errors"

	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/acl/eacl"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	"github.com/nspcc-dev/neofs-api-go/util/signature"
	"github.com/nspcc-dev/neofs-api-go/v2/acl"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	v2signature "github.com/nspcc-dev/neofs-api-go/v2/signature"
	crypto "github.com/nspcc-dev/neofs-crypto"
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

// SetLifetime sets lifetime related numeric values to the token.
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

// Exp returns epoch number of the token expiration.
func (b *BearerToken) Exp() uint64 {
	return b.token.GetBody().GetLifetime().GetExp()
}

// Nbf returns starting epoch number of the token.
func (b *BearerToken) Nbf() uint64 {
	return b.token.GetBody().GetLifetime().GetNbf()
}

// Iat returns epoch number when token was issued.
func (b *BearerToken) Iat() uint64 {
	return b.token.GetBody().GetLifetime().GetIat()
}

// SetEACLTable attaches extended ACL table to the bearer token.
func (b *BearerToken) SetEACLTable(table *eacl.Table) {
	body := b.token.GetBody()
	if body == nil {
		body = new(acl.BearerTokenBody)
	}

	body.SetEACL(table.ToV2())
	b.token.SetBody(body)
}

// EACLTable returns extended ACL table attached to the token.
func (b *BearerToken) EACLTable() *eacl.Table {
	return eacl.NewTableFromV2(b.token.GetBody().GetEACL())
}

// SetOwner sets the token owner. The same owner should be used when sending
// request with attached bearer token.
func (b *BearerToken) SetOwner(id *owner.ID) {
	body := b.token.GetBody()
	if body == nil {
		body = new(acl.BearerTokenBody)
	}

	body.SetOwnerID(id.ToV2())
	b.token.SetBody(body)
}

// Owner returns owner of the bearer token.
// Do not mistake with bearer token issuer.
func (b *BearerToken) Owner() *owner.ID {
	return owner.NewIDFromV2(b.token.GetBody().GetOwnerID())
}

// SignToken signs the token with the token issuer key.
func (b *BearerToken) SignToken(key *ecdsa.PrivateKey) error {
	err := sanityCheck(b)
	if err != nil {
		return err
	}

	signWrapper := v2signature.StableMarshalerWrapper{SM: b.token.GetBody()}

	return signature.SignDataWithHandler(key, signWrapper, func(key []byte, sig []byte) {
		bearerSignature := new(refs.Signature)
		bearerSignature.SetKey(key)
		bearerSignature.SetSign(sig)
		b.token.SetSignature(bearerSignature)
	})
}

// Signature of the session token issuer.
func (b *BearerToken) Signature() *pkg.Signature {
	return pkg.NewSignatureFromV2(b.token.GetSignature())
}

// VerifyBearerTokenSignature checks if bearer token was correctly
// signed by the key specified in the signature field.
// If signature is correct, then returns nil.
func VerifyBearerTokenSignature(b *BearerToken) error {
	signWrapper := v2signature.StableMarshalerWrapper{SM: b.token.GetBody()}

	return signature.VerifyDataWithSource(signWrapper, func() (key, sig []byte) {
		tokenSignature := b.Signature()
		return tokenSignature.Key(), tokenSignature.Sign()
	})
}

// Issuer returns owner.ID associated with the key that signed bearer token.
// To pass node validation it should be owner of requested container. Returns
// nil if token is not signed.
func (b *BearerToken) Issuer() *owner.ID {
	pubKey := crypto.UnmarshalPublicKey(b.token.GetSignature().GetKey())

	wallet, err := owner.NEO3WalletFromPublicKey(pubKey)
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
