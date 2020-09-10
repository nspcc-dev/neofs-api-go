package client

import (
	"crypto/ecdsa"
	"errors"

	"github.com/nspcc-dev/neofs-api-go/pkg"
)

type (
	Client struct {
		key        *ecdsa.PrivateKey
		remoteNode TransportInfo

		opts *clientOptions
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

var (
	unsupportedProtocolErr = errors.New("unsupported transport protocol")
)

func New(key *ecdsa.PrivateKey, opts ...ClientOption) (*Client, error) {
	clientOptions := defaultClientOptions()
	for i := range opts {
		opts[i].apply(clientOptions)
	}

	// todo: make handshake to check latest version
	return &Client{
		key: key,
		remoteNode: TransportInfo{
			Version:  pkg.SDKVersion(),
			Protocol: GRPC,
		},
		opts: clientOptions,
	}, nil
}
