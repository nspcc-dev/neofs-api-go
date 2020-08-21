package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/object"
	objectGRPC "github.com/nspcc-dev/neofs-api-go/v2/object/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/service"
	"github.com/nspcc-dev/neofs-api-go/v2/signature"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

type testGRPCClient struct {
	server objectGRPC.ObjectServiceServer
}

func (s *testGRPCClient) Get(ctx context.Context, in *objectGRPC.GetRequest, opts ...grpc.CallOption) (objectGRPC.ObjectService_GetClient, error) {
	panic("implement me")
}

func (s *testGRPCClient) Put(ctx context.Context, opts ...grpc.CallOption) (objectGRPC.ObjectService_PutClient, error) {
	panic("implement me")
}

func (s *testGRPCClient) Delete(ctx context.Context, in *objectGRPC.DeleteRequest, opts ...grpc.CallOption) (*objectGRPC.DeleteResponse, error) {
	return s.server.Delete(ctx, in)
}

func (s *testGRPCClient) Head(ctx context.Context, in *objectGRPC.HeadRequest, opts ...grpc.CallOption) (*objectGRPC.HeadResponse, error) {
	return s.server.Head(ctx, in)
}

func (s *testGRPCClient) Search(ctx context.Context, in *objectGRPC.SearchRequest, opts ...grpc.CallOption) (objectGRPC.ObjectService_SearchClient, error) {
	panic("implement me")
}

func (s *testGRPCClient) GetRange(ctx context.Context, in *objectGRPC.GetRangeRequest, opts ...grpc.CallOption) (objectGRPC.ObjectService_GetRangeClient, error) {
	panic("implement me")
}

func (s *testGRPCClient) GetRangeHash(ctx context.Context, in *objectGRPC.GetRangeHashRequest, opts ...grpc.CallOption) (*objectGRPC.GetRangeHashResponse, error) {
	return s.server.GetRangeHash(ctx, in)
}

type testGRPCServer struct {
	key              *ecdsa.PrivateKey
	headResp         *object.HeadResponse
	delResp          *object.DeleteResponse
	getRangeHashResp *object.GetRangeHashResponse
	err              error
}

func (s *testGRPCServer) Get(request *objectGRPC.GetRequest, server objectGRPC.ObjectService_GetServer) error {
	panic("implement me")
}

func (s *testGRPCServer) Put(server objectGRPC.ObjectService_PutServer) error {
	panic("implement me")
}

func (s *testGRPCServer) Delete(ctx context.Context, request *objectGRPC.DeleteRequest) (*objectGRPC.DeleteResponse, error) {
	if s.err != nil {
		return nil, s.err
	}

	// verify request structure
	if err := signature.VerifyServiceMessage(
		object.DeleteRequestFromGRPCMessage(request),
	); err != nil {
		return nil, err
	}

	// sign response structure
	if err := signature.SignServiceMessage(s.key, s.delResp); err != nil {
		return nil, err
	}

	return object.DeleteResponseToGRPCMessage(s.delResp), nil
}

func (s *testGRPCServer) Head(ctx context.Context, request *objectGRPC.HeadRequest) (*objectGRPC.HeadResponse, error) {
	if s.err != nil {
		return nil, s.err
	}

	// verify request structure
	if err := signature.VerifyServiceMessage(
		object.HeadRequestFromGRPCMessage(request),
	); err != nil {
		return nil, err
	}

	// sign response structure
	if err := signature.SignServiceMessage(s.key, s.headResp); err != nil {
		return nil, err
	}

	return object.HeadResponseToGRPCMessage(s.headResp), nil
}

func (s *testGRPCServer) Search(request *objectGRPC.SearchRequest, server objectGRPC.ObjectService_SearchServer) error {
	panic("implement me")
}

func (s *testGRPCServer) GetRange(request *objectGRPC.GetRangeRequest, server objectGRPC.ObjectService_GetRangeServer) error {
	panic("implement me")
}

func (s *testGRPCServer) GetRangeHash(ctx context.Context, request *objectGRPC.GetRangeHashRequest) (*objectGRPC.GetRangeHashResponse, error) {
	if s.err != nil {
		return nil, s.err
	}

	// verify request structure
	if err := signature.VerifyServiceMessage(
		object.GetRangeHashRequestFromGRPCMessage(request),
	); err != nil {
		return nil, err
	}

	// sign response structure
	if err := signature.SignServiceMessage(s.key, s.getRangeHashResp); err != nil {
		return nil, err
	}

	return object.GetRangeHashResponseToGRPCMessage(s.getRangeHashResp), nil
}

func testHeadRequest() *object.HeadRequest {
	cid := new(refs.ContainerID)
	cid.SetValue([]byte{1, 2, 3})

	oid := new(refs.ObjectID)
	oid.SetValue([]byte{4, 5, 6})

	addr := new(refs.Address)
	addr.SetContainerID(cid)
	addr.SetObjectID(oid)

	body := new(object.HeadRequestBody)
	body.SetAddress(addr)

	meta := new(service.RequestMetaHeader)
	meta.SetTTL(1)
	meta.SetXHeaders([]*service.XHeader{})

	req := new(object.HeadRequest)
	req.SetBody(body)
	req.SetMetaHeader(meta)

	return req
}

func testHeadResponse() *object.HeadResponse {
	shortHdr := new(object.ShortHeader)
	shortHdr.SetCreationEpoch(100)

	hdrPart := new(object.GetHeaderPartShort)
	hdrPart.SetShortHeader(shortHdr)

	body := new(object.HeadResponseBody)
	body.SetHeaderPart(hdrPart)

	meta := new(service.ResponseMetaHeader)
	meta.SetTTL(1)
	meta.SetXHeaders([]*service.XHeader{})

	resp := new(object.HeadResponse)
	resp.SetBody(body)
	resp.SetMetaHeader(meta)

	return resp
}

