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

func TestTargetInfo_StableMarshal(t *testing.T) {
	targetFrom := new(acl.TargetInfo)
	transport := new(grpc.EACLRecord_TargetInfo)

	t.Run("non empty", func(t *testing.T) {
		targetFrom.SetTarget(acl.TargetUser)
		targetFrom.SetKeyList([][]byte{
			[]byte("Public Key 1"),
			[]byte("Public Key 2"),
		})

		wire, err := targetFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		targetTo := acl.TargetInfoFromGRPCMessage(transport)
		require.Equal(t, targetFrom, targetTo)
	})
}
