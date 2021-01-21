package container

import (
	"github.com/nspcc-dev/neofs-api-go/v2/acl"
	container "github.com/nspcc-dev/neofs-api-go/v2/container/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	refsGRPC "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/session"
)

func AttributeToGRPCMessage(a *Attribute) *container.Container_Attribute {
	if a == nil {
		return nil
	}

	m := new(container.Container_Attribute)

	m.SetKey(a.GetKey())
	m.SetValue(a.GetValue())

	return m
}

func AttributeFromGRPCMessage(m *container.Container_Attribute) *Attribute {
	if m == nil {
		return nil
	}

	a := new(Attribute)

	a.SetKey(m.GetKey())
	a.SetValue(m.GetValue())

	return a
}

func ContainerToGRPCMessage(c *Container) *container.Container {
	if c == nil {
		return nil
	}

	m := new(container.Container)

	m.SetVersion(
		refs.VersionToGRPCMessage(c.GetVersion()),
	)

	m.SetOwnerId(
		refs.OwnerIDToGRPCMessage(c.GetOwnerID()),
	)

	m.SetNonce(c.GetNonce())

	m.SetBasicAcl(c.GetBasicACL())

	m.SetPlacementPolicy(
		netmap.PlacementPolicyToGRPCMessage(c.GetPlacementPolicy()),
	)

	attr := c.GetAttributes()
	attrMsg := make([]*container.Container_Attribute, 0, len(attr))

	for i := range attr {
		attrMsg = append(attrMsg, AttributeToGRPCMessage(attr[i]))
	}

	m.SetAttributes(attrMsg)

	return m
}

func ContainerFromGRPCMessage(m *container.Container) *Container {
	if m == nil {
		return nil
	}

	c := new(Container)

	c.SetVersion(
		refs.VersionFromGRPCMessage(m.GetVersion()),
	)

	c.SetOwnerID(
		refs.OwnerIDFromGRPCMessage(m.GetOwnerId()),
	)

	c.SetNonce(m.GetNonce())

	c.SetBasicACL(m.GetBasicAcl())

	c.SetPlacementPolicy(
		netmap.PlacementPolicyFromGRPCMessage(m.GetPlacementPolicy()),
	)

	attrMsg := m.GetAttributes()
	attr := make([]*Attribute, 0, len(attrMsg))

	for i := range attrMsg {
		attr = append(attr, AttributeFromGRPCMessage(attrMsg[i]))
	}

	c.SetAttributes(attr)

	return c
}

func PutRequestBodyToGRPCMessage(r *PutRequestBody) *container.PutRequest_Body {
	if r == nil {
		return nil
	}

	m := new(container.PutRequest_Body)

	m.SetContainer(
		ContainerToGRPCMessage(r.GetContainer()),
	)

	m.SetSignature(
		refs.SignatureToGRPCMessage(r.GetSignature()),
	)

	return m
}

func PutRequestBodyFromGRPCMessage(m *container.PutRequest_Body) *PutRequestBody {
	if m == nil {
		return nil
	}

	r := new(PutRequestBody)

	r.SetContainer(
		ContainerFromGRPCMessage(m.GetContainer()),
	)

	r.SetSignature(
		refs.SignatureFromGRPCMessage(m.GetSignature()),
	)

	return r
}

func PutRequestToGRPCMessage(r *PutRequest) *container.PutRequest {
	if r == nil {
		return nil
	}

	m := new(container.PutRequest)

	m.SetBody(
		PutRequestBodyToGRPCMessage(r.GetBody()),
	)

	session.RequestHeadersToGRPC(r, m)

	return m
}

func PutRequestFromGRPCMessage(m *container.PutRequest) *PutRequest {
	if m == nil {
		return nil
	}

	r := new(PutRequest)

	r.SetBody(
		PutRequestBodyFromGRPCMessage(m.GetBody()),
	)

	session.RequestHeadersFromGRPC(m, r)

	return r
}

