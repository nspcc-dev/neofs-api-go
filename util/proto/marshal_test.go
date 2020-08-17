package proto_test

import (
	"math"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/util/proto"
	"github.com/nspcc-dev/neofs-api-go/util/proto/test"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

type SomeEnum int32

type stablePrimitives struct {
	FieldA []byte
	FieldB string
	FieldC bool
	FieldD int32
	FieldE uint32
	FieldF int64
	FieldG uint64
	FieldH SomeEnum
}

type stableRepPrimitives struct {
	FieldA [][]byte
	FieldB []string
	FieldC []int32
	FieldD []uint32
	FieldE []int64
	FieldF []uint64
}

const (
	ENUM_UNKNOWN  SomeEnum = 0
	ENUM_POSITIVE          = 1
	ENUM_NEGATIVE          = -1
)

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

	fieldNum = 300
	if wrongField {
		fieldNum++
	}
	offset, err = proto.EnumMarshal(fieldNum, buf, int32(s.FieldH))
	if err != nil {
		return nil, errors.Wrap(err, "can't marshal field h")
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
		proto.UInt64Size(204, s.FieldG) +
		proto.EnumSize(300, int32(s.FieldH))
}

func (s *stableRepPrimitives) stableMarshal(buf []byte, wrongField bool) ([]byte, error) {
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
	offset, err := proto.RepeatedBytesMarshal(fieldNum, buf, s.FieldA)
	if err != nil {
		return nil, errors.Wrap(err, "can't marshal field a")
	}
	i += offset

	fieldNum = 2
	if wrongField {
		fieldNum++
	}
	offset, err = proto.RepeatedStringMarshal(fieldNum, buf, s.FieldB)
	if err != nil {
		return nil, errors.Wrap(err, "can't marshal field b")
	}
	i += offset

	fieldNum = 3
	if wrongField {
		fieldNum++
	}
	offset, err = proto.RepeatedInt32Marshal(fieldNum, buf, s.FieldC)
	if err != nil {
		return nil, errors.Wrap(err, "can't marshal field c")
	}
	i += offset

	fieldNum = 4
	if wrongField {
		fieldNum++
	}
	offset, err = proto.RepeatedUInt32Marshal(fieldNum, buf, s.FieldD)
	if err != nil {
		return nil, errors.Wrap(err, "can't marshal field d")
	}
	i += offset

	fieldNum = 5
	if wrongField {
		fieldNum++
	}
	offset, err = proto.RepeatedInt64Marshal(fieldNum, buf, s.FieldE)
	if err != nil {
		return nil, errors.Wrap(err, "can't marshal field e")
	}
	i += offset

	fieldNum = 6
	if wrongField {
		fieldNum++
	}
	offset, err = proto.RepeatedUInt64Marshal(fieldNum, buf, s.FieldF)
	if err != nil {
		return nil, errors.Wrap(err, "can't marshal field f")
	}
	i += offset

	return buf, nil
}

func (s *stableRepPrimitives) stableSize() int {
	f1 := proto.RepeatedBytesSize(1, s.FieldA)
	f2 := proto.RepeatedStringSize(2, s.FieldB)
	f3, _ := proto.RepeatedInt32Size(3, s.FieldC)
	f4, _ := proto.RepeatedUInt32Size(4, s.FieldD)
	f5, _ := proto.RepeatedInt64Size(5, s.FieldE)
	f6, _ := proto.RepeatedUInt64Size(6, s.FieldF)

	return f1 + f2 + f3 + f4 + f5 + f6
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

func TestEnumMarshal(t *testing.T) {
	testEnumMarshal(t, ENUM_UNKNOWN, false)
	testEnumMarshal(t, ENUM_POSITIVE, false)
	testEnumMarshal(t, ENUM_POSITIVE, true)
	testEnumMarshal(t, ENUM_NEGATIVE, false)
	testEnumMarshal(t, ENUM_NEGATIVE, true)
}

func TestRepeatedBytesMarshal(t *testing.T) {
	t.Run("not empty", func(t *testing.T) {
		data := [][]byte{[]byte("One"), []byte("Two"), []byte("Three")}
		testRepeatedBytesMarshal(t, data, false)
		testRepeatedBytesMarshal(t, data, true)
	})

	t.Run("empty", func(t *testing.T) {
		testRepeatedBytesMarshal(t, [][]byte{}, false)
	})

	t.Run("nil", func(t *testing.T) {
		testRepeatedBytesMarshal(t, nil, false)
	})
}

func TestRepeatedStringMarshal(t *testing.T) {
	t.Run("not empty", func(t *testing.T) {
		data := []string{"One", "Two", "Three"}
		testRepeatedStringMarshal(t, data, false)
		testRepeatedStringMarshal(t, data, true)
	})

	t.Run("empty", func(t *testing.T) {
		testRepeatedStringMarshal(t, []string{}, false)
	})

	t.Run("nil", func(t *testing.T) {
		testRepeatedStringMarshal(t, nil, false)
	})
}

func TestRepeatedInt32Marshal(t *testing.T) {
	t.Run("not empty", func(t *testing.T) {
		data := []int32{-1, 0, 1, 2, 3, 4, 5}
		testRepeatedInt32Marshal(t, data, false)
		testRepeatedInt32Marshal(t, data, true)
	})

	t.Run("empty", func(t *testing.T) {
		testRepeatedInt32Marshal(t, []int32{}, false)
	})

	t.Run("nil", func(t *testing.T) {
		testRepeatedInt32Marshal(t, nil, false)
	})
}

func TestRepeatedUInt32Marshal(t *testing.T) {
	t.Run("not empty", func(t *testing.T) {
		data := []uint32{0, 1, 2, 3, 4, 5}
		testRepeatedUInt32Marshal(t, data, false)
		testRepeatedUInt32Marshal(t, data, true)
	})

	t.Run("empty", func(t *testing.T) {
		testRepeatedUInt32Marshal(t, []uint32{}, false)
	})

	t.Run("nil", func(t *testing.T) {
		testRepeatedUInt32Marshal(t, nil, false)
	})
}

func TestRepeatedInt64Marshal(t *testing.T) {
	t.Run("not empty", func(t *testing.T) {
		data := []int64{-1, 0, 1, 2, 3, 4, 5}
		testRepeatedInt64Marshal(t, data, false)
		testRepeatedInt64Marshal(t, data, true)
	})

	t.Run("empty", func(t *testing.T) {
		testRepeatedInt64Marshal(t, []int64{}, false)
	})

	t.Run("nil", func(t *testing.T) {
		testRepeatedInt64Marshal(t, nil, false)
	})
}

func TestRepeatedUInt64Marshal(t *testing.T) {
	t.Run("not empty", func(t *testing.T) {
		data := []uint64{0, 1, 2, 3, 4, 5}
		testRepeatedUInt64Marshal(t, data, false)
		testRepeatedUInt64Marshal(t, data, true)
	})

	t.Run("empty", func(t *testing.T) {
		testRepeatedUInt64Marshal(t, []uint64{}, false)
	})

	t.Run("nil", func(t *testing.T) {
		testRepeatedUInt64Marshal(t, nil, false)
	})
}

func testMarshal(t *testing.T, c stablePrimitives, tr test.Primitives, wrongField bool) *test.Primitives {
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

func testBytesMarshal(t *testing.T, data []byte, wrongField bool) {
	var (
		custom    = stablePrimitives{FieldA: data}
		transport = test.Primitives{FieldA: data}
	)

	result := testMarshal(t, custom, transport, wrongField)

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
		custom    = stablePrimitives{FieldB: s}
		transport = test.Primitives{FieldB: s}
	)

	result := testMarshal(t, custom, transport, wrongField)

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
		custom    = stablePrimitives{FieldC: b}
		transport = test.Primitives{FieldC: b}
	)

	result := testMarshal(t, custom, transport, wrongField)

	if !wrongField {
		require.Equal(t, b, result.FieldC)
	} else {
		require.False(t, false, result.FieldC)
	}
}

func testInt32Marshal(t *testing.T, n int32, wrongField bool) {
	var (
		custom    = stablePrimitives{FieldD: n}
		transport = test.Primitives{FieldD: n}
	)

	result := testMarshal(t, custom, transport, wrongField)

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

	result := testMarshal(t, custom, transport, wrongField)

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

	result := testMarshal(t, custom, transport, wrongField)

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

	result := testMarshal(t, custom, transport, wrongField)

	if !wrongField {
		require.Equal(t, n, result.FieldG)
	} else {
		require.EqualValues(t, 0, result.FieldG)
	}
}

func testEnumMarshal(t *testing.T, e SomeEnum, wrongField bool) {
	var (
		custom    = stablePrimitives{FieldH: e}
		transport = test.Primitives{FieldH: test.Primitives_SomeEnum(e)}
	)

	result := testMarshal(t, custom, transport, wrongField)

	if !wrongField {
		require.EqualValues(t, custom.FieldH, result.FieldH)
	} else {
		require.EqualValues(t, 0, result.FieldH)
	}
}

func testRepMarshal(t *testing.T, c stableRepPrimitives, tr test.RepPrimitives, wrongField bool) *test.RepPrimitives {
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

	result := new(test.RepPrimitives)
	err = result.Unmarshal(wire)
	require.NoError(t, err)

	return result
}

func testRepeatedBytesMarshal(t *testing.T, data [][]byte, wrongField bool) {
	var (
		custom    = stableRepPrimitives{FieldA: data}
		transport = test.RepPrimitives{FieldA: data}
	)

	result := testRepMarshal(t, custom, transport, wrongField)

	if !wrongField {
		require.Len(t, result.FieldA, len(data))
		if len(data) > 0 {
			require.Equal(t, data, result.FieldA)
		}
	} else {
		require.Len(t, result.FieldA, 0)
	}
}

func testRepeatedStringMarshal(t *testing.T, s []string, wrongField bool) {
	var (
		custom    = stableRepPrimitives{FieldB: s}
		transport = test.RepPrimitives{FieldB: s}
	)

	result := testRepMarshal(t, custom, transport, wrongField)

	if !wrongField {
		require.Len(t, result.FieldB, len(s))
		if len(s) > 0 {
			require.Equal(t, s, result.FieldB)
		}
	} else {
		require.Len(t, result.FieldB, 0)
	}
}

func testRepeatedInt32Marshal(t *testing.T, n []int32, wrongField bool) {
	var (
		custom    = stableRepPrimitives{FieldC: n}
		transport = test.RepPrimitives{FieldC: n}
	)

	result := testRepMarshal(t, custom, transport, wrongField)

	if !wrongField {
		require.Len(t, result.FieldC, len(n))
		if len(n) > 0 {
			require.Equal(t, n, result.FieldC)
		}
	} else {
		require.Len(t, result.FieldC, 0)
	}
}

func testRepeatedUInt32Marshal(t *testing.T, n []uint32, wrongField bool) {
	var (
		custom    = stableRepPrimitives{FieldD: n}
		transport = test.RepPrimitives{FieldD: n}
	)

	result := testRepMarshal(t, custom, transport, wrongField)

	if !wrongField {
		require.Len(t, result.FieldD, len(n))
		if len(n) > 0 {
			require.Equal(t, n, result.FieldD)
		}
	} else {
		require.Len(t, result.FieldD, 0)
	}
}

func testRepeatedInt64Marshal(t *testing.T, n []int64, wrongField bool) {
	var (
		custom    = stableRepPrimitives{FieldE: n}
		transport = test.RepPrimitives{FieldE: n}
	)

	result := testRepMarshal(t, custom, transport, wrongField)

	if !wrongField {
		require.Len(t, result.FieldE, len(n))
		if len(n) > 0 {
			require.Equal(t, n, result.FieldE)
		}
	} else {
		require.Len(t, result.FieldE, 0)
	}
}

func testRepeatedUInt64Marshal(t *testing.T, n []uint64, wrongField bool) {
	var (
		custom    = stableRepPrimitives{FieldF: n}
		transport = test.RepPrimitives{FieldF: n}
	)

	result := testRepMarshal(t, custom, transport, wrongField)

	if !wrongField {
		require.Len(t, result.FieldF, len(n))
		if len(n) > 0 {
			require.Equal(t, n, result.FieldF)
		}
	} else {
		require.Len(t, result.FieldF, 0)
	}
}