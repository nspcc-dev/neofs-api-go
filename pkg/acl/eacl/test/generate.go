package eacltest

import (
	"github.com/nspcc-dev/neofs-api-go/pkg/acl/eacl"
	cidtest "github.com/nspcc-dev/neofs-api-go/pkg/container/id/test"
	ownertest "github.com/nspcc-dev/neofs-api-go/pkg/owner/test"
)

// Target returns random eacl.Target.
func Target() *eacl.Target {
	x := eacl.NewTarget()

	x.SetRole(eacl.RoleSystem)
	x.SetBinaryKeys([][]byte{
		{1, 2, 3},
		{4, 5, 6},
	})

	return x
}

// Record returns random eacl.Record.
func Record() *eacl.Record {
	x := eacl.NewRecord()

	x.SetAction(eacl.ActionAllow)
	x.SetOperation(eacl.OperationRangeHash)
	x.SetTargets(Target(), Target())
	x.AddObjectContainerIDFilter(eacl.MatchStringEqual, cidtest.Generate())
	x.AddObjectOwnerIDFilter(eacl.MatchStringNotEqual, ownertest.Generate())

	return x
}
