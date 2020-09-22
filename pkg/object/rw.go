package object

import (
	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	"github.com/nspcc-dev/neofs-api-go/pkg/token"
	"github.com/nspcc-dev/neofs-api-go/v2/object"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

// wrapper over v2 Object that provides
// public getter and private setters.
type rwObject object.Object

// ToV2 converts Object to v2 Object message.
func (o *rwObject) ToV2() *object.Object {
	return (*object.Object)(o)
}

func (o *rwObject) setHeaderField(setter func(*object.Header)) {
	obj := (*object.Object)(o)
	h := obj.GetHeader()

	if h == nil {
		h = new(object.Header)
		obj.SetHeader(h)
	}

	setter(h)
}

func (o *rwObject) setSplitFields(setter func(*object.SplitHeader)) {
	o.setHeaderField(func(h *object.Header) {
		split := h.GetSplit()
		if split == nil {
			split = new(object.SplitHeader)
			h.SetSplit(split)
		}

		setter(split)
	})
}

// GetID returns object identifier.
func (o *rwObject) GetID() *ID {
	return NewIDFromV2(
		(*object.Object)(o).
			GetObjectID(),
	)
}

func (o *rwObject) setID(v *ID) {
	(*object.Object)(o).
		SetObjectID(v.ToV2())
}

// GetSignature returns signature of the object identifier.
func (o *rwObject) GetSignature() *pkg.Signature {
	return pkg.NewSignatureFromV2(
		(*object.Object)(o).
			GetSignature(),
	)
}

func (o *rwObject) setSignature(v *pkg.Signature) {
	(*object.Object)(o).
		SetSignature(v.ToV2())
}

// GetPayload returns payload bytes.
func (o *rwObject) GetPayload() []byte {
	return (*object.Object)(o).
		GetPayload()
}

func (o *rwObject) setPayload(v []byte) {
	(*object.Object)(o).
		SetPayload(v)
}

// GetVersion returns version of the object.
func (o *rwObject) GetVersion() *pkg.Version {
	return pkg.NewVersionFromV2(
		(*object.Object)(o).
			GetHeader().
			GetVersion(),
	)
}

func (o *rwObject) setVersion(v *pkg.Version) {
	o.setHeaderField(func(h *object.Header) {
		h.SetVersion(v.ToV2())
	})
}

// GetPayloadSize returns payload length of the object.
func (o *rwObject) GetPayloadSize() uint64 {
	return (*object.Object)(o).
		GetHeader().
		GetPayloadLength()
}

func (o *rwObject) setPayloadSize(v uint64) {
	o.setHeaderField(func(h *object.Header) {
		h.SetPayloadLength(v)
	})
}

// GetContainerID returns identifier of the related container.
func (o *rwObject) GetContainerID() *container.ID {
	return container.NewIDFromV2(
		(*object.Object)(o).
			GetHeader().
			GetContainerID(),
	)
}

func (o *rwObject) setContainerID(v *container.ID) {
	o.setHeaderField(func(h *object.Header) {
		h.SetContainerID(v.ToV2())
	})
}

// GetOwnerID returns identifier of the object owner.
func (o *rwObject) GetOwnerID() *owner.ID {
	return owner.NewIDFromV2(
		(*object.Object)(o).
			GetHeader().
			GetOwnerID(),
	)
}

func (o *rwObject) setOwnerID(v *owner.ID) {
	o.setHeaderField(func(h *object.Header) {
		h.SetOwnerID(v.ToV2())
	})
}

// GetCreationEpoch returns epoch number in which object was created.
func (o *rwObject) GetCreationEpoch() uint64 {
	return (*object.Object)(o).
		GetHeader().
		GetCreationEpoch()
}

func (o *rwObject) setCreationEpoch(v uint64) {
	o.setHeaderField(func(h *object.Header) {
		h.SetCreationEpoch(v)
	})
}

// GetPayloadChecksum returns checksum of the object payload.
func (o *rwObject) GetPayloadChecksum() *pkg.Checksum {
	return pkg.NewChecksumFromV2(
		(*object.Object)(o).
			GetHeader().
			GetPayloadHash(),
	)
}

func (o *rwObject) setPayloadChecksum(v *pkg.Checksum) {
	o.setHeaderField(func(h *object.Header) {
		h.SetPayloadHash(v.ToV2())
	})
}

// GetPayloadHomomorphicHash returns homomorphic hash of the object payload.
func (o *rwObject) GetPayloadHomomorphicHash() *pkg.Checksum {
	return pkg.NewChecksumFromV2(
		(*object.Object)(o).
			GetHeader().
			GetHomomorphicHash(),
	)
}

func (o *rwObject) setPayloadHomomorphicHash(v *pkg.Checksum) {
	o.setHeaderField(func(h *object.Header) {
		h.SetHomomorphicHash(v.ToV2())
	})
}

// GetAttributes returns object attributes.
func (o *rwObject) GetAttributes() []*Attribute {
	attrs := (*object.Object)(o).
		GetHeader().
		GetAttributes()

	res := make([]*Attribute, 0, len(attrs))

	for i := range attrs {
		res = append(res, NewAttributeFromV2(attrs[i]))
	}

	return res
}

func (o *rwObject) setAttributes(v ...*Attribute) {
	attrs := make([]*object.Attribute, 0, len(v))

	for i := range v {
		attrs = append(attrs, v[i].ToV2())
	}

	o.setHeaderField(func(h *object.Header) {
		h.SetAttributes(attrs)
	})
}

// GetPreviousID returns identifier of the previous sibling object.
func (o *rwObject) GetPreviousID() *ID {
	return NewIDFromV2(
		(*object.Object)(o).
			GetHeader().
			GetSplit().
			GetPrevious(),
	)
}

func (o *rwObject) setPreviousID(v *ID) {
	o.setSplitFields(func(split *object.SplitHeader) {
		split.SetPrevious(v.ToV2())
	})
}

// GetChildren return list of the identifiers of the child objects.
func (o *rwObject) GetChildren() []*ID {
	ids := (*object.Object)(o).
		GetHeader().
		GetSplit().
		GetChildren()

	res := make([]*ID, 0, len(ids))

	for i := range ids {
		res = append(res, NewIDFromV2(ids[i]))
	}

	return res
}

func (o *rwObject) setChildren(v ...*ID) {
	ids := make([]*refs.ObjectID, 0, len(v))

	for i := range v {
		ids = append(ids, v[i].ToV2())
	}

	o.setSplitFields(func(split *object.SplitHeader) {
		split.SetChildren(ids)
	})
}

// GetParent returns parent object w/o payload.
func (o *rwObject) GetParent() *Object {
	h := (*object.Object)(o).
		GetHeader().
		GetSplit()

	oV2 := new(object.Object)
	oV2.SetObjectID(h.GetParent())
	oV2.SetSignature(h.GetParentSignature())
	oV2.SetHeader(h.GetParentHeader())

	return NewFromV2(oV2)
}

func (o *rwObject) setParent(v *Object) {
	o.setSplitFields(func(split *object.SplitHeader) {
		split.SetParent((*object.Object)(v.rwObject).GetObjectID())
		split.SetParentSignature((*object.Object)(v.rwObject).GetSignature())
		split.SetParentHeader((*object.Object)(v.rwObject).GetHeader())
	})
}

// GetSessionToken returns token of the session
// within which object was created.
func (o *rwObject) GetSessionToken() *token.SessionToken {
	return token.NewSessionTokenFromV2(
		(*object.Object)(o).
			GetHeader().
			GetSessionToken(),
	)
}

func (o *rwObject) setSessionToken(v *token.SessionToken) {
	o.setHeaderField(func(h *object.Header) {
		h.SetSessionToken(v.ToV2())
	})
}
