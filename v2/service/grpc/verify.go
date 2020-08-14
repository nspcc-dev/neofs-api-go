package service

// SetKey sets public key in a binary format.
func (m *Signature) SetKey(v []byte) {
	if m != nil {
		m.Key = v
	}
}

// SetSign sets signature.
func (m *Signature) SetSign(v []byte) {
	if m != nil {
		m.Sign = v
	}
}

// SetBodySignature sets signature of the request body.
func (m *RequestVerificationHeader) SetBodySignature(v *Signature) {
	if m != nil {
		m.BodySignature = v
	}
}

// SetMetaSignature sets signature of the request meta.
func (m *RequestVerificationHeader) SetMetaSignature(v *Signature) {
	if m != nil {
		m.MetaSignature = v
	}
}

// SetOriginSignature sets signature of the origin verification header of the request.
func (m *RequestVerificationHeader) SetOriginSignature(v *Signature) {
	if m != nil {
		m.OriginSignature = v
	}
}

// SetOrigin sets origin verification header of the request.
func (m *RequestVerificationHeader) SetOrigin(v *RequestVerificationHeader) {
	if m != nil {
		m.Origin = v
	}
}

// SetBodySignature sets signature of the response body.
func (m *ResponseVerificationHeader) SetBodySignature(v *Signature) {
	if m != nil {
		m.BodySignature = v
	}
}

// SetMetaSignature sets signature of the response meta.
func (m *ResponseVerificationHeader) SetMetaSignature(v *Signature) {
	if m != nil {
		m.MetaSignature = v
	}
}

// SetOriginSignature sets signature of the origin verification header of the response.
func (m *ResponseVerificationHeader) SetOriginSignature(v *Signature) {
	if m != nil {
		m.OriginSignature = v
	}
}

// SetOrigin sets origin verification header of the response.
func (m *ResponseVerificationHeader) SetOrigin(v *ResponseVerificationHeader) {
	if m != nil {
		m.Origin = v
	}
}
