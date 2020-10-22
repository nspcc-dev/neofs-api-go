package token

import (
	"crypto/ecdsa"
	"errors"

	"github.com/nspcc-dev/neofs-api-go/pkg/acl/eacl"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	"github.com/nspcc-dev/neofs-api-go/util/signature"
	"github.com/nspcc-dev/neofs-api-go/v2/acl"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	v2signature "github.com/nspcc-dev/neofs-api-go/v2/signature"
	crypto "github.com/nspcc-dev/neofs-crypto"
)

type BearerToken struct {
	token acl.BearerToken
}

func (b BearerToken) ToV2() *acl.BearerToken {
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

func NewBearerToken() *BearerToken {
	b := new(BearerToken)
	b.token = acl.BearerToken{}
	b.token.SetBody(new(acl.BearerTokenBody))

	return b
}

func NewBearerTokenFromV2(v2 *acl.BearerToken) *BearerToken {
	if v2 == nil {
		v2 = new(acl.BearerToken)
	}

	return &BearerToken{
		token: *v2,
	}
}

// sanityCheck if bearer token is ready to be issued
func sanityCheck(b *BearerToken) error {
	switch {
	case b == nil:
		return errors.New("bearer token is not set")
	case b.token.GetBody() == nil:
		return errors.New("bearer token body is not set")
	case b.token.GetBody().GetEACL() == nil:
		return errors.New("bearer token EACL table is not set")
	}

	// consider checking EACL sanity there, lifetime correctness, etc.

	return nil
}
