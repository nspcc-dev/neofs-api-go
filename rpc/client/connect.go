package client

import (
	"context"
	"errors"

	"github.com/nspcc-dev/neofs-api-go/rpc/grpc"
	grpcstd "google.golang.org/grpc"
)

func (c *Client) createGRPCClient() (err error) {
	c.gRPCClientOnce.Do(func() {
		if err = c.openGRPCConn(); err != nil {
			return
		}

		c.gRPCClient = grpc.New(grpc.WithClientConnection(c.conn))
	})

	return
}

var errInvalidEndpoint = errors.New("invalid endpoint options")

func (c *Client) openGRPCConn() error {
	if c.conn != nil {
		return nil
	}

	if c.addr == "" {
		return errInvalidEndpoint
	}

	var err error

	dialCtx, cancel := context.WithTimeout(context.Background(), c.dialTimeout)
	c.conn, err = grpcstd.DialContext(dialCtx, c.addr, grpcstd.WithInsecure())
	cancel()

	return err
}
