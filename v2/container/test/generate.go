package containertest

import (
	acltest "github.com/nspcc-dev/neofs-api-go/v2/acl/test"
	"github.com/nspcc-dev/neofs-api-go/v2/container"
	netmaptest "github.com/nspcc-dev/neofs-api-go/v2/netmap/test"
	refstest "github.com/nspcc-dev/neofs-api-go/v2/refs/test"
	sessiontest "github.com/nspcc-dev/neofs-api-go/v2/session/test"
)

func GenerateAttribute(empty bool) *container.Attribute {
	m := new(container.Attribute)

	if !empty {
		m.SetKey("key")
		m.SetValue("val")
	}

	return m
}

func GenerateAttributes(empty bool) (res []*container.Attribute) {
	if !empty {
		res = append(res,
			GenerateAttribute(false),
			GenerateAttribute(false),
		)
	}

	return
}

func GenerateContainer(empty bool) *container.Container {
	m := new(container.Container)

	if !empty {
		m.SetBasicACL(12)
		m.SetNonce([]byte{1, 2, 3})
	}

	m.SetOwnerID(refstest.GenerateOwnerID(empty))
	m.SetVersion(refstest.GenerateVersion(empty))
	m.SetAttributes(GenerateAttributes(empty))
	m.SetPlacementPolicy(netmaptest.GeneratePlacementPolicy(empty))

	return m
}

func GeneratePutRequestBody(empty bool) *container.PutRequestBody {
	m := new(container.PutRequestBody)

	m.SetContainer(GenerateContainer(empty))
	m.SetSignature(refstest.GenerateSignature(empty))

	return m
}

func GeneratePutRequest(empty bool) *container.PutRequest {
	m := new(container.PutRequest)

	m.SetBody(GeneratePutRequestBody(empty))
	m.SetMetaHeader(sessiontest.GenerateRequestMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateRequestVerificationHeader(empty))

	return m
}

func GeneratePutResponseBody(empty bool) *container.PutResponseBody {
	m := new(container.PutResponseBody)

	m.SetContainerID(refstest.GenerateContainerID(empty))

	return m
}

func GeneratePutResponse(empty bool) *container.PutResponse {
	m := new(container.PutResponse)

	m.SetBody(GeneratePutResponseBody(empty))
	m.SetMetaHeader(sessiontest.GenerateResponseMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateResponseVerificationHeader(empty))

	return m
}

func GenerateGetRequestBody(empty bool) *container.GetRequestBody {
	m := new(container.GetRequestBody)

	m.SetContainerID(refstest.GenerateContainerID(empty))

	return m
}

func GenerateGetRequest(empty bool) *container.GetRequest {
	m := new(container.GetRequest)

	m.SetBody(GenerateGetRequestBody(empty))
	m.SetMetaHeader(sessiontest.GenerateRequestMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateRequestVerificationHeader(empty))

	return m
}

func GenerateGetResponseBody(empty bool) *container.GetResponseBody {
	m := new(container.GetResponseBody)

	m.SetContainer(GenerateContainer(empty))
	m.SetSignature(refstest.GenerateSignature(empty))
	m.SetSessionToken(sessiontest.GenerateSessionToken(empty))

	return m
}

func GenerateGetResponse(empty bool) *container.GetResponse {
	m := new(container.GetResponse)

	m.SetBody(GenerateGetResponseBody(empty))
	m.SetMetaHeader(sessiontest.GenerateResponseMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateResponseVerificationHeader(empty))

	return m
}

func GenerateDeleteRequestBody(empty bool) *container.DeleteRequestBody {
	m := new(container.DeleteRequestBody)

	m.SetContainerID(refstest.GenerateContainerID(empty))
	m.SetSignature(refstest.GenerateSignature(empty))

	return m
}

func GenerateDeleteRequest(empty bool) *container.DeleteRequest {
	m := new(container.DeleteRequest)

	m.SetBody(GenerateDeleteRequestBody(empty))
	m.SetMetaHeader(sessiontest.GenerateRequestMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateRequestVerificationHeader(empty))

	return m
}

func GenerateDeleteResponseBody(empty bool) *container.DeleteResponseBody {
	m := new(container.DeleteResponseBody)

	return m
}

func GenerateDeleteResponse(empty bool) *container.DeleteResponse {
	m := new(container.DeleteResponse)

	m.SetBody(GenerateDeleteResponseBody(empty))
	m.SetMetaHeader(sessiontest.GenerateResponseMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateResponseVerificationHeader(empty))

	return m
}

func GenerateListRequestBody(empty bool) *container.ListRequestBody {
	m := new(container.ListRequestBody)

	m.SetOwnerID(refstest.GenerateOwnerID(empty))

	return m
}

func GenerateListRequest(empty bool) *container.ListRequest {
	m := new(container.ListRequest)

	m.SetBody(GenerateListRequestBody(empty))
	m.SetMetaHeader(sessiontest.GenerateRequestMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateRequestVerificationHeader(empty))

	return m
}

