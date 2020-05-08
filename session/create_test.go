package session

import (
	"context"
	"crypto/ecdsa"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/service"
	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

type testSessionClient struct {
	fn   func(*CreateRequest)
	resp *CreateResponse
	err  error
}

func (s testSessionClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	if s.fn != nil {
		s.fn(in)
	}

	return s.resp, s.err
}

func TestNewGRPCCreator(t *testing.T) {
	var (
		err  error
		conn = new(grpc.ClientConn)
		sk   = new(ecdsa.PrivateKey)
	)

	// nil client connection
	_, err = NewGRPCCreator(nil, sk)
	require.EqualError(t, err, ErrNilGPRCClientConn.Error())

	// nil private key
	_, err = NewGRPCCreator(conn, nil)
	require.EqualError(t, err, crypto.ErrEmptyPrivateKey.Error())

	// valid params
	res, err := NewGRPCCreator(conn, sk)
	require.NoError(t, err)

	v := res.(*gRPCCreator)
	require.Equal(t, conn, v.conn)
	require.Equal(t, sk, v.key)
	require.NotNil(t, v.clientFunc)
}

func TestGRPCCreator_Create(t *testing.T) {
	ctx := context.TODO()
	s := new(gRPCCreator)

	// nil CreateParamsSource
	_, err := s.Create(ctx, nil)
	require.EqualError(t, err, ErrNilCreateParamsSource.Error())

	var (
		ownerID = OwnerID{1, 2, 3}
		created = uint64(2)
		expired = uint64(4)
	)

	p := NewParams()
	p.SetOwnerID(ownerID)
	p.SetCreationEpoch(created)
	p.SetExpirationEpoch(expired)

	// nil private key
	_, err = s.Create(ctx, p)
	require.Error(t, err)

	// create test private key
	s.key = test.DecodeKey(0)

	// create test client
	c := &testSessionClient{
		fn: func(req *CreateRequest) {
			require.Equal(t, ownerID, req.GetOwnerID())
			require.Equal(t, created, req.CreationEpoch())
			require.Equal(t, expired, req.ExpirationEpoch())
			require.NoError(t, service.VerifyAccumulatedSignaturesWithToken(req))
		},
		resp: &CreateResponse{
			ID:         TokenID{1, 2, 3},
			SessionKey: []byte{1, 2, 3},
		},
		err: errors.New("test error"),
	}

	s.clientFunc = func(*grpc.ClientConn) SessionClient {
		return c
	}

	res, err := s.Create(ctx, p)
	require.EqualError(t, err, c.err.Error())
	require.Equal(t, c.resp, res)
}
