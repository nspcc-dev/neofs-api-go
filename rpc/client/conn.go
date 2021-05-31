package client

import (
	"io"
)

// Conn returns underlying connection.
//
// Returns non-nil result after the first Init() call
// completed without a connection error.
//
// Conn is NPE-safe: returns nil if Client is nil.
//
// Client should not be used after Close() call
// on the connection: behavior is undefined.
func (c *Client) Conn() io.Closer {
	if c != nil {
		return c.gRPCClient.Conn()
	}

	return nil
}
