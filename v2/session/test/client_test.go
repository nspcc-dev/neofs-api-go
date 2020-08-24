package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/session"
	sessionGRPC "github.com/nspcc-dev/neofs-api-go/v2/session/grpc"
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
	resp *session.CreateResponse
	err  error
}

func (s *testGRPCClient) Create(ctx context.Context, in *sessionGRPC.CreateRequest, opts ...grpc.CallOption) (*sessionGRPC.CreateResponse, error) {
	return s.server.Create(ctx, in)
}

func (s *testGRPCServer) Create(_ context.Context, req *sessionGRPC.CreateRequest) (*sessionGRPC.CreateResponse, error) {
	if s.err != nil {
		return nil, s.err
	}

	// verify request structure
	if err := signature.VerifyServiceMessage(
		session.CreateRequestFromGRPCMessage(req),
	); err != nil {
		return nil, err
	}

	// sign response structure
	if err := signature.SignServiceMessage(s.key, s.resp); err != nil {
		return nil, err
	}

	return session.CreateResponseToGRPCMessage(s.resp), nil
}

func testRequest() *session.CreateRequest {
	ownerID := new(refs.OwnerID)
	ownerID.SetValue([]byte{1, 2, 3})

	body := new(session.CreateRequestBody)
	body.SetOwnerID(ownerID)

	meta := new(session.RequestMetaHeader)
	meta.SetTTL(1)

	req := new(session.CreateRequest)
	req.SetBody(body)
	req.SetMetaHeader(meta)

	return req
}

func testResponse() *session.CreateResponse {
	body := new(session.CreateResponseBody)
	body.SetID([]byte{1, 2, 3})

	meta := new(session.ResponseMetaHeader)
	meta.SetTTL(1)

	resp := new(session.CreateResponse)
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

		c, err := session.NewClient(session.WithGRPCServiceClient(cli))
		require.NoError(t, err)

		resp, err := c.Create(ctx, new(session.CreateRequest))
		require.True(t, errors.Is(err, srvErr))
		require.Nil(t, resp)
	})

	t.Run("invalid request structure", func(t *testing.T) {
		req := testRequest()

		require.Error(t, signature.VerifyServiceMessage(req))

		c, err := session.NewClient(
			session.WithGRPCServiceClient(
				&testGRPCClient{
					server: new(testGRPCServer),
				},
			),
		)
		require.NoError(t, err)

		resp, err := c.Create(ctx, req)
		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("correct response", func(t *testing.T) {
		req := testRequest()

		require.NoError(t, signature.SignServiceMessage(cliKey, req))

		resp := testResponse()

		{ // w/o this require.Equal fails due to nil and []T{} difference
			meta := new(session.ResponseMetaHeader)
			meta.SetXHeaders([]*session.XHeader{})
			resp.SetMetaHeader(meta)
		}

		c, err := session.NewClient(
			session.WithGRPCServiceClient(
				&testGRPCClient{
					server: &testGRPCServer{
						key:  srvKey,
						resp: resp,
					},
				},
			),
		)
		require.NoError(t, err)

		r, err := c.Create(ctx, req)
		require.NoError(t, err)

		require.NoError(t, signature.VerifyServiceMessage(r))
		require.Equal(t, resp.GetBody(), r.GetBody())
		require.Equal(t, resp.GetMetaHeader(), r.GetMetaHeader())
	})
}
