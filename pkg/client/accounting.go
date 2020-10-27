package client

import (
	"context"

	"github.com/nspcc-dev/neofs-api-go/pkg/accounting"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	v2accounting "github.com/nspcc-dev/neofs-api-go/v2/accounting"
	"github.com/nspcc-dev/neofs-api-go/v2/client"
	v2signature "github.com/nspcc-dev/neofs-api-go/v2/signature"
	"github.com/pkg/errors"
)

func (c Client) GetSelfBalance(ctx context.Context, opts ...CallOption) (*accounting.Decimal, error) {
	w, err := owner.NEO3WalletFromPublicKey(&c.key.PublicKey)
	if err != nil {
		return nil, err
	}

	ownerID := new(owner.ID)
	ownerID.SetNeo3Wallet(w)

	return c.GetBalance(ctx, ownerID, opts...)
}

func (c Client) GetBalance(ctx context.Context, owner *owner.ID, opts ...CallOption) (*accounting.Decimal, error) {
	// check remote node version
	switch c.remoteNode.Version.GetMajor() {
	case 2:
		return c.getBalanceV2(ctx, owner, opts...)
	default:
		return nil, unsupportedProtocolErr
	}
}

func (c Client) getBalanceV2(ctx context.Context, owner *owner.ID, opts ...CallOption) (*accounting.Decimal, error) {
	// apply all available options
	callOptions := c.defaultCallOptions()
	for i := range opts {
		opts[i].apply(&callOptions)
	}

	reqBody := new(v2accounting.BalanceRequestBody)
	reqBody.SetOwnerID(owner.ToV2())

	req := new(v2accounting.BalanceRequest)
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err := v2signature.SignServiceMessage(c.key, req)
	if err != nil {
		return nil, err
	}

	switch c.remoteNode.Protocol {
	case GRPC:
		cli, err := v2AccountingClientFromOptions(c.opts)
		if err != nil {
			return nil, errors.Wrap(err, "can't create grpc client")
		}

		resp, err := cli.Balance(ctx, req)
		if err != nil {
			return nil, errors.Wrap(err, "transport error")
		}

		err = v2signature.VerifyServiceMessage(resp)
		if err != nil {
			return nil, errors.Wrap(err, "can't verify response message")
		}

		return &accounting.Decimal{
			Decimal: *resp.GetBody().GetBalance(),
		}, nil
	default:
		return nil, unsupportedProtocolErr
	}
}

func v2AccountingClientFromOptions(opts *clientOptions) (cli *v2accounting.Client, err error) {
	switch {
	case opts.grpcOpts.v2AccountingClient != nil:
		// return value from client cache
		return opts.grpcOpts.v2AccountingClient, nil

	case opts.grpcOpts.conn != nil:
		cli, err = v2accounting.NewClient(v2accounting.WithGlobalOpts(
			client.WithGRPCConn(opts.grpcOpts.conn)),
		)

	case opts.addr != "":
		cli, err = v2accounting.NewClient(v2accounting.WithGlobalOpts(
			client.WithNetworkAddress(opts.addr)),
		)

	default:
		return nil, errors.New("lack of sdk client options to create accounting client")
	}

	// check if client correct and save in cache
	if err != nil {
		return nil, err
	}

	opts.grpcOpts.v2AccountingClient = cli

	return cli, nil
}
