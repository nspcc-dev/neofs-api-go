package container_test

import (
	"fmt"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/acl"
	"github.com/nspcc-dev/neofs-api-go/v2/container"
	grpc "github.com/nspcc-dev/neofs-api-go/v2/container/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/stretchr/testify/require"
)

func TestAttribute_StableMarshal(t *testing.T) {
	attributeFrom := generateAttribute("key", "value")
	transport := new(grpc.Container_Attribute)

	t.Run("non empty", func(t *testing.T) {
		wire, err := attributeFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		attributeTo := container.AttributeFromGRPCMessage(transport)
		require.Equal(t, attributeFrom, attributeTo)
	})
}

func TestContainer_StableMarshal(t *testing.T) {
	cnrFrom := generateContainer("nonce")
	transport := new(grpc.Container)

	t.Run("non empty", func(t *testing.T) {
		wire, err := cnrFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		cnrTo := container.ContainerFromGRPCMessage(transport)
		require.Equal(t, cnrFrom, cnrTo)
	})
}

func TestPutRequestBody_StableMarshal(t *testing.T) {
	requestFrom := generatePutRequestBody("nonce")
	transport := new(grpc.PutRequest_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := requestFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		requestTo := container.PutRequestBodyFromGRPCMessage(transport)
		require.Equal(t, requestFrom, requestTo)
	})
}

func TestPutResponseBody_StableMarshal(t *testing.T) {
	responseFrom := generatePutResponseBody("Container ID")
	transport := new(grpc.PutResponse_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := responseFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		responseTo := container.PutResponseBodyFromGRPCMessage(transport)
		require.Equal(t, responseFrom, responseTo)
	})
}

func TestDeleteRequestBody_StableMarshal(t *testing.T) {
	requestFrom := generateDeleteRequestBody("Container ID")
	transport := new(grpc.DeleteRequest_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := requestFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		requestTo := container.DeleteRequestBodyFromGRPCMessage(transport)
		require.Equal(t, requestFrom, requestTo)
	})
}

func TestDeleteResponseBody_StableMarshal(t *testing.T) {
	responseFrom := generateDeleteResponseBody()
	transport := new(grpc.DeleteResponse_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := responseFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		responseTo := container.DeleteResponseBodyFromGRPCMessage(transport)
		require.Equal(t, responseFrom, responseTo)
	})
}

func TestGetRequestBody_StableMarshal(t *testing.T) {
	requestFrom := generateGetRequestBody("Container ID")
	transport := new(grpc.GetRequest_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := requestFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		requestTo := container.GetRequestBodyFromGRPCMessage(transport)
		require.Equal(t, requestFrom, requestTo)
	})
}

func TestGetResponseBody_StableMarshal(t *testing.T) {
	responseFrom := generateGetResponseBody("nonce")
	transport := new(grpc.GetResponse_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := responseFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		responseTo := container.GetResponseBodyFromGRPCMessage(transport)
		require.Equal(t, responseFrom, responseTo)
	})
}

func TestListRequestBody_StableMarshal(t *testing.T) {
	requestFrom := generateListRequestBody("Owner ID")
	transport := new(grpc.ListRequest_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := requestFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		requestTo := container.ListRequestBodyFromGRPCMessage(transport)
		require.Equal(t, requestFrom, requestTo)
	})
}

func TestListResponseBody_StableMarshal(t *testing.T) {
	responseFrom := generateListResponseBody(3)
	transport := new(grpc.ListResponse_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := responseFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		responseTo := container.ListResponseBodyFromGRPCMessage(transport)
		require.Equal(t, responseFrom, responseTo)
	})
}

func TestSetEACLRequestBody_StableMarshal(t *testing.T) {
	requestFrom := generateSetEACLRequestBody(4, "Filter Key", "Filter Value")
	transport := new(grpc.SetExtendedACLRequest_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := requestFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		requestTo := container.SetExtendedACLRequestBodyFromGRPCMessage(transport)
		require.Equal(t, requestFrom, requestTo)
	})
}

func TestSetEACLResponseBody_StableMarshal(t *testing.T) {
	responseFrom := generateSetEACLResponseBody()
	transport := new(grpc.SetExtendedACLResponse_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := responseFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		responseTo := container.SetExtendedACLResponseBodyFromGRPCMessage(transport)
		require.Equal(t, responseFrom, responseTo)
	})
}

func TestGetEACLRequestBody_StableMarshal(t *testing.T) {
	requestFrom := generateGetEACLRequestBody("Container ID")
	transport := new(grpc.GetExtendedACLRequest_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := requestFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		requestTo := container.GetExtendedACLRequestBodyFromGRPCMessage(transport)
		require.Equal(t, requestFrom, requestTo)
	})
}

func TestGetEACLResponseBody_StableMarshal(t *testing.T) {
	responseFrom := generateGetEACLResponseBody(3, "Filter Key", "Filter Value")
	transport := new(grpc.GetExtendedACLResponse_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := responseFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		responseTo := container.GetExtendedACLResponseBodyFromGRPCMessage(transport)
		require.Equal(t, responseFrom, responseTo)
	})
}

