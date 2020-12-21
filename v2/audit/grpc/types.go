package audit

import (
	refs "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
)

// SetVersion is a Version field setter.
func (x *DataAuditResult) SetVersion(v *refs.Version) {
	if x != nil {
		x.Version = v
	}
}

// SetAuditEpoch is an AuditEpoch field setter.
func (x *DataAuditResult) SetAuditEpoch(v uint64) {
	if x != nil {
		x.AuditEpoch = v
	}
}

// SetContainerId is a ContainerId field setter.
func (x *DataAuditResult) SetContainerId(v *refs.ContainerID) {
	if x != nil {
		x.ContainerId = v
	}
}

// SetPublicKey is a PublicKey field setter.
func (x *DataAuditResult) SetPublicKey(v []byte) {
	if x != nil {
		x.PublicKey = v
	}
}

// SetComplete is a Complete field setter.
func (x *DataAuditResult) SetComplete(v bool) {
	if x != nil {
		x.Complete = v
	}
}

// SetPassSg is a PassSg field setter.
func (x *DataAuditResult) SetPassSg(v []*refs.ObjectID) {
	if x != nil {
		x.PassSg = v
	}
}

// SetFailSg is a FailSg field setter.
func (x *DataAuditResult) SetFailSg(v []*refs.ObjectID) {
	if x != nil {
		x.FailSg = v
	}
}

// SetHit is a Hit field setter.
func (x *DataAuditResult) SetHit(v uint32) {
	if x != nil {
		x.Hit = v
	}
}

// SetMiss is a Miss field setter.
func (x *DataAuditResult) SetMiss(v uint32) {
	if x != nil {
		x.Miss = v
	}
}

// SetFail is a Fail field setter.
func (x *DataAuditResult) SetFail(v uint32) {
	if x != nil {
		x.Fail = v
	}
}

// SetPassNodes is a PassNodes field setter.
func (x *DataAuditResult) SetPassNodes(v [][]byte) {
	if x != nil {
		x.PassNodes = v
	}
}

// SetFailNodes is a FailNodes field setter.
func (x *DataAuditResult) SetFailNodes(v [][]byte) {
	if x != nil {
		x.FailNodes = v
	}
}
