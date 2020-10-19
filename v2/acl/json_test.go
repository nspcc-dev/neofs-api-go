package acl_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/acl"
	"github.com/stretchr/testify/require"
)

func TestRecordJSON(t *testing.T) {
	exp := generateRecord(false)

	t.Run("non empty", func(t *testing.T) {
		data := acl.RecordToJSON(exp)
		require.NotNil(t, data)

		got := acl.RecordFromJSON(data)
		require.NotNil(t, got)

		require.Equal(t, exp, got)
	})
}

func TestEACLTableJSON(t *testing.T) {
	exp := generateEACL()

	t.Run("non empty", func(t *testing.T) {
		data := acl.TableToJSON(exp)
		require.NotNil(t, data)

		got := acl.TableFromJSON(data)
		require.NotNil(t, got)

		require.Equal(t, exp, got)
	})
}

func TestBearerTokenJSON(t *testing.T) {
	exp := generateBearerToken("token")

	t.Run("non empty", func(t *testing.T) {
		data := acl.BearerTokenToJSON(exp)
		require.NotNil(t, data)

		got := acl.BearerTokenFromJSON(data)
		require.NotNil(t, got)

		require.Equal(t, exp, got)
	})
}
