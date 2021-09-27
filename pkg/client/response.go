package client

import (
	"github.com/nspcc-dev/neofs-api-go/v2/session"
)

// ResponseMetaInfo groups meta information about any NeoFS API response.
type ResponseMetaInfo struct {
	key []byte
}

// ResponderKey returns responder's public key in a binary format.
//
// Result must not be mutated.
func (x ResponseMetaInfo) ResponderKey() []byte {
	return x.key
}

// WithResponseInfoHandler allows to specify handler of response meta information for the all Client operations.
// The handler is called right after the response is received. Client returns handler's error immediately.
func WithResponseInfoHandler(f func(ResponseMetaInfo) error) Option {
	return func(opts *clientOptions) {
		opts.cbRespInfo = f
	}
}

func (c *clientImpl) handleResponseInfoV2(opts *callOptions, resp interface {
	GetVerificationHeader() *session.ResponseVerificationHeader
}) error {
	if c.opts.cbRespInfo == nil {
		return nil
	}

	return c.opts.cbRespInfo(ResponseMetaInfo{
		key: resp.GetVerificationHeader().GetBodySignature().GetKey(),
	})
}
