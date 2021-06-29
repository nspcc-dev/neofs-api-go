package container

import (
	neofsgrpc "github.com/nspcc-dev/neofs-api-go/rpc/grpc"
	"github.com/nspcc-dev/neofs-api-go/rpc/message"
	"github.com/nspcc-dev/neofs-api-go/v2/acl"
	aclGRPC "github.com/nspcc-dev/neofs-api-go/v2/acl/grpc"
	container "github.com/nspcc-dev/neofs-api-go/v2/container/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	netmapGRPC "github.com/nspcc-dev/neofs-api-go/v2/netmap/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	refsGRPC "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/session"
	sessionGRPC "github.com/nspcc-dev/neofs-api-go/v2/session/grpc"
)

func (a *Attribute) ToGRPCMessage() neofsgrpc.Message {
	var m *container.Container_Attribute

	if a != nil {
		m = new(container.Container_Attribute)

		m.SetKey(a.key)
		m.SetValue(a.val)
	}

	return m
}

func (a *Attribute) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.Container_Attribute)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	a.key = v.GetKey()
	a.val = v.GetValue()

	return nil
}

func AttributesToGRPC(xs []*Attribute) (res []*container.Container_Attribute) {
	if xs != nil {
		res = make([]*container.Container_Attribute, 0, len(xs))

		for i := range xs {
			res = append(res, xs[i].ToGRPCMessage().(*container.Container_Attribute))
		}
	}

	return
}

func AttributesFromGRPC(xs []*container.Container_Attribute) (res []*Attribute, err error) {
	if xs != nil {
		res = make([]*Attribute, 0, len(xs))

		for i := range xs {
			var x *Attribute

			if xs[i] != nil {
				x = new(Attribute)

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

func (c *Container) ToGRPCMessage() neofsgrpc.Message {
	var m *container.Container

	if c != nil {
		m = new(container.Container)

		m.SetVersion(c.version.ToGRPCMessage().(*refsGRPC.Version))
		m.SetOwnerId(c.ownerID.ToGRPCMessage().(*refsGRPC.OwnerID))
		m.SetPlacementPolicy(c.policy.ToGRPCMessage().(*netmapGRPC.PlacementPolicy))
		m.SetAttributes(AttributesToGRPC(c.attr))
		m.SetBasicAcl(c.basicACL)
		m.SetNonce(c.nonce)
	}

	return m
}

func (c *Container) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.Container)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	version := v.GetVersion()
	if version == nil {
		c.version = nil
	} else {
		if c.version == nil {
			c.version = new(refs.Version)
		}

		err = c.version.FromGRPCMessage(version)
		if err != nil {
			return err
		}
	}

	ownerID := v.GetOwnerId()
	if ownerID == nil {
		c.ownerID = nil
	} else {
		if c.ownerID == nil {
			c.ownerID = new(refs.OwnerID)
		}

		err = c.ownerID.FromGRPCMessage(ownerID)
		if err != nil {
			return err
		}
	}

	policy := v.GetPlacementPolicy()
	if policy == nil {
		c.policy = nil
	} else {
		if c.policy == nil {
			c.policy = new(netmap.PlacementPolicy)
		}

		err = c.policy.FromGRPCMessage(policy)
		if err != nil {
			return err
		}
	}

	c.attr, err = AttributesFromGRPC(v.GetAttributes())
	if err != nil {
		return err
	}

	c.basicACL = v.GetBasicAcl()
	c.nonce = v.GetNonce()

	return nil
}

func (r *PutRequestBody) ToGRPCMessage() neofsgrpc.Message {
	var m *container.PutRequest_Body

	if r != nil {
		m = new(container.PutRequest_Body)

		m.SetContainer(r.cnr.ToGRPCMessage().(*container.Container))
		m.SetSignature(r.sig.ToGRPCMessage().(*refsGRPC.Signature))
	}

	return m
}

func (r *PutRequestBody) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.PutRequest_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	cnr := v.GetContainer()
	if cnr == nil {
		r.cnr = nil
	} else {
		if r.cnr == nil {
			r.cnr = new(Container)
		}

		err = r.cnr.FromGRPCMessage(cnr)
		if err != nil {
			return err
		}
	}

	sig := v.GetSignature()
	if sig == nil {
		r.sig = nil
	} else {
		if r.sig == nil {
			r.sig = new(refs.Signature)
		}

		err = r.sig.FromGRPCMessage(sig)
	}

	return err
}

func (r *PutRequest) ToGRPCMessage() neofsgrpc.Message {
	var m *container.PutRequest

	if r != nil {
		m = new(container.PutRequest)

		m.SetBody(r.body.ToGRPCMessage().(*container.PutRequest_Body))
		r.RequestHeaders.ToMessage(m)
	}

	return m
}

func (r *PutRequest) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.PutRequest)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		r.body = nil
	} else {
		if r.body == nil {
			r.body = new(PutRequestBody)
		}

		err = r.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	return r.RequestHeaders.FromMessage(v)
}

