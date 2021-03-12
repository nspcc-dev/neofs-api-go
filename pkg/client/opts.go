package client

import (
	"crypto/ecdsa"
	"time"

	"github.com/nspcc-dev/neofs-api-go/pkg"
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
		session  *token.SessionToken
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

// WithKey sets client's key for the next request.
func WithKey(key *ecdsa.PrivateKey) CallOption {
	return func(opts *callOptions) {
		opts.key = key
	}
}

func WithEpoch(epoch uint64) CallOption {
	return func(opts *callOptions) {
		opts.epoch = epoch
	}
}

func WithSession(token *token.SessionToken) CallOption {
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
		rawOpts: make([]client.Option, 0, 3),
	}
}

func WithAddress(addr string) Option {
	return func(opts *clientOptions) {
		opts.rawOpts = append(opts.rawOpts, client.WithNetworkAddress(addr))
	}
}

func WithGRPCConnection(grpcConn *grpc.ClientConn) Option {
	return func(opts *clientOptions) {
		opts.rawOpts = append(opts.rawOpts, client.WithGRPCConn(grpcConn))
	}
}

// WithDialTimeout returns option to set connection timeout to the remote node.
func WithDialTimeout(dur time.Duration) Option {
	return func(opts *clientOptions) {
		opts.rawOpts = append(opts.rawOpts, client.WithDialTimeout(dur))
	}
}

// WithDefaultPrivateKey returns option to set default private key
// used for the work.
func WithDefaultPrivateKey(key *ecdsa.PrivateKey) Option {
	return func(opts *clientOptions) {
		opts.key = key
	}
}
