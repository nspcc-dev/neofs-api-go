package netmap

import (
	"context"

	"github.com/nspcc-dev/neofs-api-go/v2/session"
)

type Service interface {
	LocalNodeInfo(ctx context.Context, request *LocalNodeInfoRequest) (*LocalNodeInfoResponse, error)
	NetworkInfo(ctx context.Context, request *NetworkInfoRequest) (*NetworkInfoResponse, error)
}

type LocalNodeInfoRequest struct {
	body *LocalNodeInfoRequestBody

	metaHeader *session.RequestMetaHeader

	verifyHeader *session.RequestVerificationHeader
}

type LocalNodeInfoResponse struct {
	body *LocalNodeInfoResponseBody

	metaHeader *session.ResponseMetaHeader

	verifyHeader *session.ResponseVerificationHeader
}

// NetworkInfoRequest is a structure of NetworkInfo request.
type NetworkInfoRequest struct {
	body *NetworkInfoRequestBody

	metaHeader *session.RequestMetaHeader

	verifyHeader *session.RequestVerificationHeader
}

// NetworkInfoResponse is a structure of NetworkInfo response.
type NetworkInfoResponse struct {
	body *NetworkInfoResponseBody

	metaHeader *session.ResponseMetaHeader

	verifyHeader *session.ResponseVerificationHeader
}
