package protoclient_test

import (
	"testing"

	protoclient "github.com/nspcc-dev/neofs-api-go/rpc/client"
	"github.com/stretchr/testify/require"
)

func TestWithNetworkURIAddress(t *testing.T) {
	hostPort := "neofs.example.com:8080"

	testCases := []struct {
		uri string
		ok  bool
	}{
		{
			uri: "grpcs://" + hostPort,
			ok:  true,
		},
		{
			uri: "grpc://" + hostPort,
			ok:  true,
		},
		{
			uri: "unsupportedScheme://" + hostPort,
		},
		{
			uri: "invalidURI",
		},
	}

	for _, test := range testCases {
		var prm protoclient.DialPrm

		ok := protoclient.SetURIAddress(&prm, test.uri)

		require.Equal(t, test.ok, ok)
	}
}