func PutResponseBodyToGRPCMessage(r *PutResponseBody) *container.PutResponse_Body {
	if r == nil {
		return nil
	}

	m := new(container.PutResponse_Body)

	m.SetContainerId(
		refs.ContainerIDToGRPCMessage(r.GetContainerID()),
	)

	return m
}

func PutResponseBodyFromGRPCMessage(m *container.PutResponse_Body) *PutResponseBody {
	if m == nil {
		return nil
	}

	r := new(PutResponseBody)

	r.SetContainerID(
		refs.ContainerIDFromGRPCMessage(m.GetContainerId()),
	)

	return r
}

func PutResponseToGRPCMessage(r *PutResponse) *container.PutResponse {
	if r == nil {
		return nil
	}

	m := new(container.PutResponse)

	m.SetBody(
		PutResponseBodyToGRPCMessage(r.GetBody()),
	)

	session.ResponseHeadersToGRPC(r, m)

	return m
}

func PutResponseFromGRPCMessage(m *container.PutResponse) *PutResponse {
	if m == nil {
		return nil
	}

	r := new(PutResponse)

	r.SetBody(
		PutResponseBodyFromGRPCMessage(m.GetBody()),
	)

	session.ResponseHeadersFromGRPC(m, r)

	return r
}

func GetRequestBodyToGRPCMessage(r *GetRequestBody) *container.GetRequest_Body {
	if r == nil {
		return nil
	}

	m := new(container.GetRequest_Body)

	m.SetContainerId(
		refs.ContainerIDToGRPCMessage(r.GetContainerID()),
	)

	return m
}

func GetRequestBodyFromGRPCMessage(m *container.GetRequest_Body) *GetRequestBody {
	if m == nil {
		return nil
	}

	r := new(GetRequestBody)

	r.SetContainerID(
		refs.ContainerIDFromGRPCMessage(m.GetContainerId()),
	)

	return r
}

func GetRequestToGRPCMessage(r *GetRequest) *container.GetRequest {
	if r == nil {
		return nil
	}

	m := new(container.GetRequest)

	m.SetBody(
		GetRequestBodyToGRPCMessage(r.GetBody()),
	)

	session.RequestHeadersToGRPC(r, m)

	return m
}

func GetRequestFromGRPCMessage(m *container.GetRequest) *GetRequest {
	if m == nil {
		return nil
	}

	r := new(GetRequest)

	r.SetBody(
		GetRequestBodyFromGRPCMessage(m.GetBody()),
	)

	session.RequestHeadersFromGRPC(m, r)

	return r
}

func GetResponseBodyToGRPCMessage(r *GetResponseBody) *container.GetResponse_Body {
	if r == nil {
		return nil
	}

	m := new(container.GetResponse_Body)

	m.SetContainer(
		ContainerToGRPCMessage(r.GetContainer()),
	)

	return m
}

func GetResponseBodyFromGRPCMessage(m *container.GetResponse_Body) *GetResponseBody {
	if m == nil {
		return nil
	}

	r := new(GetResponseBody)

	r.SetContainer(
		ContainerFromGRPCMessage(m.GetContainer()),
	)

	return r
}

func GetResponseToGRPCMessage(r *GetResponse) *container.GetResponse {
	if r == nil {
		return nil
	}

	m := new(container.GetResponse)

	m.SetBody(
		GetResponseBodyToGRPCMessage(r.GetBody()),
	)

	session.ResponseHeadersToGRPC(r, m)

	return m
}

func GetResponseFromGRPCMessage(m *container.GetResponse) *GetResponse {
	if m == nil {
		return nil
	}

	r := new(GetResponse)

	r.SetBody(
		GetResponseBodyFromGRPCMessage(m.GetBody()),
	)

	session.ResponseHeadersFromGRPC(m, r)

	return r
}

func DeleteRequestBodyToGRPCMessage(r *DeleteRequestBody) *container.DeleteRequest_Body {
	if r == nil {
		return nil
	}

	m := new(container.DeleteRequest_Body)

	m.SetContainerId(
		refs.ContainerIDToGRPCMessage(r.GetContainerID()),
	)

	m.SetSignature(
		refs.SignatureToGRPCMessage(r.GetSignature()),
	)

	return m
}

