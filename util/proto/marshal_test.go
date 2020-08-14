package proto_test

import (
	"math"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/util/proto"
	"github.com/nspcc-dev/neofs-api-go/util/proto/test"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

type stablePrimitives struct {
	FieldA []byte
	FieldB string
	FieldC bool
	FieldD int32
	FieldE uint32
	FieldF int64
	FieldG uint64
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
		return nil, errors.Wrap(err, "can't marshal field b")
	}
	i += offset

	fieldNum = 200
	if wrongField {
		fieldNum++
	}
	offset, err = proto.BoolMarshal(fieldNum, buf, s.FieldC)
	if err != nil {
		return nil, errors.Wrap(err, "can't marshal field c")
	}
	i += offset

	fieldNum = 201
	if wrongField {
		fieldNum++
	}
	offset, err = proto.Int32Marshal(fieldNum, buf, s.FieldD)
	if err != nil {
		return nil, errors.Wrap(err, "can't marshal field d")
	}
	i += offset

	fieldNum = 202
	if wrongField {
		fieldNum++
	}
	offset, err = proto.UInt32Marshal(fieldNum, buf, s.FieldE)
	if err != nil {
		return nil, errors.Wrap(err, "can't marshal field e")
	}
	i += offset

	fieldNum = 203
	if wrongField {
		fieldNum++
	}
	offset, err = proto.Int64Marshal(fieldNum, buf, s.FieldF)
	if err != nil {
		return nil, errors.Wrap(err, "can't marshal field f")
	}
	i += offset

	fieldNum = 204
	if wrongField {
		fieldNum++
	}
	offset, err = proto.UInt64Marshal(fieldNum, buf, s.FieldG)
	if err != nil {
		return nil, errors.Wrap(err, "can't marshal field g")
	}
	i += offset

	return buf, nil
}

func (s *stablePrimitives) stableSize() int {
	return proto.BytesSize(1, s.FieldA) +
		proto.StringSize(2, s.FieldB) +
		proto.BoolSize(200, s.FieldC) +
		proto.Int32Size(201, s.FieldD) +
		proto.UInt32Size(202, s.FieldE) +
		proto.Int64Size(203, s.FieldF) +
		proto.UInt64Size(204, s.FieldG)
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

func TestBoolMarshal(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		testBoolMarshal(t, true, false)
		testBoolMarshal(t, true, true)
	})

	t.Run("false", func(t *testing.T) {
		testBoolMarshal(t, false, false)
	})
}

func TestInt32Marshal(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		testInt32Marshal(t, 0, false)
	})

	t.Run("positive", func(t *testing.T) {
		testInt32Marshal(t, math.MaxInt32, false)
		testInt32Marshal(t, math.MaxInt32, true)
	})

	t.Run("negative", func(t *testing.T) {
		testInt32Marshal(t, math.MinInt32, false)
		testInt32Marshal(t, math.MinInt32, true)
	})
}

func TestUInt32Marshal(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		testUInt32Marshal(t, 0, false)
	})

	t.Run("non zero", func(t *testing.T) {
		testUInt32Marshal(t, math.MaxUint32, false)
		testUInt32Marshal(t, math.MaxUint32, true)
	})
}

func TestInt64Marshal(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		testInt32Marshal(t, 0, false)
	})

	t.Run("positive", func(t *testing.T) {
		testInt64Marshal(t, math.MaxInt64, false)
		testInt64Marshal(t, math.MaxInt64, true)
	})

	t.Run("negative", func(t *testing.T) {
		testInt64Marshal(t, math.MinInt64, false)
		testInt64Marshal(t, math.MinInt64, true)
	})
}

func TestUInt64Marshal(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		testUInt64Marshal(t, 0, false)
	})

	t.Run("non zero", func(t *testing.T) {
		testUInt64Marshal(t, math.MaxUint64, false)
		testUInt64Marshal(t, math.MaxUint64, true)
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

func testBoolMarshal(t *testing.T, b bool, wrongField bool) {
	var (
		wire []byte
		err  error

		custom    = stablePrimitives{FieldC: b}
		transport = test.Primitives{FieldC: b}
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
		require.Equal(t, b, result.FieldC)
	} else {
		require.False(t, false, result.FieldC)
	}
}

func testIntMarshal(t *testing.T, c stablePrimitives, tr test.Primitives, wrongField bool) *test.Primitives {
	var (
		wire []byte
		err  error
	)
	wire, err = c.stableMarshal(nil, wrongField)
	require.NoError(t, err)

	wireGen, err := tr.Marshal()
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

	return result
}

func testInt32Marshal(t *testing.T, n int32, wrongField bool) {
	var (
		custom    = stablePrimitives{FieldD: n}
		transport = test.Primitives{FieldD: n}
	)

	result := testIntMarshal(t, custom, transport, wrongField)

	if !wrongField {
		require.Equal(t, n, result.FieldD)
	} else {
		require.EqualValues(t, 0, result.FieldD)
	}
}

func testUInt32Marshal(t *testing.T, n uint32, wrongField bool) {
	var (
		custom    = stablePrimitives{FieldE: n}
		transport = test.Primitives{FieldE: n}
	)

	result := testIntMarshal(t, custom, transport, wrongField)

	if !wrongField {
		require.Equal(t, n, result.FieldE)
	} else {
		require.EqualValues(t, 0, result.FieldE)
	}
}

func testInt64Marshal(t *testing.T, n int64, wrongField bool) {
	var (
		custom    = stablePrimitives{FieldF: n}
		transport = test.Primitives{FieldF: n}
	)

	result := testIntMarshal(t, custom, transport, wrongField)

	if !wrongField {
		require.Equal(t, n, result.FieldF)
	} else {
		require.EqualValues(t, 0, result.FieldF)
	}
}

func testUInt64Marshal(t *testing.T, n uint64, wrongField bool) {
	var (
		custom    = stablePrimitives{FieldG: n}
		transport = test.Primitives{FieldG: n}
	)

	result := testIntMarshal(t, custom, transport, wrongField)

	if !wrongField {
		require.Equal(t, n, result.FieldG)
	} else {
		require.EqualValues(t, 0, result.FieldG)
	}
}
