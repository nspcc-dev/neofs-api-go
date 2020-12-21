package audit

import (
	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	"github.com/nspcc-dev/neofs-api-go/pkg/object"
	"github.com/nspcc-dev/neofs-api-go/v2/audit"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

// Result represents v2-compatible data audit result.
type Result audit.DataAuditResult

// NewFromV2 wraps v2 DataAuditResult message to Result.
func NewResultFromV2(aV2 *audit.DataAuditResult) *Result {
	return (*Result)(aV2)
}

// New creates and initializes blank Result.
func NewResult() *Result {
	r := NewResultFromV2(new(audit.DataAuditResult))
	r.SetVersion(pkg.SDKVersion())

	return r
}

// ToV2 converts Result to v2 DataAuditResult message.
func (r *Result) ToV2() *audit.DataAuditResult {
	return (*audit.DataAuditResult)(r)
}

// Marshal marshals Result into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (r *Result) Marshal(b ...[]byte) ([]byte, error) {
	var buf []byte
	if len(b) > 0 {
		buf = b[0]
	}

	return (*audit.DataAuditResult)(r).
		StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of Result.
func (r *Result) Unmarshal(data []byte) error {
	return (*audit.DataAuditResult)(r).
		Unmarshal(data)
}

// MarshalJSON encodes Result to protobuf JSON format.
func (r *Result) MarshalJSON() ([]byte, error) {
	return (*audit.DataAuditResult)(r).
		MarshalJSON()
}

// UnmarshalJSON decodes Result from protobuf JSON format.
func (r *Result) UnmarshalJSON(data []byte) error {
	return (*audit.DataAuditResult)(r).
		UnmarshalJSON(data)
}

// Version returns Data Audit structure version.
func (r *Result) Version() *pkg.Version {
	return pkg.NewVersionFromV2(
		(*audit.DataAuditResult)(r).GetVersion(),
	)
}

// SetVersion sets Data Audit structure version.
func (r *Result) SetVersion(v *pkg.Version) {
	(*audit.DataAuditResult)(r).
		SetVersion(v.ToV2())
}

// AuditEpoch returns epoch number when the Data Audit was conducted.
func (r *Result) AuditEpoch() uint64 {
	return (*audit.DataAuditResult)(r).
		GetAuditEpoch()
}

// SetAuditEpoch sets epoch number when the Data Audit was conducted.
func (r *Result) SetAuditEpoch(epoch uint64) {
	(*audit.DataAuditResult)(r).
		SetAuditEpoch(epoch)
}

// ContainerID returns container under audit.
func (r *Result) ContainerID() *container.ID {
	return container.NewIDFromV2(
		(*audit.DataAuditResult)(r).
			GetContainerID(),
	)
}

// SetContainerID sets container under audit.
func (r *Result) SetContainerID(id *container.ID) {
	(*audit.DataAuditResult)(r).
		SetContainerID(id.ToV2())
}

// PublicKey returns public key of the auditing InnerRing node in a binary format.
func (r *Result) PublicKey() []byte {
	return (*audit.DataAuditResult)(r).
		GetPublicKey()
}

// SetPublicKey sets public key of the auditing InnerRing node in a binary format.
func (r *Result) SetPublicKey(key []byte) {
	(*audit.DataAuditResult)(r).
		SetPublicKey(key)
}

// Complete returns completion state of audit result.
func (r *Result) Complete() bool {
	return (*audit.DataAuditResult)(r).
		GetComplete()
}

// SetComplete sets completion state of audit result.
func (r *Result) SetComplete(v bool) {
	(*audit.DataAuditResult)(r).
		SetComplete(v)
}

// PassSG returns list of Storage Groups that passed audit PoR stage.
func (r *Result) PassSG() []*object.ID {
	mV2 := (*audit.DataAuditResult)(r).
		GetPassSG()

	if mV2 == nil {
		return nil
	}

	m := make([]*object.ID, len(mV2))

	for i := range mV2 {
		m[i] = object.NewIDFromV2(mV2[i])
	}

	return m
}

// SetPassSG sets list of Storage Groups that passed audit PoR stage.
func (r *Result) SetPassSG(list []*object.ID) {
	mV2 := (*audit.DataAuditResult)(r).
		GetPassSG()

	if list == nil {
		mV2 = nil
	} else {
		ln := len(list)

		if cap(mV2) >= ln {
			mV2 = mV2[:0]
		} else {
			mV2 = make([]*refs.ObjectID, 0, ln)
		}

		for i := 0; i < ln; i++ {
			mV2 = append(mV2, list[i].ToV2())
		}
	}

	(*audit.DataAuditResult)(r).
		SetPassSG(mV2)
}

// FailSG returns list of Storage Groups that failed audit PoR stage.
func (r *Result) FailSG() []*object.ID {
	mV2 := (*audit.DataAuditResult)(r).
		GetFailSG()

	if mV2 == nil {
		return nil
	}

	m := make([]*object.ID, len(mV2))

	for i := range mV2 {
		m[i] = object.NewIDFromV2(mV2[i])
	}

	return m
}

// SetFailSG sets list of Storage Groups that failed audit PoR stage.
func (r *Result) SetFailSG(list []*object.ID) {
	mV2 := (*audit.DataAuditResult)(r).
		GetFailSG()

	if list == nil {
		mV2 = nil
	} else {
		ln := len(list)

		if cap(mV2) >= ln {
			mV2 = mV2[:0]
		} else {
			mV2 = make([]*refs.ObjectID, 0, ln)
		}

		for i := 0; i < ln; i++ {
			mV2 = append(mV2, list[i].ToV2())
		}
	}

	(*audit.DataAuditResult)(r).
		SetFailSG(mV2)
}

// Hit returns number of sampled objects under audit placed
// in an optimal way according to the containers placement policy
// when checking PoP.
func (r *Result) Hit() uint32 {
	return (*audit.DataAuditResult)(r).
		GetHit()
}

// SetHit sets number of sampled objects under audit placed
// in an optimal way according to the containers placement policy
// when checking PoP.
func (r *Result) SetHit(hit uint32) {
	(*audit.DataAuditResult)(r).
		SetHit(hit)
}

// Miss returns number of sampled objects under audit placed
// in suboptimal way according to the containers placement policy,
// but still at a satisfactory level when checking PoP.
func (r *Result) Miss() uint32 {
	return (*audit.DataAuditResult)(r).
		GetMiss()
}

// SetMiss sets number of sampled objects under audit placed
// in suboptimal way according to the containers placement policy,
// but still at a satisfactory level when checking PoP.
func (r *Result) SetMiss(miss uint32) {
	(*audit.DataAuditResult)(r).
		SetMiss(miss)
}

// Fail returns number of sampled objects under audit stored
// in a way not confirming placement policy or not found at all
// when checking PoP.
func (r *Result) Fail() uint32 {
	return (*audit.DataAuditResult)(r).
		GetFail()
}

// SetFail sets number of sampled objects under audit stored
// in a way not confirming placement policy or not found at all
// when checking PoP.
func (r *Result) SetFail(fail uint32) {
	(*audit.DataAuditResult)(r).
		SetFail(fail)
}

// PassNodes returns list of storage node public keys that
// passed at least one PDP.
func (r *Result) PassNodes() [][]byte {
	return (*audit.DataAuditResult)(r).
		GetPassNodes()
}

// SetPassNodes sets list of storage node public keys that
// passed at least one PDP.
func (r *Result) SetPassNodes(list [][]byte) {
	(*audit.DataAuditResult)(r).
		SetPassNodes(list)
}

// FailNodes returns list of storage node public keys that
// failed at least one PDP.
func (r *Result) FailNodes() [][]byte {
	return (*audit.DataAuditResult)(r).
		GetFailNodes()
}

// SetFailNodes sets list of storage node public keys that
// failed at least one PDP.
func (r *Result) SetFailNodes(list [][]byte) {
	(*audit.DataAuditResult)(r).
		SetFailNodes(list)
}