func DeleteRequestBodyFromGRPCMessage(m *container.DeleteRequest_Body) *DeleteRequestBody {
	if m == nil {
		return nil
	}

	r := new(DeleteRequestBody)

	r.SetContainerID(
		refs.ContainerIDFromGRPCMessage(m.GetContainerId()),
	)

	r.SetSignature(
		refs.SignatureFromGRPCMessage(m.GetSignature()),
	)

	return r
}

func DeleteRequestToGRPCMessage(r *DeleteRequest) *container.DeleteRequest {
	if r == nil {
		return nil
	}

	m := new(container.DeleteRequest)

	m.SetBody(
		DeleteRequestBodyToGRPCMessage(r.GetBody()),
	)

	session.RequestHeadersToGRPC(r, m)

	return m
}

func DeleteRequestFromGRPCMessage(m *container.DeleteRequest) *DeleteRequest {
	if m == nil {
		return nil
	}

	r := new(DeleteRequest)

	r.SetBody(
		DeleteRequestBodyFromGRPCMessage(m.GetBody()),
	)

	session.RequestHeadersFromGRPC(m, r)

	return r
}

func DeleteResponseBodyToGRPCMessage(r *DeleteResponseBody) *container.DeleteResponse_Body {
	if r == nil {
		return nil
	}

	m := new(container.DeleteResponse_Body)

	return m
}

func DeleteResponseBodyFromGRPCMessage(m *container.DeleteResponse_Body) *DeleteResponseBody {
	if m == nil {
		return nil
	}

	r := new(DeleteResponseBody)

	return r
}

func DeleteResponseToGRPCMessage(r *DeleteResponse) *container.DeleteResponse {
	if r == nil {
		return nil
	}

	m := new(container.DeleteResponse)

	m.SetBody(
		DeleteResponseBodyToGRPCMessage(r.GetBody()),
	)

	session.ResponseHeadersToGRPC(r, m)

	return m
}

func DeleteResponseFromGRPCMessage(m *container.DeleteResponse) *DeleteResponse {
	if m == nil {
		return nil
	}

	r := new(DeleteResponse)

	r.SetBody(
		DeleteResponseBodyFromGRPCMessage(m.GetBody()),
	)

	session.ResponseHeadersFromGRPC(m, r)

	return r
}

func ListRequestBodyToGRPCMessage(r *ListRequestBody) *container.ListRequest_Body {
	if r == nil {
		return nil
	}

	m := new(container.ListRequest_Body)

	m.SetOwnerId(
		refs.OwnerIDToGRPCMessage(r.GetOwnerID()),
	)

	return m
}

func ListRequestBodyFromGRPCMessage(m *container.ListRequest_Body) *ListRequestBody {
	if m == nil {
		return nil
	}

	r := new(ListRequestBody)

	r.SetOwnerID(
		refs.OwnerIDFromGRPCMessage(m.GetOwnerId()),
	)

	return r
}

func ListRequestToGRPCMessage(r *ListRequest) *container.ListRequest {
	if r == nil {
		return nil
	}

	m := new(container.ListRequest)

	m.SetBody(
		ListRequestBodyToGRPCMessage(r.GetBody()),
	)

	session.RequestHeadersToGRPC(r, m)

	return m
}

func ListRequestFromGRPCMessage(m *container.ListRequest) *ListRequest {
	if m == nil {
		return nil
	}

	r := new(ListRequest)

	r.SetBody(
		ListRequestBodyFromGRPCMessage(m.GetBody()),
	)

	session.RequestHeadersFromGRPC(m, r)

	return r
}

func ListResponseBodyToGRPCMessage(r *ListResponseBody) *container.ListResponse_Body {
	if r == nil {
		return nil
	}

	m := new(container.ListResponse_Body)

	cids := r.GetContainerIDs()
	cidMsg := make([]*refsGRPC.ContainerID, 0, len(cids))

	for i := range cids {
		cidMsg = append(cidMsg, refs.ContainerIDToGRPCMessage(cids[i]))
	}

	m.SetContainerIds(cidMsg)

	return m
}

