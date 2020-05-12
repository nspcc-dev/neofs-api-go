package session

import (
	"github.com/nspcc-dev/neofs-api-go/refs"
	"github.com/nspcc-dev/neofs-api-go/service"
)

// OwnerID is a type alias of OwnerID ref.
type OwnerID = refs.OwnerID

// TokenID is a type alias of TokenID ref.
type TokenID = service.TokenID

// Token is a type alias of Token.
type Token = service.Token
