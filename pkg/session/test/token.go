package sessiontest

import (
	"math/rand"

	"github.com/google/uuid"
	neofsecdsatest "github.com/nspcc-dev/neofs-api-go/crypto/ecdsa/test"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	"github.com/nspcc-dev/neofs-api-go/pkg/session"
)

// Generate returns random session.Token.
//
// Resulting token is unsigned.
func Generate() *session.Token {
	tok := session.NewToken()

	uid, err := uuid.New().MarshalBinary()
	if err != nil {
		panic(err)
	}

	w := new(owner.NEO3Wallet)
	rand.Read(w.Bytes())

	ownerID := owner.NewID()
	ownerID.SetNeo3Wallet(w)

	tok.SetID(uid)
	tok.SetOwnerID(ownerID)
	tok.SetSessionKey(neofsecdsatest.PublicBytes())
	tok.SetExp(11)
	tok.SetNbf(22)
	tok.SetIat(33)

	return tok
}

// GenerateSigned returns signed random session.Token.
//
// Panics if token could not be signed (actually unexpected).
func GenerateSigned() *session.Token {
	tok := Generate()

	err := tok.SignECDSA(neofsecdsatest.Key())
	if err != nil {
		panic(err)
	}

	return tok
}