func ListResponseBodyFromGRPCMessage(m *container.ListResponse_Body) *ListResponseBody {
	if m == nil {
		return nil
	}

	r := new(ListResponseBody)

	cidMsg := m.GetContainerIds()
	cids := make([]*refs.ContainerID, 0, len(cidMsg))

	for i := range cidMsg {
		cids = append(cids, refs.ContainerIDFromGRPCMessage(cidMsg[i]))
	}

	r.SetContainerIDs(cids)

	return r
}

func ListResponseToGRPCMessage(r *ListResponse) *container.ListResponse {
	if r == nil {
		return nil
	}

	m := new(container.ListResponse)

	m.SetBody(
		ListResponseBodyToGRPCMessage(r.GetBody()),
	)

	session.ResponseHeadersToGRPC(r, m)

	return m
}

func ListResponseFromGRPCMessage(m *container.ListResponse) *ListResponse {
	if m == nil {
		return nil
	}

	r := new(ListResponse)

	r.SetBody(
		ListResponseBodyFromGRPCMessage(m.GetBody()),
	)

	session.ResponseHeadersFromGRPC(m, r)

	return r
}

func SetExtendedACLRequestBodyToGRPCMessage(r *SetExtendedACLRequestBody) *container.SetExtendedACLRequest_Body {
	if r == nil {
		return nil
	}

	m := new(container.SetExtendedACLRequest_Body)

	m.SetEacl(
		acl.TableToGRPCMessage(r.GetEACL()),
	)

	m.SetSignature(
		refs.SignatureToGRPCMessage(r.GetSignature()))

	return m
}

func SetExtendedACLRequestBodyFromGRPCMessage(m *container.SetExtendedACLRequest_Body) *SetExtendedACLRequestBody {
	if m == nil {
		return nil
	}

	r := new(SetExtendedACLRequestBody)

	r.SetEACL(
		acl.TableFromGRPCMessage(m.GetEacl()),
	)

	r.SetSignature(
		refs.SignatureFromGRPCMessage(m.GetSignature()),
	)

	return r
}

func SetExtendedACLRequestToGRPCMessage(r *SetExtendedACLRequest) *container.SetExtendedACLRequest {
	if r == nil {
		return nil
	}

	m := new(container.SetExtendedACLRequest)

	m.SetBody(
		SetExtendedACLRequestBodyToGRPCMessage(r.GetBody()),
	)

	session.RequestHeadersToGRPC(r, m)

	return m
}

func SetExtendedACLRequestFromGRPCMessage(m *container.SetExtendedACLRequest) *SetExtendedACLRequest {
	if m == nil {
		return nil
	}

	r := new(SetExtendedACLRequest)

	r.SetBody(
		SetExtendedACLRequestBodyFromGRPCMessage(m.GetBody()),
	)

	session.RequestHeadersFromGRPC(m, r)

	return r
}

func SetExtendedACLResponseBodyToGRPCMessage(r *SetExtendedACLResponseBody) *container.SetExtendedACLResponse_Body {
	if r == nil {
		return nil
	}

	m := new(container.SetExtendedACLResponse_Body)

	return m
}

func SetExtendedACLResponseBodyFromGRPCMessage(m *container.SetExtendedACLResponse_Body) *SetExtendedACLResponseBody {
	if m == nil {
		return nil
	}

	r := new(SetExtendedACLResponseBody)

	return r
}

func SetExtendedACLResponseToGRPCMessage(r *SetExtendedACLResponse) *container.SetExtendedACLResponse {
	if r == nil {
		return nil
	}

	m := new(container.SetExtendedACLResponse)

	m.SetBody(
		SetExtendedACLResponseBodyToGRPCMessage(r.GetBody()),
	)

	session.ResponseHeadersToGRPC(r, m)

	return m
}

func SetExtendedACLResponseFromGRPCMessage(m *container.SetExtendedACLResponse) *SetExtendedACLResponse {
	if m == nil {
		return nil
	}

	r := new(SetExtendedACLResponse)

	r.SetBody(
		SetExtendedACLResponseBodyFromGRPCMessage(m.GetBody()),
	)

	session.ResponseHeadersFromGRPC(m, r)

	return r
}

