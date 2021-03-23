package client

import (
	"context"

	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	"github.com/nspcc-dev/neofs-api-go/pkg/token"
	"github.com/nspcc-dev/neofs-api-go/rpc/client"
	rpcapi "github.com/nspcc-dev/neofs-api-go/v2/rpc"
	v2session "github.com/nspcc-dev/neofs-api-go/v2/session"
	v2signature "github.com/nspcc-dev/neofs-api-go/v2/signature"
	"github.com/pkg/errors"
)

// Session contains session-related methods.
type Session interface {
	// CreateSession creates session using provided expiration time.
	CreateSession(context.Context, uint64, ...CallOption) (*token.SessionToken, error)
}

var errMalformedResponseBody = errors.New("malformed response body")

func (c *clientImpl) CreateSession(ctx context.Context, expiration uint64, opts ...CallOption) (*token.SessionToken, error) {
	// apply all available options
	callOptions := c.defaultCallOptions()

	for i := range opts {
		opts[i](callOptions)
	}

	w, err := owner.NEO3WalletFromPublicKey(&callOptions.key.PublicKey)
	if err != nil {
		return nil, err
	}

	ownerID := new(owner.ID)
	ownerID.SetNeo3Wallet(w)

	reqBody := new(v2session.CreateRequestBody)
	reqBody.SetOwnerID(ownerID.ToV2())
	reqBody.SetExpiration(expiration)

	req := new(v2session.CreateRequest)
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err = v2signature.SignServiceMessage(callOptions.key, req)
	if err != nil {
		return nil, err
	}

	resp, err := rpcapi.CreateSession(c.Raw(), req, client.WithContext(ctx))
	if err != nil {
		return nil, errors.Wrap(err, "transport error")
	}

	err = v2signature.VerifyServiceMessage(resp, callOptions.signOpts()...)
	if err != nil {
		return nil, errors.Wrap(err, "can't verify response message")
	}

	body := resp.GetBody()
	if body == nil {
		return nil, errMalformedResponseBody
	}

	sessionToken := token.NewSessionToken()
	sessionToken.SetID(body.GetID())
	sessionToken.SetSessionKey(body.GetSessionKey())
	sessionToken.SetOwnerID(ownerID)

	return sessionToken, nil
}
