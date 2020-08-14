package container

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/stretchr/testify/require"
)

var (
	cnr = &Container{
		OwnerId:  &refs.OwnerID{Value: []byte("Owner")},
		Nonce:    []byte("Salt"),
		BasicAcl: 505,
		Attributes: []*Container_Attribute{
			{
				Key:   "Hello",
				Value: "World",
			},
			{
				Key:   "Privet",
				Value: "Mir",
			},
		},
		PlacementPolicy: &netmap.PlacementPolicy{
			ReplFactor: 4,
			FilterGroups: []*netmap.PlacementPolicy_FilterGroup{
				{
					Selectors: []*netmap.PlacementPolicy_FilterGroup_Selector{
						{
							Count: 1,
							Key:   "Node",
						},
					},
					Filters: []*netmap.PlacementPolicy_FilterGroup_Filter{
						{
							Key: "City",
						},
						{
							Key: "Datacenter",
						},
					},
					Exclude: []uint32{4, 5, 6},
				},
			},
		},
	}
)

func TestContainer_StableMarshal(t *testing.T) {
	newCnr := new(Container)

	wire, err := cnr.StableMarshal(nil)
	require.NoError(t, err)

	err = newCnr.Unmarshal(wire)
	require.NoError(t, err)

	require.Equal(t, cnr, newCnr)
}

func TestPutRequest_Body_StableMarshal(t *testing.T) {
	expectedBody := new(PutRequest_Body)
	expectedBody.Container = cnr
	expectedBody.PublicKey = []byte{1, 2, 3, 4}
	expectedBody.Signature = []byte{5, 6, 7, 8}

	wire, err := expectedBody.StableMarshal(nil)
	require.NoError(t, err)

	gotBody := new(PutRequest_Body)
	err = gotBody.Unmarshal(wire)
	require.NoError(t, err)

	require.Equal(t, expectedBody, gotBody)
}
