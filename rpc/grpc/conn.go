package grpc

import (
	"io"
)

// Conn returns underlying connection.
//
// Conn is NPE-safe: returns nil if Client is nil.
//
// Client should not be used after Close() call
// on the connection: behavior is undefined.
func (c *Client) Conn() io.Closer {
	if c != nil {
		return c.con
	}

	return nil
}
