package v2

import (
	refs "github.com/nspcc-dev/neofs-api-go/refs/v2"
	service "github.com/nspcc-dev/neofs-api-go/service/v2"
)

// SetOwnerId sets identifier of the session initiator.
func (m *CreateRequest_Body) SetOwnerId(v *refs.OwnerID) {
	if m != nil {
		m.OwnerId = v
	}
}

// SetLifetime sets lifetime of the session.
func (m *CreateRequest_Body) SetLifetime(v *service.TokenLifetime) {
	if m != nil {
		m.Lifetime = v
	}
}

// SetBody sets body of the request.
func (m *CreateRequest) SetBody(v *CreateRequest_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the request.
func (m *CreateRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the request.
func (m *CreateRequest) SetVerifyHeader(v *service.RequestVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetId sets identifier of the session token.
func (m *CreateResponse_Body) SetId(v []byte) {
	if m != nil {
		m.Id = v
	}
}

// SetSessionKey sets session public key in a binary format.
func (m *CreateResponse_Body) SetSessionKey(v []byte) {
	if m != nil {
		m.SessionKey = v
	}
}

// SetBody sets body of the response.
func (m *CreateResponse) SetBody(v *CreateResponse_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the response.
func (m *CreateResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the response.
func (m *CreateResponse) SetVerifyHeader(v *service.ResponseVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}
