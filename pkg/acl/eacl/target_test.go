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
	keys := []*ecdsa.PublicKey{
		&test.DecodeKey(1).PublicKey,
		&test.DecodeKey(2).PublicKey,
	}

	target := NewTarget()
	target.SetRole(RoleSystem)
	SetTargetECDSAKeys(target, keys...)

	v2 := target.ToV2()
	require.NotNil(t, v2)
	require.Equal(t, v2acl.RoleSystem, v2.GetRole())
	require.Len(t, v2.GetKeys(), len(keys))
	for i, key := range v2.GetKeys() {
		require.Equal(t, key, crypto.MarshalPublicKey(keys[i]))
	}

	newTarget := NewTargetFromV2(v2)
	require.Equal(t, target, newTarget)

	t.Run("from nil v2 target", func(t *testing.T) {
		require.Equal(t, new(Target), NewTargetFromV2(nil))
	})
}

func TestTargetEncoding(t *testing.T) {
	tar := NewTarget()
	tar.SetRole(RoleSystem)
	SetTargetECDSAKeys(tar, &test.DecodeKey(-1).PublicKey)

	t.Run("binary", func(t *testing.T) {
		data, err := tar.Marshal()
		require.NoError(t, err)

		tar2 := NewTarget()
		require.NoError(t, tar2.Unmarshal(data))

		require.Equal(t, tar, tar2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := tar.MarshalJSON()
		require.NoError(t, err)

		tar2 := NewTarget()
		require.NoError(t, tar2.UnmarshalJSON(data))

		require.Equal(t, tar, tar2)
	})
}

func TestTarget_ToV2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var x *Target

		require.Nil(t, x.ToV2())
	})
}
