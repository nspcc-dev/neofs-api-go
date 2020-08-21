package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/acl"
	"github.com/nspcc-dev/neofs-api-go/v2/container"
	containerGRPC "github.com/nspcc-dev/neofs-api-go/v2/container/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/session"
	"github.com/nspcc-dev/neofs-api-go/v2/signature"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

type testGRPCClient struct {
	server containerGRPC.ContainerServiceServer
}

type testGRPCServer struct {
	key       *ecdsa.PrivateKey
	putResp   *container.PutResponse
	getResp   *container.GetResponse
	delResp   *container.DeleteResponse
	listResp  *container.ListResponse
	sEaclResp *container.SetExtendedACLResponse
	gEaclResp *container.GetExtendedACLResponse
	err       error
}

func (s *testGRPCClient) Put(ctx context.Context, in *containerGRPC.PutRequest, opts ...grpc.CallOption) (*containerGRPC.PutResponse, error) {
	return s.server.Put(ctx, in)
}

func (s *testGRPCClient) Delete(ctx context.Context, in *containerGRPC.DeleteRequest, opts ...grpc.CallOption) (*containerGRPC.DeleteResponse, error) {
	return s.server.Delete(ctx, in)
}

func (s *testGRPCClient) Get(ctx context.Context, in *containerGRPC.GetRequest, opts ...grpc.CallOption) (*containerGRPC.GetResponse, error) {
	return s.server.Get(ctx, in)
}

func (s *testGRPCClient) List(ctx context.Context, in *containerGRPC.ListRequest, opts ...grpc.CallOption) (*containerGRPC.ListResponse, error) {
	return s.server.List(ctx, in)
}

func (s *testGRPCClient) SetExtendedACL(ctx context.Context, in *containerGRPC.SetExtendedACLRequest, opts ...grpc.CallOption) (*containerGRPC.SetExtendedACLResponse, error) {
	return s.server.SetExtendedACL(ctx, in)
}

func (s *testGRPCClient) GetExtendedACL(ctx context.Context, in *containerGRPC.GetExtendedACLRequest, opts ...grpc.CallOption) (*containerGRPC.GetExtendedACLResponse, error) {
	return s.server.GetExtendedACL(ctx, in)
}

func (s *testGRPCServer) Put(_ context.Context, req *containerGRPC.PutRequest) (*containerGRPC.PutResponse, error) {
	if s.err != nil {
		return nil, s.err
	}

	// verify request structure
	if err := signature.VerifyServiceMessage(
		container.PutRequestFromGRPCMessage(req),
	); err != nil {
		return nil, err
	}

	// sign response structure
	if err := signature.SignServiceMessage(s.key, s.putResp); err != nil {
		return nil, err
	}

	return container.PutResponseToGRPCMessage(s.putResp), nil
}

func (s *testGRPCServer) Delete(_ context.Context, req *containerGRPC.DeleteRequest) (*containerGRPC.DeleteResponse, error) {
	if s.err != nil {
		return nil, s.err
	}

	// verify request structure
	if err := signature.VerifyServiceMessage(
		container.DeleteRequestFromGRPCMessage(req),
	); err != nil {
		return nil, err
	}

	// sign response structure
	if err := signature.SignServiceMessage(s.key, s.delResp); err != nil {
		return nil, err
	}

	return container.DeleteResponseToGRPCMessage(s.delResp), nil
}

func (s *testGRPCServer) Get(_ context.Context, req *containerGRPC.GetRequest) (*containerGRPC.GetResponse, error) {
	if s.err != nil {
		return nil, s.err
	}

	// verify request structure
	if err := signature.VerifyServiceMessage(
		container.GetRequestFromGRPCMessage(req),
	); err != nil {
		return nil, err
	}

	// sign response structure
	if err := signature.SignServiceMessage(s.key, s.getResp); err != nil {
		return nil, err
	}

	return container.GetResponseToGRPCMessage(s.getResp), nil
}

