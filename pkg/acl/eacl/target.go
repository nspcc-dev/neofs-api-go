package eacl

import (
	"crypto/ecdsa"

	v2acl "github.com/nspcc-dev/neofs-api-go/v2/acl"
	crypto "github.com/nspcc-dev/neofs-crypto"
)

// Target is a group of request senders to match EACL. Defined by role enum
// and set of public keys.
type Target struct {
	role Role
	keys []ecdsa.PublicKey
}

func (t *Target) SetKeys(keys ...ecdsa.PublicKey) {
	t.keys = keys
}

func (t Target) Keys() []ecdsa.PublicKey {
	return t.keys
}

func (t *Target) SetRole(r Role) {
	t.role = r
}

func (t Target) Role() Role {
	return t.role
}

func (t *Target) ToV2() *v2acl.Target {
	keys := make([][]byte, 0, len(t.keys))
	for i := range t.keys {
		key := crypto.MarshalPublicKey(&t.keys[i])
		keys = append(keys, key)
	}

	target := new(v2acl.Target)

	target.SetRole(t.role.ToV2())
	target.SetKeys(keys)

	return target
}

func NewTarget() *Target {
	return NewTargetFromV2(new(v2acl.Target))
}

func NewTargetFromV2(target *v2acl.Target) *Target {
	t := new(Target)

	if target == nil {
		return t
	}

	t.role = RoleFromV2(target.GetRole())
	v2keys := target.GetKeys()
	t.keys = make([]ecdsa.PublicKey, 0, len(v2keys))
	for i := range v2keys {
		key := crypto.UnmarshalPublicKey(v2keys[i])
		t.keys = append(t.keys, *key)
	}

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
