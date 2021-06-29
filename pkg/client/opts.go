package client

import (
	"crypto/ecdsa"

	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/session"
	"github.com/nspcc-dev/neofs-api-go/pkg/token"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	v2session "github.com/nspcc-dev/neofs-api-go/v2/session"
)

type (
	CallOption func(*callOptions)

	callOptions struct {
		version  *pkg.Version
		xHeaders []*pkg.XHeader
		ttl      uint32
		epoch    uint64
		key      *ecdsa.PrivateKey
		session  *session.Token
		bearer   *token.BearerToken
	}

	v2SessionReqInfo struct {
		addr *refs.Address
		verb v2session.ObjectSessionVerb

		exp, nbf, iat uint64
	}
)

func defaultCallOptions() *callOptions {
	return &callOptions{
		version: pkg.SDKVersion(),
		ttl:     2,
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