func (r *PutResponseBody) ToGRPCMessage() neofsgrpc.Message {
	var m *container.PutResponse_Body

	if r != nil {
		m = new(container.PutResponse_Body)

		m.SetContainerId(r.cid.ToGRPCMessage().(*refsGRPC.ContainerID))
	}

	return m
}

func (r *PutResponseBody) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.PutResponse_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	cid := v.GetContainerId()
	if cid == nil {
		r.cid = nil
	} else {
		if r.cid == nil {
			r.cid = new(refs.ContainerID)
		}

		err = r.cid.FromGRPCMessage(cid)
	}

	return err
}

func (r *PutResponse) ToGRPCMessage() neofsgrpc.Message {
	var m *container.PutResponse

	if r != nil {
		m = new(container.PutResponse)

		m.SetBody(r.body.ToGRPCMessage().(*container.PutResponse_Body))
		r.ResponseHeaders.ToMessage(m)
	}

	return m
}

func (r *PutResponse) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.PutResponse)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		r.body = nil
	} else {
		if r.body == nil {
			r.body = new(PutResponseBody)
		}

		err = r.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	return r.ResponseHeaders.FromMessage(v)
}

func (r *GetRequestBody) ToGRPCMessage() neofsgrpc.Message {
	var m *container.GetRequest_Body

	if r != nil {
		m = new(container.GetRequest_Body)

		m.SetContainerId(r.cid.ToGRPCMessage().(*refsGRPC.ContainerID))
	}

	return m
}

func (r *GetRequestBody) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.GetRequest_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	cid := v.GetContainerId()
	if cid == nil {
		r.cid = nil
	} else {
		if r.cid == nil {
			r.cid = new(refs.ContainerID)
		}

		err = r.cid.FromGRPCMessage(cid)
	}

	return err
}

func (r *GetRequest) ToGRPCMessage() neofsgrpc.Message {
	var m *container.GetRequest

	if r != nil {
		m = new(container.GetRequest)

		m.SetBody(r.body.ToGRPCMessage().(*container.GetRequest_Body))
		r.RequestHeaders.ToMessage(m)
	}

	return m
}

func (r *GetRequest) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.GetRequest)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		r.body = nil
	} else {
		if r.body == nil {
			r.body = new(GetRequestBody)
		}

		err = r.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	return r.RequestHeaders.FromMessage(v)
}

func (r *GetResponseBody) ToGRPCMessage() neofsgrpc.Message {
	var m *container.GetResponse_Body

	if r != nil {
		m = new(container.GetResponse_Body)

		m.SetContainer(r.cnr.ToGRPCMessage().(*container.Container))
		m.SetSessionToken(r.token.ToGRPCMessage().(*sessionGRPC.SessionToken))
		m.SetSignature(r.sig.ToGRPCMessage().(*refsGRPC.Signature))
	}

	return m
}

func (r *GetResponseBody) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.GetResponse_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	cnr := v.GetContainer()
	if cnr == nil {
		r.cnr = nil
	} else {
		if r.cnr == nil {
			r.cnr = new(Container)
		}

		err = r.cnr.FromGRPCMessage(cnr)
	}

	sig := v.GetSignature()
	if sig == nil {
		r.sig = nil
	} else {
		if r.sig == nil {
			r.sig = new(refs.Signature)
		}

		err = r.sig.FromGRPCMessage(sig)
	}

	token := v.GetSessionToken()
	if token == nil {
		r.token = nil
	} else {
		if r.token == nil {
			r.token = new(session.SessionToken)
		}

		err = r.token.FromGRPCMessage(token)
	}

	return err
}