func (s *testGRPCServer) List(_ context.Context, req *containerGRPC.ListRequest) (*containerGRPC.ListResponse, error) {
	if s.err != nil {
		return nil, s.err
	}

	// verify request structure
	if err := signature.VerifyServiceMessage(
		container.ListRequestFromGRPCMessage(req),
	); err != nil {
		return nil, err
	}

	// sign response structure
	if err := signature.SignServiceMessage(s.key, s.listResp); err != nil {
		return nil, err
	}

	return container.ListResponseToGRPCMessage(s.listResp), nil
}

func (s *testGRPCServer) SetExtendedACL(_ context.Context, req *containerGRPC.SetExtendedACLRequest) (*containerGRPC.SetExtendedACLResponse, error) {
	if s.err != nil {
		return nil, s.err
	}

	// verify request structure
	if err := signature.VerifyServiceMessage(
		container.SetExtendedACLRequestFromGRPCMessage(req),
	); err != nil {
		return nil, err
	}

	// sign response structure
	if err := signature.SignServiceMessage(s.key, s.sEaclResp); err != nil {
		return nil, err
	}

	return container.SetExtendedACLResponseToGRPCMessage(s.sEaclResp), nil
}

func (s *testGRPCServer) GetExtendedACL(_ context.Context, req *containerGRPC.GetExtendedACLRequest) (*containerGRPC.GetExtendedACLResponse, error) {
	if s.err != nil {
		return nil, s.err
	}

	// verify request structure
	if err := signature.VerifyServiceMessage(
		container.GetExtendedACLRequestFromGRPCMessage(req),
	); err != nil {
		return nil, err
	}

	// sign response structure
	if err := signature.SignServiceMessage(s.key, s.gEaclResp); err != nil {
		return nil, err
	}

	return container.GetExtendedACLResponseToGRPCMessage(s.gEaclResp), nil
}

func testPutRequest() *container.PutRequest {
	cnr := new(container.Container)
	cnr.SetBasicACL(1)

	body := new(container.PutRequestBody)
	body.SetContainer(cnr)

	meta := new(session.RequestMetaHeader)
	meta.SetTTL(1)

	req := new(container.PutRequest)
	req.SetBody(body)
	req.SetMetaHeader(meta)

	return req
}

func testPutResponse() *container.PutResponse {
	cid := new(refs.ContainerID)
	cid.SetValue([]byte{1, 2, 3})

	body := new(container.PutResponseBody)
	body.SetContainerID(cid)

	meta := new(session.ResponseMetaHeader)
	meta.SetTTL(1)
	meta.SetXHeaders([]*session.XHeader{}) // w/o this require.Equal fails due to nil and []T{} difference

	resp := new(container.PutResponse)
	resp.SetBody(body)
	resp.SetMetaHeader(meta)

	return resp
}

func testGetRequest() *container.GetRequest {
	cid := new(refs.ContainerID)
	cid.SetValue([]byte{1, 2, 3})

	body := new(container.GetRequestBody)
	body.SetContainerID(cid)

	meta := new(session.RequestMetaHeader)
	meta.SetTTL(1)

	req := new(container.GetRequest)
	req.SetBody(body)
	req.SetMetaHeader(meta)

	return req
}

func testGetResponse() *container.GetResponse {
	cnr := new(container.Container)
	cnr.SetAttributes([]*container.Attribute{}) // w/o this require.Equal fails due to nil and []T{} difference

	body := new(container.GetResponseBody)
	body.SetContainer(cnr)

	meta := new(session.ResponseMetaHeader)
	meta.SetTTL(1)
	meta.SetXHeaders([]*session.XHeader{}) // w/o this require.Equal fails due to nil and []T{} difference

	resp := new(container.GetResponse)
	resp.SetBody(body)
	resp.SetMetaHeader(meta)

	return resp
}

func testDelRequest() *container.DeleteRequest {
	cid := new(refs.ContainerID)
	cid.SetValue([]byte{1, 2, 3})

	body := new(container.DeleteRequestBody)
	body.SetContainerID(cid)

	meta := new(session.RequestMetaHeader)
	meta.SetTTL(1)

	req := new(container.DeleteRequest)
	req.SetBody(body)
	req.SetMetaHeader(meta)

	return req
}

