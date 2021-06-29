package rpc

import (
	"context"

	protoclient "github.com/nspcc-dev/neofs-api-go/rpc/client"
	"github.com/nspcc-dev/neofs-api-go/v2/session"
)

const serviceSession = serviceNamePrefix + "session.SessionService"

const (
	rpcSessionCreate = "Create"
)

// CreateSessionPrm groups the parameters of CreateSession call.
type CreateSessionPrm struct {
	req session.CreateRequest
}

// SetRequest sets request to send to the server.
//
// Required parameter.
func (x *CreateSessionPrm) SetRequest(req session.CreateRequest) {
	x.req = req
}

// CreateSessionRes groups the results of CreateSession call.
type CreateSessionRes struct {
	resp session.CreateResponse
}

// Response returns the server response.
func (x *CreateSessionRes) Response() session.CreateResponse {
	return x.resp
}

// CreateSession executes SessionService.Create RPC.
//
// Client connection should be established.
//
// If there is no error, res.Response() contains the server response.
//
// Context and res must not be nil.
func CreateSession(ctx context.Context, cli protoclient.Client, prm CreateSessionPrm, res *CreateSessionRes) error {
	return sendUnaryRPC(ctx, cli, &prm.req, &res.resp, serviceSession, rpcSessionCreate)
}
