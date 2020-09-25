package object

import (
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
