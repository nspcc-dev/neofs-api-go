package refs

import (
	refs "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
)

// String returns string representation of ChecksumType.
func (t ChecksumType) String() string {
	return ChecksumTypeToGRPC(t).String()
}

// FromString parses ChecksumType from a string representation.
// It is a reverse action to String().
//
// Returns true if s was parsed successfully.
func (t *ChecksumType) FromString(s string) bool {
	var g refs.ChecksumType

	ok := g.FromString(s)

	if ok {
		*t = ChecksumTypeFromGRPC(g)
	}

	return ok
}
