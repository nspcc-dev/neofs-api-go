package service

type (
	// MetaHeader contains meta information of request.
	// It provides methods to get or set meta information meta header.
	// Also contains methods to reset and restore meta header.
	// Also contains methods to get or set request protocol version
	MetaHeader interface {
		ResetMeta() RequestMetaHeader
		RestoreMeta(RequestMetaHeader)

		// TTLHeader allows to get and set TTL value of request.
		TTLHeader

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
)

// SetVersion sets protocol version to ResponseMetaHeader.
func (m *ResponseMetaHeader) SetVersion(v uint32) { m.Version = v }

// SetEpoch sets Epoch to ResponseMetaHeader.
func (m *ResponseMetaHeader) SetEpoch(v uint64) { m.Epoch = v }

// SetVersion sets protocol version to RequestMetaHeader.
func (m *RequestMetaHeader) SetVersion(v uint32) { m.Version = v }

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
