package protoclient

import (
	"crypto/tls"
	"net/url"
	"time"

	neofsgrpc "github.com/nspcc-dev/neofs-api-go/rpc/grpc"
)

// DialPrm groups the Dial parameters.
type DialPrm struct {
	g neofsgrpc.DialPrm
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
	var r neofsgrpc.DialRes

	err := x.g.Dial(&r)
	if err != nil {
		return err
	}

	res.c.g = r.Client()

	return nil
}

// SetAddress sets network address as a connection target.
func (x *DialPrm) SetAddress(addr string) {
	x.g.SetAddress(addr)
}

// SetTimeout sets timeout for dialing.
func (x *DialPrm) SetTimeout(timeout time.Duration) {
	x.g.SetTimeout(timeout)
}

// SetTLSConfig sets TLS client configuration.
func (x *DialPrm) SetTLSConfig(c *tls.Config) {
	x.g.SetTLSConfig(c)
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
func SetURIAddress(prm *DialPrm, address string) bool {
	uri, err := url.ParseRequestURI(address)

	// check if passed string was parsed correctly
	// URIs that do not start with a slash after the scheme are interpreted as:
	// `scheme:opaque` => if `opaque` is not empty, then it is supposed that URI
	// is in `host:port` format
	if err == nil && uri.Opaque == "" {
		switch uri.Scheme {
		case "grpcs":
			prm.SetTLSConfig(new(tls.Config))
			fallthrough
		case "grpc":
			prm.SetAddress(uri.Host)
			return true
		}
	}

	return false
}

// Client returns initialized Client instance.
//
// Should be called only after successful Dial.
func (x DialRes) Client() Client {
	return x.c
}
