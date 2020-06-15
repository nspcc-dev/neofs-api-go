package service

import (
	"io"
)

type extHdrWrapper struct {
	extHdr *RequestExtendedHeader_KV
}

type extHdrSrcWrapper struct {
	extHdrSrc ExtendedHeadersSource
}

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

// SetK is a K field setter.
func (m *RequestExtendedHeader_KV) SetK(v string) {
	m.K = v
}

// SetV is a V field setter.
func (m *RequestExtendedHeader_KV) SetV(v string) {
	m.V = v
}

// Key returns the result of K field getter.
func (m extHdrWrapper) Key() string {
	return m.extHdr.GetK()
}

// Value returns the result of V field getter.
func (m extHdrWrapper) Value() string {
	return m.extHdr.GetV()
}

// SetHeaders is a Headers field setter.
func (m *RequestExtendedHeader) SetHeaders(v []RequestExtendedHeader_KV) {
	m.Headers = v
}

// ExtendedHeaders composes ExtendedHeader list from the Headers field getter result.
func (m RequestExtendedHeader) ExtendedHeaders() []ExtendedHeader {
	hs := m.GetHeaders()

	res := make([]ExtendedHeader, 0, len(hs))

	for i := range hs {
		res = append(res, extHdrWrapper{
			extHdr: &hs[i],
		})
	}

	return res
}

// ExtendedHeadersSignedData wraps passed ExtendedHeadersSource and returns SignedDataSource.
func ExtendedHeadersSignedData(headers ExtendedHeadersSource) SignedDataSource {
	return &extHdrSrcWrapper{
		extHdrSrc: headers,
	}
}

// SignedData returns extended headers in a binary representation.
func (s extHdrSrcWrapper) SignedData() ([]byte, error) {
	return SignedDataFromReader(s)
}

// SignedDataSize returns the length of extended headers slice.
func (s extHdrSrcWrapper) SignedDataSize() (res int) {
	if s.extHdrSrc != nil {
		for _, h := range s.extHdrSrc.ExtendedHeaders() {
			if h != nil {
				res += len(h.Key()) + len(h.Value())
			}
		}
	}

	return
}

// ReadSignedData copies a binary representation of the extended headers to passed buffer.
//
// If buffer length is less than required, io.ErrUnexpectedEOF returns.
func (s extHdrSrcWrapper) ReadSignedData(p []byte) (int, error) {
	sz := s.SignedDataSize()
	if len(p) < sz {
		return 0, io.ErrUnexpectedEOF
	}

	if s.extHdrSrc != nil {
		off := 0
		for _, h := range s.extHdrSrc.ExtendedHeaders() {
			if h == nil {
				continue
			}

			off += copy(p[off:], []byte(h.Key()))

			off += copy(p[off:], []byte(h.Value()))
		}
	}

	return sz, nil
}
