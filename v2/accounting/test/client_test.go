package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/accounting"
	accountingGRPC "github.com/nspcc-dev/neofs-api-go/v2/accounting/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/service"
	"github.com/nspcc-dev/neofs-api-go/v2/signature"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

type testGRPCClient struct {
	server *testGRPCServer
}

type testGRPCServer struct {
	key  *ecdsa.PrivateKey
	resp *accounting.BalanceResponse
	err  error
}

func (s *testGRPCClient) Balance(ctx context.Context, in *accountingGRPC.BalanceRequest, opts ...grpc.CallOption) (*accountingGRPC.BalanceResponse, error) {
	return s.server.Balance(ctx, in)
}

func (s *testGRPCServer) Balance(_ context.Context, req *accountingGRPC.BalanceRequest) (*accountingGRPC.BalanceResponse, error) {
	if s.err != nil {
		return nil, s.err
	}

	// verify request structure
	if err := signature.VerifyServiceMessage(
		accounting.BalanceRequestFromGRPCMessage(req),
	); err != nil {
		return nil, err
	}

	// sign response structure
	if err := signature.SignServiceMessage(s.key, s.resp); err != nil {
		return nil, err
	}

	return accounting.BalanceResponseToGRPCMessage(s.resp), nil
}

func testRequest() *accounting.BalanceRequest {
	ownerID := new(refs.OwnerID)
	ownerID.SetValue([]byte{1, 2, 3})

	body := new(accounting.BalanceRequestBody)
	body.SetOwnerID(ownerID)

	meta := new(service.RequestMetaHeader)
	meta.SetTTL(1)

	req := new(accounting.BalanceRequest)
	req.SetBody(body)
	req.SetMetaHeader(meta)

	return req
}

func testResponse() *accounting.BalanceResponse {
	dec := new(accounting.Decimal)
	dec.SetValue(10)

	body := new(accounting.BalanceResponseBody)
	body.SetBalance(dec)

	meta := new(service.ResponseMetaHeader)
	meta.SetTTL(1)

	resp := new(accounting.BalanceResponse)
	resp.SetBody(body)
	resp.SetMetaHeader(meta)

	return resp
}

func TestGRPCClient(t *testing.T) {
	ctx := context.TODO()

	cliKey := test.DecodeKey(0)
	srvKey := test.DecodeKey(1)

	t.Run("gRPC server error", func(t *testing.T) {
		srvErr := errors.New("test server error")

		srv := &testGRPCServer{
			err: srvErr,
		}

		cli := &testGRPCClient{
			server: srv,
		}

		c, err := accounting.New(accounting.WithGRPCServiceClient(cli))
		require.NoError(t, err)

		resp, err := c.Balance(ctx, new(accounting.BalanceRequest))
		require.True(t, errors.Is(err, srvErr))
		require.Nil(t, resp)
	})

	t.Run("invalid request structure", func(t *testing.T) {
		req := testRequest()

		require.Error(t, signature.VerifyServiceMessage(req))

		c, err := accounting.New(
			accounting.WithGRPCServiceClient(
				&testGRPCClient{
					server: new(testGRPCServer),
				},
			),
		)
		require.NoError(t, err)

		resp, err := c.Balance(ctx, req)
		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("correct response", func(t *testing.T) {
		req := testRequest()

		require.NoError(t, signature.SignServiceMessage(cliKey, req))

		resp := testResponse()

		{ // w/o this require.Equal fails due to nil and []T{} difference
			meta := new(service.ResponseMetaHeader)
			meta.SetXHeaders([]*service.XHeader{})
			resp.SetMetaHeader(meta)
		}

		c, err := accounting.New(
			accounting.WithGRPCServiceClient(
				&testGRPCClient{
					server: &testGRPCServer{
						key:  srvKey,
						resp: resp,
					},
				},
			),
		)
		require.NoError(t, err)

		r, err := c.Balance(ctx, req)
		require.NoError(t, err)

		require.NoError(t, signature.VerifyServiceMessage(r))
		require.Equal(t, resp.GetBody(), r.GetBody())
		require.Equal(t, resp.GetMetaHeader(), r.GetMetaHeader())
	})
}
