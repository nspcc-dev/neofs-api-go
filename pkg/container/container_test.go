package container_test

import (
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/acl"
	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	"github.com/nspcc-dev/neofs-api-go/pkg/netmap"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
)

func TestNewContainer(t *testing.T) {
	c := container.New()

	nonce := uuid.New()

	wallet, err := owner.NEO3WalletFromPublicKey(&test.DecodeKey(1).PublicKey)
	require.NoError(t, err)

	ownerID := owner.NewIDFromNeo3Wallet(wallet)
	policy := generatePlacementPolicy()

	c.SetBasicACL(acl.PublicBasicRule)
	c.SetAttributes(generateAttributes(5))
	c.SetPlacementPolicy(policy)
	c.SetNonceUUID(nonce)
	c.SetOwnerID(ownerID)
	c.SetVersion(pkg.SDKVersion())

	v2 := c.ToV2()
	newContainer := container.NewContainerFromV2(v2)

	require.EqualValues(t, newContainer.PlacementPolicy(), policy)
	require.EqualValues(t, newContainer.Attributes(), generateAttributes(5))
	require.EqualValues(t, newContainer.BasicACL(), acl.PublicBasicRule)

	newNonce, err := newContainer.NonceUUID()
	require.NoError(t, err)

	require.EqualValues(t, newNonce, nonce)
	require.EqualValues(t, newContainer.OwnerID(), ownerID)
	require.EqualValues(t, newContainer.Version(), pkg.SDKVersion())
}

func generateAttributes(n int) container.Attributes {
	attrs := make(container.Attributes, 0, n)

	for i := 0; i < n; i++ {
		strN := strconv.Itoa(n)

		attr := container.NewAttribute()
		attr.SetKey("key" + strN)
		attr.SetValue("val" + strN)

		attrs = append(attrs, attr)
	}

	return attrs
}

func generatePlacementPolicy() *netmap.PlacementPolicy {
	p := new(netmap.PlacementPolicy)
	p.SetContainerBackupFactor(10)

	return p
}

func TestContainerEncoding(t *testing.T) {
	c := container.New(
		container.WithAttribute("key", "value"),
	)

	t.Run("binary", func(t *testing.T) {
		data, err := c.Marshal()
		require.NoError(t, err)

		c2 := container.New()
		require.NoError(t, c2.Unmarshal(data))

		require.Equal(t, c, c2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := c.MarshalJSON()
		require.NoError(t, err)

		c2 := container.New()
		require.NoError(t, c2.UnmarshalJSON(data))

		require.Equal(t, c, c2)
	})
}
