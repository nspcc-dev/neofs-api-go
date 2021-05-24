package container_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/rpc/message"
	messagetest "github.com/nspcc-dev/neofs-api-go/rpc/message/test"
	containertest "github.com/nspcc-dev/neofs-api-go/v2/container/test"
)

func TestMessageConvert(t *testing.T) {
	messagetest.TestRPCMessage(t,
		func(empty bool) message.Message { return containertest.GenerateAttribute(empty) },
		func(empty bool) message.Message { return containertest.GenerateContainer(empty) },
		func(empty bool) message.Message { return containertest.GeneratePutRequestBody(empty) },
		func(empty bool) message.Message { return containertest.GeneratePutRequest(empty) },
		func(empty bool) message.Message { return containertest.GeneratePutResponseBody(empty) },
		func(empty bool) message.Message { return containertest.GeneratePutResponse(empty) },
		func(empty bool) message.Message { return containertest.GenerateGetRequestBody(empty) },
		func(empty bool) message.Message { return containertest.GenerateGetRequest(empty) },
		func(empty bool) message.Message { return containertest.GenerateGetResponseBody(empty) },
		func(empty bool) message.Message { return containertest.GenerateGetResponse(empty) },
		func(empty bool) message.Message { return containertest.GenerateDeleteRequestBody(empty) },
		func(empty bool) message.Message { return containertest.GenerateDeleteRequest(empty) },
		func(empty bool) message.Message { return containertest.GenerateDeleteResponseBody(empty) },
		func(empty bool) message.Message { return containertest.GenerateDeleteResponse(empty) },
		func(empty bool) message.Message { return containertest.GenerateListRequestBody(empty) },
		func(empty bool) message.Message { return containertest.GenerateListRequest(empty) },
		func(empty bool) message.Message { return containertest.GenerateListResponseBody(empty) },
		func(empty bool) message.Message { return containertest.GenerateListResponse(empty) },
		func(empty bool) message.Message { return containertest.GenerateSetExtendedACLRequestBody(empty) },
		func(empty bool) message.Message { return containertest.GenerateSetExtendedACLRequest(empty) },
		func(empty bool) message.Message { return containertest.GenerateGetRequestBody(empty) },
		func(empty bool) message.Message { return containertest.GenerateGetRequest(empty) },
		func(empty bool) message.Message { return containertest.GenerateGetResponseBody(empty) },
		func(empty bool) message.Message { return containertest.GenerateGetResponse(empty) },
		func(empty bool) message.Message { return containertest.GenerateGetExtendedACLRequestBody(empty) },
		func(empty bool) message.Message { return containertest.GenerateGetExtendedACLRequest(empty) },
		func(empty bool) message.Message { return containertest.GenerateGetExtendedACLResponseBody(empty) },
		func(empty bool) message.Message { return containertest.GenerateGetExtendedACLResponse(empty) },
		func(empty bool) message.Message { return containertest.GenerateUsedSpaceAnnouncement(empty) },
		func(empty bool) message.Message { return containertest.GenerateAnnounceUsedSpaceRequestBody(empty) },
		func(empty bool) message.Message { return containertest.GenerateAnnounceUsedSpaceRequest(empty) },
		func(empty bool) message.Message { return containertest.GenerateAnnounceUsedSpaceResponseBody(empty) },
		func(empty bool) message.Message { return containertest.GenerateAnnounceUsedSpaceResponse(empty) },
	)
}
