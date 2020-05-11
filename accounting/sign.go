package accounting

import "io"

// SignedData returns payload bytes of the request.
func (m BalanceRequest) SignedData() ([]byte, error) {
	data := make([]byte, m.SignedDataSize())

	if _, err := m.ReadSignedData(data); err != nil {
		return nil, err
	}

	return data, nil
}

// SignedDataSize returns payload size of the request.
func (m BalanceRequest) SignedDataSize() int {
	return m.GetOwnerID().Size()
}

// ReadSignedData copies payload bytes to passed buffer.
//
// If the buffer size is insufficient, io.ErrUnexpectedEOF returns.
func (m BalanceRequest) ReadSignedData(p []byte) (int, error) {
	sz := m.SignedDataSize()
	if len(p) < sz {
		return 0, io.ErrUnexpectedEOF
	}

	copy(p, m.GetOwnerID().Bytes())

	return sz, nil
}
