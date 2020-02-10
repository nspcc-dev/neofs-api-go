package state

import (
	"encoding/json"
	"expvar"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncodeVariables(t *testing.T) {
	dump := make(map[string]interface{})

	expvar.NewString("test1").Set("test1")
	expvar.NewString("test2").Set("test2")

	res := EncodeVariables()

	require.NoError(t, json.Unmarshal(res.Variables, &dump))
	require.NotEmpty(t, dump)

	// dump should contains keys `test1` and `test2`
	require.Contains(t, dump, "test1")
	require.Equal(t, "test1", dump["test1"])

	require.Contains(t, dump, "test2")
	require.Equal(t, "test2", dump["test2"])
}
