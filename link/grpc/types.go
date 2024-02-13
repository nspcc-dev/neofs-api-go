package link

import grpc "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"

// SetId sets object ID.
func (x *Link_MeasuredObject) SetId(v *grpc.ObjectID) {
	x.Id = v
}

// SetSize sets object size.
func (x *Link_MeasuredObject) SetSize(v uint32) {
	x.Size = v
}

// SetChildren sets object's children.
func (x *Link) SetChildren(v []*Link_MeasuredObject) {
	x.Children = v
}
