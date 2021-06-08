package audit_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/audit"
	audittest "github.com/nspcc-dev/neofs-api-go/pkg/audit/test"
	cidtest "github.com/nspcc-dev/neofs-api-go/pkg/container/id/test"
	"github.com/nspcc-dev/neofs-api-go/pkg/object"
	objecttest "github.com/nspcc-dev/neofs-api-go/pkg/object/test"
	"github.com/stretchr/testify/require"
)

func TestResult(t *testing.T) {
	r := audit.NewResult()
	require.Equal(t, pkg.SDKVersion(), r.Version())

	epoch := uint64(13)
	r.SetAuditEpoch(epoch)
	require.Equal(t, epoch, r.AuditEpoch())

	cid := cidtest.Generate()
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

	passSG := []*object.ID{objecttest.ID(), objecttest.ID()}
	r.SetPassSG(passSG)
	require.Equal(t, passSG, r.PassSG())

	failSG := []*object.ID{objecttest.ID(), objecttest.ID()}
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
	r := audittest.Generate()

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
