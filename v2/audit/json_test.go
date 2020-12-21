package audit_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/audit"
	"github.com/stretchr/testify/require"
)

func TestDataAuditResultJSON(t *testing.T) {
	a := generateDataAuditResult()

	data, err := a.MarshalJSON()
	require.NoError(t, err)

	a2 := new(audit.DataAuditResult)
	require.NoError(t, a2.UnmarshalJSON(data))

	require.Equal(t, a, a2)
}
