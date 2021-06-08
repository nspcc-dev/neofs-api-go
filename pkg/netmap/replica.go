package netmap

import (
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
)

// Replica represents v2-compatible object replica descriptor.
type Replica netmap.Replica

// NewReplica creates and returns new Replica instance.
func NewReplica() *Replica {
	return NewReplicaFromV2(new(netmap.Replica))
}

// NewReplicaFromV2 converts v2 Replica to Replica.
//
// Nil netmap.Replica converts to nil.
func NewReplicaFromV2(f *netmap.Replica) *Replica {
	return (*Replica)(f)
}

// ToV2 converts Replica to v2 Replica.
//
// Nil Replica converts to nil.
func (r *Replica) ToV2() *netmap.Replica {
	return (*netmap.Replica)(r)
}

// Count returns number of object replicas.
func (r *Replica) Count() uint32 {
	return (*netmap.Replica)(r).
		GetCount()
}

// SetCount sets number of object replicas.
func (r *Replica) SetCount(c uint32) {
	(*netmap.Replica)(r).
		SetCount(c)
}

// Selector returns name of selector bucket to put replicas.
func (r *Replica) Selector() string {
	return (*netmap.Replica)(r).
		GetSelector()
}

// SetSelector sets name of selector bucket to put replicas.
func (r *Replica) SetSelector(s string) {
	(*netmap.Replica)(r).
		SetSelector(s)
}

// Marshal marshals Replica into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (r *Replica) Marshal(b ...[]byte) ([]byte, error) {
	var buf []byte
	if len(b) > 0 {
		buf = b[0]
	}

	return (*netmap.Replica)(r).
		StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of Replica.
func (r *Replica) Unmarshal(data []byte) error {
	return (*netmap.Replica)(r).
		Unmarshal(data)
}

// MarshalJSON encodes Replica to protobuf JSON format.
func (r *Replica) MarshalJSON() ([]byte, error) {
	return (*netmap.Replica)(r).
		MarshalJSON()
}

// UnmarshalJSON decodes Replica from protobuf JSON format.
func (r *Replica) UnmarshalJSON(data []byte) error {
	return (*netmap.Replica)(r).
		UnmarshalJSON(data)
}
