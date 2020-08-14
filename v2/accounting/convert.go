package accounting

import (
	accounting "github.com/nspcc-dev/neofs-api-go/v2/accounting/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/service"
	grpcService "github.com/nspcc-dev/neofs-api-go/v2/service/grpc"
)

func BalanceRequestBodyToGRPCMessage(b *BalanceRequestBody) *accounting.BalanceRequest_Body {
	if b == nil {
		return nil
	}

	m := new(accounting.BalanceRequest_Body)

	m.SetOwnerId(
		refs.OwnerIDToGRPCMessage(b.GetOwnerID()),
	)

	return m
}

func BalanceRequestBodyFromGRPCMessage(m *accounting.BalanceRequest_Body) *BalanceRequestBody {
	if m == nil {
		return nil
	}

	b := new(BalanceRequestBody)

	b.SetOwnerID(
		refs.OwnerIDFromGRPCMessage(m.GetOwnerId()),
	)

	return b
}

func headersToGRPC(
	src interface {
		GetRequestMetaHeader() *service.RequestMetaHeader
		GetRequestVerificationHeader() *service.RequestVerificationHeader
	},
	dst interface {
		SetMetaHeader(*grpcService.RequestMetaHeader)
		SetVerifyHeader(*grpcService.RequestVerificationHeader)
	},
) {
	dst.SetMetaHeader(
		service.RequestMetaHeaderToGRPCMessage(src.GetRequestMetaHeader()),
	)

	dst.SetVerifyHeader(
		service.RequestVerificationHeaderToGRPCMessage(src.GetRequestVerificationHeader()),
	)
}

func headersFromGRPC(
	src interface {
		GetMetaHeader() *grpcService.RequestMetaHeader
		GetVerifyHeader() *grpcService.RequestVerificationHeader
	},
	dst interface {
		SetRequestMetaHeader(*service.RequestMetaHeader)
		SetRequestVerificationHeader(*service.RequestVerificationHeader)
	},
) {
	dst.SetRequestMetaHeader(
		service.RequestMetaHeaderFromGRPCMessage(src.GetMetaHeader()),
	)

	dst.SetRequestVerificationHeader(
		service.RequestVerificationHeaderFromGRPCMessage(src.GetVerifyHeader()),
	)
}

func BalanceRequestToGRPCMessage(b *BalanceRequest) *accounting.BalanceRequest {
	if b == nil {
		return nil
	}

	m := new(accounting.BalanceRequest)

	m.SetBody(
		BalanceRequestBodyToGRPCMessage(b.GetBody()),
	)

	headersToGRPC(b, m)

	return m
}

func BalanceRequestFromGRPCMessage(m *accounting.BalanceRequest) *BalanceRequest {
	if m == nil {
		return nil
	}

	b := new(BalanceRequest)

	b.SetBody(
		BalanceRequestBodyFromGRPCMessage(m.GetBody()),
	)

	headersFromGRPC(m, b)

	return b
}
