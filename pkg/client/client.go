package client

import (
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
		raw:  client.New(),
		opts: clientOptions,
	}, nil
}
