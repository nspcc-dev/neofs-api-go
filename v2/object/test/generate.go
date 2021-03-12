package objecttest

import (
	"github.com/nspcc-dev/neofs-api-go/v2/object"
	refstest "github.com/nspcc-dev/neofs-api-go/v2/refs/test"
	sessiontest "github.com/nspcc-dev/neofs-api-go/v2/session/test"
)

func GenerateShortHeader(empty bool) *object.ShortHeader {
	m := new(object.ShortHeader)

	if !empty {
		m.SetObjectType(13)
		m.SetCreationEpoch(100)
		m.SetPayloadLength(12321)
	}

	m.SetVersion(refstest.GenerateVersion(empty))
	m.SetOwnerID(refstest.GenerateOwnerID(empty))
	m.SetHomomorphicHash(refstest.GenerateChecksum(empty))
	m.SetPayloadHash(refstest.GenerateChecksum(empty))

	return m
}

func GenerateAttribute(empty bool) *object.Attribute {
	m := new(object.Attribute)

	if !empty {
		m.SetKey("object key")
		m.SetValue("object value")
	}

	return m
}

func GenerateAttributes(empty bool) (res []*object.Attribute) {
	if !empty {
		res = append(res,
			GenerateAttribute(false),
			GenerateAttribute(false),
		)
	}

	return
}

func GenerateSplitHeader(empty bool) *object.SplitHeader {
	return generateSplitHeader(empty, true)
}

func generateSplitHeader(empty, withPar bool) *object.SplitHeader {
	m := new(object.SplitHeader)

	if !empty {
		m.SetSplitID([]byte{1, 3, 5})
	}

	m.SetParent(refstest.GenerateObjectID(empty))
	m.SetPrevious(refstest.GenerateObjectID(empty))
	m.SetParentSignature(refstest.GenerateSignature(empty))
	m.SetChildren(refstest.GenerateObjectIDs(empty))

	if withPar {
		m.SetParentHeader(generateHeader(empty, false))
	}

	return m
}

func GenerateHeader(empty bool) *object.Header {
	return generateHeader(empty, true)
}

func generateHeader(empty, withSplit bool) *object.Header {
	m := new(object.Header)

	if !empty {
		m.SetPayloadLength(777)
		m.SetCreationEpoch(432)
		m.SetObjectType(111)
	}

	m.SetVersion(refstest.GenerateVersion(empty))
	m.SetPayloadHash(refstest.GenerateChecksum(empty))
	m.SetOwnerID(refstest.GenerateOwnerID(empty))
	m.SetHomomorphicHash(refstest.GenerateChecksum(empty))
	m.SetContainerID(refstest.GenerateContainerID(empty))
	m.SetSessionToken(sessiontest.GenerateSessionToken(empty))
	m.SetAttributes(GenerateAttributes(empty))

	if withSplit {
		m.SetSplit(generateSplitHeader(empty, false))
	}

	return m
}

func GenerateHeaderWithSignature(empty bool) *object.HeaderWithSignature {
	m := new(object.HeaderWithSignature)

	m.SetSignature(refstest.GenerateSignature(empty))
	m.SetHeader(GenerateHeader(empty))

	return m
}

func GenerateObject(empty bool) *object.Object {
	m := new(object.Object)

	if !empty {
		m.SetPayload([]byte{7, 8, 9})
	}

	m.SetObjectID(refstest.GenerateObjectID(empty))
	m.SetSignature(refstest.GenerateSignature(empty))
	m.SetHeader(GenerateHeader(empty))

	return m
}

func GenerateSplitInfo(empty bool) *object.SplitInfo {
	m := new(object.SplitInfo)

	if !empty {
		m.SetSplitID([]byte("splitID"))
	}

	m.SetLastPart(refstest.GenerateObjectID(empty))
	m.SetLink(refstest.GenerateObjectID(empty))

	return m
}

func GenerateGetRequestBody(empty bool) *object.GetRequestBody {
	m := new(object.GetRequestBody)

	if !empty {
		m.SetRaw(true)
	}

	m.SetAddress(refstest.GenerateAddress(empty))

	return m
}

func GenerateGetRequest(empty bool) *object.GetRequest {
	m := new(object.GetRequest)

	m.SetBody(GenerateGetRequestBody(empty))
	m.SetMetaHeader(sessiontest.GenerateRequestMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateRequestVerificationHeader(empty))

	return m
}

func GenerateGetObjectPartInit(empty bool) *object.GetObjectPartInit {
	m := new(object.GetObjectPartInit)

	m.SetObjectID(refstest.GenerateObjectID(empty))
	m.SetSignature(refstest.GenerateSignature(empty))
	m.SetHeader(GenerateHeader(empty))

	return m
}

