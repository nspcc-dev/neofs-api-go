package object

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/object"
	"github.com/stretchr/testify/require"
)

func TestType_ToV2(t *testing.T) {
	typs := []struct {
		t  Type
		t2 object.Type
	}{
		{
			t:  TypeRegular,
			t2: object.TypeRegular,
		},
		{
			t:  TypeTombstone,
			t2: object.TypeTombstone,
		},
		{
			t:  TypeStorageGroup,
			t2: object.TypeStorageGroup,
		},
	}

	for _, item := range typs {
		t2 := item.t.ToV2()

		require.Equal(t, item.t2, t2)

		require.Equal(t, item.t, TypeFromV2(item.t2))
	}
}
