package protoclient

import (
	"context"

	"github.com/nspcc-dev/neofs-api-go/rpc/message"
)

// SendUnaryPrm inherits RPCPrm with additional
// parameters for SendUnary.
type SendUnaryPrm struct {
	RPCPrm

	req, resp message.Message
}

// OpenClientStreamRes groups the results of OpenClientStream.
type SendUnaryRes struct{}

// SetMessages sets request and response messages for the unary RPC communication.
//
// Must not be nil.
func (x *SendUnaryPrm) SetMessages(req, resp message.Message) {
	x.req, x.resp = req, resp
}

// SendUnary initializes communication session by RPC info, performs unary RPC
// and closes the session.
//
// Client connection should be established.
//
// Context and client must not be nil.
// SendUnaryRes is currently ignored and can be nil. It is left for potential occasion of results.
func SendUnary(ctx context.Context, cli Client, prm SendUnaryPrm, _ *SendUnaryRes) error {
	var res RPCRes

	err := cli.RPC(ctx, prm.RPCPrm, &res)
	if err != nil {
		return err
	}

	m := res.Messager()

	err = m.WriteMessage(prm.req)
	if err != nil {
		return err
	}

	err = m.CloseSend()
	if err != nil {
		return err
	}

	return m.ReadMessage(prm.resp)
}

// MessageWriterCloser is a wrapper over MessageReadWriter
// which shadows ReadMessage and CloseSend methods to combine
// them in single CloseSend method.
type MessageWriterCloser struct {
	m MessageReadWriter

	resp message.Message
}

// WriteMessage calls WriteMessage on underlying MessageReadWriter.
func (x MessageWriterCloser) WriteMessage(m message.Message) error {
	return x.m.WriteMessage(m)
}

// CloseSend calls CloseSend and reads the response message on underlying MessageReadWriter.
func (x MessageWriterCloser) CloseSend() error {
	err := x.m.CloseSend()
	if err != nil {
		return err
	}

	return x.m.ReadMessage(x.resp)
}

// OpenClientStreamPrm inherits RPCPrm with additional
// parameters for OpenClientStream.
type OpenClientStreamPrm struct {
	RPCPrm

	resp message.Message
}

// OpenClientStreamRes groups the results of OpenClientStream.
type OpenClientStreamRes struct {
	msgr MessageWriterCloser
}

// OpenClientStream initializes communication session by RPC info, opens client-side stream
// and wraps it into .
//
// All stream writes must be performed before the closing. CloseSend must be called once.
//
// Context, client and res must not be nil.
func OpenClientStream(ctx context.Context, cli Client, prm OpenClientStreamPrm, res *OpenClientStreamRes) error {
	var r RPCRes

	err := cli.RPC(ctx, prm.RPCPrm, &r)
	if err != nil {
		return err
	}

	res.msgr = MessageWriterCloser{
		m:    r.Messager(),
		resp: prm.resp,
	}

	return nil
}

// SetResponse sets response message for the client-streaming RPC communication.
func (x *OpenClientStreamPrm) SetResponse(resp message.Message) {
	x.resp = resp
}

// Messager initialized MessageReaderCloser which can be used for request writing.
//
// Should be closed after communication is completed. Parameterized response
// is after the successful Close.
func (x OpenClientStreamRes) Messager() MessageWriterCloser {
	return x.msgr
}

// MessageReaderCloser is a wrapper over MessageReadWriter
// which shadows WriteMessage and CloseSend methods method in redefined ReadMessage.
type MessageReaderCloser struct {
	m MessageReadWriter

	reqSent bool

	req message.Message
}

// ReadMessage calls WriteMessage on underlying MessageReadWriter on first call,
// and then performs ReadMessage. Calls CloseSend if ReadMessage returned io.EOF.
func (s *MessageReaderCloser) ReadMessage(msg message.Message) error {
	if !s.reqSent {
		if err := s.m.WriteMessage(s.req); err != nil {
			return err
		} else if err = s.m.CloseSend(); err != nil {
			return err
		}

		s.reqSent = true
	}

	return s.m.ReadMessage(msg)
}

// OpenServerStreamPrm inherits RPCPrm with additional
// parameters for OpenServerStream.
type OpenServerStreamPrm struct {
	RPCPrm

	req message.Message
}

// OpenServerStreamRes groups the results of OpenServerStream.
type OpenServerStreamRes struct {
	msgr MessageReaderCloser
}

// OpenServerStream initializes communication session by RPC info, opens server-side stream
// and returns its interface.
//
// All stream reads must be performed before the closing. CloseSend must be called once.
//
// Context, client and res must not be nil.
func OpenServerStream(ctx context.Context, cli Client, prm OpenServerStreamPrm, res *OpenServerStreamRes) error {
	var r RPCRes

	err := cli.RPC(ctx, prm.RPCPrm, &r)
	if err != nil {
		return err
	}

	res.msgr = MessageReaderCloser{
		m:   r.Messager(),
		req: prm.req,
	}

	return nil
}

// SetResponse sets request message for the server-streaming RPC communication.
func (x *OpenServerStreamPrm) SetRequest(resp message.Message) {
	x.req = resp
}

// Messager initialized MessageReaderCloser which can be used for response reading.
//
// Should be closed after communication is completed. Parameterized response
// is after the successful Close.
func (x OpenServerStreamRes) Messager() MessageReaderCloser {
	return x.msgr
}
