package object

import (
	refs "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
	service "github.com/nspcc-dev/neofs-api-go/v2/service/grpc"
)

// SetKey sets key to the object attribute.
func (m *Header_Attribute) SetKey(v string) {
	if m != nil {
		m.Key = v
	}
}

// SetValue sets value of the object attribute.
func (m *Header_Attribute) SetValue(v string) {
	if m != nil {
		m.Value = v
	}
}

// SetParent sets identifier of the parent object.
func (m *Header_Split) SetParent(v *refs.ObjectID) {
	if m != nil {
		m.Parent = v
	}
}

// SetPrevious sets identifier of the previous object in split-chain.
func (m *Header_Split) SetPrevious(v *refs.ObjectID) {
	if m != nil {
		m.Previous = v
	}
}

// SetParentSignature sets signature of the parent object header.
func (m *Header_Split) SetParentSignature(v *refs.Signature) {
	if m != nil {
		m.ParentSignature = v
	}
}

// SetParentHeader sets parent header structure.
func (m *Header_Split) SetParentHeader(v *Header) {
	if m != nil {
		m.ParentHeader = v
	}
}

// SetChildren sets list of the identifiers of the child objects.
func (m *Header_Split) SetChildren(v []*refs.ObjectID) {
	if m != nil {
		m.Children = v
	}
}

// SetContainerId sets identifier of the container.
func (m *Header) SetContainerId(v *refs.ContainerID) {
	if m != nil {
		m.ContainerId = v
	}
}

// SetOwnerId sets identifier of the object owner.
func (m *Header) SetOwnerId(v *refs.OwnerID) {
	if m != nil {
		m.OwnerId = v
	}
}

// SetCreationEpoch sets creation epoch number.
func (m *Header) SetCreationEpoch(v uint64) {
	if m != nil {
		m.CreationEpoch = v
	}
}

// SetVersion sets version of the object format.
func (m *Header) SetVersion(v *refs.Version) {
	if m != nil {
		m.Version = v
	}
}

// SetPayloadLength sets length of the object payload.
func (m *Header) SetPayloadLength(v uint64) {
	if m != nil {
		m.PayloadLength = v
	}
}

// SetPayloadHash sets hash of the object payload.
func (m *Header) SetPayloadHash(v *refs.Checksum) {
	if m != nil {
		m.PayloadHash = v
	}
}

// SetObjectType sets type of the object.
func (m *Header) SetObjectType(v ObjectType) {
	if m != nil {
		m.ObjectType = v
	}
}

// SetHomomorphicHash sets homomorphic hash of the object payload.
func (m *Header) SetHomomorphicHash(v *refs.Checksum) {
	if m != nil {
		m.HomomorphicHash = v
	}
}

// SetSessionToken sets session token.
func (m *Header) SetSessionToken(v *service.SessionToken) {
	if m != nil {
		m.SessionToken = v
	}
}

// SetAttributes sets list of the object attributes.
func (m *Header) SetAttributes(v []*Header_Attribute) {
	if m != nil {
		m.Attributes = v
	}
}

// SetSplit sets split header.
func (m *Header) SetSplit(v *Header_Split) {
	if m != nil {
		m.Split = v
	}
}

// SetObjectId sets identifier of the object.
func (m *Object) SetObjectId(v *refs.ObjectID) {
	if m != nil {
		m.ObjectId = v
	}
}

// SetSignature sets signature of the object identifier.
func (m *Object) SetSignature(v *refs.Signature) {
	if m != nil {
		m.Signature = v
	}
}

// SetHeader sets header of the object.
func (m *Object) SetHeader(v *Header) {
	if m != nil {
		m.Header = v
	}
}

// SetPayload sets payload bytes of the object.
func (m *Object) SetPayload(v []byte) {
	if m != nil {
		m.Payload = v
	}
}

// SetVersion sets version of the object.
func (m *ShortHeader) SetVersion(v *refs.Version) {
	if m != nil {
		m.Version = v
	}
}

// SetCreationEpoch sets creation epoch number.
func (m *ShortHeader) SetCreationEpoch(v uint64) {
	if m != nil {
		m.CreationEpoch = v
	}
}

// SetOwnerId sets identifier of the object owner.
func (m *ShortHeader) SetOwnerId(v *refs.OwnerID) {
	if m != nil {
		m.OwnerId = v
	}
}

// SetObjectType sets type of the object.
func (m *ShortHeader) SetObjectType(v ObjectType) {
	if m != nil {
		m.ObjectType = v
	}
}

// SetPayloadLength sets length of the object payload.
func (m *ShortHeader) SetPayloadLength(v uint64) {
	if m != nil {
		m.PayloadLength = v
	}
}
