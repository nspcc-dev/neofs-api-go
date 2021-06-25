package client

import (
	"crypto/ecdsa"
	"crypto/tls"
	"time"

	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/session"
	"github.com/nspcc-dev/neofs-api-go/pkg/token"
	"github.com/nspcc-dev/neofs-api-go/rpc/client"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	v2session "github.com/nspcc-dev/neofs-api-go/v2/session"
	"google.golang.org/grpc"
)

type (
	CallOption func(*callOptions)

	Option func(*clientOptions)

	callOptions struct {
		version  *pkg.Version
		xHeaders []*pkg.XHeader
		ttl      uint32
		epoch    uint64
		key      *ecdsa.PrivateKey
		session  *session.Token
		bearer   *token.BearerToken
	}

	clientOptions struct {
		key *ecdsa.PrivateKey

		rawOpts []client.Option
	}

	v2SessionReqInfo struct {
		addr *refs.Address
		verb v2session.ObjectSessionVerb

		exp, nbf, iat uint64
	}
)

func (c *clientImpl) defaultCallOptions() *callOptions {
	return &callOptions{
		version: pkg.SDKVersion(),
		ttl:     2,
		key:     c.opts.key,
	}
}

func WithXHeader(x *pkg.XHeader) CallOption {
	return func(opts *callOptions) {
		opts.xHeaders = append(opts.xHeaders, x)
	}
}

func WithTTL(ttl uint32) CallOption {
	return func(opts *callOptions) {
		opts.ttl = ttl
	}
}

// WithECDSAKey sets client's ECDSA key for the next request.
func WithECDSAKey(key *ecdsa.PrivateKey) CallOption {
	return func(opts *callOptions) {
		opts.key = key
	}
}

func WithEpoch(epoch uint64) CallOption {
	return func(opts *callOptions) {
		opts.epoch = epoch
	}
}

func WithSession(token *session.Token) CallOption {
	return func(opts *callOptions) {
		opts.session = token
	}
}

func WithBearer(token *token.BearerToken) CallOption {
	return func(opts *callOptions) {
		opts.bearer = token
	}
}

func v2MetaHeaderFromOpts(options *callOptions) *v2session.RequestMetaHeader {
	meta := new(v2session.RequestMetaHeader)
	meta.SetVersion(options.version.ToV2())
	meta.SetTTL(options.ttl)
	meta.SetEpoch(options.epoch)

	xhdrs := make([]*v2session.XHeader, len(options.xHeaders))
	for i := range options.xHeaders {
		xhdrs[i] = options.xHeaders[i].ToV2()
	}

	meta.SetXHeaders(xhdrs)

	if options.bearer != nil {
		meta.SetBearerToken(options.bearer.ToV2())
	}

	meta.SetSessionToken(options.session.ToV2())

	return meta
}

func defaultClientOptions() *clientOptions {
	return &clientOptions{
		rawOpts: make([]client.Option, 0, 4),
	}
}

// WithAddress returns option to specify
// network address of the remote server.
//
// Ignored if WithGRPCConnection is provided.
func WithAddress(addr string) Option {
	return func(opts *clientOptions) {
		opts.rawOpts = append(opts.rawOpts, client.WithNetworkAddress(addr))
	}
}

// WithDialTimeout returns option to set connection timeout to the remote node.
//
// Ignored if WithGRPCConn is provided.
func WithDialTimeout(dur time.Duration) Option {
	return func(opts *clientOptions) {
		opts.rawOpts = append(opts.rawOpts, client.WithDialTimeout(dur))
	}
}

// WithTLSConfig returns option to set connection's TLS config to the remote node.
//
// Ignored if WithGRPCConnection is provided.
func WithTLSConfig(cfg *tls.Config) Option {
	return func(opts *clientOptions) {
		opts.rawOpts = append(opts.rawOpts, client.WithTLSCfg(cfg))
	}
}

// WithDefaultECDSAPrivateKey returns option to set default ECDSA private key
// used for the work.
func WithDefaultECDSAPrivateKey(key *ecdsa.PrivateKey) Option {
	return func(opts *clientOptions) {
		opts.key = key
	}
}

// WithURIAddress returns option to specify
// network address of a remote server and connection
// scheme for it.
//
// Format of the URI:
//
//		[scheme://]host:port
//
// Supported schemes:
//  - grpc;
//  - grpcs.
//
// tls.Cfg second argument is optional and is taken into
// account only in case of `grpcs` scheme.
//
// Falls back to WithNetworkAddress if address is not a valid URI.
//
// Do not use along with WithAddress and WithTLSConfig.
//
// Ignored if WithGRPCConnection is provided.
func WithURIAddress(addr string, tlsCfg *tls.Config) Option {
	return func(opts *clientOptions) {
		opts.rawOpts = append(opts.rawOpts, client.WithNetworkURIAddress(addr, tlsCfg)...)
	}
}

// WithGRPCConnection returns option to set GRPC connection to
// the remote node.
func WithGRPCConnection(grpcConn *grpc.ClientConn) Option {
	return func(opts *clientOptions) {
		opts.rawOpts = append(opts.rawOpts, client.WithGRPCConn(grpcConn))
	}
}
