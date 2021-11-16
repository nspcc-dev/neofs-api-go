package tombstone

import (
	refs "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
)

// SetExpirationEpoch sets number of tombstone expiration epoch.
func (x *Tombstone) SetExpirationEpoch(v uint64) {
	if x != nil {
		x.ExpirationEpoch = v
	}
}

// SetSplitId sets identifier of split object hierarchy.
func (x *Tombstone) SetSplitId(v []byte) {
	if x != nil {
		x.SplitId = v
	}
}

// SetMembers sets list of objects to be deleted.
func (x *Tombstone) SetMembers(v []*refs.ObjectID) {
	if x != nil {
		x.Members = v
	}
}