func testDelResponse() *container.DeleteResponse {
	body := new(container.DeleteResponseBody)

	meta := new(session.ResponseMetaHeader)
	meta.SetTTL(1)
	meta.SetXHeaders([]*session.XHeader{}) // w/o this require.Equal fails due to nil and []T{} difference

	resp := new(container.DeleteResponse)
	resp.SetBody(body)
	resp.SetMetaHeader(meta)

	return resp
}

func testListRequest() *container.ListRequest {
	ownerID := new(refs.OwnerID)
	ownerID.SetValue([]byte{1, 2, 3})

	body := new(container.ListRequestBody)
	body.SetOwnerID(ownerID)

	meta := new(session.RequestMetaHeader)
	meta.SetTTL(1)

	req := new(container.ListRequest)
	req.SetBody(body)
	req.SetMetaHeader(meta)

	return req
}

func testListResponse() *container.ListResponse {
	cid := new(refs.ContainerID)
	cid.SetValue([]byte{1, 2, 3})

	body := new(container.ListResponseBody)
	body.SetContainerIDs([]*refs.ContainerID{cid})

	meta := new(session.ResponseMetaHeader)
	meta.SetTTL(1)
	meta.SetXHeaders([]*session.XHeader{}) // w/o this require.Equal fails due to nil and []T{} difference

	resp := new(container.ListResponse)
	resp.SetBody(body)
	resp.SetMetaHeader(meta)

	return resp
}

func testSetEACLRequest() *container.SetExtendedACLRequest {
	cid := new(refs.ContainerID)
	cid.SetValue([]byte{1, 2, 3})

	eacl := new(acl.Table)
	eacl.SetContainerID(cid)

	body := new(container.SetExtendedACLRequestBody)
	body.SetEACL(eacl)

	meta := new(session.RequestMetaHeader)
	meta.SetTTL(1)

	req := new(container.SetExtendedACLRequest)
	req.SetBody(body)
	req.SetMetaHeader(meta)

	return req
}

func testSetEACLResponse() *container.SetExtendedACLResponse {
	body := new(container.SetExtendedACLResponseBody)

	meta := new(session.ResponseMetaHeader)
	meta.SetTTL(1)
	meta.SetXHeaders([]*session.XHeader{}) // w/o this require.Equal fails due to nil and []T{} difference

	resp := new(container.SetExtendedACLResponse)
	resp.SetBody(body)
	resp.SetMetaHeader(meta)

	return resp
}

func testGetEACLRequest() *container.GetExtendedACLRequest {
	cid := new(refs.ContainerID)
	cid.SetValue([]byte{1, 2, 3})

	body := new(container.GetExtendedACLRequestBody)
	body.SetContainerID(cid)

	meta := new(session.RequestMetaHeader)
	meta.SetTTL(1)

	req := new(container.GetExtendedACLRequest)
	req.SetBody(body)
	req.SetMetaHeader(meta)

	return req
}

func testGetEACLResponse() *container.GetExtendedACLResponse {
	cid := new(refs.ContainerID)
	cid.SetValue([]byte{1, 2, 3})

	eacl := new(acl.Table)
	eacl.SetContainerID(cid)
	eacl.SetRecords([]*acl.Record{}) // w/o this require.Equal fails due to nil and []T{} difference

	body := new(container.GetExtendedACLResponseBody)
	body.SetEACL(eacl)

	meta := new(session.ResponseMetaHeader)
	meta.SetTTL(1)
	meta.SetXHeaders([]*session.XHeader{}) // w/o this require.Equal fails due to nil and []T{} difference

	resp := new(container.GetExtendedACLResponse)
	resp.SetBody(body)
	resp.SetMetaHeader(meta)

	return resp
}

