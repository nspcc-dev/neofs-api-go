package neofsgrpc

import (
	"context"
	"crypto/tls"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// DialPrm groups the Dial parameters.
type DialPrm struct {
	addr string

	timeout time.Duration

	withTLS bool
	tls     *tls.Config
}

// DialPrm groups the Dial results.
type DialRes struct {
	c Client
}

// Dial opens client connection and initializes Client with it.
//
// Blocks until successful connection or timeout.
// If TLS is configured, connection is opened using it.
//
// DialRes must not be nil. Client can be received using res.Client().
func (x DialPrm) Dial(res *DialRes) error {
	ctx, cancel := context.WithTimeout(context.Background(), x.timeout)
	defer cancel()

	var credOpt grpc.DialOption

	if x.withTLS {
		credOpt = grpc.WithTransportCredentials(credentials.NewTLS(x.tls))
	} else {
		credOpt = grpc.WithInsecure()
	}

	conn, err := grpc.DialContext(ctx, x.addr, credOpt, grpc.WithBlock())
	if err != nil {
		return err
	}

	res.c.conn = conn

	return nil
}

// SetAddress sets network address as a connection target.
func (x *DialPrm) SetAddress(addr string) {
	x.addr = addr
}

// SetTimeout sets timeout for dialing.
func (x *DialPrm) SetTimeout(timeout time.Duration) {
	x.timeout = timeout
}

// SetTLSConfig sets TLS client configuration.
func (x *DialPrm) SetTLSConfig(c *tls.Config) {
	x.tls = c
	x.withTLS = true
}

// Client returns initialized Client instance.
//
// Should be called only after successful Dial.
func (x DialRes) Client() Client {
	return x.c
}
