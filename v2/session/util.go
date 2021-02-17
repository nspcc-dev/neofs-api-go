package session

// RequestHeaders represents common part of
// all NeoFS requests including headers.
type RequestHeaders struct {
	metaHeader *RequestMetaHeader

	verifyHeader *RequestVerificationHeader
}

// GetMetaHeader returns meta header of the request.
func (c *RequestHeaders) GetMetaHeader() *RequestMetaHeader {
	if c != nil {
		return c.metaHeader
	}

	return nil
}

// SetMetaHeader sets meta header of the request.
func (c *RequestHeaders) SetMetaHeader(v *RequestMetaHeader) {
	if c != nil {
		c.metaHeader = v
	}
}

// GetVerificationHeader returns verification header of the request.
func (c *RequestHeaders) GetVerificationHeader() *RequestVerificationHeader {
	if c != nil {
		return c.verifyHeader
	}

	return nil
}

// SetVerificationHeader sets verification header of the request.
func (c *RequestHeaders) SetVerificationHeader(v *RequestVerificationHeader) {
	if c != nil {
		c.verifyHeader = v
	}
}

// ResponseHeaders represents common part of
// all NeoFS responses including headers.
type ResponseHeaders struct {
	metaHeader *ResponseMetaHeader

	verifyHeader *ResponseVerificationHeader
}

// GetMetaHeader returns meta header of the response.
func (c *ResponseHeaders) GetMetaHeader() *ResponseMetaHeader {
	if c != nil {
		return c.metaHeader
	}

	return nil
}

// SetMetaHeader sets meta header of the response.
func (c *ResponseHeaders) SetMetaHeader(v *ResponseMetaHeader) {
	if c != nil {
		c.metaHeader = v
	}
}

// GetVerificationHeader returns verification header of the response.
func (c *ResponseHeaders) GetVerificationHeader() *ResponseVerificationHeader {
	if c != nil {
		return c.verifyHeader
	}

	return nil
}

// SetVerificationHeader sets verification header of the response.
func (c *ResponseHeaders) SetVerificationHeader(v *ResponseVerificationHeader) {
	if c != nil {
		c.verifyHeader = v
	}
}
