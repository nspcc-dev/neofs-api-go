package service

// CutMeta returns current value and sets RequestMetaHeader to empty value.
func (m *RequestMetaHeader) CutMeta() RequestMetaHeader {
	cp := *m
	m.Reset()
	return cp
}

// RestoreMeta sets current RequestMetaHeader to passed value.
func (m *RequestMetaHeader) RestoreMeta(v RequestMetaHeader) {
	*m = v
}
