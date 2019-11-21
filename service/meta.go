package service

import (
	"github.com/nspcc-dev/neofs-proto/internal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	// MetaHeader contains meta information of request.
	// It provides methods to get or set meta information meta header.
	// Also contains methods to reset and restore meta header.
	// Also contains methods to get or set request protocol version
	MetaHeader interface {
		ResetMeta() RequestMetaHeader
		RestoreMeta(RequestMetaHeader)

		// TTLRequest to verify and update ttl requests.
		GetTTL() uint32
		SetTTL(uint32)

		// EpochRequest gives possibility to get or set epoch in RPC Requests.
		GetEpoch() uint64
		SetEpoch(uint64)

		// VersionHeader allows get or set version of protocol request
		VersionHeader
	}

	// VersionHeader allows get or set version of protocol request
	VersionHeader interface {
		GetVersion() uint32
		SetVersion(uint32)
	}

	// TTLCondition is closure, that allows to validate request with ttl.
	TTLCondition func(ttl uint32) error
)

const (
	// ZeroTTL is empty ttl, should produce ErrZeroTTL.
	ZeroTTL = iota

	// NonForwardingTTL is a ttl that allows direct connections only.
	NonForwardingTTL

	// SingleForwardingTTL is a ttl that allows connections through another node.
	SingleForwardingTTL
)

const (
	// ErrZeroTTL is raised when zero ttl is passed.
	ErrZeroTTL = internal.Error("zero ttl")

	// ErrIncorrectTTL is raised when NonForwardingTTL is passed and NodeRole != InnerRingNode.
	ErrIncorrectTTL = internal.Error("incorrect ttl")
)

// SetVersion sets protocol version to RequestMetaHeader.
func (m *RequestMetaHeader) SetVersion(v uint32) { m.Version = v }

// SetTTL sets TTL to RequestMetaHeader.
func (m *RequestMetaHeader) SetTTL(v uint32) { m.TTL = v }

// SetEpoch sets Epoch to RequestMetaHeader.
func (m *RequestMetaHeader) SetEpoch(v uint64) { m.Epoch = v }

// ResetMeta returns current value and sets RequestMetaHeader to empty value.
func (m *RequestMetaHeader) ResetMeta() RequestMetaHeader {
	cp := *m
	m.Reset()
	return cp
}

// RestoreMeta sets current RequestMetaHeader to passed value.
func (m *RequestMetaHeader) RestoreMeta(v RequestMetaHeader) { *m = v }

// IRNonForwarding condition that allows NonForwardingTTL only for IR
func IRNonForwarding(role NodeRole) TTLCondition {
	return func(ttl uint32) error {
		if ttl == NonForwardingTTL && role != InnerRingNode {
			return ErrIncorrectTTL
		}

		return nil
	}
}

// ProcessRequestTTL validates and update ttl requests.
func ProcessRequestTTL(req MetaHeader, cond ...TTLCondition) error {
	var ttl = req.GetTTL()

	if ttl == ZeroTTL {
		return status.New(codes.InvalidArgument, ErrZeroTTL.Error()).Err()
	}

	for i := range cond {
		if cond[i] == nil {
			continue
		}

		// check specific condition:
		if err := cond[i](ttl); err != nil {
			return status.New(codes.InvalidArgument, err.Error()).Err()
		}
	}

	req.SetTTL(ttl - 1)

	return nil
}
