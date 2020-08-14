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
	FieldB string
}

func (s *stablePrimitives) stableMarshal(buf []byte, wrongField bool) ([]byte, error) {
	if s == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, s.stableSize())
	}

	var (
		i, offset, fieldNum int
	)

	fieldNum = 1
	if wrongField {
		fieldNum++
	}
	offset, err := proto.BytesMarshal(fieldNum, buf, s.FieldA)
	if err != nil {
		return nil, errors.Wrap(err, "can't marshal field a")
	}
	i += offset

	fieldNum = 2
	if wrongField {
		fieldNum++
	}
	offset, err = proto.StringMarshal(fieldNum, buf, s.FieldB)
	if err != nil {
		return nil, errors.Wrap(err, "can't marshal field a")
	}
	i += offset

	return buf, nil
}

func (s *stablePrimitives) stableSize() int {
	return proto.BytesSize(1, s.FieldA) +
		proto.StringSize(2, s.FieldB)
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

	wire, err = custom.stableMarshal(nil, wrongField)
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

func TestStringMarshal(t *testing.T) {
	t.Run("not empty", func(t *testing.T) {
		data := "Hello World"
		testStringMarshal(t, data, false)
		testStringMarshal(t, data, true)
	})

	t.Run("empty", func(t *testing.T) {
		testStringMarshal(t, "", false)
	})
}

func testStringMarshal(t *testing.T, s string, wrongField bool) {
	var (
		wire []byte
		err  error

		custom    = stablePrimitives{FieldB: s}
		transport = test.Primitives{FieldB: s}
	)

	wire, err = custom.stableMarshal(nil, wrongField)
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
		require.Len(t, result.FieldB, len(s))
		if len(s) > 0 {
			require.Equal(t, s, result.FieldB)
		}
	} else {
		require.Len(t, result.FieldB, 0)
	}
}
