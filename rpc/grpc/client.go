package neofsgrpc

import (
	"google.golang.org/grpc"
)

// Client represents client for exchanging messages
// with a remote server using gRPC protocol.
//
// Should be created using DialPrm.Dial.
//
// Client is one-time use, and should only be used within an open connection.
type Client struct {
	conn *grpc.ClientConn
}

// Close closes opened connection.
//
// Must be called only after successful initialization.
func (x Client) Close() error {
	return x.conn.Close()
}
