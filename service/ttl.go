package service

import (
	"github.com/nspcc-dev/neofs-proto/internal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TTLRequest to verify and update ttl requests.
type TTLRequest interface {
	GetTTL() uint32
	SetTTL(uint32)
}

const (
	// ZeroTTL is empty ttl, should produce ErrZeroTTL.
	ZeroTTL = iota

	// NonForwardingTTL is a ttl that allows direct connections only.
	NonForwardingTTL

	// SingleForwardingTTL is a ttl that allows connections through another node.
	SingleForwardingTTL

	// ErrZeroTTL is raised when zero ttl is passed.
	ErrZeroTTL = internal.Error("zero ttl")

	// ErrIncorrectTTL is raised when NonForwardingTTL is passed and NodeRole != InnerRingNode.
	ErrIncorrectTTL = internal.Error("incorrect ttl")
)

// CheckTTLRequest validates and update ttl requests.
func CheckTTLRequest(req TTLRequest, role NodeRole) error {
	var ttl = req.GetTTL()

	if ttl == ZeroTTL {
		return status.New(codes.InvalidArgument, ErrZeroTTL.Error()).Err()
	} else if ttl == NonForwardingTTL && role != InnerRingNode {
		return status.New(codes.InvalidArgument, ErrIncorrectTTL.Error()).Err()
	}

	req.SetTTL(ttl - 1)

	return nil
}
