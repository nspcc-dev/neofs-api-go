package netmap_test

import (
	"strconv"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/stretchr/testify/require"
)

func subnetAttrKey(val string) string {
	return "__NEOFS__SUBNET_" + val
}

func assertSubnetAttrKey(t *testing.T, attr *netmap.Attribute, num uint32) {
	require.Equal(t, subnetAttrKey(strconv.FormatUint(uint64(num), 10)), attr.GetKey())
}

func TestWriteSubnetInfo(t *testing.T) {
	t.Run("entry", func(t *testing.T) {
		t.Run("zero subnet", func(t *testing.T) {
			var (
				node netmap.NodeInfo
				info netmap.NodeSubnetInfo
			)

			netmap.WriteSubnetInfo(&node, info)

			// entry to zero subnet does not require an attribute
			attrs := node.GetAttributes()
			require.Empty(t, attrs)

			// exit the subnet
			info.SetEntryFlag(false)

			netmap.WriteSubnetInfo(&node, info)

			// exit from zero subnet should be clearly reflected in attributes
			attrs = node.GetAttributes()
			require.Len(t, attrs, 1)

			attr := attrs[0]
			assertSubnetAttrKey(t, attr, 0)
			require.Equal(t, "False", attr.GetValue())

			// again enter to zero subnet
			info.SetEntryFlag(true)

			netmap.WriteSubnetInfo(&node, info)

			// attribute should be removed
			attrs = node.GetAttributes()
			require.Empty(t, attrs)
		})

		t.Run("non-zero subnet", func(t *testing.T) {
			var (
				node netmap.NodeInfo
				info netmap.NodeSubnetInfo
				id   refs.SubnetID
			)

			// create non-zero subnet ID
			const num = 15

			id.SetValue(num)

			// enter to the subnet
			info.SetID(&id)
			info.SetEntryFlag(true)

			netmap.WriteSubnetInfo(&node, info)

			// check attribute format
			attrs := node.GetAttributes()
			require.Len(t, attrs, 1)

			attr := attrs[0]
			assertSubnetAttrKey(t, attr, num)
			require.Equal(t, "True", attr.GetValue())

			// again exit the subnet
			info.SetEntryFlag(false)

			netmap.WriteSubnetInfo(&node, info)

			// attribute should be removed
			attrs = node.GetAttributes()
			require.Empty(t, attrs)
		})
	})
}

func TestSubnets(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var node netmap.NodeInfo

		called := 0

		err := netmap.IterateSubnets(&node, func(id refs.SubnetID) error {
			called++

			require.True(t, refs.IsZeroSubnet(&id))

			return nil
		})

		require.NoError(t, err)
		require.EqualValues(t, 1, called)
	})

	t.Run("with correct attribute", func(t *testing.T) {
		var (
			node netmap.NodeInfo
			attr netmap.Attribute
		)

		attr.SetKey(subnetAttrKey("13"))
		attr.SetValue("True")

		attrs := []*netmap.Attribute{&attr}

		node.SetAttributes(attrs)

		called := 0

		err := netmap.IterateSubnets(&node, func(id refs.SubnetID) error {
			if !refs.IsZeroSubnet(&id) {
				called++
				require.EqualValues(t, 13, id.GetValue())
			}

			return nil
		})

		require.NoError(t, err)
		require.EqualValues(t, 1, called)
	})

	t.Run("with incorrect attribute", func(t *testing.T) {
		assertErr := func(attr netmap.Attribute) {
			var node netmap.NodeInfo

			node.SetAttributes([]*netmap.Attribute{&attr})

			require.Error(t, netmap.IterateSubnets(&node, func(refs.SubnetID) error {
				return nil
			}))
		}

		t.Run("incorrect key", func(t *testing.T) {
			var attr netmap.Attribute

			attr.SetKey(subnetAttrKey("one-two-three"))

			assertErr(attr)
		})

		t.Run("incorrect value", func(t *testing.T) {
			var attr netmap.Attribute

			attr.SetKey(subnetAttrKey("1"))

			for _, invalidVal := range []string{
				"",
				"Troo",
				"Fols",
			} {
				attr.SetValue(invalidVal)
				assertErr(attr)
			}

			assertErr(attr)
		})
	})

	t.Run("remove entry", func(t *testing.T) {
		t.Run("zero", func(t *testing.T) {
			var node netmap.NodeInfo

			// enter to some non-zero subnet so that zero is not the only one
			var attr netmap.Attribute

			attr.SetKey(subnetAttrKey("321"))
			attr.SetValue("True")

			attrs := []*netmap.Attribute{&attr}
			node.SetAttributes(attrs)

			err := netmap.IterateSubnets(&node, func(id refs.SubnetID) error {
				if refs.IsZeroSubnet(&id) {
					return netmap.ErrRemoveSubnet
				}

				return nil
			})

			require.NoError(t, err)

			attrs = node.GetAttributes()
			require.Len(t, attrs, 2)

			found := false

			for i := range attrs {
				if attrs[i].GetKey() == subnetAttrKey("0") {
					require.Equal(t, "False", attrs[i].GetValue())
					found = true
				}
			}

			require.True(t, found)
		})

		t.Run("non-zero", func(t *testing.T) {
			var (
				node netmap.NodeInfo
				attr netmap.Attribute
			)

			attr.SetKey(subnetAttrKey("99"))
			attr.SetValue("True")

			attrs := []*netmap.Attribute{&attr}
			node.SetAttributes(attrs)

			err := netmap.IterateSubnets(&node, func(id refs.SubnetID) error {
				if !refs.IsZeroSubnet(&id) {
					return netmap.ErrRemoveSubnet
				}

				return nil
			})

			require.NoError(t, err)

			attrs = node.GetAttributes()
			require.Empty(t, attrs)
		})

		t.Run("all", func(t *testing.T) {
			var (
				node  netmap.NodeInfo
				attrs []*netmap.Attribute
			)

			// enter to some non-zero subnet so that zero is not the only one
			for i := 1; i <= 5; i++ {
				var attr netmap.Attribute

				attr.SetKey(subnetAttrKey(strconv.Itoa(i)))
				attr.SetValue("True")

				attrs = append(attrs, &attr)
			}

			node.SetAttributes(attrs)

			err := netmap.IterateSubnets(&node, func(id refs.SubnetID) error {
				return netmap.ErrRemoveSubnet
			})

			require.Error(t, err)
		})
	})
}
