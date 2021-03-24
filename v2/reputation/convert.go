package reputation

import (
	"github.com/nspcc-dev/neofs-api-go/rpc/grpc"
	"github.com/nspcc-dev/neofs-api-go/rpc/message"
	reputation "github.com/nspcc-dev/neofs-api-go/v2/reputation/grpc"
)

// ToGRPCMessage converts Trust to gRPC-generated
// reputation.Trust message.
func (x *Trust) ToGRPCMessage() grpc.Message {
	var m *reputation.Trust

	if x != nil {
		m = new(reputation.Trust)

		m.SetValue(x.val)
		m.SetPeer(x.peer)
	}

	return m
}

// FromGRPCMessage tries to restore Trust from grpc.Message.
//
// Returns message.ErrUnexpectedMessageType if m is not
// a gRPC-generated reputation.Trust message.
func (x *Trust) FromGRPCMessage(m grpc.Message) error {
	v, ok := m.(*reputation.Trust)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	x.val = v.GetValue()
	x.peer = v.GetPeer()

	return nil
}

// TrustsToGRPC converts slice of Trust structures
// to slice of gRPC-generated Trust messages.
func TrustsToGRPC(xs []*Trust) (res []*reputation.Trust) {
	if xs != nil {
		res = make([]*reputation.Trust, 0, len(xs))

		for i := range xs {
			res = append(res, xs[i].ToGRPCMessage().(*reputation.Trust))
		}
	}

	return
}

// TrustsFromGRPC tries to restore slice of Trust structures from
// slice of gRPC-generated reputation.Trust messages.
func TrustsFromGRPC(xs []*reputation.Trust) (res []*Trust, err error) {
	if xs != nil {
		res = make([]*Trust, 0, len(xs))

		for i := range xs {
			var x *Trust

			if xs[i] != nil {
				x = new(Trust)

				err = x.FromGRPCMessage(xs[i])
				if err != nil {
					return
				}
			}

			res = append(res, x)
		}
	}

	return
}

// ToGRPCMessage converts SendLocalTrustRequestBody to gRPC-generated
// reputation.SendLocalTrustRequest_Body message.
func (x *SendLocalTrustRequestBody) ToGRPCMessage() grpc.Message {
	var m *reputation.SendLocalTrustRequest_Body

	if x != nil {
		m = new(reputation.SendLocalTrustRequest_Body)

		m.SetEpoch(x.epoch)
		m.SetTrusts(TrustsToGRPC(x.trusts))
	}

	return m
}

// FromGRPCMessage tries to restore SendLocalTrustRequestBody from grpc.Message.
//
// Returns message.ErrUnexpectedMessageType if m is not
// a gRPC-generated reputation.SendLocalTrustRequest_Body message.
func (x *SendLocalTrustRequestBody) FromGRPCMessage(m grpc.Message) error {
	v, ok := m.(*reputation.SendLocalTrustRequest_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	x.trusts, err = TrustsFromGRPC(v.GetTrusts())
	if err != nil {
		return err
	}

	x.epoch = v.GetEpoch()

	return nil
}

// ToGRPCMessage converts SendLocalTrustRequest to gRPC-generated
// reputation.SendLocalTrustRequest message.
func (x *SendLocalTrustRequest) ToGRPCMessage() grpc.Message {
	var m *reputation.SendLocalTrustRequest

	if x != nil {
		m = new(reputation.SendLocalTrustRequest)

		m.SetBody(x.body.ToGRPCMessage().(*reputation.SendLocalTrustRequest_Body))
		x.RequestHeaders.ToMessage(m)
	}

	return m
}

// FromGRPCMessage tries to restore SendLocalTrustRequest from grpc.Message.
//
// Returns message.ErrUnexpectedMessageType if m is not
// a gRPC-generated reputation.SendLocalTrustRequest message.
func (x *SendLocalTrustRequest) FromGRPCMessage(m grpc.Message) error {
	v, ok := m.(*reputation.SendLocalTrustRequest)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		x.body = nil
	} else {
		if x.body == nil {
			x.body = new(SendLocalTrustRequestBody)
		}

		err = x.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	return x.RequestHeaders.FromMessage(v)
}

// ToGRPCMessage converts SendLocalTrustResponseBody to gRPC-generated
// reputation.SendLocalTrustResponse_Body message.
func (x *SendLocalTrustResponseBody) ToGRPCMessage() grpc.Message {
	var m *reputation.SendLocalTrustResponse_Body

	if x != nil {
		m = new(reputation.SendLocalTrustResponse_Body)
	}

	return m
}

// FromGRPCMessage tries to restore SendLocalTrustResponseBody from grpc.Message.
//
// Returns message.ErrUnexpectedMessageType if m is not
// a gRPC-generated reputation.SendLocalTrustResponse_Body message.
func (x *SendLocalTrustResponseBody) FromGRPCMessage(m grpc.Message) error {
	v, ok := m.(*reputation.SendLocalTrustResponse_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	return nil
}

// ToGRPCMessage converts SendLocalTrustResponse to gRPC-generated
// reputation.SendLocalTrustResponse message.
func (x *SendLocalTrustResponse) ToGRPCMessage() grpc.Message {
	var m *reputation.SendLocalTrustResponse

	if x != nil {
		m = new(reputation.SendLocalTrustResponse)

		m.SetBody(x.body.ToGRPCMessage().(*reputation.SendLocalTrustResponse_Body))
		x.ResponseHeaders.ToMessage(m)
	}

	return m
}

// FromGRPCMessage tries to restore SendLocalTrustResponse from grpc.Message.
//
// Returns message.ErrUnexpectedMessageType if m is not
// a gRPC-generated reputation.SendLocalTrustResponse message.
func (x *SendLocalTrustResponse) FromGRPCMessage(m grpc.Message) error {
	v, ok := m.(*reputation.SendLocalTrustResponse)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		x.body = nil
	} else {
		if x.body == nil {
			x.body = new(SendLocalTrustResponseBody)
		}

		err = x.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	return x.ResponseHeaders.FromMessage(v)
}
