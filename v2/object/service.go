package object

import (
	"context"

	"github.com/nspcc-dev/neofs-api-go/v2/session"
)

type Service interface {
	Get(context.Context, *GetRequest) (GetObjectStreamer, error)
	Put(context.Context) (PutObjectStreamer, error)
	Head(context.Context, *HeadRequest) (*HeadResponse, error)
	Search(context.Context, *SearchRequest) (SearchObjectStreamer, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	GetRange(context.Context, *GetRangeRequest) (GetRangeObjectStreamer, error)
	GetRangeHash(context.Context, *GetRangeHashRequest) (*GetRangeHashResponse, error)
}

type GetRequest struct {
	body *GetRequestBody

	metaHeader *session.RequestMetaHeader

	verifyHeader *session.RequestVerificationHeader
}

type GetResponse struct {
	body *GetResponseBody

	metaHeader *session.ResponseMetaHeader

	verifyHeader *session.ResponseVerificationHeader
}

type PutRequest struct {
	body *PutRequestBody

	metaHeader *session.RequestMetaHeader

	verifyHeader *session.RequestVerificationHeader
}

type PutResponse struct {
	body *PutResponseBody

	metaHeader *session.ResponseMetaHeader

	verifyHeader *session.ResponseVerificationHeader
}

type DeleteRequest struct {
	body *DeleteRequestBody

	metaHeader *session.RequestMetaHeader

	verifyHeader *session.RequestVerificationHeader
}

type DeleteResponse struct {
	body *DeleteResponseBody

	metaHeader *session.ResponseMetaHeader

	verifyHeader *session.ResponseVerificationHeader
}

type HeadRequest struct {
	body *HeadRequestBody

	metaHeader *session.RequestMetaHeader

	verifyHeader *session.RequestVerificationHeader
}

type HeadResponse struct {
	body *HeadResponseBody

	metaHeader *session.ResponseMetaHeader

	verifyHeader *session.ResponseVerificationHeader
}

type SearchRequest struct {
	body *SearchRequestBody

	metaHeader *session.RequestMetaHeader

	verifyHeader *session.RequestVerificationHeader
}

type SearchResponse struct {
	body *SearchResponseBody

	metaHeader *session.ResponseMetaHeader

	verifyHeader *session.ResponseVerificationHeader
}

type GetRangeRequest struct {
	body *GetRangeRequestBody

	metaHeader *session.RequestMetaHeader

	verifyHeader *session.RequestVerificationHeader
}

type GetRangeResponse struct {
	body *GetRangeResponseBody

	metaHeader *session.ResponseMetaHeader

	verifyHeader *session.ResponseVerificationHeader
}

type GetRangeHashRequest struct {
	body *GetRangeHashRequestBody

	metaHeader *session.RequestMetaHeader

	verifyHeader *session.RequestVerificationHeader
}

type GetRangeHashResponse struct {
	body *GetRangeHashResponseBody

	metaHeader *session.ResponseMetaHeader

	verifyHeader *session.ResponseVerificationHeader
}
