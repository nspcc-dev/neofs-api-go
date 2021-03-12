package client

import (
	"crypto/ecdsa"
	"errors"

	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/token"
)

type (
	// Client represents NeoFS client.
	Client interface {
		Accounting
		Container
		Netmap
		Object
		Session
	}

	clientImpl struct {
		key        *ecdsa.PrivateKey
		remoteNode TransportInfo

		opts *clientOptions

		sessionToken *token.SessionToken

		bearerToken *token.BearerToken
	}

	TransportProtocol uint32

	TransportInfo struct {
		Version  *pkg.Version
		Protocol TransportProtocol
	}
)

const (
	Unknown TransportProtocol = iota
	GRPC
)

var errUnsupportedProtocol = errors.New("unsupported transport protocol")

// New returns new client which uses key as default signing key.
func New(key *ecdsa.PrivateKey, opts ...Option) (Client, error) {
	clientOptions := defaultClientOptions()

	for i := range opts {
		opts[i].apply(clientOptions)
	}

	// todo: make handshake to check latest version
	return &clientImpl{
		key: key,
		remoteNode: TransportInfo{
			Version:  pkg.SDKVersion(),
			Protocol: GRPC,
		},
		opts: clientOptions,
	}, nil
}
