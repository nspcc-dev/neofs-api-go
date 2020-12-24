package audit

import (
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

// DataAuditResult is a unified structure of
// DataAuditResult message from proto definition.
type DataAuditResult struct {
	auditEpoch uint64

	hit, miss, fail uint32

	cid *refs.ContainerID

	pubKey []byte

	passSG, failSG []*refs.ObjectID

	failNodes, passNodes [][]byte
}

// GetAuditEpoch returns epoch number when the Data Audit was conducted.
func (a *DataAuditResult) GetAuditEpoch() uint64 {
	if a != nil {
		return a.auditEpoch
	}

	return 0
}

// SetAuditEpoch sets epoch number when the Data Audit was conducted.
func (a *DataAuditResult) SetAuditEpoch(v uint64) {
	if a != nil {
		a.auditEpoch = v
	}
}

// GetContainerID returns container under audit.
func (a *DataAuditResult) GetContainerID() *refs.ContainerID {
	if a != nil {
		return a.cid
	}

	return nil
}

// SetContainerID sets container under audit.
func (a *DataAuditResult) SetContainerID(v *refs.ContainerID) {
	if a != nil {
		a.cid = v
	}
}

// GetPublicKey returns public key of the auditing InnerRing node in a binary format.
func (a *DataAuditResult) GetPublicKey() []byte {
	if a != nil {
		return a.pubKey
	}

	return nil
}

// SetPublicKey sets public key of the auditing InnerRing node in a binary format.
func (a *DataAuditResult) SetPublicKey(v []byte) {
	if a != nil {
		a.pubKey = v
	}
}

// GetPassSG returns list of Storage Groups that passed audit PoR stage.
func (a *DataAuditResult) GetPassSG() []*refs.ObjectID {
	if a != nil {
		return a.passSG
	}

	return nil
}

// SetPassSG sets list of Storage Groups that passed audit PoR stage.
func (a *DataAuditResult) SetPassSG(v []*refs.ObjectID) {
	if a != nil {
		a.passSG = v
	}
}

// GetFailSG returns list of Storage Groups that failed audit PoR stage.
func (a *DataAuditResult) GetFailSG() []*refs.ObjectID {
	if a != nil {
		return a.failSG
	}

	return nil
}

// SetFailSG sets list of Storage Groups that failed audit PoR stage.
func (a *DataAuditResult) SetFailSG(v []*refs.ObjectID) {
	if a != nil {
		a.failSG = v
	}
}

// GetHit returns number of sampled objects under audit placed
// in an optimal way according to the containers placement policy
// when checking PoP.
func (a *DataAuditResult) GetHit() uint32 {
	if a != nil {
		return a.hit
	}

	return 0
}

// SetHit sets number of sampled objects under audit placed
// in an optimal way according to the containers placement policy
// when checking PoP.
func (a *DataAuditResult) SetHit(v uint32) {
	if a != nil {
		a.hit = v
	}
}

// GetMiss returns number of sampled objects under audit placed
// in suboptimal way according to the containers placement policy,
// but still at a satisfactory level when checking PoP.
func (a *DataAuditResult) GetMiss() uint32 {
	if a != nil {
		return a.miss
	}

	return 0
}

// SetMiss sets number of sampled objects under audit placed
// in suboptimal way according to the containers placement policy,
// but still at a satisfactory level when checking PoP.
func (a *DataAuditResult) SetMiss(v uint32) {
	if a != nil {
		a.miss = v
	}
}

// GetFail returns number of sampled objects under audit stored
// in a way not confirming placement policy or not found at all
// when checking PoP.
func (a *DataAuditResult) GetFail() uint32 {
	if a != nil {
		return a.fail
	}

	return 0
}

// SetFail sets number of sampled objects under audit stored
// in a way not confirming placement policy or not found at all
// when checking PoP.
func (a *DataAuditResult) SetFail(v uint32) {
	if a != nil {
		a.fail = v
	}
}

// GetPassNodes returns list of storage node public keys that
// passed at least one PDP.
func (a *DataAuditResult) GetPassNodes() [][]byte {
	if a != nil {
		return a.passNodes
	}

	return nil
}

// SetPassNodes sets list of storage node public keys that
// passed at least one PDP.
func (a *DataAuditResult) SetPassNodes(v [][]byte) {
	if a != nil {
		a.passNodes = v
	}
}

// GetFailNodes returns list of storage node public keys that
// failed at least one PDP.
func (a *DataAuditResult) GetFailNodes() [][]byte {
	if a != nil {
		return a.failNodes
	}

	return nil
}

// SetFailNodes sets list of storage node public keys that
// failed at least one PDP.
func (a *DataAuditResult) SetFailNodes(v [][]byte) {
	if a != nil {
		a.failNodes = v
	}
}