func GenerateGetObjectPartChunk(empty bool) *object.GetObjectPartChunk {
	m := new(object.GetObjectPartChunk)

	if !empty {
		m.SetChunk([]byte("get chunk"))
	}

	return m
}

func GenerateGetResponseBody(empty bool) *object.GetResponseBody {
	m := new(object.GetResponseBody)

	m.SetObjectPart(GenerateGetObjectPartInit(empty))

	return m
}

func GenerateGetResponse(empty bool) *object.GetResponse {
	m := new(object.GetResponse)

	m.SetBody(GenerateGetResponseBody(empty))
	m.SetMetaHeader(sessiontest.GenerateResponseMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateResponseVerificationHeader(empty))

	return m
}

func GeneratePutObjectPartInit(empty bool) *object.PutObjectPartInit {
	m := new(object.PutObjectPartInit)

	if !empty {
		m.SetCopiesNumber(234)
	}

	m.SetObjectID(refstest.GenerateObjectID(empty))
	m.SetSignature(refstest.GenerateSignature(empty))
	m.SetHeader(GenerateHeader(empty))

	return m
}

func GeneratePutObjectPartChunk(empty bool) *object.PutObjectPartChunk {
	m := new(object.PutObjectPartChunk)

	if !empty {
		m.SetChunk([]byte("put chunk"))
	}

	return m
}

func GeneratePutRequestBody(empty bool) *object.PutRequestBody {
	m := new(object.PutRequestBody)

	m.SetObjectPart(GeneratePutObjectPartInit(empty))

	return m
}

func GeneratePutRequest(empty bool) *object.PutRequest {
	m := new(object.PutRequest)

	m.SetBody(GeneratePutRequestBody(empty))
	m.SetMetaHeader(sessiontest.GenerateRequestMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateRequestVerificationHeader(empty))

	return m
}

func GeneratePutResponseBody(empty bool) *object.PutResponseBody {
	m := new(object.PutResponseBody)

	m.SetObjectID(refstest.GenerateObjectID(empty))

	return m
}

func GeneratePutResponse(empty bool) *object.PutResponse {
	m := new(object.PutResponse)

	m.SetBody(GeneratePutResponseBody(empty))
	m.SetMetaHeader(sessiontest.GenerateResponseMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateResponseVerificationHeader(empty))

	return m
}

func GenerateDeleteRequestBody(empty bool) *object.DeleteRequestBody {
	m := new(object.DeleteRequestBody)

	m.SetAddress(refstest.GenerateAddress(empty))

	return m
}

func GenerateDeleteRequest(empty bool) *object.DeleteRequest {
	m := new(object.DeleteRequest)

	m.SetBody(GenerateDeleteRequestBody(empty))
	m.SetMetaHeader(sessiontest.GenerateRequestMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateRequestVerificationHeader(empty))

	return m
}

func GenerateDeleteResponseBody(empty bool) *object.DeleteResponseBody {
	m := new(object.DeleteResponseBody)

	m.SetTombstone(refstest.GenerateAddress(empty))

	return m
}

func GenerateDeleteResponse(empty bool) *object.DeleteResponse {
	m := new(object.DeleteResponse)

	m.SetBody(GenerateDeleteResponseBody(empty))
	m.SetMetaHeader(sessiontest.GenerateResponseMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateResponseVerificationHeader(empty))

	return m
}

func GenerateHeadRequestBody(empty bool) *object.HeadRequestBody {
	m := new(object.HeadRequestBody)

	if !empty {
		m.SetRaw(true)
		m.SetMainOnly(true)
	}

	m.SetAddress(refstest.GenerateAddress(empty))

	return m
}

func GenerateHeadRequest(empty bool) *object.HeadRequest {
	m := new(object.HeadRequest)

	m.SetBody(GenerateHeadRequestBody(empty))
	m.SetMetaHeader(sessiontest.GenerateRequestMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateRequestVerificationHeader(empty))

	return m
}

func GenerateHeadResponseBody(empty bool) *object.HeadResponseBody {
	m := new(object.HeadResponseBody)

	m.SetHeaderPart(GenerateHeaderWithSignature(empty))

	return m
}

func GenerateHeadResponse(empty bool) *object.HeadResponse {
	m := new(object.HeadResponse)

	m.SetBody(GenerateHeadResponseBody(empty))
	m.SetMetaHeader(sessiontest.GenerateResponseMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateResponseVerificationHeader(empty))

	return m
}

