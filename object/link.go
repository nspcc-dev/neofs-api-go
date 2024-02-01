package object

import (
	"errors"
	"fmt"

	link "github.com/nspcc-dev/neofs-api-go/v2/link/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	refsGRPC "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/rpc/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/rpc/message"
	"github.com/nspcc-dev/neofs-api-go/v2/util/proto"
)

// Link represents object Link message from NeoFS API V2 protocol.
type Link struct {
	children []MeasuredObject
}

// NumberOfChildren returns length of children list.
func (x *Link) NumberOfChildren() int {
	if x != nil {
		return len(x.children)
	}

	return 0
}

// IterateChildren passes members of the link list to f.
func (x *Link) IterateChildren(f func(MeasuredObject)) {
	if x != nil {
		for i := range x.children {
			f(x.children[i])
		}
	}
}

// SetChildren sets a list of the child objects.
// Arg must not be mutated for the duration of the Link.
func (x *Link) SetChildren(chs []MeasuredObject) {
	x.children = chs
}

const (
	_ = iota
	linkFNumChildren
)

const (
	_ = iota
	measuredObjectFNumID
	measuredObjectFNumSize
)

// StableMarshal encodes the Link into Protocol Buffers binary format
// with direct field order.
func (x *Link) StableMarshal(buf []byte) []byte {
	if x == nil || len(x.children) == 0 {
		return []byte{}
	}

	if buf == nil {
		buf = make([]byte, x.StableSize())
	}

	var offset int

	for i := range x.children {
		offset += proto.NestedStructureMarshal(linkFNumChildren, buf[offset:], &x.children[i])
	}

	return buf
}

// StableSize size of the buffer required to write the Link in Protocol Buffers
// binary format.
func (x *Link) StableSize() int {
	var size int

	if x != nil {
		for i := range x.children {
			size += proto.NestedStructureSize(linkFNumChildren, &x.children[i])
		}
	}

	return size
}

// Unmarshal decodes the Link from its Protocol Buffers binary format.
func (x *Link) Unmarshal(data []byte) error {
	return message.Unmarshal(x, data, new(link.Link))
}

func (x *Link) ToGRPCMessage() grpc.Message {
	var m *link.Link

	if x != nil {
		m = new(link.Link)

		var children []*link.Link_MeasuredObject

		if x.children != nil {
			children = make([]*link.Link_MeasuredObject, len(x.children))

			for i := range x.children {
				children[i] = x.children[i].ToGRPCMessage().(*link.Link_MeasuredObject)
			}
		}

		m.Children = children
	}

	return m
}

func (x *Link) FromGRPCMessage(m grpc.Message) error {
	v, ok := m.(*link.Link)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	children := v.GetChildren()
	if children == nil {
		x.children = nil
	} else {
		x.children = make([]MeasuredObject, len(children))
		var err error

		for i := range x.children {
			err = x.children[i].FromGRPCMessage(children[i])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// MeasuredObject groups object descriptor and object's length.
type MeasuredObject struct {
	ID   refs.ObjectID
	Size uint32
}

// StableSize size of the buffer required to write the MeasuredObject in Protocol Buffers
// binary format.
func (x *MeasuredObject) StableSize() int {
	var size int

	size += proto.NestedStructureSize(measuredObjectFNumID, &x.ID)
	size += proto.UInt32Size(measuredObjectFNumSize, x.Size)

	return size
}

// StableMarshal encodes the MeasuredObject into Protocol Buffers binary format
// with direct field order.
func (x *MeasuredObject) StableMarshal(buf []byte) []byte {
	if buf == nil {
		buf = make([]byte, x.StableSize())
	}

	var offset int

	offset += proto.NestedStructureMarshal(measuredObjectFNumID, buf[offset:], &x.ID)
	proto.UInt32Marshal(measuredObjectFNumSize, buf[offset:], x.Size)

	return buf
}

// Unmarshal decodes the Link from its Protocol Buffers binary format.
func (x *MeasuredObject) Unmarshal(data []byte) error {
	return message.Unmarshal(x, data, new(link.Link_MeasuredObject))
}

func (x *MeasuredObject) FromGRPCMessage(m grpc.Message) error {
	v, ok := m.(*link.Link_MeasuredObject)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	if v.Id != nil {
		err := x.ID.FromGRPCMessage(v.Id)
		if err != nil {
			return err
		}
	}

	x.Size = v.Size

	return nil
}

func (x *MeasuredObject) ToGRPCMessage() grpc.Message {
	m := new(link.Link_MeasuredObject)

	m.Id = x.ID.ToGRPCMessage().(*refsGRPC.ObjectID)
	m.Size = x.Size

	return m
}

// WriteLink writes Link to the Object as a payload content.
// The object must not be nil.
func WriteLink(obj *Object, link Link) {
	hdr := obj.GetHeader()
	if hdr == nil {
		hdr = new(Header)
		obj.SetHeader(hdr)
	}

	payload := link.StableMarshal(nil)
	obj.SetPayload(payload)
}

// ReadLink reads Link from the Object payload content.
func ReadLink(link *Link, obj Object) error {
	payload := obj.GetPayload()
	if len(payload) == 0 {
		return errors.New("empty payload")
	}

	err := link.Unmarshal(payload)
	if err != nil {
		return fmt.Errorf("decode link content from payload: %w", err)
	}

	return nil
}