func (r *GetResponse) ToGRPCMessage() neofsgrpc.Message {
	var m *container.GetResponse

	if r != nil {
		m = new(container.GetResponse)

		m.SetBody(r.body.ToGRPCMessage().(*container.GetResponse_Body))
		r.ResponseHeaders.ToMessage(m)
	}

	return m
}

func (r *GetResponse) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.GetResponse)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		r.body = nil
	} else {
		if r.body == nil {
			r.body = new(GetResponseBody)
		}

		err = r.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	return r.ResponseHeaders.FromMessage(v)
}

func (r *DeleteRequestBody) ToGRPCMessage() neofsgrpc.Message {
	var m *container.DeleteRequest_Body

	if r != nil {
		m = new(container.DeleteRequest_Body)

		m.SetContainerId(r.cid.ToGRPCMessage().(*refsGRPC.ContainerID))
		m.SetSignature(r.sig.ToGRPCMessage().(*refsGRPC.Signature))
	}

	return m
}

func (r *DeleteRequestBody) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.DeleteRequest_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	cid := v.GetContainerId()
	if cid == nil {
		r.cid = nil
	} else {
		if r.cid == nil {
			r.cid = new(refs.ContainerID)
		}

		err = r.cid.FromGRPCMessage(cid)
		if err != nil {
			return err
		}
	}

	sig := v.GetSignature()
	if sig == nil {
		r.sig = nil
	} else {
		if r.sig == nil {
			r.sig = new(refs.Signature)
		}

		err = r.sig.FromGRPCMessage(sig)
	}

	return err
}

func (r *DeleteRequest) ToGRPCMessage() neofsgrpc.Message {
	var m *container.DeleteRequest

	if r != nil {
		m = new(container.DeleteRequest)

		m.SetBody(r.body.ToGRPCMessage().(*container.DeleteRequest_Body))
		r.RequestHeaders.ToMessage(m)
	}

	return m
}

func (r *DeleteRequest) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.DeleteRequest)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		r.body = nil
	} else {
		if r.body == nil {
			r.body = new(DeleteRequestBody)
		}

		err = r.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	return r.RequestHeaders.FromMessage(v)
}

func (r *DeleteResponseBody) ToGRPCMessage() neofsgrpc.Message {
	var m *container.DeleteResponse_Body

	if r != nil {
		m = new(container.DeleteResponse_Body)
	}

	return m
}

func (r *DeleteResponseBody) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.DeleteResponse_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	return nil
}

func (r *DeleteResponse) ToGRPCMessage() neofsgrpc.Message {
	var m *container.DeleteResponse

	if r != nil {
		m = new(container.DeleteResponse)

		m.SetBody(r.body.ToGRPCMessage().(*container.DeleteResponse_Body))
		r.ResponseHeaders.ToMessage(m)
	}

	return m
}

func (r *DeleteResponse) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.DeleteResponse)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		r.body = nil
	} else {
		if r.body == nil {
			r.body = new(DeleteResponseBody)
		}

		err = r.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	return r.ResponseHeaders.FromMessage(v)
}

func (r *ListRequestBody) ToGRPCMessage() neofsgrpc.Message {
	var m *container.ListRequest_Body

	if r != nil {
		m = new(container.ListRequest_Body)

		m.SetOwnerId(r.ownerID.ToGRPCMessage().(*refsGRPC.OwnerID))
	}

	return m
}

func (r *ListRequestBody) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.ListRequest_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	ownerID := v.GetOwnerId()
	if ownerID == nil {
		r.ownerID = nil
	} else {
		if r.ownerID == nil {
			r.ownerID = new(refs.OwnerID)
		}

		err = r.ownerID.FromGRPCMessage(ownerID)
	}

	return err
}

