package client

import (
	"context"
	"crypto/tls"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestClient_Init(t *testing.T) {
	t.Run("TLS handshake failure", func(t *testing.T) {
		lis := bufconn.Listen(1024) // size does not matter in this test

		srv := grpc.NewServer()
		t.Cleanup(srv.Stop)
		go func() { _ = srv.Serve(lis) }()

		c := New(WithNetworkURIAddress("grpcs://any:54321", new(tls.Config))...)

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		t.Cleanup(cancel)

		err := c.openGRPCConn(ctx, grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return lis.DialContext(ctx)
		}))
		// error is not wrapped properly, so we can do nothing more to check it.
		// Text from stdlib tls.Conn.HandshakeContext.
		require.ErrorContains(t, err, "first record does not look like a TLS handshake")
	})
}
