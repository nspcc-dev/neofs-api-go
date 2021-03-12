package client

import (
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
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

	Option interface {
		apply(*clientOptions)
	}

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
		addr string

		dialTimeout time.Duration

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

type errOptionsLack string

func (e errOptionsLack) Error() string {
	return fmt.Sprintf("lack of sdk client options to create %s client", string(e))
}

func (c clientImpl) defaultCallOptions() callOptions {
	return callOptions{
		ttl:     2,
		version: pkg.SDKVersion(),
		key:     c.key,
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

// WithKey sets client's key for the next request.
func WithKey(key *ecdsa.PrivateKey) CallOption {
	return newFuncCallOption(func(option *callOptions) {
		option.key = key
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

func WithAddress(addr string) Option {
	return newFuncClientOption(func(option *clientOptions) {
		option.addr = addr
	})
}

func WithGRPCConnection(grpcConn *grpc.ClientConn) Option {
	return newFuncClientOption(func(option *clientOptions) {
		option.grpcOpts.conn = grpcConn
	})
}

// WithDialTimeout returns option to set connection timeout to the remote node.
func WithDialTimeout(dur time.Duration) Option {
	return newFuncClientOption(func(option *clientOptions) {
		option.dialTimeout = dur
	})
}

func newOwnerIDFromKey(key *ecdsa.PublicKey) (*owner.ID, error) {
	w, err := owner.NEO3WalletFromPublicKey(key)
	if err != nil {
		return nil, err
	}

	ownerID := new(owner.ID)
	ownerID.SetNeo3Wallet(w)
	return ownerID, nil
}
