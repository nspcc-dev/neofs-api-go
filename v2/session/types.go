package session

import (
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/service"
)

type CreateRequestBody struct {
	ownerID *refs.OwnerID

	lifetime *service.TokenLifetime
}

type CreateRequest struct {
	body *CreateRequestBody

	metaHeader *service.RequestMetaHeader

	verifyHeader *service.RequestVerificationHeader
}

type CreateResponseBody struct {
	id []byte

	sessionKey []byte
}

type CreateResponse struct {
	body *CreateResponseBody

	metaHeader *service.ResponseMetaHeader

	verifyHeader *service.ResponseVerificationHeader
}

func (c *CreateRequestBody) GetOwnerID() *refs.OwnerID {
	if c != nil {
		return c.ownerID
	}

	return nil
}

func (c *CreateRequestBody) SetOwnerID(v *refs.OwnerID) {
	if c != nil {
		c.ownerID = v
	}
}

func (c *CreateRequestBody) GetLifetime() *service.TokenLifetime {
	if c != nil {
		return c.lifetime
	}

	return nil
}

func (c *CreateRequestBody) SetLifetime(v *service.TokenLifetime) {
	if c != nil {
		c.lifetime = v
	}
}

func (c *CreateRequest) GetBody() *CreateRequestBody {
	if c != nil {
		return c.body
	}

	return nil
}

func (c *CreateRequest) SetBody(v *CreateRequestBody) {
	if c != nil {
		c.body = v
	}
}

func (c *CreateRequest) GetMetaHeader() *service.RequestMetaHeader {
	if c != nil {
		return c.metaHeader
	}

	return nil
}

func (c *CreateRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if c != nil {
		c.metaHeader = v
	}
}

func (c *CreateRequest) GetVerificationHeader() *service.RequestVerificationHeader {
	if c != nil {
		return c.verifyHeader
	}

	return nil
}

func (c *CreateRequest) SetVerificationHeader(v *service.RequestVerificationHeader) {
	if c != nil {
		c.verifyHeader = v
	}
}

func (c *CreateResponseBody) GetID() []byte {
	if c != nil {
		return c.id
	}

	return nil
}

func (c *CreateResponseBody) SetID(v []byte) {
	if c != nil {
		c.id = v
	}
}

func (c *CreateResponseBody) GetSessionKey() []byte {
	if c != nil {
		return c.sessionKey
	}

	return nil
}

func (c *CreateResponseBody) SetSessionKey(v []byte) {
	if c != nil {
		c.sessionKey = v
	}
}

func (c *CreateResponse) GetBody() *CreateResponseBody {
	if c != nil {
		return c.body
	}

	return nil
}

func (c *CreateResponse) SetBody(v *CreateResponseBody) {
	if c != nil {
		c.body = v
	}
}

func (c *CreateResponse) GetMetaHeader() *service.ResponseMetaHeader {
	if c != nil {
		return c.metaHeader
	}

	return nil
}

func (c *CreateResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if c != nil {
		c.metaHeader = v
	}
}

func (c *CreateResponse) GetVerificationHeader() *service.ResponseVerificationHeader {
	if c != nil {
		return c.verifyHeader
	}

	return nil
}

func (c *CreateResponse) SetVerificationHeader(v *service.ResponseVerificationHeader) {
	if c != nil {
		c.verifyHeader = v
	}
}
