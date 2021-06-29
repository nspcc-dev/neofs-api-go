package protoclient

import (
	neofsgrpc "github.com/nspcc-dev/neofs-api-go/rpc/grpc"
)

// Client represents client for exchanging messages
// with a remote server using Protobuf RPC.
//
// Should be created using DialPrm.Dial.
//
// Client is one-time use, and should only be used within an open connection.
type Client struct {
	g neofsgrpc.Client
}

// Close closes opened connection.
//
// Must be called only after successful initialization.
func (x Client) Close() error {
	return x.g.Close()
}
