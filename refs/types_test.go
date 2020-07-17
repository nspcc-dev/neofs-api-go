package refs

import (
	"strings"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
)

func TestSGID(t *testing.T) {
	t.Run("check that marshal/unmarshal works like expected", func(t *testing.T) {
		var sgid1, sgid2 UUID

		sgid1, err := NewSGID()
		require.NoError(t, err)

		data, err := proto.Marshal(&sgid1)
		require.NoError(t, err)

		require.NoError(t, sgid2.Unmarshal(data))
		require.Equal(t, sgid1, sgid2)
	})

	t.Run("check that proto.Clone works like expected", func(t *testing.T) {
		var (
			sgid1 UUID
			sgid2 *UUID
		)

		sgid1, err := NewSGID()
		require.NoError(t, err)

		sgid2 = proto.Clone(&sgid1).(*SGID)
		require.Equal(t, sgid1, *sgid2)
	})
}

func TestUUID(t *testing.T) {
	t.Run("parse should work like expected", func(t *testing.T) {
		var u UUID

		id, err := uuid.NewRandom()
		require.NoError(t, err)

		require.NoError(t, u.Parse(id.String()))
		require.Equal(t, id.String(), u.String())
	})

	t.Run("check that marshal/unmarshal works like expected", func(t *testing.T) {
		var u1, u2 UUID

		u1 = UUID{0x8f, 0xe4, 0xeb, 0xa0, 0xb8, 0xfb, 0x49, 0x3b, 0xbb, 0x1d, 0x1d, 0x13, 0x6e, 0x69, 0xfc, 0xf7}

		data, err := proto.Marshal(&u1)
		require.NoError(t, err)

		require.NoError(t, u2.Unmarshal(data))
		require.Equal(t, u1, u2)
	})

	t.Run("check that marshal/unmarshal works like expected even for msg id", func(t *testing.T) {
		var u2 MessageID

		u1, err := NewMessageID()
		require.NoError(t, err)

		data, err := proto.Marshal(&u1)
		require.NoError(t, err)

		require.NoError(t, u2.Unmarshal(data))
		require.Equal(t, u1, u2)
	})
}

func TestOwnerID(t *testing.T) {
	t.Run("check that marshal/unmarshal works like expected", func(t *testing.T) {
		var u1, u2 OwnerID

		owner, err := NewOwnerID()
		require.NoError(t, err)
		require.True(t, owner.Empty())

		key := test.DecodeKey(0)

		u1, err = NewOwnerID(&key.PublicKey)
		require.NoError(t, err)
		data, err := proto.Marshal(&u1)
		require.NoError(t, err)

		require.NoError(t, u2.Unmarshal(data))
		require.Equal(t, u1, u2)
	})

	t.Run("check that proto.Clone works like expected", func(t *testing.T) {
		var u2 *OwnerID

		key := test.DecodeKey(0)

		u1, err := NewOwnerID(&key.PublicKey)
		require.NoError(t, err)

		u2 = proto.Clone(&u1).(*OwnerID)
		require.Equal(t, u1, *u2)
	})
}

func TestAddress(t *testing.T) {
	cid := CIDForBytes([]byte("test"))

	id, err := NewObjectID()
	require.NoError(t, err)

	expect := strings.Join([]string{
		cid.String(),
		id.String(),
	}, joinSeparator)

	require.NotPanics(t, func() {
		actual := (Address{
			ObjectID: id,
			CID:      cid,
		}).String()

		require.Equal(t, expect, actual)
	})

	var temp Address
	require.NoError(t, temp.Parse(expect))
	require.Equal(t, expect, temp.String())

	actual, err := ParseAddress(expect)
	require.NoError(t, err)
	require.Equal(t, expect, actual.String())

	addr := proto.Clone(actual).(*Address)
	require.Equal(t, actual, addr)
	require.Equal(t, expect, addr.String())
}
