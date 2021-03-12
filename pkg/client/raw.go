package client

import (
	"github.com/nspcc-dev/neofs-api-go/rpc/client"
)

// Raw returns underlying raw protobuf client.
func (c *clientImpl) Raw() *client.Client {
	c.onceInit.Do(func() {
		c.raw = client.New(c.opts.rawOpts...)
	})

	return c.raw
}
