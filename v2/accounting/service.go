package accounting

import (
	"context"

	"github.com/nspcc-dev/neofs-api-go/v2/session"
)

type Service interface {
	Balance(context.Context, *BalanceRequest) (*BalanceResponse, error)
}

type BalanceRequest struct {
	body *BalanceRequestBody

	metaHeader *session.RequestMetaHeader

	verifyHeader *session.RequestVerificationHeader
}

type BalanceResponse struct {
	body *BalanceResponseBody

	metaHeader *session.ResponseMetaHeader

	verifyHeader *session.ResponseVerificationHeader
}
