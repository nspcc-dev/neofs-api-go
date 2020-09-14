package eacl

import (
	"testing"

	v2acl "github.com/nspcc-dev/neofs-api-go/v2/acl"
	"github.com/stretchr/testify/require"
)

func TestFilter(t *testing.T) {
	filter := &Filter{
		from:    HeaderFromObject,
		name:    HdrObjSysNamePayloadLength,
		matcher: MatchStringEqual,
		value:   "200",
	}

	v2 := filter.ToV2()
	require.NotNil(t, v2)
	require.Equal(t, v2acl.HeaderTypeObject, v2.GetHeaderType())
	require.EqualValues(t, v2acl.MatchTypeStringEqual, v2.GetMatchType())
	require.Equal(t, filter.Name(), v2.GetName())
	require.Equal(t, filter.Value(), v2.GetValue())

	newFilter := NewFilterFromV2(v2)
	require.Equal(t, filter, newFilter)

	t.Run("from nil v2 filter", func(t *testing.T) {
		require.Equal(t, new(Filter), NewFilterFromV2(nil))
	})
}
