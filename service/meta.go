package service

import (
	"github.com/pkg/errors"
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

		// EpochHeader gives possibility to get or set epoch in RPC Requests.
		EpochHeader

		// VersionHeader allows get or set version of protocol request
		VersionHeader

		// RawHeader allows to get and set raw option of request
		RawHeader
	}

	// EpochHeader interface gives possibility to get or set epoch in RPC Requests.
	EpochHeader interface {
		GetEpoch() uint64
		SetEpoch(v uint64)
	}

	// VersionHeader allows get or set version of protocol request
	VersionHeader interface {
		GetVersion() uint32
		SetVersion(uint32)
	}

	// RawHeader is an interface of the container of a boolean Raw value
	RawHeader interface {
		GetRaw() bool
		SetRaw(bool)
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

// SetVersion sets protocol version to ResponseMetaHeader.
func (m *ResponseMetaHeader) SetVersion(v uint32) { m.Version = v }

// SetEpoch sets Epoch to ResponseMetaHeader.
func (m *ResponseMetaHeader) SetEpoch(v uint64) { m.Epoch = v }

// SetVersion sets protocol version to RequestMetaHeader.
func (m *RequestMetaHeader) SetVersion(v uint32) { m.Version = v }

// SetTTL sets TTL to RequestMetaHeader.
func (m *RequestMetaHeader) SetTTL(v uint32) { m.TTL = v }

// SetEpoch sets Epoch to RequestMetaHeader.
func (m *RequestMetaHeader) SetEpoch(v uint64) { m.Epoch = v }

// SetRaw is a Raw field setter.
func (m *RequestMetaHeader) SetRaw(raw bool) {
	m.Raw = raw
}

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
			return ErrInvalidTTL
		}

		return nil
	}
}

// ProcessRequestTTL validates and update ttl requests.
func ProcessRequestTTL(req MetaHeader, cond ...TTLCondition) error {
	ttl := req.GetTTL()

	if ttl == ZeroTTL {
		return status.New(codes.InvalidArgument, ErrInvalidTTL.Error()).Err()
	}

	for i := range cond {
		if cond[i] == nil {
			continue
		}

		// check specific condition:
		if err := cond[i](ttl); err != nil {
			if st, ok := status.FromError(errors.Cause(err)); ok {
				return st.Err()
			}

			return status.New(codes.InvalidArgument, err.Error()).Err()
		}
	}

	req.SetTTL(ttl - 1)

	return nil
}
