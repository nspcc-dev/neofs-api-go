package sessiontest

import (
	acltest "github.com/nspcc-dev/neofs-api-go/v2/acl/test"
	refstest "github.com/nspcc-dev/neofs-api-go/v2/refs/test"
	"github.com/nspcc-dev/neofs-api-go/v2/session"
)

func GenerateCreateRequestBody(empty bool) *session.CreateRequestBody {
	m := new(session.CreateRequestBody)

	if !empty {
		m.SetExpiration(555)
	}

	m.SetOwnerID(refstest.GenerateOwnerID(empty))

	return m
}

func GenerateCreateRequest(empty bool) *session.CreateRequest {
	m := new(session.CreateRequest)

	m.SetBody(GenerateCreateRequestBody(empty))
	m.SetMetaHeader(GenerateRequestMetaHeader(empty))
	m.SetVerificationHeader(GenerateRequestVerificationHeader(empty))

	return m
}

func GenerateCreateResponseBody(empty bool) *session.CreateResponseBody {
	m := new(session.CreateResponseBody)

	if !empty {
		m.SetID([]byte{1, 2, 3})
		m.SetSessionKey([]byte{4, 5, 6})
	}

	return m
}

func GenerateCreateResponse(empty bool) *session.CreateResponse {
	m := new(session.CreateResponse)

	m.SetBody(GenerateCreateResponseBody(empty))
	m.SetMetaHeader(GenerateResponseMetaHeader(empty))
	m.SetVerificationHeader(GenerateResponseVerificationHeader(empty))

	return m
}

func GenerateResponseVerificationHeader(empty bool) *session.ResponseVerificationHeader {
	return generateResponseVerificationHeader(empty, true)
}

func generateResponseVerificationHeader(empty, withOrigin bool) *session.ResponseVerificationHeader {
	m := new(session.ResponseVerificationHeader)

	m.SetBodySignature(refstest.GenerateSignature(empty))
	m.SetMetaSignature(refstest.GenerateSignature(empty))
	m.SetOriginSignature(refstest.GenerateSignature(empty))

	if withOrigin {
		m.SetOrigin(generateResponseVerificationHeader(empty, false))
	}

	return m
}

func GenerateResponseMetaHeader(empty bool) *session.ResponseMetaHeader {
	return generateResponseMetaHeader(empty, true)
}

func generateResponseMetaHeader(empty, withOrigin bool) *session.ResponseMetaHeader {
	m := new(session.ResponseMetaHeader)

	if !empty {
		m.SetEpoch(13)
		m.SetTTL(100)
	}

	m.SetXHeaders(GenerateXHeaders(empty))
	m.SetVersion(refstest.GenerateVersion(empty))

	if withOrigin {
		m.SetOrigin(generateResponseMetaHeader(empty, false))
	}

	return m
}

func GenerateRequestVerificationHeader(empty bool) *session.RequestVerificationHeader {
	return generateRequestVerificationHeader(empty, true)
}

func generateRequestVerificationHeader(empty, withOrigin bool) *session.RequestVerificationHeader {
	m := new(session.RequestVerificationHeader)

	m.SetBodySignature(refstest.GenerateSignature(empty))
	m.SetMetaSignature(refstest.GenerateSignature(empty))
	m.SetOriginSignature(refstest.GenerateSignature(empty))

	if withOrigin {
		m.SetOrigin(generateRequestVerificationHeader(empty, false))
	}

	return m
}

func GenerateRequestMetaHeader(empty bool) *session.RequestMetaHeader {
	return generateRequestMetaHeader(empty, true)
}

func generateRequestMetaHeader(empty, withOrigin bool) *session.RequestMetaHeader {
	m := new(session.RequestMetaHeader)

	if !empty {
		m.SetEpoch(13)
		m.SetTTL(100)
	}

	m.SetXHeaders(GenerateXHeaders(empty))
	m.SetVersion(refstest.GenerateVersion(empty))
	m.SetSessionToken(GenerateSessionToken(empty))
	m.SetBearerToken(acltest.GenerateBearerToken(empty))

	if withOrigin {
		m.SetOrigin(generateRequestMetaHeader(empty, false))
	}

	return m
}

func GenerateSessionToken(empty bool) *session.SessionToken {
	m := new(session.SessionToken)

	m.SetBody(GenerateSessionTokenBody(empty))
	m.SetSignature(refstest.GenerateSignature(empty))

	return m
}

func GenerateSessionTokenBody(empty bool) *session.SessionTokenBody {
	m := new(session.SessionTokenBody)

	if !empty {
		m.SetID([]byte{1})
		m.SetSessionKey([]byte{2})
	}

	m.SetOwnerID(refstest.GenerateOwnerID(empty))
	m.SetLifetime(GenerateTokenLifetime(empty))
	m.SetContext(GenerateObjectSessionContext(empty))

	return m
}

func GenerateTokenLifetime(empty bool) *session.TokenLifetime {
	m := new(session.TokenLifetime)

	if !empty {
		m.SetExp(1)
		m.SetIat(2)
		m.SetExp(3)
	}

	return m
}

func GenerateObjectSessionContext(empty bool) *session.ObjectSessionContext {
	m := new(session.ObjectSessionContext)

	if !empty {
		m.SetVerb(session.ObjectVerbHead)
	}

	m.SetAddress(refstest.GenerateAddress(empty))

	return m
}

func GenerateXHeader(empty bool) *session.XHeader {
	m := new(session.XHeader)

	if !empty {
		m.SetKey("key")
		m.SetValue("val")
	}

	return m
}

func GenerateXHeaders(empty bool) []*session.XHeader {
	xs := make([]*session.XHeader, 0)

	if !empty {
		xs = append(xs,
			GenerateXHeader(false),
			GenerateXHeader(false),
		)
	}

	return xs
}
