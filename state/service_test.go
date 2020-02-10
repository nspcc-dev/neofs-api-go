package state

import (
	"encoding/json"
	"expvar"
	"testing"

	"github.com/spf13/viper"
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

func TestEncodeConfig(t *testing.T) {
	v := viper.New()
	v.Set("test1", "test1")
	v.Set("test2", "test2")

	res, err := EncodeConfig(v)
	require.NoError(t, err)

	dump := make(map[string]interface{})
	require.NoError(t, json.Unmarshal(res.Config, &dump))

	require.NotEmpty(t, dump)

	require.Contains(t, dump, "test1")
	require.Equal(t, dump["test1"], "test1")

	require.Contains(t, dump, "test2")
	require.Equal(t, dump["test2"], "test2")
}
