package reputationtest

import (
	refstest "github.com/nspcc-dev/neofs-api-go/v2/refs/test"
	"github.com/nspcc-dev/neofs-api-go/v2/reputation"
	sessiontest "github.com/nspcc-dev/neofs-api-go/v2/session/test"
)

func GeneratePeerID(empty bool) *reputation.PeerID {
	m := new(reputation.PeerID)

	if !empty {
		m.SetPublicKey([]byte{1, 2, 3})
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

func GeneratePeerToPeerTrust(empty bool) *reputation.PeerToPeerTrust {
	m := new(reputation.PeerToPeerTrust)

	m.SetTrustingPeer(GeneratePeerID(empty))
	m.SetTrust(GenerateTrust(empty))

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

func GenerateAnnounceLocalTrustRequestBody(empty bool) *reputation.AnnounceLocalTrustRequestBody {
	m := new(reputation.AnnounceLocalTrustRequestBody)

	if !empty {
		m.SetEpoch(13)
	}

	m.SetTrusts(GenerateTrusts(empty))

	return m
}

func GenerateAnnounceLocalTrustRequest(empty bool) *reputation.AnnounceLocalTrustRequest {
	m := new(reputation.AnnounceLocalTrustRequest)

	m.SetBody(GenerateAnnounceLocalTrustRequestBody(empty))
	m.SetMetaHeader(sessiontest.GenerateRequestMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateRequestVerificationHeader(empty))

	return m
}

func GenerateAnnounceLocalTrustResponseBody(empty bool) *reputation.AnnounceLocalTrustResponseBody {
	m := new(reputation.AnnounceLocalTrustResponseBody)

	return m
}

func GenerateAnnounceLocalTrustResponse(empty bool) *reputation.AnnounceLocalTrustResponse {
	m := new(reputation.AnnounceLocalTrustResponse)

	m.SetBody(GenerateAnnounceLocalTrustResponseBody(empty))
	m.SetMetaHeader(sessiontest.GenerateResponseMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateResponseVerificationHeader(empty))

	return m
}

func GenerateAnnounceIntermediateResultRequestBody(empty bool) *reputation.AnnounceIntermediateResultRequestBody {
	m := new(reputation.AnnounceIntermediateResultRequestBody)

	if !empty {
		m.SetEpoch(123)
		m.SetIteration(564)
		m.SetTrust(GeneratePeerToPeerTrust(empty))
	}

	return m
}

func GenerateAnnounceIntermediateResultRequest(empty bool) *reputation.AnnounceIntermediateResultRequest {
	m := new(reputation.AnnounceIntermediateResultRequest)

	m.SetBody(GenerateAnnounceIntermediateResultRequestBody(empty))
	m.SetMetaHeader(sessiontest.GenerateRequestMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateRequestVerificationHeader(empty))

	return m
}

func GenerateAnnounceIntermediateResultResponseBody(empty bool) *reputation.AnnounceIntermediateResultResponseBody {
	m := new(reputation.AnnounceIntermediateResultResponseBody)

	return m
}

func GenerateAnnounceIntermediateResultResponse(empty bool) *reputation.AnnounceIntermediateResultResponse {
	m := new(reputation.AnnounceIntermediateResultResponse)

	m.SetBody(GenerateAnnounceIntermediateResultResponseBody(empty))
	m.SetMetaHeader(sessiontest.GenerateResponseMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateResponseVerificationHeader(empty))

	return m
}
