package audit_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/audit"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/stretchr/testify/require"
)

func TestDataAuditResult_StableMarshal(t *testing.T) {
	from := generateDataAuditResult()

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		to := new(audit.DataAuditResult)
		require.NoError(t, to.Unmarshal(wire))

		require.Equal(t, from, to)
	})
}

func generateDataAuditResult() *audit.DataAuditResult {
	a := new(audit.DataAuditResult)

	v := new(refs.Version)
	v.SetMajor(2)
	v.SetMinor(1)

	oid1 := new(refs.ObjectID)
	oid1.SetValue([]byte("Object ID 1"))

	oid2 := new(refs.ObjectID)
	oid2.SetValue([]byte("Object ID 2"))

	cid := new(refs.ContainerID)
	cid.SetValue([]byte("Container ID"))

	a.SetVersion(v)
	a.SetAuditEpoch(13)
	a.SetContainerID(cid)
	a.SetPublicKey([]byte("Public key"))
	a.SetComplete(true)
	a.SetPassSG([]*refs.ObjectID{oid1, oid2})
	a.SetFailSG([]*refs.ObjectID{oid2, oid1})
	a.SetHit(1)
	a.SetMiss(2)
	a.SetFail(3)
	a.SetPassNodes([][]byte{
		{1, 2},
		{3, 4},
	})
	a.SetFailNodes([][]byte{
		{5, 6},
		{7, 8},
	})

	return a
}
