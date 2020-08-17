package session

import (
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/service"
	session "github.com/nspcc-dev/neofs-api-go/v2/session/grpc"
)

func CreateRequestBodyToGRPCMessage(c *CreateRequestBody) *session.CreateRequest_Body {
	if c == nil {
		return nil
	}

	m := new(session.CreateRequest_Body)

	m.SetOwnerId(
		refs.OwnerIDToGRPCMessage(c.GetOwnerID()),
	)

	m.SetLifetime(
		service.TokenLifetimeToGRPCMessage(c.GetLifetime()),
	)

	return m
}

func CreateRequestBodyFromGRPCMessage(m *session.CreateRequest_Body) *CreateRequestBody {
	if m == nil {
		return nil
	}

	c := new(CreateRequestBody)

	c.SetOwnerID(
		refs.OwnerIDFromGRPCMessage(m.GetOwnerId()),
	)

	c.SetLifetime(
		service.TokenLifetimeFromGRPCMessage(m.GetLifetime()),
	)

	return c
}

func CreateRequestToGRPCMessage(c *CreateRequest) *session.CreateRequest {
	if c == nil {
		return nil
	}

	m := new(session.CreateRequest)

	m.SetBody(
		CreateRequestBodyToGRPCMessage(c.GetBody()),
	)

	service.RequestHeadersToGRPC(c, m)

	return m
}

func CreateRequestFromGRPCMessage(m *session.CreateRequest) *CreateRequest {
	if m == nil {
		return nil
	}

	c := new(CreateRequest)

	c.SetBody(
		CreateRequestBodyFromGRPCMessage(m.GetBody()),
	)

	service.RequestHeadersFromGRPC(m, c)

	return c
}

func CreateResponseBodyToGRPCMessage(c *CreateResponseBody) *session.CreateResponse_Body {
	if c == nil {
		return nil
	}

	m := new(session.CreateResponse_Body)

	m.SetId(c.GetID())
	m.SetSessionKey(c.GetSessionKey())

	return m
}

func CreateResponseBodyFromGRPCMessage(m *session.CreateResponse_Body) *CreateResponseBody {
	if m == nil {
		return nil
	}

	c := new(CreateResponseBody)

	c.SetID(m.GetId())
	c.SetSessionKey(m.GetSessionKey())

	return c
}

func CreateResponseToGRPCMessage(c *CreateResponse) *session.CreateResponse {
	if c == nil {
		return nil
	}

	m := new(session.CreateResponse)

	m.SetBody(
		CreateResponseBodyToGRPCMessage(c.GetBody()),
	)

	service.ResponseHeadersToGRPC(c, m)

	return m
}

func CreateResponseFromGRPCMessage(m *session.CreateResponse) *CreateResponse {
	if m == nil {
		return nil
	}

	c := new(CreateResponse)

	c.SetBody(
		CreateResponseBodyFromGRPCMessage(m.GetBody()),
	)

	service.ResponseHeadersFromGRPC(m, c)

	return c
}
