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

// Accounting contains methods related to balance querying.
type Accounting interface {
	// GetSelfBalance returns balance of the account deduced from client's key.
	GetSelfBalance(context.Context, ...CallOption) (*accounting.Decimal, error)
	// GetBalance returns balance of provided account.
	GetBalance(context.Context, *owner.ID, ...CallOption) (*accounting.Decimal, error)
}

func (c clientImpl) GetSelfBalance(ctx context.Context, opts ...CallOption) (*accounting.Decimal, error) {
	return c.GetBalance(ctx, nil, opts...)
}

func (c clientImpl) GetBalance(ctx context.Context, owner *owner.ID, opts ...CallOption) (*accounting.Decimal, error) {
	// check remote node version
	switch c.remoteNode.Version.Major() {
	case 2:
		return c.getBalanceV2(ctx, owner, opts...)
	default:
		return nil, errUnsupportedProtocol
	}
}

func (c clientImpl) getBalanceV2(ctx context.Context, ownerID *owner.ID, opts ...CallOption) (*accounting.Decimal, error) {
	// apply all available options
	callOptions := c.defaultCallOptions()

	for i := range opts {
		opts[i].apply(&callOptions)
	}

	if ownerID == nil {
		w, err := owner.NEO3WalletFromPublicKey(&callOptions.key.PublicKey)
		if err != nil {
			return nil, err
		}

		ownerID = new(owner.ID)
		ownerID.SetNeo3Wallet(w)
	}

	reqBody := new(v2accounting.BalanceRequestBody)
	reqBody.SetOwnerID(ownerID.ToV2())

	req := new(v2accounting.BalanceRequest)
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err := v2signature.SignServiceMessage(callOptions.key, req)
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

		return accounting.NewDecimalFromV2(resp.GetBody().GetBalance()), nil
	default:
		return nil, errUnsupportedProtocol
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
			client.WithNetworkAddress(opts.addr),
			client.WithDialTimeout(opts.dialTimeout),
		))

	default:
		return nil, errOptionsLack("Accounting")
	}

	// check if client correct and save in cache
	if err != nil {
		return nil, err
	}

	opts.grpcOpts.v2AccountingClient = cli

	return cli, nil
}
