package token

import (
	"github.com/nspcc-dev/neofs-api-go/pkg/session"
)

// SessionToken represents NeoFS API v2-compatible
// session token.
//
// Deprecated: use session.Token instead
type SessionToken = session.Token

// NewSessionTokenFromV2 wraps session.SessionToken message structure
// into Token.
//
// Deprecated: use session.NewTokenFromV2 instead.
var NewSessionTokenFromV2 = session.NewTokenFromV2

// NewSessionToken creates and returns blank session token.
//
// Deprecated: use session.NewToken instead.
var NewSessionToken = session.NewToken
