package container

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/nspcc-dev/neofs-api-go/refs"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/nspcc-dev/netmap"
	"github.com/stretchr/testify/require"
)

func TestCID(t *testing.T) {
	t.Run("check that marshal/unmarshal works like expected", func(t *testing.T) {
		var (
			c2   Container
			cid2 CID
			key  = test.DecodeKey(0)
		)

		rules := netmap.PlacementRule{
			ReplFactor: 2,
			SFGroups: []netmap.SFGroup{
				{
					Selectors: []netmap.Select{
						{Key: "Country", Count: 1},
						{Key: netmap.NodesBucket, Count: 2},
					},
					Filters: []netmap.Filter{
						{Key: "Country", F: netmap.FilterIn("USA")},
					},
				},
			},
		}

		owner, err := refs.NewOwnerID(&key.PublicKey)
		require.NoError(t, err)

		c1, err := New(10, owner, 0xDEADBEEF, rules)
		require.NoError(t, err)

		data, err := proto.Marshal(c1)
		require.NoError(t, err)

		require.NoError(t, c2.Unmarshal(data))
		require.Equal(t, c1, &c2)

		cid1, err := c1.ID()
		require.NoError(t, err)

		data, err = proto.Marshal(&cid1)
		require.NoError(t, err)
		require.NoError(t, cid2.Unmarshal(data))

		require.Equal(t, cid1, cid2)
	})
}
