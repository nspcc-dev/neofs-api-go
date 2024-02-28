package client

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/url"

	"github.com/nspcc-dev/neofs-api-go/v2/rpc/grpc"
	grpcstd "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func (c *Client) createGRPCClient(ctx context.Context) (err error) {
	c.gRPCClientOnce.Do(func() {
		if err = c.openGRPCConn(ctx); err != nil {
			err = fmt.Errorf("open gRPC connection: %w", err)
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

func (c *Client) openGRPCConn(ctx context.Context, extraDialOpts ...grpcstd.DialOption) error {
	if c.conn != nil {
		return nil
	}

	if c.addr == "" {
		return errInvalidEndpoint
	}

	var creds credentials.TransportCredentials

	if c.tlsCfg != nil {
		creds = credentials.NewTLS(c.tlsCfg)
	} else {
		creds = insecure.NewCredentials()
	}

	dialCtx, cancel := context.WithTimeout(ctx, c.dialTimeout)
	var err error

	c.conn, err = grpcstd.DialContext(dialCtx, c.addr, append([]grpcstd.DialOption{
		grpcstd.WithTransportCredentials(creds),
		grpcstd.WithReturnConnectionError(),
		grpcstd.FailOnNonTempDialError(true),
	}, extraDialOpts...)...)

	cancel()

	if err != nil {
		return fmt.Errorf("gRPC dial: %w", err)
	}

	return nil
}

// ParseURI parses s as address and returns a host and a flag
// indicating that TLS is enabled. If multi-address is provided
// the argument is returned unchanged.
func ParseURI(s string) (string, bool, error) {
	uri, err := url.ParseRequestURI(s)
	if err != nil {
		return s, false, nil
	}

	// check if passed string was parsed correctly
	// URIs that do not start with a slash after the scheme are interpreted as:
	// `scheme:opaque` => if `opaque` is not empty, then it is supposed that URI
	// is in `host:port` format
	if uri.Host == "" {
		uri.Host = uri.Scheme
		uri.Scheme = grpcScheme // assume GRPC by default
		if uri.Opaque != "" {
			uri.Host = net.JoinHostPort(uri.Host, uri.Opaque)
		}
	}

	switch uri.Scheme {
	case grpcTLSScheme, grpcScheme:
	default:
		return "", false, fmt.Errorf("unsupported scheme: %s", uri.Scheme)
	}

	return uri.Host, uri.Scheme == grpcTLSScheme, nil
}
