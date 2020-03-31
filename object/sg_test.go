package object

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/hash"
	"github.com/nspcc-dev/neofs-api-go/storagegroup"
	"github.com/stretchr/testify/require"
)

func TestObject_StorageGroup(t *testing.T) {
	t.Run("group method", func(t *testing.T) {
		var linkCount byte = 100

		obj := &Object{Headers: make([]Header, 0, linkCount)}
		require.Empty(t, obj.Group())

		idList := make([]ID, linkCount)
		for i := byte(0); i < linkCount; i++ {
			idList[i] = ID{i}
			obj.Headers = append(obj.Headers, Header{
				Value: &Header_Link{Link: &Link{
					Type: Link_StorageGroup,
					ID:   idList[i],
				}},
			})
		}

		rand.Shuffle(len(obj.Headers), func(i, j int) { obj.Headers[i], obj.Headers[j] = obj.Headers[j], obj.Headers[i] })
		sort.Sort(storagegroup.IDList(idList))
		require.Equal(t, idList, obj.Group())
	})
	t.Run("identification method", func(t *testing.T) {
		oid, cid, owner := ID{1}, CID{2}, OwnerID{3}
		obj := &Object{
			SystemHeader: SystemHeader{
				ID:      oid,
				OwnerID: owner,
				CID:     cid,
			},
		}

		idInfo := obj.IDInfo()
		require.Equal(t, oid, idInfo.SGID)
		require.Equal(t, cid, idInfo.CID)
		require.Equal(t, owner, idInfo.OwnerID)
	})
	t.Run("zones method", func(t *testing.T) {
		sgSize := uint64(100)

		d := make([]byte, sgSize)
		_, err := rand.Read(d)
		require.NoError(t, err)
		sgHash := hash.Sum(d)

		obj := &Object{
			Headers: []Header{
				{
					Value: &Header_StorageGroup{
						StorageGroup: &storagegroup.StorageGroup{
							ValidationDataSize: sgSize,
							ValidationHash:     sgHash,
						},
					},
				},
			},
		}

		var (
			sumSize uint64
			zones   = obj.Zones()
			hashes  = make([]Hash, len(zones))
		)

		for i := range zones {
			sumSize += zones[i].Size
			hashes[i] = zones[i].Hash
		}

		sumHash, err := hash.Concat(hashes)
		require.NoError(t, err)

		require.Equal(t, sgSize, sumSize)
		require.Equal(t, sgHash, sumHash)
	})
}
