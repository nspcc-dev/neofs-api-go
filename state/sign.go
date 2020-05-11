package state

// SignedData returns payload bytes of the request.
//
// Always returns an empty slice.
func (m NetmapRequest) SignedData() ([]byte, error) {
	return make([]byte, 0), nil
}

// SignedData returns payload bytes of the request.
//
// Always returns an empty slice.
func (m MetricsRequest) SignedData() ([]byte, error) {
	return make([]byte, 0), nil
}

// SignedData returns payload bytes of the request.
//
// Always returns an empty slice.
func (m HealthRequest) SignedData() ([]byte, error) {
	return make([]byte, 0), nil
}

// SignedData returns payload bytes of the request.
//
// Always returns an empty slice.
func (m DumpRequest) SignedData() ([]byte, error) {
	return make([]byte, 0), nil
}
