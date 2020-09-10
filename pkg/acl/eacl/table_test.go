package eacl_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg/acl/eacl"
	"github.com/nspcc-dev/neofs-crypto/test"
)

func TextExample(t *testing.T) {
	record := eacl.CreateRecord(eacl.ActionDeny, eacl.OperationPut)
	record.AddFilter(eacl.HeaderFromObject, eacl.MatchStringEqual, "filename", "cat.jpg")
	record.AddTarget(eacl.RoleOthers, test.DecodeKey(1).PublicKey, test.DecodeKey(2).PublicKey)

	table := eacl.NewTable()
	table.AddRecord(record)
}