func testDeleteRequest() *object.DeleteRequest {
	cid := new(refs.ContainerID)
	cid.SetValue([]byte{1, 2, 3})

	oid := new(refs.ObjectID)
	oid.SetValue([]byte{4, 5, 6})

	addr := new(refs.Address)
	addr.SetContainerID(cid)
	addr.SetObjectID(oid)

	body := new(object.DeleteRequestBody)
	body.SetAddress(addr)

	meta := new(service.RequestMetaHeader)
	meta.SetTTL(1)
	meta.SetXHeaders([]*service.XHeader{})

	req := new(object.DeleteRequest)
	req.SetBody(body)
	req.SetMetaHeader(meta)

	return req
}

func testDeleteResponse() *object.DeleteResponse {
	body := new(object.DeleteResponseBody)

	meta := new(service.ResponseMetaHeader)
	meta.SetTTL(1)
	meta.SetXHeaders([]*service.XHeader{})

	resp := new(object.DeleteResponse)
	resp.SetBody(body)
	resp.SetMetaHeader(meta)

	return resp
}

func testGetRangeHashRequest() *object.GetRangeHashRequest {
	cid := new(refs.ContainerID)
	cid.SetValue([]byte{1, 2, 3})

	oid := new(refs.ObjectID)
	oid.SetValue([]byte{4, 5, 6})

	addr := new(refs.Address)
	addr.SetContainerID(cid)
	addr.SetObjectID(oid)

	body := new(object.GetRangeHashRequestBody)
	body.SetAddress(addr)

	meta := new(service.RequestMetaHeader)
	meta.SetTTL(1)
	meta.SetXHeaders([]*service.XHeader{})

	req := new(object.GetRangeHashRequest)
	req.SetBody(body)
	req.SetMetaHeader(meta)

	return req
}

func testGetRangeHashResponse() *object.GetRangeHashResponse {
	body := new(object.GetRangeHashResponseBody)
	body.SetHashList([][]byte{{7, 8, 9}})

	meta := new(service.ResponseMetaHeader)
	meta.SetTTL(1)
	meta.SetXHeaders([]*service.XHeader{})

	resp := new(object.GetRangeHashResponse)
	resp.SetBody(body)
	resp.SetMetaHeader(meta)

	return resp
}

func TestGRPCClient_Head(t *testing.T) {
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

		c, err := object.New(object.WithGRPCServiceClient(cli))
		require.NoError(t, err)

		resp, err := c.Head(ctx, new(object.HeadRequest))
		require.True(t, errors.Is(err, srvErr))
		require.Nil(t, resp)
	})

	t.Run("invalid request structure", func(t *testing.T) {
		req := testHeadRequest()

		require.Error(t, signature.VerifyServiceMessage(req))

		c, err := object.New(
			object.WithGRPCServiceClient(
				&testGRPCClient{
					server: new(testGRPCServer),
				},
			),
		)
		require.NoError(t, err)

		resp, err := c.Head(ctx, req)
		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("correct response", func(t *testing.T) {
		req := testHeadRequest()

		require.NoError(t, signature.SignServiceMessage(cliKey, req))

		resp := testHeadResponse()

		c, err := object.New(
			object.WithGRPCServiceClient(
				&testGRPCClient{
					server: &testGRPCServer{
						key:      srvKey,
						headResp: resp,
					},
				},
			),
		)
		require.NoError(t, err)

		r, err := c.Head(ctx, req)
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

		c, err := object.New(object.WithGRPCServiceClient(cli))
		require.NoError(t, err)

		resp, err := c.Delete(ctx, new(object.DeleteRequest))
		require.True(t, errors.Is(err, srvErr))
		require.Nil(t, resp)
	})

	t.Run("invalid request structure", func(t *testing.T) {
		req := testDeleteRequest()

		require.Error(t, signature.VerifyServiceMessage(req))

		c, err := object.New(
			object.WithGRPCServiceClient(
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
		req := testDeleteRequest()

		require.NoError(t, signature.SignServiceMessage(cliKey, req))

		resp := testDeleteResponse()

		c, err := object.New(
			object.WithGRPCServiceClient(
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

func TestGRPCClient_GetRangeHash(t *testing.T) {
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

		c, err := object.New(object.WithGRPCServiceClient(cli))
		require.NoError(t, err)

		resp, err := c.GetRangeHash(ctx, new(object.GetRangeHashRequest))
		require.True(t, errors.Is(err, srvErr))
		require.Nil(t, resp)
	})

	t.Run("invalid request structure", func(t *testing.T) {
		req := testGetRangeHashRequest()

		require.Error(t, signature.VerifyServiceMessage(req))

		c, err := object.New(
			object.WithGRPCServiceClient(
				&testGRPCClient{
					server: new(testGRPCServer),
				},
			),
		)
		require.NoError(t, err)

		resp, err := c.GetRangeHash(ctx, req)
		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("correct response", func(t *testing.T) {
		req := testGetRangeHashRequest()

		require.NoError(t, signature.SignServiceMessage(cliKey, req))

		resp := testGetRangeHashResponse()

		c, err := object.New(
			object.WithGRPCServiceClient(
				&testGRPCClient{
					server: &testGRPCServer{
						key:              srvKey,
						getRangeHashResp: resp,
					},
				},
			),
		)
		require.NoError(t, err)

		r, err := c.GetRangeHash(ctx, req)
		require.NoError(t, err)

		require.NoError(t, signature.VerifyServiceMessage(r))
		require.Equal(t, resp.GetBody(), r.GetBody())
		require.Equal(t, resp.GetMetaHeader(), r.GetMetaHeader())
	})
}
