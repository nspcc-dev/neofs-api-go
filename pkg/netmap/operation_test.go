package netmap

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	"github.com/stretchr/testify/require"
)

func TestOperationFromV2(t *testing.T) {
	for _, item := range []struct {
		op   Operation
		opV2 netmap.Operation
	}{
		{
			op:   0,
			opV2: netmap.UnspecifiedOperation,
		},
		{
			op:   OpEQ,
			opV2: netmap.EQ,
		},
		{
			op:   OpNE,
			opV2: netmap.NE,
		},
		{
			op:   OpOR,
			opV2: netmap.OR,
		},
		{
			op:   OpAND,
			opV2: netmap.AND,
		},
		{
			op:   OpLE,
			opV2: netmap.LE,
		},
		{
			op:   OpLT,
			opV2: netmap.LT,
		},
		{
			op:   OpGT,
			opV2: netmap.GT,
		},
		{
			op:   OpGE,
			opV2: netmap.GE,
		},
	} {
		require.Equal(t, item.op, OperationFromV2(item.opV2))
		require.Equal(t, item.opV2, item.op.ToV2())
	}
}

func TestOperation_String(t *testing.T) {
	toPtr := func(v Operation) *Operation {
		return &v
	}

	testEnumStrings(t, new(Operation), []enumStringItem{
		{val: toPtr(OpEQ), str: "EQ"},
		{val: toPtr(OpNE), str: "NE"},
		{val: toPtr(OpGT), str: "GT"},
		{val: toPtr(OpGE), str: "GE"},
		{val: toPtr(OpLT), str: "LT"},
		{val: toPtr(OpLE), str: "LE"},
		{val: toPtr(OpAND), str: "AND"},
		{val: toPtr(OpOR), str: "OR"},
		{val: toPtr(0), str: "OPERATION_UNSPECIFIED"},
	})
}
