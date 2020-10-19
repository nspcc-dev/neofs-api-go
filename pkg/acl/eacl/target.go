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

func (t Target) Keys() []ecdsa.PublicKey {
	return t.keys
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
