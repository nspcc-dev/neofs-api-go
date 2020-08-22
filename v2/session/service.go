package session

import (
	"context"
)

type Service interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
}

type CreateRequest struct {
	body *CreateRequestBody

	metaHeader *RequestMetaHeader

	verifyHeader *RequestVerificationHeader
}

type CreateResponse struct {
	body *CreateResponseBody

	metaHeader *ResponseMetaHeader

	verifyHeader *ResponseVerificationHeader
}
