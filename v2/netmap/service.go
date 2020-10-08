package netmap

import (
	"context"

	"github.com/nspcc-dev/neofs-api-go/v2/session"
)

type Service interface {
	LocalNodeInfo(ctx context.Context, request *LocalNodeInfoRequest) (*LocalNodeInfoResponse, error)
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
