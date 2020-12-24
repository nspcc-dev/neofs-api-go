package audit_test

import (
	"crypto/rand"
	"crypto/sha256"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/audit"
	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	"github.com/nspcc-dev/neofs-api-go/pkg/object"
	"github.com/stretchr/testify/require"
)

func testSHA256() (cs [sha256.Size]byte) {
	_, _ = rand.Read(cs[:])
	return
}

func testCID() *container.ID {
	cid := container.NewID()
	cid.SetSHA256(testSHA256())

	return cid
}

func testOID() *object.ID {
	id := object.NewID()
	id.SetSHA256(testSHA256())

	return id
}

func TestResult(t *testing.T) {
	r := audit.NewResult()
	require.Equal(t, pkg.SDKVersion(), r.Version())

	epoch := uint64(13)
	r.SetAuditEpoch(epoch)
	require.Equal(t, epoch, r.AuditEpoch())

	cid := testCID()
	r.SetContainerID(cid)
	require.Equal(t, cid, r.ContainerID())

	key := []byte{1, 2, 3}
	r.SetPublicKey(key)
	require.Equal(t, key, r.PublicKey())

	r.SetComplete(true)
	require.True(t, r.Complete())

	requests := uint32(2)
	r.SetRequests(requests)
	require.Equal(t, requests, r.Requests())

	retries := uint32(1)
	r.SetRetries(retries)
	require.Equal(t, retries, r.Retries())

	passSG := []*object.ID{testOID(), testOID()}
	r.SetPassSG(passSG)
	require.Equal(t, passSG, r.PassSG())

	failSG := []*object.ID{testOID(), testOID()}
	r.SetFailSG(failSG)
	require.Equal(t, failSG, r.FailSG())

	hit := uint32(1)
	r.SetHit(hit)
	require.Equal(t, hit, r.Hit())

	miss := uint32(2)
	r.SetMiss(miss)
	require.Equal(t, miss, r.Miss())

	fail := uint32(3)
	r.SetFail(fail)
	require.Equal(t, fail, r.Fail())

	passNodes := [][]byte{{1}, {2}}
	r.SetPassNodes(passNodes)
	require.Equal(t, passNodes, r.PassNodes())

	failNodes := [][]byte{{3}, {4}}
	r.SetFailNodes(failNodes)
	require.Equal(t, failNodes, r.FailNodes())
}

func TestStorageGroupEncoding(t *testing.T) {
	r := audit.NewResult()
	r.SetAuditEpoch(13)
	r.SetContainerID(testCID())
	r.SetPublicKey([]byte{1, 2, 3})
	r.SetPassSG([]*object.ID{testOID(), testOID()})
	r.SetFailSG([]*object.ID{testOID(), testOID()})
	r.SetRequests(3)
	r.SetRetries(2)
	r.SetHit(1)
	r.SetMiss(2)
	r.SetFail(3)
	r.SetPassNodes([][]byte{{1}, {2}})
	r.SetFailNodes([][]byte{{3}, {4}})

	t.Run("binary", func(t *testing.T) {
		data, err := r.Marshal()
		require.NoError(t, err)

		r2 := audit.NewResult()
		require.NoError(t, r2.Unmarshal(data))

		require.Equal(t, r, r2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := r.MarshalJSON()
		require.NoError(t, err)

		r2 := audit.NewResult()
		require.NoError(t, r2.UnmarshalJSON(data))

		require.Equal(t, r, r2)
	})
}
