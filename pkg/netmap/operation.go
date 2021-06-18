package netmap

import (
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
)

// Operation is an enumeration of v2-compatible filtering operations.
type Operation uint32

const (
	_ Operation = iota

	// OpEQ is an "Equal" operation.
	OpEQ

	// OpNE is a "Not equal" operation.
	OpNE

	// OpGT is a "Greater than" operation.
	OpGT

	// OpGE is a "Greater than or equal to" operation.
	OpGE

	// OpLT is a "Less than" operation.
	OpLT

	// OpLE is a "Less than or equal to" operation.
	OpLE

	// OpOR is an "OR" operation.
	OpOR

	// OpAND is an "AND" operation.
	OpAND
)

// OperationFromV2 converts v2 Operation to Operation.
func OperationFromV2(op netmap.Operation) Operation {
	switch op {
	default:
		return 0
	case netmap.OR:
		return OpOR
	case netmap.AND:
		return OpAND
	case netmap.GE:
		return OpGE
	case netmap.GT:
		return OpGT
	case netmap.LE:
		return OpLE
	case netmap.LT:
		return OpLT
	case netmap.EQ:
		return OpEQ
	case netmap.NE:
		return OpNE
	}
}

// ToV2 converts Operation to v2 Operation.
func (op Operation) ToV2() netmap.Operation {
	switch op {
	default:
		return netmap.UnspecifiedOperation
	case OpOR:
		return netmap.OR
	case OpAND:
		return netmap.AND
	case OpGE:
		return netmap.GE
	case OpGT:
		return netmap.GT
	case OpLE:
		return netmap.LE
	case OpLT:
		return netmap.LT
	case OpEQ:
		return netmap.EQ
	case OpNE:
		return netmap.NE
	}
}

// String returns string representation of Operation.
//
// String mapping:
//  * OpNE: NE;
//  * OpEQ: EQ;
//  * OpLT: LT;
//  * OpLE: LE;
//  * OpGT: GT;
//  * OpGE: GE;
//  * OpAND: AND;
//  * OpOR: OR;
//  * default: OPERATION_UNSPECIFIED.
func (op Operation) String() string {
	return op.ToV2().String()
}

// FromString parses Operation from a string representation.
// It is a reverse action to String().
//
// Returns true if s was parsed successfully.
func (op *Operation) FromString(s string) bool {
	var g netmap.Operation

	ok := g.FromString(s)

	if ok {
		*op = OperationFromV2(g)
	}

	return ok
}
