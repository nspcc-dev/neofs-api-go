package state

// SignedData returns payload bytes of the request.
//
// Always returns empty slice.
func (m NetmapRequest) SignedData() ([]byte, error) {
	return make([]byte, 0), nil
}

// SignedData returns payload bytes of the request.
//
// Always returns empty slice.
func (m MetricsRequest) SignedData() ([]byte, error) {
	return make([]byte, 0), nil
}
