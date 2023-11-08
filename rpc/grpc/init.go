package grpc

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/nspcc-dev/neofs-api-go/v2/rpc/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/encoding/proto"
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
	timeout time.Duration
	cancel  context.CancelFunc
}

func (w streamWrapper) ReadMessage(m Message) error {
	return w.withTimeout(func() error {
		return w.ClientStream.RecvMsg(m)
	})
}

func (w streamWrapper) WriteMessage(m Message) error {
	return w.withTimeout(func() error {
		return w.ClientStream.SendMsg(m)
	})
}

func (w *streamWrapper) Close() error {
	return w.withTimeout(w.ClientStream.CloseSend)
}

func (w *streamWrapper) withTimeout(closure func() error) error {
	ch := make(chan error, 1)
	go func() {
		ch <- closure()
		close(ch)
	}()

	tt := time.NewTimer(w.timeout)

	select {
	case err := <-ch:
		tt.Stop()
		return err
	case <-tt.C:
		w.cancel()
		return context.DeadlineExceeded
	}
}

type onlyBinarySendingCodec struct{}

func (x onlyBinarySendingCodec) Name() string {
	// may be any non-empty, conflicts are unlikely to arise
	return "neofs_binary_sender"
}

func (x onlyBinarySendingCodec) Marshal(msg interface{}) ([]byte, error) {
	bMsg, ok := msg.([]byte)
	if ok {
		return bMsg, nil
	}

	return nil, fmt.Errorf("message is not of type %T", bMsg)
}

func (x onlyBinarySendingCodec) Unmarshal(raw []byte, msg interface{}) error {
	return encoding.GetCodec(proto.Name).Unmarshal(raw, msg)
}

// Init initiates a messaging session within the RPC configured by options.
func (c *Client) Init(info common.CallMethodInfo, opts ...CallOption) (MessageReadWriter, error) {
	prm := defaultCallParameters()

	for _, opt := range opts {
		opt(prm)
	}

	var grpcCallOpts []grpc.CallOption
	if prm.allowBinarySendingOnly {
		grpcCallOpts = []grpc.CallOption{grpc.ForceCodec(onlyBinarySendingCodec{})}
	}

	ctx, cancel := context.WithCancel(prm.ctx)
	stream, err := c.con.NewStream(ctx, &grpc.StreamDesc{
		StreamName:    info.Name,
		ServerStreams: info.ServerStream(),
		ClientStreams: info.ClientStream(),
	}, toMethodName(info), grpcCallOpts...)
	if err != nil {
		cancel()
		return nil, err
	}

	return &streamWrapper{
		ClientStream: stream,
		cancel:       cancel,
		timeout:      c.rwTimeout,
	}, nil
}
