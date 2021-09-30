package netmap

import (
	"errors"

	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
)

// Operation is an enumeration of v2-compatible filtering operations.
type Operation uint32

const (
	// OpUnknown is a Operation value used to mark header type as undefined.
	OpUnknown Operation = iota

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
	if op2, ok := operationToV2(op); ok {
		return op2
	}

	return netmap.UnspecifiedOperation
}

// converts Operation to v2 Operation enum value. Returns false if value is not a named constant.
func operationToV2(op Operation) (netmap.Operation, bool) {
	switch op {
	default:
		return 0, false
	case OpUnknown:
		return netmap.UnspecifiedOperation, true
	case OpOR:
		return netmap.OR, true
	case OpAND:
		return netmap.AND, true
	case OpGE:
		return netmap.GE, true
	case OpGT:
		return netmap.GT, true
	case OpLE:
		return netmap.LE, true
	case OpLT:
		return netmap.LT, true
	case OpEQ:
		return netmap.EQ, true
	case OpNE:
		return netmap.NE, true
	}
}

// String implements fmt.Stringer.
//
// Use MarshalText to get the canonical text format.
func (op Operation) String() string {
	// TODO: simplify stringer after FromString will be removed (neofs-api-go#346)
	txt, _ := op.MarshalText()
	return string(txt)
}

var errUnsupportedOp = errors.New("unsupported Operation")

// MarshalText implements encoding.TextMarshaler.
//
// Text mapping:
//  * OpNE: NE;
//  * OpEQ: EQ;
//  * OpLT: LT;
//  * OpLE: LE;
//  * OpGT: GT;
//  * OpGE: GE;
//  * OpAND: AND;
//  * OpOR: OR;
//  * OpUnknown: OPERATION_UNSPECIFIED.
func (op Operation) MarshalText() ([]byte, error) {
	op2, ok := operationToV2(op)
	if !ok {
		return nil, errUnsupportedOp
	}

	return []byte(op2.String()), nil
}

func (op *Operation) UnmarshalText(text []byte) error {
	var op2 netmap.Operation

	ok := op2.FromString(string(text))
	if !ok {
		return errUnsupportedOp
	}

	*op = OperationFromV2(op2)

	return nil
}

// FromString parses Operation from a string representation.
// It is a reverse action to String().
//
// Returns true if s was parsed successfully.
//
// Deprecated: use UnmarshalText instead.
func (op *Operation) FromString(s string) bool {
	return op.UnmarshalText([]byte(s)) == nil
}
