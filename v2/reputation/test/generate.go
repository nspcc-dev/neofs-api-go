package reputationtest

import (
	refstest "github.com/nspcc-dev/neofs-api-go/v2/refs/test"
	"github.com/nspcc-dev/neofs-api-go/v2/reputation"
	sessiontest "github.com/nspcc-dev/neofs-api-go/v2/session/test"
)

func GeneratePeerID(empty bool) *reputation.PeerID {
	m := new(reputation.PeerID)

	if !empty {
		m.SetValue([]byte{1, 2, 3})
	}

	return m
}

func GenerateTrust(empty bool) *reputation.Trust {
	m := new(reputation.Trust)

	if !empty {
		m.SetValue(1)
	}

	m.SetPeer(GeneratePeerID(empty))

	return m
}

func GenerateGlobalTrustBody(empty bool) *reputation.GlobalTrustBody {
	m := new(reputation.GlobalTrustBody)

	m.SetManager(GeneratePeerID(empty))
	m.SetTrust(GenerateTrust(empty))

	return m
}

func GenerateGlobalTrust(empty bool) *reputation.GlobalTrust {
	m := new(reputation.GlobalTrust)

	m.SetVersion(refstest.GenerateVersion(empty))
	m.SetBody(GenerateGlobalTrustBody(empty))
	m.SetSignature(refstest.GenerateSignature(empty))

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

func GenerateSendIntermediateResultRequestBody(empty bool) *reputation.SendIntermediateResultRequestBody {
	m := new(reputation.SendIntermediateResultRequestBody)

	if !empty {
		m.SetIteration(564)
	}

	m.SetTrust(GenerateTrust(empty))

	return m
}

func GenerateSendIntermediateResultRequest(empty bool) *reputation.SendIntermediateResultRequest {
	m := new(reputation.SendIntermediateResultRequest)

	m.SetBody(GenerateSendIntermediateResultRequestBody(empty))
	m.SetMetaHeader(sessiontest.GenerateRequestMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateRequestVerificationHeader(empty))

	return m
}

func GenerateSendIntermediateResultResponseBody(empty bool) *reputation.SendIntermediateResultResponseBody {
	m := new(reputation.SendIntermediateResultResponseBody)

	return m
}

func GenerateSendIntermediateResultResponse(empty bool) *reputation.SendIntermediateResultResponse {
	m := new(reputation.SendIntermediateResultResponse)

	m.SetBody(GenerateSendIntermediateResultResponseBody(empty))
	m.SetMetaHeader(sessiontest.GenerateResponseMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateResponseVerificationHeader(empty))

	return m
}
