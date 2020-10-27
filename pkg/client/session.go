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

func (c Client) CreateSession(ctx context.Context, expiration uint64, opts ...CallOption) (*token.SessionToken, error) {
	switch c.remoteNode.Version.GetMajor() {
	case 2:
		return c.createSessionV2(ctx, expiration, opts...)
	default:
		return nil, unsupportedProtocolErr
	}
}

func (c Client) createSessionV2(ctx context.Context, expiration uint64, opts ...CallOption) (*token.SessionToken, error) {
	// apply all available options
	callOptions := c.defaultCallOptions()
	for i := range opts {
		opts[i].apply(&callOptions)
	}

	w, err := owner.NEO3WalletFromPublicKey(&c.key.PublicKey)
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

	err = v2signature.SignServiceMessage(c.key, req)
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
			return nil, errors.New("malformed response body")
		}

		sessionToken := token.NewSessionToken()
		sessionToken.SetID(body.GetID())
		sessionToken.SetSessionKey(body.GetSessionKey())
		sessionToken.SetOwnerID(ownerID)

		return sessionToken, nil
	default:
		return nil, unsupportedProtocolErr
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
			client.WithNetworkAddress(opts.addr)),
		)

	default:
		return nil, errors.New("lack of sdk client options to create accounting client")
	}

	// check if client correct and save in cache
	if err != nil {
		return nil, err
	}

	opts.grpcOpts.v2SessionClient = cli

	return cli, nil
}
