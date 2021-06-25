package eacl

import (
	v2acl "github.com/nspcc-dev/neofs-api-go/v2/acl"
)

// Target is a group of request senders to match EACL. Defined by role enum
// and set of public keys.
//
// Target is compatible with v2 acl.EACLRecord.Target message.
type Target struct {
	role Role
	keys [][]byte
}

// BinaryKeys returns list of public keys to identify
// target subject in a binary format.
func (t *Target) BinaryKeys() [][]byte {
	return t.keys
}

// SetBinaryKeys sets list of binary public keys to identify
// target subject.
func (t *Target) SetBinaryKeys(keys [][]byte) {
	t.keys = keys
}

// SetRole sets target subject's role class.
func (t *Target) SetRole(r Role) {
	t.role = r
}

// Role returns target subject's role class.
func (t Target) Role() Role {
	return t.role
}

// ToV2 converts Target to v2 acl.EACLRecord.Target message.
//
// Nil Target converts to nil.
func (t *Target) ToV2() *v2acl.Target {
	if t == nil {
		return nil
	}

	target := new(v2acl.Target)

	target.SetRole(t.role.ToV2())
	target.SetKeys(t.keys)

	return target
}

// NewTarget creates, initializes and returns blank Target instance.
//
// Defaults:
//  - role: RoleUnknown;
//  - keys: nil.
func NewTarget() *Target {
	return NewTargetFromV2(new(v2acl.Target))
}

// NewTargetFromV2 converts v2 acl.EACLRecord.Target message to Target.
func NewTargetFromV2(target *v2acl.Target) *Target {
	t := new(Target)

	if target == nil {
		return t
	}

	t.role = RoleFromV2(target.GetRole())
	t.keys = target.GetKeys()

	return t
}

// Marshal marshals Target into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (t *Target) Marshal(b ...[]byte) ([]byte, error) {
	var buf []byte
	if len(b) > 0 {
		buf = b[0]
	}

	return t.ToV2().
		StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of Target.
func (t *Target) Unmarshal(data []byte) error {
	fV2 := new(v2acl.Target)
	if err := fV2.Unmarshal(data); err != nil {
		return err
	}

	*t = *NewTargetFromV2(fV2)

	return nil
}

// MarshalJSON encodes Target to protobuf JSON format.
func (t *Target) MarshalJSON() ([]byte, error) {
	return t.ToV2().
		MarshalJSON()
}

// UnmarshalJSON decodes Target from protobuf JSON format.
func (t *Target) UnmarshalJSON(data []byte) error {
	tV2 := new(v2acl.Target)
	if err := tV2.UnmarshalJSON(data); err != nil {
		return err
	}

	*t = *NewTargetFromV2(tV2)

	return nil
}
