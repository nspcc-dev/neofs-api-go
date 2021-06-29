package client

import protoclient "github.com/nspcc-dev/neofs-api-go/rpc/client"

// Client represents NeoFS API client.
// It is a wrapper over protoclient.Client.
//
// Should be initialized using DialPrm.Dial or SetProtoClient method.
//
// Client is one-time use, and should only be used within an open connection.
type Client struct {
	c protoclient.Client
}

// SetProtoClient sets underlying protoclient.Client.
func (x *Client) SetProtoClient(c protoclient.Client) {
	x.c = c
}

// Close closes opened connection.
//
// Must be called only after successful initialization.
func (x Client) Close() error {
	return x.c.Close()
}
