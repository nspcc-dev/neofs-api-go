package storagegroup_test

import (
	"crypto/rand"
	"crypto/sha256"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/object"
	objecttest "github.com/nspcc-dev/neofs-api-go/pkg/object/test"
	"github.com/nspcc-dev/neofs-api-go/pkg/storagegroup"
	"github.com/stretchr/testify/require"
)

func testSHA256() (cs [sha256.Size]byte) {
	_, _ = rand.Read(cs[:])
	return
}

func testChecksum() *pkg.Checksum {
	h := pkg.NewChecksum()
	h.SetSHA256(testSHA256())

	return h
}

func TestStorageGroup(t *testing.T) {
	sg := storagegroup.New()

	sz := uint64(13)
	sg.SetValidationDataSize(sz)
	require.Equal(t, sz, sg.ValidationDataSize())

	cs := testChecksum()
	sg.SetValidationDataHash(cs)
	require.Equal(t, cs, sg.ValidationDataHash())

	exp := uint64(33)
	sg.SetExpirationEpoch(exp)
	require.Equal(t, exp, sg.ExpirationEpoch())

	members := []*object.ID{objecttest.GenerateID(), objecttest.GenerateID()}
	sg.SetMembers(members)
	require.Equal(t, members, sg.Members())
}

func TestStorageGroupEncoding(t *testing.T) {
	sg := storagegroup.New()
	sg.SetValidationDataSize(13)
	sg.SetValidationDataHash(testChecksum())
	sg.SetExpirationEpoch(33)
	sg.SetMembers([]*object.ID{objecttest.GenerateID(), objecttest.GenerateID()})

	t.Run("binary", func(t *testing.T) {
		data, err := sg.Marshal()
		require.NoError(t, err)

		sg2 := storagegroup.New()
		require.NoError(t, sg2.Unmarshal(data))

		require.Equal(t, sg, sg2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := sg.MarshalJSON()
		require.NoError(t, err)

		sg2 := storagegroup.New()
		require.NoError(t, sg2.UnmarshalJSON(data))

		require.Equal(t, sg, sg2)
	})
}