func (r *ListRequest) ToGRPCMessage() neofsgrpc.Message {
	var m *container.ListRequest

	if r != nil {
		m = new(container.ListRequest)

		m.SetBody(r.body.ToGRPCMessage().(*container.ListRequest_Body))
		r.RequestHeaders.ToMessage(m)
	}

	return m
}

func (r *ListRequest) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.ListRequest)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		r.body = nil
	} else {
		if r.body == nil {
			r.body = new(ListRequestBody)
		}

		err = r.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	return r.RequestHeaders.FromMessage(v)
}

func (r *ListResponseBody) ToGRPCMessage() neofsgrpc.Message {
	var m *container.ListResponse_Body

	if r != nil {
		m = new(container.ListResponse_Body)

		m.SetContainerIds(refs.ContainerIDsToGRPCMessage(r.cidList))
	}

	return m
}

func (r *ListResponseBody) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.ListResponse_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	r.cidList, err = refs.ContainerIDsFromGRPCMessage(v.GetContainerIds())

	return err
}

func (r *ListResponse) ToGRPCMessage() neofsgrpc.Message {
	var m *container.ListResponse

	if r != nil {
		m = new(container.ListResponse)

		m.SetBody(r.body.ToGRPCMessage().(*container.ListResponse_Body))
		r.ResponseHeaders.ToMessage(m)
	}

	return m
}

func (r *ListResponse) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.ListResponse)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		r.body = nil
	} else {
		if r.body == nil {
			r.body = new(ListResponseBody)
		}

		err = r.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	return r.ResponseHeaders.FromMessage(v)
}

func (r *SetExtendedACLRequestBody) ToGRPCMessage() neofsgrpc.Message {
	var m *container.SetExtendedACLRequest_Body

	if r != nil {
		m = new(container.SetExtendedACLRequest_Body)

		m.SetEacl(r.eacl.ToGRPCMessage().(*aclGRPC.EACLTable))
		m.SetSignature(r.sig.ToGRPCMessage().(*refsGRPC.Signature))
	}

	return m
}

func (r *SetExtendedACLRequestBody) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.SetExtendedACLRequest_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	eacl := v.GetEacl()
	if eacl == nil {
		r.eacl = nil
	} else {
		if r.eacl == nil {
			r.eacl = new(acl.Table)
		}

		err = r.eacl.FromGRPCMessage(eacl)
		if err != nil {
			return err
		}
	}

	sig := v.GetSignature()
	if sig == nil {
		r.sig = nil
	} else {
		if r.sig == nil {
			r.sig = new(refs.Signature)
		}

		err = r.sig.FromGRPCMessage(sig)
	}

	return err
}

func (r *SetExtendedACLRequest) ToGRPCMessage() neofsgrpc.Message {
	var m *container.SetExtendedACLRequest

	if r != nil {
		m = new(container.SetExtendedACLRequest)

		m.SetBody(r.body.ToGRPCMessage().(*container.SetExtendedACLRequest_Body))
		r.RequestHeaders.ToMessage(m)
	}

	return m
}

func (r *SetExtendedACLRequest) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.SetExtendedACLRequest)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		r.body = nil
	} else {
		if r.body == nil {
			r.body = new(SetExtendedACLRequestBody)
		}

		err = r.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	return r.RequestHeaders.FromMessage(v)
}

func (r *SetExtendedACLResponseBody) ToGRPCMessage() neofsgrpc.Message {
	var m *container.SetExtendedACLResponse_Body

	if r != nil {
		m = new(container.SetExtendedACLResponse_Body)
	}

	return m
}

func (r *SetExtendedACLResponseBody) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.SetExtendedACLResponse_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	return nil
}

func (r *SetExtendedACLResponse) ToGRPCMessage() neofsgrpc.Message {
	var m *container.SetExtendedACLResponse

	if r != nil {
		m = new(container.SetExtendedACLResponse)

		m.SetBody(r.body.ToGRPCMessage().(*container.SetExtendedACLResponse_Body))
		r.ResponseHeaders.ToMessage(m)
	}

	return m
}

