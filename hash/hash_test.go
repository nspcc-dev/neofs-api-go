package hash

import (
	"bytes"
	"crypto/rand"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func Test_Sum(t *testing.T) {
	var (
		data = []byte("Hello world")
		sum  = Sum(data)
		hash = []byte{0, 0, 0, 0, 1, 79, 16, 173, 134, 90, 176, 77, 114, 165, 253, 114, 0, 0, 0, 0, 0, 148,
			172, 222, 98, 248, 15, 99, 205, 129, 66, 91, 0, 0, 0, 0, 0, 138, 173, 39, 228, 231, 239, 123,
			170, 96, 186, 61, 0, 0, 0, 0, 0, 90, 69, 237, 131, 90, 161, 73, 38, 164, 185, 55}
	)

	require.Equal(t, hash, sum.Bytes())
}

func Test_Validate(t *testing.T) {
	var (
		data   = []byte("Hello world")
		hash   = Sum(data)
		pieces = splitData(data, 2)
		ln     = len(pieces)
		hashes = make([]Hash, 0, ln)
	)

	for i := 0; i < ln; i++ {
		hashes = append(hashes, Sum(pieces[i]))
	}

	require.True(t, hash.Validate(hashes))
}

func Test_Concat(t *testing.T) {
	var (
		data   = []byte("Hello world")
		hash   = Sum(data)
		pieces = splitData(data, 2)
		ln     = len(pieces)
		hashes = make([]Hash, 0, ln)
	)

	for i := 0; i < ln; i++ {
		hashes = append(hashes, Sum(pieces[i]))
	}

	res, err := Concat(hashes)
	require.NoError(t, err)
	require.Equal(t, hash, res)
}

func Test_HashChunks(t *testing.T) {
	var (
		chars = []byte("+")
		size  = 1400
		data  = bytes.Repeat(chars, size)
		hash  = Sum(data)
		count = 150
	)

	hashes, err := dataHashes(data, count)
	require.NoError(t, err)
	require.Len(t, hashes, count)

	require.True(t, hash.Validate(hashes))

	// 100 / 150 = 0
	hashes, err = dataHashes(data[:100], count)
	require.Error(t, err)
	require.Nil(t, hashes)
}

func TestXOR(t *testing.T) {
	var (
		dl   = 10
		data = make([]byte, dl)
	)

	_, err := rand.Read(data)
	require.NoError(t, err)

	t.Run("XOR with <nil> salt", func(t *testing.T) {
		res := SaltXOR(data, nil)
		require.Equal(t, res, data)
	})

	t.Run("XOR with empty salt", func(t *testing.T) {
		xorWithSalt(t, data, 0)
	})

	t.Run("XOR with salt same data size", func(t *testing.T) {
		xorWithSalt(t, data, dl)
	})

	t.Run("XOR with salt shorter than data aliquot", func(t *testing.T) {
		xorWithSalt(t, data, dl/2)
	})

	t.Run("XOR with salt shorter than data aliquant", func(t *testing.T) {
		xorWithSalt(t, data, dl/3/+1)
	})

	t.Run("XOR with salt longer than data aliquot", func(t *testing.T) {
		xorWithSalt(t, data, dl*2)
	})

	t.Run("XOR with salt longer than data aliquant", func(t *testing.T) {
		xorWithSalt(t, data, dl*2-1)
	})
}

func xorWithSalt(t *testing.T, data []byte, saltSize int) {
	var (
		direct, reverse []byte
		salt            = make([]byte, saltSize)
	)

	_, err := rand.Read(salt)
	require.NoError(t, err)

	direct = SaltXOR(data, salt)
	require.Len(t, direct, len(data))

	reverse = SaltXOR(direct, salt)
	require.Len(t, reverse, len(data))

	require.Equal(t, reverse, data)
}

func splitData(buf []byte, lim int) [][]byte {
	var piece []byte
	pieces := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		piece, buf = buf[:lim], buf[lim:]
		pieces = append(pieces, piece)
	}
	if len(buf) > 0 {
		pieces = append(pieces, buf)
	}
	return pieces
}

func dataHashes(data []byte, count int) ([]Hash, error) {
	var (
		ln     = len(data)
		mis    = ln / count
		off    = (count - 1) * mis
		hashes = make([]Hash, 0, count)
	)
	if mis == 0 {
		return nil, errors.Errorf("could not split %d bytes to %d pieces", ln, count)
	}

	pieces := splitData(data[:off], mis)
	pieces = append(pieces, data[off:])
	for i := 0; i < count; i++ {
		hashes = append(hashes, Sum(pieces[i]))
	}
	return hashes, nil
}
