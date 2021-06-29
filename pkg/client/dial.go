package client

import (
	"crypto/tls"
	"time"

	protoclient "github.com/nspcc-dev/neofs-api-go/rpc/client"
)

// DialPrm groups the Dial parameters.
type DialPrm struct {
	p protoclient.DialPrm
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
	var r protoclient.DialRes

	err := x.p.Dial(&r)
	if err != nil {
		return err
	}

	res.c.SetProtoClient(r.Client())

	return nil
}

// SetTCPAddress sets network address as a connection target.
func (x *DialPrm) SetAddress(addr string) {
	x.p.SetAddress(addr)
}

// SetTimeout sets timeout for dialing.
func (x *DialPrm) SetTimeout(timeout time.Duration) {
	x.p.SetTimeout(timeout)
}

// SetTLSConfig sets TLS client configuration.
func (x *DialPrm) SetTLSConfig(c *tls.Config) {
	x.p.SetTLSConfig(c)
}

// SetURIAddress tries to parse URI from string and
// pass host to prm.SetAddress. Also updates TLS config.
// Returns true if prm is modified.
//
// If address is not a valid URI, returns false.
//
// Format of the URI:
//
// 		[scheme://]host:port
//
// Supported schemes:
//  - grpc;
//  - grpcs.
//
// If scheme is not supported, returns false.
// If URI has grpcs scheme, empty TLS config is used.
//
// Should not be used with SetAddress and SetTLSConfig if result is true.
func (x *DialPrm) SetURIAddress(address string) bool {
	return protoclient.SetURIAddress(&x.p, address)
}

// Client returns initialized Client instance.
//
// Should be called only after successful Dial.
func (x DialRes) Client() Client {
	return x.c
}
