package object

import (
	"github.com/nspcc-dev/neofs-api-go/v2/object"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRange_SetOffset(t *testing.T) {
	r := NewRange()

	off := uint64(13)
	r.SetOffset(off)

	require.Equal(t, off, r.GetOffset())
}

func TestRange_SetLength(t *testing.T) {
	r := NewRange()

	ln := uint64(7)
	r.SetLength(ln)

	require.Equal(t, ln, r.GetLength())
}

func TestNewRangeFromV2(t *testing.T) {
	t.Run("from nil", func(t *testing.T) {
		var x *object.Range

		require.Nil(t, NewRangeFromV2(x))
	})
}

func TestRange_ToV2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var x *Range

		require.Nil(t, x.ToV2())
	})
}