func GenerateSearchFilter(empty bool) *object.SearchFilter {
	m := new(object.SearchFilter)

	if !empty {
		m.SetKey("search filter key")
		m.SetValue("search filter val")
		m.SetMatchType(987)
	}

	return m
}

func GenerateSearchFilters(empty bool) (res []*object.SearchFilter) {
	if !empty {
		res = append(res,
			GenerateSearchFilter(false),
			GenerateSearchFilter(false),
		)
	}

	return
}

func GenerateSearchRequestBody(empty bool) *object.SearchRequestBody {
	m := new(object.SearchRequestBody)

	if !empty {
		m.SetVersion(555)
	}

	m.SetContainerID(refstest.GenerateContainerID(empty))
	m.SetFilters(GenerateSearchFilters(empty))

	return m
}

func GenerateSearchRequest(empty bool) *object.SearchRequest {
	m := new(object.SearchRequest)

	m.SetBody(GenerateSearchRequestBody(empty))
	m.SetMetaHeader(sessiontest.GenerateRequestMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateRequestVerificationHeader(empty))

	return m
}

func GenerateSearchResponseBody(empty bool) *object.SearchResponseBody {
	m := new(object.SearchResponseBody)

	m.SetIDList(refstest.GenerateObjectIDs(empty))

	return m
}

func GenerateSearchResponse(empty bool) *object.SearchResponse {
	m := new(object.SearchResponse)

	m.SetBody(GenerateSearchResponseBody(empty))
	m.SetMetaHeader(sessiontest.GenerateResponseMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateResponseVerificationHeader(empty))

	return m
}

func GenerateRange(empty bool) *object.Range {
	m := new(object.Range)

	if !empty {
		m.SetLength(11)
		m.SetOffset(22)
	}

	return m
}

func GenerateRanges(empty bool) (res []*object.Range) {
	if !empty {
		res = append(res,
			GenerateRange(false),
			GenerateRange(false),
		)
	}

	return
}

func GenerateGetRangeRequestBody(empty bool) *object.GetRangeRequestBody {
	m := new(object.GetRangeRequestBody)

	if !empty {
		m.SetRaw(true)
	}

	m.SetAddress(refstest.GenerateAddress(empty))
	m.SetRange(GenerateRange(empty))

	return m
}

func GenerateGetRangeRequest(empty bool) *object.GetRangeRequest {
	m := new(object.GetRangeRequest)

	m.SetBody(GenerateGetRangeRequestBody(empty))
	m.SetMetaHeader(sessiontest.GenerateRequestMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateRequestVerificationHeader(empty))

	return m
}

func GenerateGetRangePartChunk(empty bool) *object.GetRangePartChunk {
	m := new(object.GetRangePartChunk)

	if !empty {
		m.SetChunk([]byte("get range chunk"))
	}

	return m
}

func GenerateGetRangeResponseBody(empty bool) *object.GetRangeResponseBody {
	m := new(object.GetRangeResponseBody)

	m.SetRangePart(GenerateGetRangePartChunk(empty))

	return m
}

func GenerateGetRangeResponse(empty bool) *object.GetRangeResponse {
	m := new(object.GetRangeResponse)

	m.SetBody(GenerateGetRangeResponseBody(empty))
	m.SetMetaHeader(sessiontest.GenerateResponseMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateResponseVerificationHeader(empty))

	return m
}

func GenerateGetRangeHashRequestBody(empty bool) *object.GetRangeHashRequestBody {
	m := new(object.GetRangeHashRequestBody)

	if !empty {
		m.SetSalt([]byte("range hash salt"))
		m.SetType(455)
	}

	m.SetAddress(refstest.GenerateAddress(empty))
	m.SetRanges(GenerateRanges(empty))

	return m
}

func GenerateGetRangeHashRequest(empty bool) *object.GetRangeHashRequest {
	m := new(object.GetRangeHashRequest)

	m.SetBody(GenerateGetRangeHashRequestBody(empty))
	m.SetMetaHeader(sessiontest.GenerateRequestMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateRequestVerificationHeader(empty))

	return m
}

func GenerateGetRangeHashResponseBody(empty bool) *object.GetRangeHashResponseBody {
	m := new(object.GetRangeHashResponseBody)

	if !empty {
		m.SetType(678)
		m.SetHashList([][]byte{{1}, {2}})
	}

	return m
}

func GenerateGetRangeHashResponse(empty bool) *object.GetRangeHashResponse {
	m := new(object.GetRangeHashResponse)

	m.SetBody(GenerateGetRangeHashResponseBody(empty))
	m.SetMetaHeader(sessiontest.GenerateResponseMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateResponseVerificationHeader(empty))

	return m
}
