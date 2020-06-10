package session

import (
	"context"
	"crypto/ecdsa"

	"github.com/nspcc-dev/neofs-api-go/service"
	crypto "github.com/nspcc-dev/neofs-crypto"
	"google.golang.org/grpc"
)

type gRPCCreator struct {
	conn *grpc.ClientConn

	key *ecdsa.PrivateKey

	clientFunc func(*grpc.ClientConn) SessionClient
}

// NewGRPCCreator unites virtual gRPC client with private ket and returns Creator interface.
//
// If passed ClientConn is nil, ErrNilGPRCClientConn returns.
// If passed private key is nil, crypto.ErrEmptyPrivateKey returns.
func NewGRPCCreator(conn *grpc.ClientConn, key *ecdsa.PrivateKey) (Creator, error) {
	if conn == nil {
		return nil, ErrNilGPRCClientConn
	} else if key == nil {
		return nil, crypto.ErrEmptyPrivateKey
	}

	return &gRPCCreator{
		conn: conn,

		key: key,

		clientFunc: NewSessionClient,
	}, nil
}

// Create constructs message, signs it with private key and sends it to a gRPC client.
//
// If passed CreateParamsSource is nil, ErrNilCreateParamsSource returns.
// If message could not be signed, an error returns.
func (s gRPCCreator) Create(ctx context.Context, p CreateParamsSource) (CreateResult, error) {
	if p == nil {
		return nil, ErrNilCreateParamsSource
	}

	// create and fill a message
	req := new(CreateRequest)
	req.SetOwnerID(p.GetOwnerID())
	req.SetCreationEpoch(p.CreationEpoch())
	req.SetExpirationEpoch(p.ExpirationEpoch())

	// sign with private key
	if err := service.SignRequestData(s.key, req); err != nil {
		return nil, err
	}

	// make gRPC call
	return s.clientFunc(s.conn).Create(ctx, req)
}
