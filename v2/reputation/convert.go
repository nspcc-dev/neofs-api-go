package reputation

import (
	"github.com/nspcc-dev/neofs-api-go/rpc/grpc"
	"github.com/nspcc-dev/neofs-api-go/rpc/message"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	refsGRPC "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
	reputation "github.com/nspcc-dev/neofs-api-go/v2/reputation/grpc"
)

// ToGRPCMessage converts PeerID to gRPC-generated
// reputation.PeerID message.
func (x *PeerID) ToGRPCMessage() grpc.Message {
	var m *reputation.PeerID

	if x != nil {
		m = new(reputation.PeerID)

		m.SetValue(x.val)
	}

	return m
}

// FromGRPCMessage tries to restore PeerID from grpc.Message.
//
// Returns message.ErrUnexpectedMessageType if m is not
// a gRPC-generated reputation.PeerID message.
func (x *PeerID) FromGRPCMessage(m grpc.Message) error {
	v, ok := m.(*reputation.PeerID)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	x.val = v.GetValue()

	return nil
}

// ToGRPCMessage converts Trust to gRPC-generated
// reputation.Trust message.
func (x *Trust) ToGRPCMessage() grpc.Message {
	var m *reputation.Trust

	if x != nil {
		m = new(reputation.Trust)

		m.SetValue(x.val)
		m.SetPeer(x.peer.ToGRPCMessage().(*reputation.PeerID))
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

	peer := v.GetPeer()
	if peer == nil {
		x.peer = nil
	} else {
		if x.peer == nil {
			x.peer = new(PeerID)
		}

		err := x.peer.FromGRPCMessage(peer)
		if err != nil {
			return err
		}
	}

	x.val = v.GetValue()

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

// ToGRPCMessage converts GlobalTrustBody to gRPC-generated
// reputation.GlobalTrust_Body message.
func (x *GlobalTrustBody) ToGRPCMessage() grpc.Message {
	var m *reputation.GlobalTrust_Body

	if x != nil {
		m = new(reputation.GlobalTrust_Body)

		m.SetManager(x.manager.ToGRPCMessage().(*reputation.PeerID))
		m.SetTrust(x.trust.ToGRPCMessage().(*reputation.Trust))
	}

	return m
}

// FromGRPCMessage tries to restore GlobalTrustBody from grpc.Message.
//
// Returns message.ErrUnexpectedMessageType if m is not
// a gRPC-generated reputation.GlobalTrust_Body message.
func (x *GlobalTrustBody) FromGRPCMessage(m grpc.Message) error {
	v, ok := m.(*reputation.GlobalTrust_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	manager := v.GetManager()
	if manager == nil {
		x.manager = nil
	} else {
		if x.manager == nil {
			x.manager = new(PeerID)
		}

		err = x.manager.FromGRPCMessage(manager)
		if err != nil {
			return err
		}
	}

	trust := v.GetTrust()
	if trust == nil {
		x.trust = nil
	} else {
		if x.trust == nil {
			x.trust = new(Trust)
		}

		err = x.trust.FromGRPCMessage(trust)
	}

	return err
}

// ToGRPCMessage converts GlobalTrust to gRPC-generated
// reputation.GlobalTrust message.
func (x *GlobalTrust) ToGRPCMessage() grpc.Message {
	var m *reputation.GlobalTrust

	if x != nil {
		m = new(reputation.GlobalTrust)

		m.SetVersion(x.version.ToGRPCMessage().(*refsGRPC.Version))
		m.SetBody(x.body.ToGRPCMessage().(*reputation.GlobalTrust_Body))
		m.SetSignature(x.sig.ToGRPCMessage().(*refsGRPC.Signature))
	}

	return m
}

// FromGRPCMessage tries to restore GlobalTrust from grpc.Message.
//
// Returns message.ErrUnexpectedMessageType if m is not
// a gRPC-generated reputation.GlobalTrust message.
func (x *GlobalTrust) FromGRPCMessage(m grpc.Message) error {
	v, ok := m.(*reputation.GlobalTrust)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	version := v.GetVersion()
	if version == nil {
		x.version = nil
	} else {
		if x.version == nil {
			x.version = new(refs.Version)
		}

		err = x.version.FromGRPCMessage(version)
		if err != nil {
			return err
		}
	}

	body := v.GetBody()
	if body == nil {
		x.body = nil
	} else {
		if x.body == nil {
			x.body = new(GlobalTrustBody)
		}

		err = x.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	sig := v.GetSignature()
	if sig == nil {
		x.sig = nil
	} else {
		if x.sig == nil {
			x.sig = new(refs.Signature)
		}

		err = x.sig.FromGRPCMessage(sig)
	}

	return err
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

// ToGRPCMessage converts SendIntermediateResultRequestBody to gRPC-generated
// reputation.SendIntermediateResultRequest_Body message.
func (x *SendIntermediateResultRequestBody) ToGRPCMessage() grpc.Message {
	var m *reputation.SendIntermediateResultRequest_Body

	if x != nil {
		m = new(reputation.SendIntermediateResultRequest_Body)

		m.SetIteration(x.iter)
		m.SetTrust(x.trust.ToGRPCMessage().(*reputation.Trust))
	}

	return m
}

// FromGRPCMessage tries to restore SendIntermediateResultRequestBody from grpc.Message.
//
// Returns message.ErrUnexpectedMessageType if m is not
// a gRPC-generated reputation.SendIntermediateResultRequest_Body message.
func (x *SendIntermediateResultRequestBody) FromGRPCMessage(m grpc.Message) error {
	v, ok := m.(*reputation.SendIntermediateResultRequest_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	err := x.trust.FromGRPCMessage(v.GetTrust())
	if err != nil {
		return err
	}

	x.iter = v.GetIteration()

	return nil
}

// ToGRPCMessage converts SendIntermediateResultRequest to gRPC-generated
// reputation.SendIntermediateResultRequest message.
func (x *SendIntermediateResultRequest) ToGRPCMessage() grpc.Message {
	var m *reputation.SendIntermediateResultRequest

	if x != nil {
		m = new(reputation.SendIntermediateResultRequest)

		m.SetBody(x.body.ToGRPCMessage().(*reputation.SendIntermediateResultRequest_Body))
		x.RequestHeaders.ToMessage(m)
	}

	return m
}

// FromGRPCMessage tries to restore SendIntermediateResultRequest from grpc.Message.
//
// Returns message.ErrUnexpectedMessageType if m is not
// a gRPC-generated reputation.SendIntermediateResultRequest message.
func (x *SendIntermediateResultRequest) FromGRPCMessage(m grpc.Message) error {
	v, ok := m.(*reputation.SendIntermediateResultRequest)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		x.body = nil
	} else {
		if x.body == nil {
			x.body = new(SendIntermediateResultRequestBody)
		}

		err = x.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	return x.RequestHeaders.FromMessage(v)
}

// ToGRPCMessage converts SendIntermediateResultResponseBody to gRPC-generated
// reputation.SendIntermediateResultResponse_Body message.
func (x *SendIntermediateResultResponseBody) ToGRPCMessage() grpc.Message {
	var m *reputation.SendIntermediateResultResponse_Body

	if x != nil {
		m = new(reputation.SendIntermediateResultResponse_Body)
	}

	return m
}

// FromGRPCMessage tries to restore SendIntermediateResultResponseBody from grpc.Message.
//
// Returns message.ErrUnexpectedMessageType if m is not
// a gRPC-generated reputation.SendIntermediateResultResponse_Body message.
func (x *SendIntermediateResultResponseBody) FromGRPCMessage(m grpc.Message) error {
	v, ok := m.(*reputation.SendIntermediateResultResponse_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	return nil
}

// ToGRPCMessage converts SendIntermediateResultResponse to gRPC-generated
// reputation.SendIntermediateResultResponse message.
func (x *SendIntermediateResultResponse) ToGRPCMessage() grpc.Message {
	var m *reputation.SendIntermediateResultResponse

	if x != nil {
		m = new(reputation.SendIntermediateResultResponse)

		m.SetBody(x.body.ToGRPCMessage().(*reputation.SendIntermediateResultResponse_Body))
		x.ResponseHeaders.ToMessage(m)
	}

	return m
}

// FromGRPCMessage tries to restore SendIntermediateResultResponse from grpc.Message.
//
// Returns message.ErrUnexpectedMessageType if m is not
// a gRPC-generated reputation.SendIntermediateResultResponse message.
func (x *SendIntermediateResultResponse) FromGRPCMessage(m grpc.Message) error {
	v, ok := m.(*reputation.SendIntermediateResultResponse)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		x.body = nil
	} else {
		if x.body == nil {
			x.body = new(SendIntermediateResultResponseBody)
		}

		err = x.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	return x.ResponseHeaders.FromMessage(v)
}
