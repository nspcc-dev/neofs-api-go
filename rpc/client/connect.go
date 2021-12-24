package client

import (
	"context"
	"errors"

	"github.com/nspcc-dev/neofs-api-go/v2/rpc/grpc"
	grpcstd "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func (c *Client) createGRPCClient() (err error) {
	c.gRPCClientOnce.Do(func() {
		if err = c.openGRPCConn(); err != nil {
			return
		}

		c.gRPCClient = grpc.New(
			grpc.WithClientConnection(c.conn),
			grpc.WithRWTimeout(c.rwTimeout),
		)
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

	var credOpt grpcstd.DialOption

	if c.tlsCfg != nil {
		creds := credentials.NewTLS(c.tlsCfg)
		credOpt = grpcstd.WithTransportCredentials(creds)
	} else {
		credOpt = grpcstd.WithInsecure()
	}

	dialCtx, cancel := context.WithTimeout(context.Background(), c.dialTimeout)
	c.conn, err = grpcstd.DialContext(dialCtx, c.addr, credOpt)
	cancel()

	return err
}
