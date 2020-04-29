package session

// NewInitRequest returns new initialization CreateRequest from passed Token.
func NewInitRequest(t *Token) *CreateRequest {
	return &CreateRequest{Message: &CreateRequest_Init{Init: t}}
}

// NewSignedRequest returns new signed CreateRequest from passed Token.
func NewSignedRequest(t *Token) *CreateRequest {
	return &CreateRequest{Message: &CreateRequest_Signed{Signed: t}}
}
