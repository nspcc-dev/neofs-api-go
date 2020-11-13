package accounting_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/accounting"
	"github.com/stretchr/testify/require"
)

func TestDecimalJSON(t *testing.T) {
	i := generateDecimal(10)

	data, err := i.MarshalJSON()
	require.NoError(t, err)

	i2 := new(accounting.Decimal)
	require.NoError(t, i2.UnmarshalJSON(data))

	require.Equal(t, i, i2)
}
