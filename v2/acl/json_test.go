package acl_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/acl"
	"github.com/stretchr/testify/require"
)

func TestRecordJSON(t *testing.T) {
	exp := generateRecord(false)

	t.Run("non empty", func(t *testing.T) {
		data, err := acl.RecordToJSON(exp)
		require.NoError(t, err)

		got, err := acl.RecordFromJSON(data)
		require.NoError(t, err)

		require.Equal(t, exp, got)
	})

	t.Run("empty", func(t *testing.T) {
		_, err := acl.RecordToJSON(nil)
		require.Error(t, err)

		_, err = acl.RecordFromJSON(nil)
		require.Error(t, err)
	})
}

func TestEACLTableJSON(t *testing.T) {
	exp := generateEACL()

	t.Run("non empty", func(t *testing.T) {
		data, err := acl.TableToJSON(exp)
		require.NoError(t, err)

		got, err := acl.TableFromJSON(data)
		require.NoError(t, err)

		require.Equal(t, exp, got)
	})

	t.Run("empty", func(t *testing.T) {
		_, err := acl.TableToJSON(nil)
		require.Error(t, err)

		_, err = acl.TableFromJSON(nil)
		require.Error(t, err)
	})
}

func TestBearerTokenJSON(t *testing.T) {
	exp := generateBearerToken("token")

	t.Run("non empty", func(t *testing.T) {
		data, err := acl.BearerTokenToJSON(exp)
		require.NoError(t, err)

		got, err := acl.BearerTokenFromJSON(data)
		require.NoError(t, err)

		require.Equal(t, exp, got)
	})
}
