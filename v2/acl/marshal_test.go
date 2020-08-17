package acl_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/acl"
	grpc "github.com/nspcc-dev/neofs-api-go/v2/acl/grpc"
	"github.com/stretchr/testify/require"
)

func TestHeaderFilter_StableMarshal(t *testing.T) {
	filterFrom := new(acl.HeaderFilter)
	transport := new(grpc.EACLRecord_FilterInfo)

	t.Run("non empty", func(t *testing.T) {
		filterFrom.SetHeaderType(acl.HeaderTypeObject)
		filterFrom.SetMatchType(acl.MatchTypeStringEqual)
		filterFrom.SetName("Hello")
		filterFrom.SetValue("World")

		wire, err := filterFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		filterTo := acl.HeaderFilterFromGRPCMessage(transport)
		require.Equal(t, filterFrom, filterTo)
	})
}
