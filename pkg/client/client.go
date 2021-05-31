package client

import (
	"io"
	"sync"

	"github.com/nspcc-dev/neofs-api-go/rpc/client"
)

// Client represents NeoFS client.
type Client interface {
	Accounting
	Container
	Netmap
	Object
	Session
	Reputation

	// Raw must return underlying raw protobuf client.
	Raw() *client.Client

	// Conn must return underlying connection.
	//
	// Must return a non-nil result after the first RPC call
	// completed without a connection error.
	Conn() io.Closer
}

type clientImpl struct {
	onceInit sync.Once

	raw *client.Client

	opts *clientOptions
}

func New(opts ...Option) (Client, error) {
	clientOptions := defaultClientOptions()

	for i := range opts {
		opts[i](clientOptions)
	}

	return &clientImpl{
		opts: clientOptions,
	}, nil
}
