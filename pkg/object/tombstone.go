package object

import (
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/tombstone"
)

// Tombstone represents v2-compatible tombstone structure.
type Tombstone tombstone.Tombstone

// NewTombstoneFromV2 wraps v2 Tombstone message to Tombstone.
//
// Nil tombstone.Tombstone converts to nil.
func NewTombstoneFromV2(tV2 *tombstone.Tombstone) *Tombstone {
	return (*Tombstone)(tV2)
}

// NewTombstone creates and initializes blank Tombstone.
//
// Defaults:
//  - exp: 0;
//  - splitID: nil;
//  - members: nil.
func NewTombstone() *Tombstone {
	return NewTombstoneFromV2(new(tombstone.Tombstone))
}

// ToV2 converts Tombstone to v2 Tombstone message.
//
// Nil Tombstone converts to nil.
func (ts *Tombstone) ToV2() *tombstone.Tombstone {
	return (*tombstone.Tombstone)(ts)
}

// ExpirationEpoch return number of tombstone expiration epoch.
func (t *Tombstone) ExpirationEpoch() uint64 {
	return (*tombstone.Tombstone)(t).
		GetExpirationEpoch()
}

// SetExpirationEpoch sets number of tombstone expiration epoch.
func (t *Tombstone) SetExpirationEpoch(v uint64) {
	(*tombstone.Tombstone)(t).
		SetExpirationEpoch(v)
}

// SplitID returns identifier of object split hierarchy.
func (t *Tombstone) SplitID() *SplitID {
	return NewSplitIDFromV2(
		(*tombstone.Tombstone)(t).
			GetSplitID(),
	)
}

// SetSplitID sets identifier of object split hierarchy.
func (t *Tombstone) SetSplitID(v *SplitID) {
	(*tombstone.Tombstone)(t).
		SetSplitID(v.ToV2())
}

// Members returns list of objects to be deleted.
func (t *Tombstone) Members() []*ID {
	msV2 := (*tombstone.Tombstone)(t).
		GetMembers()

	if msV2 == nil {
		return nil
	}

	ms := make([]*ID, 0, len(msV2))

	for i := range msV2 {
		ms = append(ms, NewIDFromV2(msV2[i]))
	}

	return ms
}

// SetMembers sets list of objects to be deleted.
func (t *Tombstone) SetMembers(v []*ID) {
	var ms []*refs.ObjectID

	if v != nil {
		ms = (*tombstone.Tombstone)(t).
			GetMembers()

		if ln := len(v); cap(ms) >= ln {
			ms = ms[:0]
		} else {
			ms = make([]*refs.ObjectID, 0, ln)
		}

		for i := range v {
			ms = append(ms, v[i].ToV2())
		}
	}

	(*tombstone.Tombstone)(t).
		SetMembers(ms)
}

// Marshal marshals Tombstone into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (t *Tombstone) Marshal(b ...[]byte) ([]byte, error) {
	var buf []byte
	if len(b) > 0 {
		buf = b[0]
	}

	return (*tombstone.Tombstone)(t).
		StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of Tombstone.
func (t *Tombstone) Unmarshal(data []byte) error {
	return (*tombstone.Tombstone)(t).
		Unmarshal(data)
}

// MarshalJSON encodes Tombstone to protobuf JSON format.
func (t *Tombstone) MarshalJSON() ([]byte, error) {
	return (*tombstone.Tombstone)(t).
		MarshalJSON()
}

// UnmarshalJSON decodes Tombstone from protobuf JSON format.
func (t *Tombstone) UnmarshalJSON(data []byte) error {
	return (*tombstone.Tombstone)(t).
		UnmarshalJSON(data)
}
