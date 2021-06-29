package protoclient

import (
	"context"
	"errors"
	"io"

	"github.com/nspcc-dev/neofs-api-go/rpc/common"
	neofsgrpc "github.com/nspcc-dev/neofs-api-go/rpc/grpc"
	"github.com/nspcc-dev/neofs-api-go/rpc/message"
)

// MessageReadWriter is a component for transmitting raw Protobuf messages.
type MessageReadWriter struct {
	g neofsgrpc.MessageReadWriter
}

// ReadMessage reads the next message from the remote server,
// and writes it to the message.Message.
//
// Returns io.EOF if there are no more messages to read.
// ReadMessage should not be called after io.EOF occasion.
func (x MessageReadWriter) ReadMessage(m message.Message) error {
	// Can be optimized: we can create blank message here.
	gm := m.ToGRPCMessage()

	if err := x.g.ReadMessage(gm); err != nil {
		if errors.Is(err, io.EOF) {
			return io.EOF
		}

		return err
	}

	return m.FromGRPCMessage(gm)
}

// WriteMessage writes the next message.Message.
//
// WriteMessage should not be called after any error or after CloseSend.
func (x MessageReadWriter) WriteMessage(m message.Message) error {
	return x.g.WriteMessage(m.ToGRPCMessage())
}

// CloseSend ends writing messages to the server.
//
// All WriteMessage calls must be done before closing.
func (x MessageReadWriter) CloseSend() error {
	return x.g.CloseSend()
}

// RPCPrm groups the parameters of Client.RPC call.
type RPCPrm struct {
	g neofsgrpc.RPCPrm
}

// RPCRes groups the results of Client.RPC call.
type RPCRes struct {
	g neofsgrpc.RPCRes
}

// Communicate initiates a messaging session within the parameterized RPC.
// Connection must be opened before (Dial).
//
// Context is used for message exchange. Timeout can be configured using context.WithTimeout.
//
// If there is no error, res.Messager() can be used for communication. It should be closed finally.
//
// Panics if ctx or res is nil, CallMethodInfo has empty service or method name.
func (x Client) RPC(ctx context.Context, prm RPCPrm, res *RPCRes) error {
	return x.g.RPC(ctx, prm.g, &res.g)
}

// SetCallMethodInfo sets information about the RPC service and method.
//
// Required parameter, service and names must not be empty.
func (x *RPCPrm) SetCallMethodInfo(mInfo common.CallMethodInfo) {
	x.g.SetCallMethodInfo(mInfo)
}

// Messager returns initialized MessageReadWriter which can be used for message exchange.
//
// Should be called only after successful Client.RPC call.
// MessageReadWriter should be closed after communication is completed.
func (x RPCRes) Messager() MessageReadWriter {
	return MessageReadWriter{
		g: x.g.Messager(),
	}
}