func generateAttribute(k, v string) *container.Attribute {
	attr := new(container.Attribute)
	attr.SetKey(k)
	attr.SetValue(v)

	return attr
}

func generateContainer(n string) *container.Container {
	owner := new(refs.OwnerID)
	owner.SetValue([]byte("Owner ID"))

	version := new(refs.Version)
	version.SetMajor(2)
	version.SetMinor(0)

	// todo: add placement rule

	cnr := new(container.Container)
	cnr.SetOwnerID(owner)
	cnr.SetVersion(version)
	cnr.SetAttributes([]*container.Attribute{
		generateAttribute("one", "two"),
		generateAttribute("three", "four"),
	})
	cnr.SetBasicACL(100)
	cnr.SetNonce([]byte(n))

	return cnr
}

func generateSignature(k, v string) *refs.Signature {
	sig := new(refs.Signature)
	sig.SetKey([]byte(k))
	sig.SetSign([]byte(v))

	return sig
}

func generatePutRequestBody(n string) *container.PutRequestBody {
	req := new(container.PutRequestBody)
	req.SetContainer(generateContainer(n))
	req.SetSignature(generateSignature("public key", "signature"))

	return req
}

func generatePutResponseBody(id string) *container.PutResponseBody {
	cid := new(refs.ContainerID)
	cid.SetValue([]byte(id))

	resp := new(container.PutResponseBody)
	resp.SetContainerID(cid)

	return resp
}

func generateDeleteRequestBody(id string) *container.DeleteRequestBody {
	cid := new(refs.ContainerID)
	cid.SetValue([]byte(id))

	req := new(container.DeleteRequestBody)
	req.SetContainerID(cid)
	req.SetSignature(generateSignature("public key", "signature"))

	return req
}

func generateDeleteResponseBody() *container.DeleteResponseBody {
	return new(container.DeleteResponseBody)
}

func generateGetRequestBody(id string) *container.GetRequestBody {
	cid := new(refs.ContainerID)
	cid.SetValue([]byte(id))

	req := new(container.GetRequestBody)
	req.SetContainerID(cid)

	return req
}

func generateGetResponseBody(n string) *container.GetResponseBody {
	resp := new(container.GetResponseBody)
	resp.SetContainer(generateContainer(n))

	return resp
}

func generateListRequestBody(id string) *container.ListRequestBody {
	owner := new(refs.OwnerID)
	owner.SetValue([]byte(id))

	req := new(container.ListRequestBody)
	req.SetOwnerID(owner)

	return req
}

func generateListResponseBody(n int) *container.ListResponseBody {
	resp := new(container.ListResponseBody)

	ids := make([]*refs.ContainerID, n)
	for i := 0; i < n; i++ {
		cid := new(refs.ContainerID)
		cid.SetValue([]byte(fmt.Sprintf("Container ID %d", n+1)))
		ids[i] = cid
	}

	resp.SetContainerIDs(ids)

	return resp
}

func generateEACL(n int, k, v string) *acl.Table {
	target := new(acl.TargetInfo)
	target.SetRole(acl.RoleUser)

	keys := make([][]byte, n)

	for i := 0; i < n; i++ {
		s := fmt.Sprintf("Public Key %d", i+1)
		keys[i] = []byte(s)
	}

	filter := new(acl.HeaderFilter)
	filter.SetHeaderType(acl.HeaderTypeObject)
	filter.SetMatchType(acl.MatchTypeStringEqual)
	filter.SetName(k)
	filter.SetValue(v)

	record := new(acl.Record)
	record.SetOperation(acl.OperationHead)
	record.SetAction(acl.ActionDeny)
	record.SetTargets([]*acl.TargetInfo{target})
	record.SetFilters([]*acl.HeaderFilter{filter})

	table := new(acl.Table)
	cid := new(refs.ContainerID)
	cid.SetValue([]byte("Container ID"))

	table.SetContainerID(cid)
	table.SetRecords([]*acl.Record{record})

	return table
}

func generateSetEACLRequestBody(n int, k, v string) *container.SetExtendedACLRequestBody {
	req := new(container.SetExtendedACLRequestBody)
	req.SetEACL(generateEACL(n, k, v))
	req.SetSignature(generateSignature("public key", "signature"))

	return req
}

func generateSetEACLResponseBody() *container.SetExtendedACLResponseBody {
	return new(container.SetExtendedACLResponseBody)
}

func generateGetEACLRequestBody(id string) *container.GetExtendedACLRequestBody {
	cid := new(refs.ContainerID)
	cid.SetValue([]byte(id))

	req := new(container.GetExtendedACLRequestBody)
	req.SetContainerID(cid)

	return req
}

func generateGetEACLResponseBody(n int, k, v string) *container.GetExtendedACLResponseBody {
	resp := new(container.GetExtendedACLResponseBody)
	resp.SetEACL(generateEACL(n, k, v))
	resp.SetSignature(generateSignature("public key", "signature"))

	return resp
}
