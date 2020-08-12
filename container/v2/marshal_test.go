package v2

import (
	"testing"

	netmap "github.com/nspcc-dev/neofs-api-go/netmap/v2"
	refs "github.com/nspcc-dev/neofs-api-go/refs/v2"
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
		Rules: &netmap.PlacementRule{
			ReplFactor: 4,
			SfGroups: []*netmap.PlacementRule_SFGroup{
				{
					Selectors: []*netmap.PlacementRule_SFGroup_Selector{
						{
							Count: 1,
							Key:   "Node",
						},
					},
					Filters: []*netmap.PlacementRule_SFGroup_Filter{
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
