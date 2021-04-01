package reputation_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/rpc/message"
	messagetest "github.com/nspcc-dev/neofs-api-go/rpc/message/test"
	reputationtest "github.com/nspcc-dev/neofs-api-go/v2/reputation/test"
)

func TestMessageConvert(t *testing.T) {
	messagetest.TestRPCMessage(t,
		func(empty bool) message.Message { return reputationtest.GenerateTrust(empty) },
		func(empty bool) message.Message { return reputationtest.GenerateSendLocalTrustRequestBody(empty) },
		func(empty bool) message.Message { return reputationtest.GenerateSendLocalTrustRequest(empty) },
		func(empty bool) message.Message { return reputationtest.GenerateSendLocalTrustResponseBody(empty) },
		func(empty bool) message.Message { return reputationtest.GenerateSendLocalTrustResponse(empty) },
		func(empty bool) message.Message {
			return reputationtest.GenerateSendIntermediateResultRequestBody(empty)
		},
		func(empty bool) message.Message { return reputationtest.GenerateSendIntermediateResultRequest(empty) },
		func(empty bool) message.Message {
			return reputationtest.GenerateSendIntermediateResultResponseBody(empty)
		},
		func(empty bool) message.Message { return reputationtest.GenerateSendIntermediateResultResponse(empty) },
		func(empty bool) message.Message { return reputationtest.GenerateGlobalTrustBody(empty) },
		func(empty bool) message.Message { return reputationtest.GenerateGlobalTrust(empty) },
	)
}
