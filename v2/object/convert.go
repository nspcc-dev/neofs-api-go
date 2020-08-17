package object

import (
	object "github.com/nspcc-dev/neofs-api-go/v2/object/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	refsGRPC "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/service"
)

func TypeToGRPCField(t Type) object.ObjectType {
	return object.ObjectType(t)
}

func TypeFromGRPCField(t object.ObjectType) Type {
	return Type(t)
}

func MatchTypeToGRPCField(t MatchType) object.MatchType {
	return object.MatchType(t)
}

func MatchTypeFromGRPCField(t object.MatchType) MatchType {
	return MatchType(t)
}

func ShortHeaderToGRPCMessage(h *ShortHeader) *object.ShortHeader {
	if h == nil {
		return nil
	}

	m := new(object.ShortHeader)

	m.SetVersion(
		service.VersionToGRPCMessage(h.GetVersion()),
	)

	m.SetCreationEpoch(h.GetCreationEpoch())

	m.SetOwnerId(
		refs.OwnerIDToGRPCMessage(h.GetOwnerID()),
	)

	m.SetObjectType(
		TypeToGRPCField(h.GetObjectType()),
	)

	m.SetPayloadLength(h.GeyPayloadLength())

	return m
}

func ShortHeaderFromGRPCMessage(m *object.ShortHeader) *ShortHeader {
	if m == nil {
		return nil
	}

	h := new(ShortHeader)

	h.SetVersion(
		service.VersionFromGRPCMessage(m.GetVersion()),
	)

	h.SetCreationEpoch(m.GetCreationEpoch())

	h.SetOwnerID(
		refs.OwnerIDFromGRPCMessage(m.GetOwnerId()),
	)

	h.SetObjectType(
		TypeFromGRPCField(m.GetObjectType()),
	)

	h.SetPayloadLength(m.GetPayloadLength())

	return h
}

func AttributeToGRPCMessage(a *Attribute) *object.Header_Attribute {
	if a == nil {
		return nil
	}

	m := new(object.Header_Attribute)

	m.SetKey(a.GetKey())
	m.SetValue(a.GetValue())

	return m
}

func AttributeFromGRPCMessage(m *object.Header_Attribute) *Attribute {
	if m == nil {
		return nil
	}

	h := new(Attribute)

	h.SetKey(m.GetKey())
	h.SetValue(m.GetValue())

	return h
}

func SplitHeaderToGRPCMessage(h *SplitHeader) *object.Header_Split {
	if h == nil {
		return nil
	}

	m := new(object.Header_Split)

	m.SetParent(
		refs.ObjectIDToGRPCMessage(h.GetParent()),
	)

	m.SetPrevious(
		refs.ObjectIDToGRPCMessage(h.GetPrevious()),
	)

	m.SetParentSignature(
		service.SignatureToGRPCMessage(h.GetParentSignature()),
	)

	m.SetParentHeader(
		HeaderToGRPCMessage(h.GetParentHeader()),
	)

	children := h.GetChildren()
	childMsg := make([]*refsGRPC.ObjectID, 0, len(children))

	for i := range children {
		childMsg = append(childMsg, refs.ObjectIDToGRPCMessage(children[i]))
	}

	m.SetChildren(childMsg)

	return m
}

func SplitHeaderFromGRPCMessage(m *object.Header_Split) *SplitHeader {
	if m == nil {
		return nil
	}

	h := new(SplitHeader)

	h.SetParent(
		refs.ObjectIDFromGRPCMessage(m.GetParent()),
	)

	h.SetPrevious(
		refs.ObjectIDFromGRPCMessage(m.GetPrevious()),
	)

	h.SetParentSignature(
		service.SignatureFromGRPCMessage(m.GetParentSignature()),
	)

	h.SetParentHeader(
		HeaderFromGRPCMessage(m.GetParentHeader()),
	)

	childMsg := m.GetChildren()
	children := make([]*refs.ObjectID, 0, len(childMsg))

	for i := range childMsg {
		children = append(children, refs.ObjectIDFromGRPCMessage(childMsg[i]))
	}

	h.SetChildren(children)

	return h
}

func HeaderToGRPCMessage(h *Header) *object.Header {
	if h == nil {
		return nil
	}

	m := new(object.Header)

	m.SetVersion(
		service.VersionToGRPCMessage(h.GetVersion()),
	)

	m.SetContainerId(
		refs.ContainerIDToGRPCMessage(h.GetContainerID()),
	)

	m.SetOwnerId(
		refs.OwnerIDToGRPCMessage(h.GetOwnerID()),
	)

	m.SetCreationEpoch(h.GetCreationEpoch())

	m.SetPayloadLength(h.GetPayloadLength())

	m.SetPayloadHash(h.GetPayloadHash())

	m.SetHomomorphicHash(h.GetHomomorphicHash())

	m.SetObjectType(
		TypeToGRPCField(h.GetObjectType()),
	)

	m.SetSessionToken(
		service.SessionTokenToGRPCMessage(h.GetSessionToken()),
	)

	attr := h.GetAttributes()
	attrMsg := make([]*object.Header_Attribute, 0, len(attr))

	for i := range attr {
		attrMsg = append(attrMsg, AttributeToGRPCMessage(attr[i]))
	}

	m.SetAttributes(attrMsg)

	m.SetSplit(
		SplitHeaderToGRPCMessage(h.GetSplit()),
	)

	return m
}

func HeaderFromGRPCMessage(m *object.Header) *Header {
	if m == nil {
		return nil
	}

	h := new(Header)

	h.SetVersion(
		service.VersionFromGRPCMessage(m.GetVersion()),
	)

	h.SetContainerID(
		refs.ContainerIDFromGRPCMessage(m.GetContainerId()),
	)

	h.SetOwnerID(
		refs.OwnerIDFromGRPCMessage(m.GetOwnerId()),
	)

	h.SetCreationEpoch(m.GetCreationEpoch())

	h.SetPayloadLength(m.GetPayloadLength())

	h.SetPayloadHash(m.GetPayloadHash())

	h.SetHomomorphicHash(m.GetHomomorphicHash())

	h.SetObjectType(
		TypeFromGRPCField(m.GetObjectType()),
	)

	h.SetSessionToken(
		service.SessionTokenFromGRPCMessage(m.GetSessionToken()),
	)

	attrMsg := m.GetAttributes()
	attr := make([]*Attribute, 0, len(attrMsg))

	for i := range attrMsg {
		attr = append(attr, AttributeFromGRPCMessage(attrMsg[i]))
	}

	h.SetAttributes(attr)

	h.SetSplit(
		SplitHeaderFromGRPCMessage(m.GetSplit()),
	)

	return h
}

func ObjectToGRPCMessage(o *Object) *object.Object {
	if o == nil {
		return nil
	}

	m := new(object.Object)

	m.SetObjectId(
		refs.ObjectIDToGRPCMessage(o.GetObjectID()),
	)

	m.SetSignature(
		service.SignatureToGRPCMessage(o.GetSignature()),
	)

	m.SetHeader(
		HeaderToGRPCMessage(o.GetHeader()),
	)

	m.SetPayload(o.GetPayload())

	return m
}

func ObjectFromGRPCMessage(m *object.Object) *Object {
	if m == nil {
		return nil
	}

	o := new(Object)

	o.SetObjectID(
		refs.ObjectIDFromGRPCMessage(m.GetObjectId()),
	)

	o.SetSignature(
		service.SignatureFromGRPCMessage(m.GetSignature()),
	)

	o.SetHeader(
		HeaderFromGRPCMessage(m.GetHeader()),
	)

	o.SetPayload(m.GetPayload())

	return o
}
