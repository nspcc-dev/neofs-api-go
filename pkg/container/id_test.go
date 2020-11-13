package container

import (
	"crypto/rand"
	"crypto/sha256"
	"strconv"
	"testing"

	"github.com/mr-tron/base58"
	"github.com/stretchr/testify/require"
)

func TestIDV2_0(t *testing.T) {
	cid := NewID()

	checksum := [sha256.Size]byte{}

	_, err := rand.Read(checksum[:])
	require.NoError(t, err)

	cid.SetSHA256(checksum)

	cidV2 := cid.ToV2()

	require.Equal(t, checksum[:], cidV2.GetValue())
}

func randSHA256Checksum(t *testing.T) (cs [sha256.Size]byte) {
	_, err := rand.Read(cs[:])
	require.NoError(t, err)

	return
}

func TestID_Equal(t *testing.T) {
	cs := randSHA256Checksum(t)

	id1 := NewID()
	id1.SetSHA256(cs)

	id2 := NewID()
	id2.SetSHA256(cs)

	id3 := NewID()
	id3.SetSHA256(randSHA256Checksum(t))

	require.True(t, id1.Equal(id2))
	require.False(t, id1.Equal(id3))
}

func TestID_Parse(t *testing.T) {
	t.Run("should parse successful", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			t.Run(strconv.Itoa(i), func(t *testing.T) {
				cs := randSHA256Checksum(t)
				str := base58.Encode(cs[:])
				cid := NewID()

				require.NoError(t, cid.Parse(str))
				require.Equal(t, cs[:], cid.ToV2().GetValue())
			})
		}
	})

	t.Run("should failure on parse", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			j := i
			t.Run(strconv.Itoa(j), func(t *testing.T) {
				cs := []byte{1, 2, 3, 4, 5, byte(j)}
				str := base58.Encode(cs)
				cid := NewID()

				require.EqualError(t, cid.Parse(str), ErrBadID.Error())
			})
		}
	})
}

func TestID_String(t *testing.T) {
	t.Run("should be equal", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			t.Run(strconv.Itoa(i), func(t *testing.T) {
				cs := randSHA256Checksum(t)
				str := base58.Encode(cs[:])
				cid := NewID()

				require.NoError(t, cid.Parse(str))
				require.Equal(t, str, cid.String())
			})
		}
	})
}

func TestContainerIDEncoding(t *testing.T) {
	id := NewID()
	id.SetSHA256(randSHA256Checksum(t))

	t.Run("binary", func(t *testing.T) {
		data, err := id.Marshal()
		require.NoError(t, err)

		id2 := NewID()
		require.NoError(t, id2.Unmarshal(data))

		require.Equal(t, id, id2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := id.MarshalJSON()
		require.NoError(t, err)

		a2 := NewID()
		require.NoError(t, a2.UnmarshalJSON(data))

		require.Equal(t, id, a2)
	})
}
