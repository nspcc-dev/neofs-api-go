package proto_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/util/proto"
	"github.com/nspcc-dev/neofs-api-go/util/proto/test"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

type stablePrimitives struct {
	FieldA []byte
}

func (s *stablePrimitives) stableMarshal(buf []byte) ([]byte, error) {
	if s == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, s.stableSize())
	}

	var (
		i, offset int
	)

	offset, err := proto.BytesMarshal(1, buf, s.FieldA)
	if err != nil {
		return nil, errors.Wrap(err, "can't marshal field a")
	}
	i += offset

	return buf, nil
}

func (s *stablePrimitives) stableMarshalWrongFieldNum(buf []byte) ([]byte, error) {
	if s == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, s.stableSize())
	}

	var (
		i, offset int
	)

	offset, err := proto.BytesMarshal(1+1, buf, s.FieldA)
	if err != nil {
		return nil, errors.Wrap(err, "can't marshal field a")
	}
	i += offset

	return buf, nil
}

func (s *stablePrimitives) stableSize() int {
	return proto.BytesSize(1, s.FieldA)
}

func TestBytesMarshal(t *testing.T) {
	t.Run("not empty", func(t *testing.T) {
		data := []byte("Hello World")
		testBytesMarshal(t, data, false)
		testBytesMarshal(t, data, true)
	})

	t.Run("empty", func(t *testing.T) {
		testBytesMarshal(t, []byte{}, false)
	})

	t.Run("nil", func(t *testing.T) {
		testBytesMarshal(t, nil, false)
	})
}

func testBytesMarshal(t *testing.T, data []byte, wrongField bool) {
	var (
		wire []byte
		err  error

		custom    = stablePrimitives{FieldA: data}
		transport = test.Primitives{FieldA: data}
	)

	if !wrongField {
		wire, err = custom.stableMarshal(nil)
	} else {
		wire, err = custom.stableMarshalWrongFieldNum(nil)
	}
	require.NoError(t, err)

	wireGen, err := transport.Marshal()
	require.NoError(t, err)

	if !wrongField {
		// we can check equality because single field cannot be unstable marshalled
		require.Equal(t, wireGen, wire)
	} else {
		require.NotEqual(t, wireGen, wire)
	}

	result := new(test.Primitives)
	err = result.Unmarshal(wire)
	require.NoError(t, err)

	if !wrongField {
		require.Len(t, result.FieldA, len(data))
		if len(data) > 0 {
			require.Equal(t, data, result.FieldA)
		}
	} else {
		require.Len(t, result.FieldA, 0)
	}
}
