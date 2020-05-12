package service

// SignedDataFromReader allocates buffer and reads bytes from passed reader to it.
//
// If passed SignedDataReader is nil, ErrNilSignedDataReader returns.
func SignedDataFromReader(r SignedDataReader) ([]byte, error) {
	if r == nil {
		return nil, ErrNilSignedDataReader
	}

	data := make([]byte, r.SignedDataSize())

	if _, err := r.ReadSignedData(data); err != nil {
		return nil, err
	}

	return data, nil
}
