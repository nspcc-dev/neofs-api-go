package container

import (
	"context"

	"github.com/nspcc-dev/neofs-api-go/v2/session"
)

type Service interface {
	Put(context.Context, *PutRequest) (*PutResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	SetExtendedACL(context.Context, *SetExtendedACLRequest) (*SetExtendedACLResponse, error)
	GetExtendedACL(context.Context, *GetExtendedACLRequest) (*GetExtendedACLResponse, error)
	AnnounceUsedSpace(context.Context, *AnnounceUsedSpaceRequest) (*AnnounceUsedSpaceResponse, error)
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

type ListRequest struct {
	body *ListRequestBody

	metaHeader *session.RequestMetaHeader

	verifyHeader *session.RequestVerificationHeader
}

type ListResponse struct {
	body *ListResponseBody

	metaHeader *session.ResponseMetaHeader

	verifyHeader *session.ResponseVerificationHeader
}

type SetExtendedACLRequest struct {
	body *SetExtendedACLRequestBody

	metaHeader *session.RequestMetaHeader

	verifyHeader *session.RequestVerificationHeader
}

type SetExtendedACLResponse struct {
	body *SetExtendedACLResponseBody

	metaHeader *session.ResponseMetaHeader

	verifyHeader *session.ResponseVerificationHeader
}

type GetExtendedACLRequest struct {
	body *GetExtendedACLRequestBody

	metaHeader *session.RequestMetaHeader

	verifyHeader *session.RequestVerificationHeader
}

type GetExtendedACLResponse struct {
	body *GetExtendedACLResponseBody

	metaHeader *session.ResponseMetaHeader

	verifyHeader *session.ResponseVerificationHeader
}

type AnnounceUsedSpaceRequest struct {
	body *AnnounceUsedSpaceRequestBody

	metaHeader *session.RequestMetaHeader

	verifyHeader *session.RequestVerificationHeader
}

type AnnounceUsedSpaceResponse struct {
	body *AnnounceUsedSpaceResponseBody

	metaHeader *session.ResponseMetaHeader

	verifyHeader *session.ResponseVerificationHeader
}