func (r *SetExtendedACLResponse) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.SetExtendedACLResponse)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		r.body = nil
	} else {
		if r.body == nil {
			r.body = new(SetExtendedACLResponseBody)
		}

		err = r.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	return r.ResponseHeaders.FromMessage(v)
}

func (r *GetExtendedACLRequestBody) ToGRPCMessage() neofsgrpc.Message {
	var m *container.GetExtendedACLRequest_Body

	if r != nil {
		m = new(container.GetExtendedACLRequest_Body)

		m.SetContainerId(r.cid.ToGRPCMessage().(*refsGRPC.ContainerID))
	}

	return m
}

func (r *GetExtendedACLRequestBody) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.GetExtendedACLRequest_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	cid := v.GetContainerId()
	if cid == nil {
		r.cid = nil
	} else {
		if r.cid == nil {
			r.cid = new(refs.ContainerID)
		}

		err = r.cid.FromGRPCMessage(cid)
	}

	return err
}

func (r *GetExtendedACLRequest) ToGRPCMessage() neofsgrpc.Message {
	var m *container.GetExtendedACLRequest

	if r != nil {
		m = new(container.GetExtendedACLRequest)

		m.SetBody(r.body.ToGRPCMessage().(*container.GetExtendedACLRequest_Body))
		r.RequestHeaders.ToMessage(m)
	}

	return m
}

func (r *GetExtendedACLRequest) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.GetExtendedACLRequest)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		r.body = nil
	} else {
		if r.body == nil {
			r.body = new(GetExtendedACLRequestBody)
		}

		err = r.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	return r.RequestHeaders.FromMessage(v)
}

func (r *GetExtendedACLResponseBody) ToGRPCMessage() neofsgrpc.Message {
	var m *container.GetExtendedACLResponse_Body

	if r != nil {
		m = new(container.GetExtendedACLResponse_Body)

		m.SetEacl(r.eacl.ToGRPCMessage().(*aclGRPC.EACLTable))
		m.SetSignature(r.sig.ToGRPCMessage().(*refsGRPC.Signature))
		m.SetSessionToken(r.token.ToGRPCMessage().(*sessionGRPC.SessionToken))
	}

	return m
}

func (r *GetExtendedACLResponseBody) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.GetExtendedACLResponse_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	eacl := v.GetEacl()
	if eacl == nil {
		r.eacl = nil
	} else {
		if r.eacl == nil {
			r.eacl = new(acl.Table)
		}

		err = r.eacl.FromGRPCMessage(eacl)
		if err != nil {
			return err
		}
	}

	sig := v.GetSignature()
	if sig == nil {
		r.sig = nil
	} else {
		if r.sig == nil {
			r.sig = new(refs.Signature)
		}

		err = r.sig.FromGRPCMessage(sig)
	}

	token := v.GetSessionToken()
	if token == nil {
		r.token = nil
	} else {
		if r.token == nil {
			r.token = new(session.SessionToken)
		}

		err = r.token.FromGRPCMessage(token)
	}

	return err
}

func (r *GetExtendedACLResponse) ToGRPCMessage() neofsgrpc.Message {
	var m *container.GetExtendedACLResponse

	if r != nil {
		m = new(container.GetExtendedACLResponse)

		m.SetBody(r.body.ToGRPCMessage().(*container.GetExtendedACLResponse_Body))
		r.ResponseHeaders.ToMessage(m)
	}

	return m
}

func (r *GetExtendedACLResponse) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.GetExtendedACLResponse)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		r.body = nil
	} else {
		if r.body == nil {
			r.body = new(GetExtendedACLResponseBody)
		}

		err = r.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	return r.ResponseHeaders.FromMessage(v)
}

func (a *UsedSpaceAnnouncement) ToGRPCMessage() neofsgrpc.Message {
	var m *container.AnnounceUsedSpaceRequest_Body_Announcement

	if a != nil {
		m = new(container.AnnounceUsedSpaceRequest_Body_Announcement)

		m.SetContainerId(a.cid.ToGRPCMessage().(*refsGRPC.ContainerID))
		m.SetEpoch(a.epoch)
		m.SetUsedSpace(a.usedSpace)
	}

	return m
}

