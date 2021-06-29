package client

import (
	"context"
	"fmt"

	neofsecdsa "github.com/nspcc-dev/neofs-api-go/crypto/ecdsa"
	"github.com/nspcc-dev/neofs-api-go/pkg/accounting"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	v2accounting "github.com/nspcc-dev/neofs-api-go/v2/accounting"
	rpcapi "github.com/nspcc-dev/neofs-api-go/v2/rpc"
	v2signature "github.com/nspcc-dev/neofs-api-go/v2/signature"
)

func (x Client) GetBalance(ctx context.Context, owner *owner.ID, opts ...CallOption) (*accounting.Decimal, error) {
	// apply all available options
	callOptions := defaultCallOptions()

	for i := range opts {
		opts[i](callOptions)
	}

	reqBody := new(v2accounting.BalanceRequestBody)
	reqBody.SetOwnerID(owner.ToV2())

	var req v2accounting.BalanceRequest
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err := v2signature.SignServiceMessage(neofsecdsa.Signer(callOptions.key), &req)
	if err != nil {
		return nil, err
	}

	var prm rpcapi.BalancePrm

	prm.SetRequest(req)

	var res rpcapi.BalanceRes

	err = rpcapi.Balance(ctx, x.c, prm, &res)
	if err != nil {
		return nil, fmt.Errorf("transport error: %w", err)
	}

	resp := res.Response()

	err = v2signature.VerifyServiceMessage(&resp)
	if err != nil {
		return nil, fmt.Errorf("can't verify response message: %w", err)
	}

	return accounting.NewDecimalFromV2(resp.GetBody().GetBalance()), nil
}
