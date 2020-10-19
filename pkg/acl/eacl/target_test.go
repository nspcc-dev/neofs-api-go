package eacl

import (
	"crypto/ecdsa"
	"testing"

	v2acl "github.com/nspcc-dev/neofs-api-go/v2/acl"
	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
)

func TestTarget(t *testing.T) {
	keys := []ecdsa.PublicKey{
		test.DecodeKey(1).PublicKey,
		test.DecodeKey(2).PublicKey,
	}

	target := &Target{
		role: RoleSystem,
		keys: keys,
	}

	v2 := target.ToV2()
	require.NotNil(t, v2)
	require.Equal(t, v2acl.RoleSystem, v2.GetRole())
	require.Len(t, v2.GetKeys(), len(keys))
	for i, key := range v2.GetKeys() {
		require.Equal(t, key, crypto.MarshalPublicKey(&keys[i]))
	}

	newTarget := NewTargetFromV2(v2)
	require.Equal(t, target, newTarget)

	t.Run("from nil v2 target", func(t *testing.T) {
		require.Equal(t, new(Target), NewTargetFromV2(nil))
	})
}
