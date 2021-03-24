package reputationtest

import (
	"github.com/nspcc-dev/neofs-api-go/v2/reputation"
	sessiontest "github.com/nspcc-dev/neofs-api-go/v2/session/test"
)

func GenerateTrust(empty bool) *reputation.Trust {
	m := new(reputation.Trust)

	if !empty {
		m.SetPeer([]byte{1, 2, 3})
		m.SetValue(1)
	}

	return m
}

func GenerateTrusts(empty bool) (res []*reputation.Trust) {
	if !empty {
		res = append(res,
			GenerateTrust(false),
			GenerateTrust(false),
		)
	}

	return
}

func GenerateSendLocalTrustRequestBody(empty bool) *reputation.SendLocalTrustRequestBody {
	m := new(reputation.SendLocalTrustRequestBody)

	if !empty {
		m.SetEpoch(13)
	}

	m.SetTrusts(GenerateTrusts(empty))

	return m
}

func GenerateSendLocalTrustRequest(empty bool) *reputation.SendLocalTrustRequest {
	m := new(reputation.SendLocalTrustRequest)

	m.SetBody(GenerateSendLocalTrustRequestBody(empty))
	m.SetMetaHeader(sessiontest.GenerateRequestMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateRequestVerificationHeader(empty))

	return m
}

func GenerateSendLocalTrustResponseBody(empty bool) *reputation.SendLocalTrustResponseBody {
	m := new(reputation.SendLocalTrustResponseBody)

	return m
}

func GenerateSendLocalTrustResponse(empty bool) *reputation.SendLocalTrustResponse {
	m := new(reputation.SendLocalTrustResponse)

	m.SetBody(GenerateSendLocalTrustResponseBody(empty))
	m.SetMetaHeader(sessiontest.GenerateResponseMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateResponseVerificationHeader(empty))

	return m
}
