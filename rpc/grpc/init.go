package grpc

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/nspcc-dev/neofs-api-go/v2/rpc/common"
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
	timeout time.Duration
	cancel  context.CancelFunc
	binOnly bool
}

func (w streamWrapper) ReadMessage(m Message) error {
	return w.withTimeout(func() error {
		return w.ClientStream.RecvMsg(m)
	})
}

func (w streamWrapper) WriteMessage(m Message) error {
	if w.binOnly {
		dr, ok := m.(DataReader)
		if !ok {
			return fmt.Errorf("only %T is allowed with AllowBinarySendingOnly option", dr)
		}
		m = grpc.DataReader{
			R:    dr.R,
			Size: dr.Size,
		}
	}
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

// DataReader combines [io.Reader] and size.
type DataReader struct {
	R    io.Reader
	Size int
}

// Init initiates a messaging session within the RPC configured by options.
func (c *Client) Init(info common.CallMethodInfo, opts ...CallOption) (MessageReadWriter, error) {
	prm := defaultCallParameters()

	for _, opt := range opts {
		opt(prm)
	}

	if prm.allowBinarySendingOnly && (info.ClientStream() || info.ServerStream()) {
		return nil, fmt.Errorf("AllowBinarySendingOnly option is allowed for unary RPC only")
	}

	ctx, cancel := context.WithCancel(prm.ctx)
	stream, err := c.con.NewStream(ctx, &grpc.StreamDesc{
		StreamName:    info.Name,
		ServerStreams: info.ServerStream(),
		ClientStreams: info.ClientStream(),
	}, toMethodName(info))
	if err != nil {
		cancel()
		return nil, err
	}

	return &streamWrapper{
		ClientStream: stream,
		timeout:      c.rwTimeout,
		cancel:       cancel,
		binOnly:      prm.allowBinarySendingOnly,
	}, nil
}
