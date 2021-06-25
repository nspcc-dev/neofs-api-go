package tokentest

import (
	neofsecdsatest "github.com/nspcc-dev/neofs-api-go/crypto/ecdsa/test"
	eacltest "github.com/nspcc-dev/neofs-api-go/pkg/acl/eacl/test"
	ownertest "github.com/nspcc-dev/neofs-api-go/pkg/owner/test"
	"github.com/nspcc-dev/neofs-api-go/pkg/token"
)

// Generate returns random token.BearerToken.
//
// Resulting token is unsigned.
func Generate() *token.BearerToken {
	x := token.NewBearerToken()

	x.SetLifetime(3, 2, 1)
	x.SetOwner(ownertest.Generate())
	x.SetEACLTable(eacltest.Table())

	return x
}

// GenerateSigned returns signed random token.BearerToken.
//
// Panics if token could not be signed (actually unexpected).
func GenerateSigned() *token.BearerToken {
	tok := Generate()

	err := tok.SignTokenECDSA(neofsecdsatest.Key())
	if err != nil {
		panic(err)
	}

	return tok
}
