package common_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/rpc/common"
	"github.com/stretchr/testify/require"
)

const (
	testServiceName = "test service"
	testRPCName     = "test RPC"
)

func TestCallMethodInfo_ServiceName(t *testing.T) {
	var i common.CallMethodInfo

	i.SetServiceName(testServiceName)

	require.Equal(t, testServiceName, i.ServiceName())
}

func TestCallMethodInfo_MethodName(t *testing.T) {
	var i common.CallMethodInfo

	i.SetMethodName(testRPCName)

	require.Equal(t, testRPCName, i.MethodName())
}

func TestCallMethodInfo_ServerStream(t *testing.T) {
	var i common.CallMethodInfo

	require.False(t, i.ClientStream())
	require.False(t, i.ServerStream())

	i.SetServerStream()

	require.False(t, i.ClientStream())
	require.True(t, i.ServerStream())
}

func TestCallMethodInfo_ClientStream(t *testing.T) {
	var i common.CallMethodInfo

	require.False(t, i.ClientStream())
	require.False(t, i.ServerStream())

	i.SetClientStream()

	require.True(t, i.ClientStream())
	require.False(t, i.ServerStream())
}