func GetExtendedACLRequestBodyToGRPCMessage(r *GetExtendedACLRequestBody) *container.GetExtendedACLRequest_Body {
	if r == nil {
		return nil
	}

	m := new(container.GetExtendedACLRequest_Body)

	m.SetContainerId(
		refs.ContainerIDToGRPCMessage(r.GetContainerID()),
	)

	return m
}

func GetExtendedACLRequestBodyFromGRPCMessage(m *container.GetExtendedACLRequest_Body) *GetExtendedACLRequestBody {
	if m == nil {
		return nil
	}

	r := new(GetExtendedACLRequestBody)

	r.SetContainerID(
		refs.ContainerIDFromGRPCMessage(m.GetContainerId()),
	)

	return r
}

func GetExtendedACLRequestToGRPCMessage(r *GetExtendedACLRequest) *container.GetExtendedACLRequest {
	if r == nil {
		return nil
	}

	m := new(container.GetExtendedACLRequest)

	m.SetBody(
		GetExtendedACLRequestBodyToGRPCMessage(r.GetBody()),
	)

	session.RequestHeadersToGRPC(r, m)

	return m
}

func GetExtendedACLRequestFromGRPCMessage(m *container.GetExtendedACLRequest) *GetExtendedACLRequest {
	if m == nil {
		return nil
	}

	r := new(GetExtendedACLRequest)

	r.SetBody(
		GetExtendedACLRequestBodyFromGRPCMessage(m.GetBody()),
	)

	session.RequestHeadersFromGRPC(m, r)

	return r
}

func GetExtendedACLResponseBodyToGRPCMessage(r *GetExtendedACLResponseBody) *container.GetExtendedACLResponse_Body {
	if r == nil {
		return nil
	}

	m := new(container.GetExtendedACLResponse_Body)

	m.SetEacl(
		acl.TableToGRPCMessage(r.GetEACL()),
	)

	m.SetSignature(
		refs.SignatureToGRPCMessage(r.GetSignature()),
	)

	return m
}

func GetExtendedACLResponseBodyFromGRPCMessage(m *container.GetExtendedACLResponse_Body) *GetExtendedACLResponseBody {
	if m == nil {
		return nil
	}

	r := new(GetExtendedACLResponseBody)

	r.SetEACL(
		acl.TableFromGRPCMessage(m.GetEacl()),
	)

	r.SetSignature(
		refs.SignatureFromGRPCMessage(m.GetSignature()),
	)

	return r
}

func GetExtendedACLResponseToGRPCMessage(r *GetExtendedACLResponse) *container.GetExtendedACLResponse {
	if r == nil {
		return nil
	}

	m := new(container.GetExtendedACLResponse)

	m.SetBody(
		GetExtendedACLResponseBodyToGRPCMessage(r.GetBody()),
	)

	session.ResponseHeadersToGRPC(r, m)

	return m
}

func GetExtendedACLResponseFromGRPCMessage(m *container.GetExtendedACLResponse) *GetExtendedACLResponse {
	if m == nil {
		return nil
	}

	r := new(GetExtendedACLResponse)

	r.SetBody(
		GetExtendedACLResponseBodyFromGRPCMessage(m.GetBody()),
	)

	session.ResponseHeadersFromGRPC(m, r)

	return r
}

func UsedSpaceAnnouncementToGRPCMessage(a *UsedSpaceAnnouncement) *container.AnnounceUsedSpaceRequest_Body_Announcement {
	if a == nil {
		return nil
	}

	m := new(container.AnnounceUsedSpaceRequest_Body_Announcement)

	m.SetContainerId(
		refs.ContainerIDToGRPCMessage(a.GetContainerID()),
	)

	m.SetUsedSpace(a.GetUsedSpace())

	return m
}

