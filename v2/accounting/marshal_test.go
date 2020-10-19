package accounting_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/accounting"
	grpc "github.com/nspcc-dev/neofs-api-go/v2/accounting/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/stretchr/testify/require"
	goproto "google.golang.org/protobuf/proto"
)

func TestDecimal_StableMarshal(t *testing.T) {
	decimalFrom := generateDecimal(888)
	transport := new(grpc.Decimal)

	t.Run("non empty", func(t *testing.T) {
		wire, err := decimalFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = goproto.Unmarshal(wire, transport)
		require.NoError(t, err)

		decimalTo := accounting.DecimalFromGRPCMessage(transport)
		require.Equal(t, decimalFrom, decimalTo)
	})
}

func TestBalanceRequestBody_StableMarshal(t *testing.T) {
	requestBodyFrom := generateBalanceRequestBody("Owner ID")
	transport := new(grpc.BalanceRequest_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := requestBodyFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = goproto.Unmarshal(wire, transport)
		require.NoError(t, err)

		requestBodyTo := accounting.BalanceRequestBodyFromGRPCMessage(transport)
		require.Equal(t, requestBodyFrom, requestBodyTo)
	})
}

func TestBalanceResponseBody_StableMarshal(t *testing.T) {
	responseBodyFrom := generateBalanceResponseBody(444)
	transport := new(grpc.BalanceResponse_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := responseBodyFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = goproto.Unmarshal(wire, transport)
		require.NoError(t, err)

		responseBodyTo := accounting.BalanceResponseBodyFromGRPCMessage(transport)
		require.Equal(t, responseBodyFrom, responseBodyTo)
	})
}

func generateDecimal(val int64) *accounting.Decimal {
	decimal := new(accounting.Decimal)
	decimal.SetValue(val)
	decimal.SetPrecision(1000)

	return decimal
}

func generateBalanceRequestBody(id string) *accounting.BalanceRequestBody {
	owner := new(refs.OwnerID)
	owner.SetValue([]byte(id))

	request := new(accounting.BalanceRequestBody)
	request.SetOwnerID(owner)

	return request
}

func generateBalanceResponseBody(val int64) *accounting.BalanceResponseBody {
	response := new(accounting.BalanceResponseBody)
	response.SetBalance(generateDecimal(val))

	return response
}
