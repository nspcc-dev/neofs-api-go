package client

import (
	"context"

	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	"github.com/nspcc-dev/neofs-api-go/pkg/token"
	"github.com/nspcc-dev/neofs-api-go/v2/client"
	v2session "github.com/nspcc-dev/neofs-api-go/v2/session"
	v2signature "github.com/nspcc-dev/neofs-api-go/v2/signature"
	"github.com/pkg/errors"
)

var errMalformedResponseBody = errors.New("malformed response body")

func (c Client) CreateSession(ctx context.Context, expiration uint64, opts ...CallOption) (*token.SessionToken, error) {
	switch c.remoteNode.Version.Major() {
	case 2:
		return c.createSessionV2(ctx, expiration, opts...)
	default:
		return nil, errUnsupportedProtocol
	}
}

func (c Client) createSessionV2(ctx context.Context, expiration uint64, opts ...CallOption) (*token.SessionToken, error) {
	// apply all available options
	callOptions := c.defaultCallOptions()

	for i := range opts {
		opts[i].apply(&callOptions)
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

	switch c.remoteNode.Protocol {
	case GRPC:
		cli, err := v2SessionClientFromOptions(c.opts)
		if err != nil {
			return nil, errors.Wrap(err, "can't create grpc client")
		}

		resp, err := cli.Create(ctx, req)
		if err != nil {
			return nil, errors.Wrap(err, "transport error")
		}

		err = v2signature.VerifyServiceMessage(resp)
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
	default:
		return nil, errUnsupportedProtocol
	}
}

func v2SessionClientFromOptions(opts *clientOptions) (cli *v2session.Client, err error) {
	switch {
	case opts.grpcOpts.v2SessionClient != nil:
		// return value from client cache
		return opts.grpcOpts.v2SessionClient, nil

	case opts.grpcOpts.conn != nil:
		cli, err = v2session.NewClient(v2session.WithGlobalOpts(
			client.WithGRPCConn(opts.grpcOpts.conn)),
		)

	case opts.addr != "":
		cli, err = v2session.NewClient(v2session.WithGlobalOpts(
			client.WithNetworkAddress(opts.addr),
			client.WithDialTimeout(opts.dialTimeout),
		))

	default:
		return nil, errOptionsLack("Session")
	}

	// check if client correct and save in cache
	if err != nil {
		return nil, err
	}

	opts.grpcOpts.v2SessionClient = cli

	return cli, nil
}

// AttachSessionToken attaches session token to client.
//
// Provided token is attached to all requests without WithSession option.
// Use WithSession(nil) option in order to send request without session token.
func (c *Client) AttachSessionToken(token *token.SessionToken) {
	c.sessionToken = token
}

// AttachBearerToken attaches bearer token to client.
//
// Provided bearer is attached to all requests without WithBearer option.
// Use WithBearer(nil) option in order to send request without bearer token.
func (c *Client) AttachBearerToken(token *token.BearerToken) {
	c.bearerToken = token
}