func GenerateListResponseBody(empty bool) *container.ListResponseBody {
	m := new(container.ListResponseBody)

	m.SetContainerIDs(refstest.GenerateContainerIDs(empty))

	return m
}

func GenerateListResponse(empty bool) *container.ListResponse {
	m := new(container.ListResponse)

	m.SetBody(GenerateListResponseBody(empty))
	m.SetMetaHeader(sessiontest.GenerateResponseMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateResponseVerificationHeader(empty))

	return m
}

func GenerateSetExtendedACLRequestBody(empty bool) *container.SetExtendedACLRequestBody {
	m := new(container.SetExtendedACLRequestBody)

	m.SetEACL(acltest.GenerateTable(empty))
	m.SetSignature(refstest.GenerateSignature(empty))

	return m
}

func GenerateSetExtendedACLRequest(empty bool) *container.SetExtendedACLRequest {
	m := new(container.SetExtendedACLRequest)

	m.SetBody(GenerateSetExtendedACLRequestBody(empty))
	m.SetMetaHeader(sessiontest.GenerateRequestMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateRequestVerificationHeader(empty))

	return m
}

func GenerateSetExtendedACLResponseBody(empty bool) *container.SetExtendedACLResponseBody {
	m := new(container.SetExtendedACLResponseBody)

	return m
}

func GenerateSetExtendedACLResponse(empty bool) *container.SetExtendedACLResponse {
	m := new(container.SetExtendedACLResponse)

	m.SetBody(GenerateSetExtendedACLResponseBody(empty))
	m.SetMetaHeader(sessiontest.GenerateResponseMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateResponseVerificationHeader(empty))

	return m
}

func GenerateGetExtendedACLRequestBody(empty bool) *container.GetExtendedACLRequestBody {
	m := new(container.GetExtendedACLRequestBody)

	m.SetContainerID(refstest.GenerateContainerID(empty))

	return m
}

func GenerateGetExtendedACLRequest(empty bool) *container.GetExtendedACLRequest {
	m := new(container.GetExtendedACLRequest)

	m.SetBody(GenerateGetExtendedACLRequestBody(empty))
	m.SetMetaHeader(sessiontest.GenerateRequestMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateRequestVerificationHeader(empty))

	return m
}

func GenerateGetExtendedACLResponseBody(empty bool) *container.GetExtendedACLResponseBody {
	m := new(container.GetExtendedACLResponseBody)

	m.SetEACL(acltest.GenerateTable(empty))
	m.SetSignature(refstest.GenerateSignature(empty))
	m.SetSessionToken(sessiontest.GenerateSessionToken(empty))

	return m
}

func GenerateGetExtendedACLResponse(empty bool) *container.GetExtendedACLResponse {
	m := new(container.GetExtendedACLResponse)

	m.SetBody(GenerateGetExtendedACLResponseBody(empty))
	m.SetMetaHeader(sessiontest.GenerateResponseMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateResponseVerificationHeader(empty))

	return m
}

func GenerateUsedSpaceAnnouncement(empty bool) *container.UsedSpaceAnnouncement {
	m := new(container.UsedSpaceAnnouncement)

	m.SetContainerID(refstest.GenerateContainerID(empty))
	m.SetEpoch(1)
	m.SetUsedSpace(2)

	return m
}

func GenerateUsedSpaceAnnouncements(empty bool) (res []*container.UsedSpaceAnnouncement) {
	if !empty {
		res = append(res,
			GenerateUsedSpaceAnnouncement(false),
			GenerateUsedSpaceAnnouncement(false),
		)
	}

	return
}

func GenerateAnnounceUsedSpaceRequestBody(empty bool) *container.AnnounceUsedSpaceRequestBody {
	m := new(container.AnnounceUsedSpaceRequestBody)

	m.SetAnnouncements(GenerateUsedSpaceAnnouncements(empty))

	return m
}

func GenerateAnnounceUsedSpaceRequest(empty bool) *container.AnnounceUsedSpaceRequest {
	m := new(container.AnnounceUsedSpaceRequest)

	m.SetBody(GenerateAnnounceUsedSpaceRequestBody(empty))
	m.SetMetaHeader(sessiontest.GenerateRequestMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateRequestVerificationHeader(empty))

	return m
}

func GenerateAnnounceUsedSpaceResponseBody(empty bool) *container.AnnounceUsedSpaceResponseBody {
	m := new(container.AnnounceUsedSpaceResponseBody)

	return m
}

func GenerateAnnounceUsedSpaceResponse(empty bool) *container.AnnounceUsedSpaceResponse {
	m := new(container.AnnounceUsedSpaceResponse)

	m.SetBody(GenerateAnnounceUsedSpaceResponseBody(empty))
	m.SetMetaHeader(sessiontest.GenerateResponseMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateResponseVerificationHeader(empty))

	return m
}
