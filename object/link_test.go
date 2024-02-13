package object_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/object"
	objecttest "github.com/nspcc-dev/neofs-api-go/v2/object/test"
	"github.com/stretchr/testify/require"
)

func TestLinkRW(t *testing.T) {
	var l object.Link
	var obj object.Object

	require.Error(t, object.ReadLink(&l, obj))

	l = *objecttest.GenerateLink(false)

	object.WriteLink(&obj, l)

	var l2 object.Link

	require.NoError(t, object.ReadLink(&l2, obj))

	require.Equal(t, l, l2)
}