func (a *UsedSpaceAnnouncement) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.AnnounceUsedSpaceRequest_Body_Announcement)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	cid := v.GetContainerId()
	if cid == nil {
		a.cid = nil
	} else {
		if a.cid == nil {
			a.cid = new(refs.ContainerID)
		}

		err = a.cid.FromGRPCMessage(cid)
		if err != nil {
			return err
		}
	}

	a.epoch = v.GetEpoch()
	a.usedSpace = v.GetUsedSpace()

	return nil
}

func UsedSpaceAnnouncementsToGRPCMessage(
	ids []*UsedSpaceAnnouncement,
) (res []*container.AnnounceUsedSpaceRequest_Body_Announcement) {
	if ids != nil {
		res = make([]*container.AnnounceUsedSpaceRequest_Body_Announcement, 0, len(ids))

		for i := range ids {
			res = append(res, ids[i].ToGRPCMessage().(*container.AnnounceUsedSpaceRequest_Body_Announcement))
		}
	}

	return
}

func UsedSpaceAnnouncementssFromGRPCMessage(
	asV2 []*container.AnnounceUsedSpaceRequest_Body_Announcement,
) (res []*UsedSpaceAnnouncement, err error) {
	if asV2 != nil {
		res = make([]*UsedSpaceAnnouncement, 0, len(asV2))

		for i := range asV2 {
			var a *UsedSpaceAnnouncement

			if asV2[i] != nil {
				a = new(UsedSpaceAnnouncement)

				err = a.FromGRPCMessage(asV2[i])
				if err != nil {
					return
				}
			}

			res = append(res, a)
		}
	}

	return
}

func (r *AnnounceUsedSpaceRequestBody) ToGRPCMessage() neofsgrpc.Message {
	var m *container.AnnounceUsedSpaceRequest_Body

	if r != nil {
		m = new(container.AnnounceUsedSpaceRequest_Body)

		m.SetAnnouncements(UsedSpaceAnnouncementsToGRPCMessage(r.announcements))
	}

	return m
}

func (r *AnnounceUsedSpaceRequestBody) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.AnnounceUsedSpaceRequest_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	r.announcements, err = UsedSpaceAnnouncementssFromGRPCMessage(v.GetAnnouncements())

	return err
}

func (r *AnnounceUsedSpaceRequest) ToGRPCMessage() neofsgrpc.Message {
	var m *container.AnnounceUsedSpaceRequest

	if r != nil {
		m = new(container.AnnounceUsedSpaceRequest)

		m.SetBody(r.body.ToGRPCMessage().(*container.AnnounceUsedSpaceRequest_Body))
		r.RequestHeaders.ToMessage(m)
	}

	return m
}

func (r *AnnounceUsedSpaceRequest) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.AnnounceUsedSpaceRequest)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		r.body = nil
	} else {
		if r.body == nil {
			r.body = new(AnnounceUsedSpaceRequestBody)
		}

		err = r.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	return r.RequestHeaders.FromMessage(v)
}

func (r *AnnounceUsedSpaceResponseBody) ToGRPCMessage() neofsgrpc.Message {
	var m *container.AnnounceUsedSpaceResponse_Body

	if r != nil {
		m = new(container.AnnounceUsedSpaceResponse_Body)
	}

	return m
}

func (r *AnnounceUsedSpaceResponseBody) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.AnnounceUsedSpaceResponse_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	return nil
}

func (r *AnnounceUsedSpaceResponse) ToGRPCMessage() neofsgrpc.Message {
	var m *container.AnnounceUsedSpaceResponse

	if r != nil {
		m = new(container.AnnounceUsedSpaceResponse)

		m.SetBody(r.body.ToGRPCMessage().(*container.AnnounceUsedSpaceResponse_Body))
		r.ResponseHeaders.ToMessage(m)
	}

	return m
}

func (r *AnnounceUsedSpaceResponse) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*container.AnnounceUsedSpaceResponse)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		r.body = nil
	} else {
		if r.body == nil {
			r.body = new(AnnounceUsedSpaceResponseBody)
		}

		err = r.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	return r.ResponseHeaders.FromMessage(v)
}