func UsedSpaceAnnouncementFromGRPCMessage(m *container.AnnounceUsedSpaceRequest_Body_Announcement) *UsedSpaceAnnouncement {
	if m == nil {
		return nil
	}

	a := new(UsedSpaceAnnouncement)

	a.SetContainerID(
		refs.ContainerIDFromGRPCMessage(m.GetContainerId()),
	)

	a.SetUsedSpace(m.GetUsedSpace())

	return a
}

func AnnounceUsedSpaceRequestBodyToGRPCMessage(r *AnnounceUsedSpaceRequestBody) *container.AnnounceUsedSpaceRequest_Body {
	if r == nil {
		return nil
	}

	m := new(container.AnnounceUsedSpaceRequest_Body)

	announcements := r.GetAnnouncements()
	msgAnnouncements := make([]*container.AnnounceUsedSpaceRequest_Body_Announcement, 0, len(announcements))

	for i := range announcements {
		msgAnnouncements = append(
			msgAnnouncements,
			UsedSpaceAnnouncementToGRPCMessage(announcements[i]),
		)
	}

	m.SetAnnouncements(msgAnnouncements)

	return m
}

func AnnounceUsedSpaceRequestBodyFromGRPCMessage(m *container.AnnounceUsedSpaceRequest_Body) *AnnounceUsedSpaceRequestBody {
	if m == nil {
		return nil
	}

	r := new(AnnounceUsedSpaceRequestBody)

	msgAnnouncements := m.GetAnnouncements()
	announcements := make([]*UsedSpaceAnnouncement, 0, len(msgAnnouncements))

	for i := range msgAnnouncements {
		announcements = append(
			announcements,
			UsedSpaceAnnouncementFromGRPCMessage(msgAnnouncements[i]),
		)
	}

	r.SetAnnouncements(announcements)

	return r
}

func AnnounceUsedSpaceRequestToGRPCMessage(r *AnnounceUsedSpaceRequest) *container.AnnounceUsedSpaceRequest {
	if r == nil {
		return nil
	}

	m := new(container.AnnounceUsedSpaceRequest)

	m.SetBody(
		AnnounceUsedSpaceRequestBodyToGRPCMessage(r.GetBody()),
	)

	session.RequestHeadersToGRPC(r, m)

	return m
}

func AnnounceUsedSpaceRequestFromGRPCMessage(m *container.AnnounceUsedSpaceRequest) *AnnounceUsedSpaceRequest {
	if m == nil {
		return nil
	}

	r := new(AnnounceUsedSpaceRequest)

	r.SetBody(
		AnnounceUsedSpaceRequestBodyFromGRPCMessage(m.GetBody()),
	)

	session.RequestHeadersFromGRPC(m, r)

	return r
}

func AnnounceUsedSpaceResponseBodyToGRPCMessage(r *AnnounceUsedSpaceResponseBody) *container.AnnounceUsedSpaceResponse_Body {
	if r == nil {
		return nil
	}

	m := new(container.AnnounceUsedSpaceResponse_Body)

	return m
}

func AnnounceUsedSpaceResponseBodyFromGRPCMessage(m *container.AnnounceUsedSpaceResponse_Body) *AnnounceUsedSpaceResponseBody {
	if m == nil {
		return nil
	}

	r := new(AnnounceUsedSpaceResponseBody)

	return r
}

func AnnounceUsedSpaceResponseToGRPCMessage(r *AnnounceUsedSpaceResponse) *container.AnnounceUsedSpaceResponse {
	if r == nil {
		return nil
	}

	m := new(container.AnnounceUsedSpaceResponse)

	m.SetBody(
		AnnounceUsedSpaceResponseBodyToGRPCMessage(r.GetBody()),
	)

	session.ResponseHeadersToGRPC(r, m)

	return m
}

func AnnounceUsedSpaceResponseFromGRPCMessage(m *container.AnnounceUsedSpaceResponse) *AnnounceUsedSpaceResponse {
	if m == nil {
		return nil
	}

	r := new(AnnounceUsedSpaceResponse)

	r.SetBody(
		AnnounceUsedSpaceResponseBodyFromGRPCMessage(m.GetBody()),
	)

	session.ResponseHeadersFromGRPC(m, r)

	return r
}
