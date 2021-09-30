package object_test

import (
	"encoding"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg/object"
	v2object "github.com/nspcc-dev/neofs-api-go/v2/object"
	"github.com/stretchr/testify/require"
)

func TestType_ToV2(t *testing.T) {
	typs := []struct {
		t  object.Type
		t2 v2object.Type
	}{
		{
			t:  object.TypeRegular,
			t2: v2object.TypeRegular,
		},
		{
			t:  object.TypeTombstone,
			t2: v2object.TypeTombstone,
		},
		{
			t:  object.TypeStorageGroup,
			t2: v2object.TypeStorageGroup,
		},
	}

	for _, item := range typs {
		t2 := item.t.ToV2()

		require.Equal(t, item.t2, t2)

		require.Equal(t, item.t, object.TypeFromV2(item.t2))
	}
}

func TestType_String(t *testing.T) {
	toPtr := func(v object.Type) *object.Type {
		return &v
	}

	testEnumStrings(t, new(object.Type), []enumStringItem{
		{val: toPtr(object.TypeTombstone), str: "TOMBSTONE"},
		{val: toPtr(object.TypeStorageGroup), str: "STORAGE_GROUP"},
		{val: toPtr(object.TypeRegular), str: "REGULAR"},
	})
}

type enumIface interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
}

type enumStringItem struct {
	val enumIface
	str string
}

func testEnumStrings(t *testing.T, e enumIface, items []enumStringItem) {
	for _, item := range items {
		txt, err := item.val.MarshalText()
		require.NoError(t, err)

		require.Equal(t, item.str, string(txt))

		err = e.UnmarshalText(txt)
		require.NoError(t, err)

		require.EqualValues(t, item.val, e, item.val)
	}

	// incorrect strings
	for _, str := range []string{
		"some string",
		"undefined",
	} {
		require.Error(t, e.UnmarshalText([]byte(str)))
	}
}
