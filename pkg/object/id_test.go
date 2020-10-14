package object

import (
	"crypto/rand"
	"crypto/sha256"
	"strconv"
	"testing"

	"github.com/mr-tron/base58"
	"github.com/stretchr/testify/require"
)

func TestIDV2(t *testing.T) {
	id := NewID()

	checksum := [sha256.Size]byte{}

	_, err := rand.Read(checksum[:])
	require.NoError(t, err)

	id.SetSHA256(checksum)

	idV2 := id.ToV2()

	require.Equal(t, checksum[:], idV2.GetValue())
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
				oid := NewID()

				require.NoError(t, oid.Parse(str))
				require.Equal(t, cs[:], oid.ToV2().GetValue())
			})
		}
	})

	t.Run("should failure on parse", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			j := i
			t.Run(strconv.Itoa(j), func(t *testing.T) {
				cs := []byte{1, 2, 3, 4, 5, byte(j)}
				str := base58.Encode(cs)
				oid := NewID()

				require.EqualError(t, oid.Parse(str), ErrBadID.Error())
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
				oid := NewID()

				require.NoError(t, oid.Parse(str))
				require.Equal(t, str, oid.String())
			})
		}
	})
}
