package grpc

import (
	"io"

	"github.com/nspcc-dev/neofs-api-go/rpc/common"
	"google.golang.org/grpc"
)

// Message represents raw gRPC message.
type Message interface{}

// MessageReadWriter is a component interface
// for transmitting raw messages over gRPC protocol.
type MessageReadWriter interface {
	// ReadMessage reads the next message from the remote server,
	// and writes it to the argument.
	ReadMessage(Message) error

	// WriteMessage sends message from argument to remote server.
	WriteMessage(Message) error

	// Closes the communication session with the remote server.
	//
	// All calls to send/receive messages must be done before closing.
	io.Closer
}

type streamWrapper struct {
	grpc.ClientStream
}

func (w streamWrapper) ReadMessage(m Message) error {
	return w.ClientStream.RecvMsg(m)
}

func (w streamWrapper) WriteMessage(m Message) error {
	return w.ClientStream.SendMsg(m)
}

func (w *streamWrapper) Close() error {
	return w.ClientStream.CloseSend()
}

// Init initiates a messaging session within the RPC configured by options.
func (c *Client) Init(info common.CallMethodInfo, opts ...CallOption) (MessageReadWriter, error) {
	prm := defaultCallParameters()

	for _, opt := range opts {
		opt(prm)
	}

	stream, err := c.con.NewStream(prm.ctx, &grpc.StreamDesc{
		StreamName:    info.Name,
		ServerStreams: info.ServerStream(),
		ClientStreams: info.ClientStream(),
	}, toMethodName(info))
	if err != nil {
		return nil, err
	}

	return &streamWrapper{
		ClientStream: stream,
	}, nil
}
