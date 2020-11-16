package client

import (
	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/token"
	v2accounting "github.com/nspcc-dev/neofs-api-go/v2/accounting"
	v2container "github.com/nspcc-dev/neofs-api-go/v2/container"
	v2netmap "github.com/nspcc-dev/neofs-api-go/v2/netmap"
	v2object "github.com/nspcc-dev/neofs-api-go/v2/object"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	v2session "github.com/nspcc-dev/neofs-api-go/v2/session"
	"google.golang.org/grpc"
)

type (
	CallOption interface {
		apply(*callOptions)
	}

	ClientOption interface {
		apply(*clientOptions)
	}

	callOptions struct {
		version  *pkg.Version
		xHeaders []*pkg.XHeader
		ttl      uint32
		epoch    uint64
		session  *token.SessionToken
		bearer   *token.BearerToken
	}

	clientOptions struct {
		addr string

		grpcOpts *grpcOptions
	}

	grpcOptions struct {
		conn               *grpc.ClientConn
		v2ContainerClient  *v2container.Client
		v2AccountingClient *v2accounting.Client
		v2SessionClient    *v2session.Client
		v2NetmapClient     *v2netmap.Client

		objectClientV2 *v2object.Client
	}

	v2SessionReqInfo struct {
		addr *refs.Address
		verb v2session.ObjectSessionVerb

		exp, nbf, iat uint64
	}
)

func (c Client) defaultCallOptions() callOptions {
	return callOptions{
		ttl:     2,
		version: pkg.SDKVersion(),
		session: c.sessionToken,
		bearer:  c.bearerToken,
	}
}

type funcCallOption struct {
	f func(*callOptions)
}

func (fco *funcCallOption) apply(co *callOptions) {
	fco.f(co)
}

func newFuncCallOption(f func(option *callOptions)) *funcCallOption {
	return &funcCallOption{
		f: f,
	}
}

func WithXHeader(x *pkg.XHeader) CallOption {
	return newFuncCallOption(func(option *callOptions) {
		option.xHeaders = append(option.xHeaders, x)
	})
}

func WithTTL(ttl uint32) CallOption {
	return newFuncCallOption(func(option *callOptions) {
		option.ttl = ttl
	})
}

func WithEpoch(epoch uint64) CallOption {
	return newFuncCallOption(func(option *callOptions) {
		option.epoch = epoch
	})
}

func WithSession(token *token.SessionToken) CallOption {
	return newFuncCallOption(func(option *callOptions) {
		option.session = token
	})
}

func WithBearer(token *token.BearerToken) CallOption {
	return newFuncCallOption(func(option *callOptions) {
		option.bearer = token
	})
}

func v2MetaHeaderFromOpts(options callOptions) *v2session.RequestMetaHeader {
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
		grpcOpts: new(grpcOptions),
	}
}

type funcClientOption struct {
	f func(*clientOptions)
}

func (fco *funcClientOption) apply(co *clientOptions) {
	fco.f(co)
}

func newFuncClientOption(f func(option *clientOptions)) *funcClientOption {
	return &funcClientOption{
		f: f,
	}
}

func WithAddress(addr string) ClientOption {
	return newFuncClientOption(func(option *clientOptions) {
		option.addr = addr
	})
}

func WithGRPCConnection(grpcConn *grpc.ClientConn) ClientOption {
	return newFuncClientOption(func(option *clientOptions) {
		option.grpcOpts.conn = grpcConn
	})
}
