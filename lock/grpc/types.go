package lock

import refs "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"

// SetMembers sets `members` field.
func (x *Lock) SetMembers(ids []*refs.ObjectID) {
	x.Members = ids
}
