package eacl

import (
	"crypto/ecdsa"

	v2acl "github.com/nspcc-dev/neofs-api-go/v2/acl"
	crypto "github.com/nspcc-dev/neofs-crypto"
)

// Target is a group of request senders to match EACL. Defined by role enum
// and set of public keys.
//
// Target is compatible with v2 acl.EACLRecord.Target message.
type Target struct {
	role Role
	keys [][]byte
}

func ecdsaKeysToPtrs(keys []ecdsa.PublicKey) []*ecdsa.PublicKey {
	keysPtr := make([]*ecdsa.PublicKey, len(keys))

	for i := range keys {
		keysPtr[i] = &keys[i]
	}

	return keysPtr
}

// SetKeys sets list of ECDSA public keys to identify target subject.
//
// Deprecated: use SetTargetECDSAKeys instead.
func (t *Target) SetKeys(keys ...ecdsa.PublicKey) {
	SetTargetECDSAKeys(t, ecdsaKeysToPtrs(keys)...)
}

// Keys returns list of ECDSA public keys to identify target subject.
// If some key has a different format, it is ignored.
//
// Deprecated: use TargetECDSAKeys instead.
func (t *Target) Keys() []ecdsa.PublicKey {
	keysPtr := TargetECDSAKeys(t)
	keys := make([]ecdsa.PublicKey, 0, len(keysPtr))

	for i := range keysPtr {
		if keysPtr[i] != nil {
			keys = append(keys, *keysPtr[i])
		}
	}

	return keys
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

// SetTargetECDSAKeys converts ECDSA public keys to a binary
// format and stores them in Target.
func SetTargetECDSAKeys(t *Target, keys ...*ecdsa.PublicKey) {
	binKeys := t.BinaryKeys()
	ln := len(keys)

	if cap(binKeys) >= ln {
		binKeys = binKeys[:0]
	} else {
		binKeys = make([][]byte, 0, ln)
	}

	for i := 0; i < ln; i++ {
		binKeys = append(binKeys, crypto.MarshalPublicKey(keys[i]))
	}

	t.SetBinaryKeys(binKeys)
}

// TargetECDSAKeys interprets binary public keys of Target
// as ECDSA public keys. If any key has a different format,
// the corresponding element will be nil.
func TargetECDSAKeys(t *Target) []*ecdsa.PublicKey {
	binKeys := t.BinaryKeys()
	ln := len(binKeys)

	keys := make([]*ecdsa.PublicKey, ln)

	for i := 0; i < ln; i++ {
		keys[i] = crypto.UnmarshalPublicKey(binKeys[i])
	}

	return keys
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
func (t *Target) ToV2() *v2acl.Target {
	target := new(v2acl.Target)

	target.SetRole(t.role.ToV2())
	target.SetKeys(t.keys)

	return target
}

// NewTarget creates, initializes and returns blank Target instance.
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
