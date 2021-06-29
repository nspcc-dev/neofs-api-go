package neofsgrpc

import (
	"context"

	"github.com/nspcc-dev/neofs-api-go/rpc/common"
	"google.golang.org/grpc"
)

// Message represents raw gRPC message.
type Message interface{}

// MessageReadWriter is a component for transmitting raw messages over gRPC protocol.
type MessageReadWriter struct {
	s grpc.ClientStream
}

// ReadMessage reads the next message from the remote server,
// and writes it to the Message.
func (w MessageReadWriter) ReadMessage(m Message) error {
	return w.s.RecvMsg(m)
}

// WriteMessage sends Message to the remote server.
//
// Should not be called after the CloseSend.
func (w MessageReadWriter) WriteMessage(m Message) error {
	return w.s.SendMsg(m)
}

// CloseSend ends writing messages to the server.
//
// All WriteMessage calls should be done before closing.
func (w MessageReadWriter) CloseSend() error {
	return w.s.CloseSend()
}

// RPCPrm groups the parameters of Client.RPC call.
type RPCPrm struct {
	mInfo common.CallMethodInfo
}

// RPCRes groups the results of Client.RPC call.
type RPCRes struct {
	msgr MessageReadWriter
}

// RPC initiates a messaging session within the parameterized RPC.
// Underlying connection must be initialized before.
//
// Context is used for message exchange. Timeout can be configured using context.WithTimeout.
//
// If there is no error, res.Messager() can be used for communication. It should be closed finally.
//
// Panics if ctx or res is nil, CallMethodInfo has empty service or method name.
func (x Client) RPC(ctx context.Context, prm RPCPrm, res *RPCRes) error {
	svcName, mtdName := prm.mInfo.ServiceName(), prm.mInfo.MethodName()

	switch {
	case svcName == "" || mtdName == "":
		panic("invalid RPC method info")
	case ctx == nil:
		panic("no context provided")
	}

	stream, err := x.conn.NewStream(ctx, &grpc.StreamDesc{
		StreamName:    mtdName,
		ServerStreams: prm.mInfo.ServerStream(),
		ClientStreams: prm.mInfo.ClientStream(),
	}, toMethodName(svcName, mtdName))
	if err != nil {
		return err
	}

	res.msgr.s = stream

	return nil
}

// SetCallMethodInfo sets information about the RPC service and method.
//
// Required parameter, service and names must not be empty.
func (x *RPCPrm) SetCallMethodInfo(mInfo common.CallMethodInfo) {
	x.mInfo = mInfo
}

// Messager returns initialized MessageReadWriter which can be used for message exchange.
//
// Should be called only after successful Client.RPC call.
// MessageReadWriter should be closed after communication is completed.
func (x RPCRes) Messager() MessageReadWriter {
	return x.msgr
}
