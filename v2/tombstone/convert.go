package tombstone

import (
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	refsGRPC "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
	tombstone "github.com/nspcc-dev/neofs-api-go/v2/tombstone/grpc"
)

// TombstoneToGRPCMessage converts unified tombstone message into gRPC message.
func TombstoneToGRPCMessage(t *Tombstone) *tombstone.Tombstone {
	if t == nil {
		return nil
	}

	m := new(tombstone.Tombstone)

	m.SetExpirationEpoch(t.GetExpirationEpoch())
	m.SetSplitId(t.GetSplitID())

	members := t.GetMembers()
	memberMsg := make([]*refsGRPC.ObjectID, 0, len(members))

	for i := range members {
		memberMsg = append(memberMsg, refs.ObjectIDToGRPCMessage(members[i]))
	}

	m.SetMembers(memberMsg)

	return m
}

// TombstoneFromGRPCMessage converts gRPC message into unified tombstone message.
func TombstoneFromGRPCMessage(m *tombstone.Tombstone) *Tombstone {
	if m == nil {
		return nil
	}

	t := new(Tombstone)

	t.SetExpirationEpoch(m.GetExpirationEpoch())
	t.SetSplitID(m.GetSplitId())

	memberMsg := m.GetMembers()
	members := make([]*refs.ObjectID, 0, len(memberMsg))

	for i := range memberMsg {
		members = append(members, refs.ObjectIDFromGRPCMessage(memberMsg[i]))
	}

	t.SetMembers(members)

	return t
}