func TestGRPCClient_Put(t *testing.T) {
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

		c, err := container.New(container.WithGRPCServiceClient(cli))
		require.NoError(t, err)

		resp, err := c.Put(ctx, new(container.PutRequest))
		require.True(t, errors.Is(err, srvErr))
		require.Nil(t, resp)
	})

	t.Run("invalid request structure", func(t *testing.T) {
		req := testPutRequest()

		require.Error(t, signature.VerifyServiceMessage(req))

		c, err := container.New(
			container.WithGRPCServiceClient(
				&testGRPCClient{
					server: new(testGRPCServer),
				},
			),
		)
		require.NoError(t, err)

		resp, err := c.Put(ctx, req)
		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("correct response", func(t *testing.T) {
		req := testPutRequest()

		require.NoError(t, signature.SignServiceMessage(cliKey, req))

		resp := testPutResponse()

		c, err := container.New(
			container.WithGRPCServiceClient(
				&testGRPCClient{
					server: &testGRPCServer{
						key:     srvKey,
						putResp: resp,
					},
				},
			),
		)
		require.NoError(t, err)

		r, err := c.Put(ctx, req)
		require.NoError(t, err)

		require.NoError(t, signature.VerifyServiceMessage(r))
		require.Equal(t, resp.GetBody(), r.GetBody())
		require.Equal(t, resp.GetMetaHeader(), r.GetMetaHeader())
	})
}

func TestGRPCClient_Get(t *testing.T) {
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

		c, err := container.New(container.WithGRPCServiceClient(cli))
		require.NoError(t, err)

		resp, err := c.Get(ctx, new(container.GetRequest))
		require.True(t, errors.Is(err, srvErr))
		require.Nil(t, resp)
	})

	t.Run("invalid request structure", func(t *testing.T) {
		req := testGetRequest()

		require.Error(t, signature.VerifyServiceMessage(req))

		c, err := container.New(
			container.WithGRPCServiceClient(
				&testGRPCClient{
					server: new(testGRPCServer),
				},
			),
		)
		require.NoError(t, err)

		resp, err := c.Get(ctx, req)
		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("correct response", func(t *testing.T) {
		req := testGetRequest()

		require.NoError(t, signature.SignServiceMessage(cliKey, req))

		resp := testGetResponse()

		c, err := container.New(
			container.WithGRPCServiceClient(
				&testGRPCClient{
					server: &testGRPCServer{
						key:     srvKey,
						getResp: resp,
					},
				},
			),
		)
		require.NoError(t, err)

		r, err := c.Get(ctx, req)
		require.NoError(t, err)

		require.NoError(t, signature.VerifyServiceMessage(r))
		require.Equal(t, resp.GetBody(), r.GetBody())
		require.Equal(t, resp.GetMetaHeader(), r.GetMetaHeader())
	})
}

func TestGRPCClient_Delete(t *testing.T) {
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

		c, err := container.New(container.WithGRPCServiceClient(cli))
		require.NoError(t, err)

		resp, err := c.Delete(ctx, new(container.DeleteRequest))
		require.True(t, errors.Is(err, srvErr))
		require.Nil(t, resp)
	})

	t.Run("invalid request structure", func(t *testing.T) {
		req := testDelRequest()

		require.Error(t, signature.VerifyServiceMessage(req))

		c, err := container.New(
			container.WithGRPCServiceClient(
				&testGRPCClient{
					server: new(testGRPCServer),
				},
			),
		)
		require.NoError(t, err)

		resp, err := c.Delete(ctx, req)
		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("correct response", func(t *testing.T) {
		req := testDelRequest()

		require.NoError(t, signature.SignServiceMessage(cliKey, req))

		resp := testDelResponse()

		c, err := container.New(
			container.WithGRPCServiceClient(
				&testGRPCClient{
					server: &testGRPCServer{
						key:     srvKey,
						delResp: resp,
					},
				},
			),
		)
		require.NoError(t, err)

		r, err := c.Delete(ctx, req)
		require.NoError(t, err)

		require.NoError(t, signature.VerifyServiceMessage(r))
		require.Equal(t, resp.GetBody(), r.GetBody())
		require.Equal(t, resp.GetMetaHeader(), r.GetMetaHeader())
	})
}

func TestGRPCClient_List(t *testing.T) {
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

		c, err := container.New(container.WithGRPCServiceClient(cli))
		require.NoError(t, err)

		resp, err := c.List(ctx, new(container.ListRequest))
		require.True(t, errors.Is(err, srvErr))
		require.Nil(t, resp)
	})

	t.Run("invalid request structure", func(t *testing.T) {
		req := testListRequest()

		require.Error(t, signature.VerifyServiceMessage(req))

		c, err := container.New(
			container.WithGRPCServiceClient(
				&testGRPCClient{
					server: new(testGRPCServer),
				},
			),
		)
		require.NoError(t, err)

		resp, err := c.List(ctx, req)
		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("correct response", func(t *testing.T) {
		req := testListRequest()

		require.NoError(t, signature.SignServiceMessage(cliKey, req))

		resp := testListResponse()

		c, err := container.New(
			container.WithGRPCServiceClient(
				&testGRPCClient{
					server: &testGRPCServer{
						key:      srvKey,
						listResp: resp,
					},
				},
			),
		)
		require.NoError(t, err)

		r, err := c.List(ctx, req)
		require.NoError(t, err)

		require.NoError(t, signature.VerifyServiceMessage(r))
		require.Equal(t, resp.GetBody(), r.GetBody())
		require.Equal(t, resp.GetMetaHeader(), r.GetMetaHeader())
	})
}

func TestGRPCClient_SetEACL(t *testing.T) {
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

		c, err := container.New(container.WithGRPCServiceClient(cli))
		require.NoError(t, err)

		resp, err := c.SetExtendedACL(ctx, new(container.SetExtendedACLRequest))
		require.True(t, errors.Is(err, srvErr))
		require.Nil(t, resp)
	})
	t.Run("invalid request structure", func(t *testing.T) {
		req := testSetEACLRequest()

		require.Error(t, signature.VerifyServiceMessage(req))

		c, err := container.New(
			container.WithGRPCServiceClient(
				&testGRPCClient{
					server: new(testGRPCServer),
				},
			),
		)
		require.NoError(t, err)

		resp, err := c.SetExtendedACL(ctx, req)
		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("correct response", func(t *testing.T) {
		req := testSetEACLRequest()

		require.NoError(t, signature.SignServiceMessage(cliKey, req))

		resp := testSetEACLResponse()

		c, err := container.New(
			container.WithGRPCServiceClient(
				&testGRPCClient{
					server: &testGRPCServer{
						key:       srvKey,
						sEaclResp: resp,
					},
				},
			),
		)
		require.NoError(t, err)

		r, err := c.SetExtendedACL(ctx, req)
		require.NoError(t, err)

		require.NoError(t, signature.VerifyServiceMessage(r))
		require.Equal(t, resp.GetBody(), r.GetBody())
		require.Equal(t, resp.GetMetaHeader(), r.GetMetaHeader())
	})
}

func TestGRPCClient_GetEACL(t *testing.T) {
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

		c, err := container.New(container.WithGRPCServiceClient(cli))
		require.NoError(t, err)

		resp, err := c.GetExtendedACL(ctx, new(container.GetExtendedACLRequest))
		require.True(t, errors.Is(err, srvErr))
		require.Nil(t, resp)
	})
	t.Run("invalid request structure", func(t *testing.T) {
		req := testGetEACLRequest()

		require.Error(t, signature.VerifyServiceMessage(req))

		c, err := container.New(
			container.WithGRPCServiceClient(
				&testGRPCClient{
					server: new(testGRPCServer),
				},
			),
		)
		require.NoError(t, err)

		resp, err := c.GetExtendedACL(ctx, req)
		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("correct response", func(t *testing.T) {
		req := testGetEACLRequest()

		require.NoError(t, signature.SignServiceMessage(cliKey, req))

		resp := testGetEACLResponse()

		c, err := container.New(
			container.WithGRPCServiceClient(
				&testGRPCClient{
					server: &testGRPCServer{
						key:       srvKey,
						gEaclResp: resp,
					},
				},
			),
		)
		require.NoError(t, err)

		r, err := c.GetExtendedACL(ctx, req)
		require.NoError(t, err)

		require.NoError(t, signature.VerifyServiceMessage(r))
		require.Equal(t, resp.GetBody(), r.GetBody())
		require.Equal(t, resp.GetMetaHeader(), r.GetMetaHeader())
	})
}
