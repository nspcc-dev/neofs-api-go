package service

type (
	// MetaHeader contains meta information of request.
	// It provides methods to get or set meta information meta header.
	// Also contains methods to reset and restore meta header.
	// Also contains methods to get or set request protocol version
	MetaHeader interface {
		ResetMeta() RequestMetaHeader
		RestoreMeta(RequestMetaHeader)

		TTLContainer
		EpochContainer
		VersionContainer
		RawContainer
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
