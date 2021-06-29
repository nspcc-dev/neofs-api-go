package rpc

import (
	"context"

	protoclient "github.com/nspcc-dev/neofs-api-go/rpc/client"
	"github.com/nspcc-dev/neofs-api-go/v2/accounting"
)

const serviceAccounting = serviceNamePrefix + "accounting.AccountingService"

const (
	rpcAccountingBalance = "Balance"
)

// BalancePrm groups the parameters of Balance call.
type BalancePrm struct {
	req accounting.BalanceRequest
}

// SetRequest sets request to send to the server.
//
// Required parameter.
func (x *BalancePrm) SetRequest(req accounting.BalanceRequest) {
	x.req = req
}

// BalanceRes groups the results of Balance call.
type BalanceRes struct {
	resp accounting.BalanceResponse
}

// Response returns the server response.
func (x *BalanceRes) Response() accounting.BalanceResponse {
	return x.resp
}

// Balance executes AccountingService.Balance RPC.
//
// Client connection should be established.
//
// If there is no error, res.Response() contains the server response.
//
// Context and res must not be nil.
func Balance(ctx context.Context, cli protoclient.Client, prm BalancePrm, res *BalanceRes) error {
	return sendUnaryRPC(ctx, cli, &prm.req, &res.resp, serviceAccounting, rpcAccountingBalance)
}
