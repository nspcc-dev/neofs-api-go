package session_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/session"
	"github.com/stretchr/testify/require"
)

func TestChecksumJSON(t *testing.T) {
	ctx := generateObjectCtx("id")

	data, err := ctx.MarshalJSON()
	require.NoError(t, err)

	ctx2 := new(session.ObjectSessionContext)
	require.NoError(t, ctx2.UnmarshalJSON(data))

	require.Equal(t, ctx, ctx2)
}
