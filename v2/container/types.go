package container

import (
	"github.com/nspcc-dev/neofs-api-go/v2/acl"
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/service"
)

type Attribute struct {
	key, val string
}

type Container struct {
	version *refs.Version

	ownerID *refs.OwnerID

	nonce []byte

	basicACL uint32

	attr []*Attribute

	policy *netmap.PlacementPolicy
}

type PutRequestBody struct {
	cnr *Container

	sig *refs.Signature
}

type PutRequest struct {
	body *PutRequestBody

	metaHeader *service.RequestMetaHeader

	verifyHeader *service.RequestVerificationHeader
}

type PutResponseBody struct {
	cid *refs.ContainerID
}

type PutResponse struct {
	body *PutResponseBody

	metaHeader *service.ResponseMetaHeader

	verifyHeader *service.ResponseVerificationHeader
}

type GetRequestBody struct {
	cid *refs.ContainerID
}

type GetRequest struct {
	body *GetRequestBody

	metaHeader *service.RequestMetaHeader

	verifyHeader *service.RequestVerificationHeader
}

type GetResponseBody struct {
	cnr *Container
}

type GetResponse struct {
	body *GetResponseBody

	metaHeader *service.ResponseMetaHeader

	verifyHeader *service.ResponseVerificationHeader
}

type DeleteRequestBody struct {
	cid *refs.ContainerID

	sig *refs.Signature
}

type DeleteRequest struct {
	body *DeleteRequestBody

	metaHeader *service.RequestMetaHeader

	verifyHeader *service.RequestVerificationHeader
}

type DeleteResponseBody struct{}

type DeleteResponse struct {
	body *DeleteResponseBody

	metaHeader *service.ResponseMetaHeader

	verifyHeader *service.ResponseVerificationHeader
}

type ListRequestBody struct {
	ownerID *refs.OwnerID
}

type ListRequest struct {
	body *ListRequestBody

	metaHeader *service.RequestMetaHeader

	verifyHeader *service.RequestVerificationHeader
}

type ListResponseBody struct {
	cidList []*refs.ContainerID
}

type ListResponse struct {
	body *ListResponseBody

	metaHeader *service.ResponseMetaHeader

	verifyHeader *service.ResponseVerificationHeader
}

type SetExtendedACLRequestBody struct {
	eacl *acl.Table

	sig *refs.Signature
}

type SetExtendedACLRequest struct {
	body *SetExtendedACLRequestBody

	metaHeader *service.RequestMetaHeader

	verifyHeader *service.RequestVerificationHeader
}

type SetExtendedACLResponseBody struct{}

type SetExtendedACLResponse struct {
	body *SetExtendedACLResponseBody

	metaHeader *service.ResponseMetaHeader

	verifyHeader *service.ResponseVerificationHeader
}

type GetExtendedACLRequestBody struct {
	cid *refs.ContainerID
}

type GetExtendedACLRequest struct {
	body *GetExtendedACLRequestBody

	metaHeader *service.RequestMetaHeader

	verifyHeader *service.RequestVerificationHeader
}

type GetExtendedACLResponseBody struct {
	eacl *acl.Table

	sig *refs.Signature
}

type GetExtendedACLResponse struct {
	body *GetExtendedACLResponseBody

	metaHeader *service.ResponseMetaHeader

	verifyHeader *service.ResponseVerificationHeader
}

func (a *Attribute) GetKey() string {
	if a != nil {
		return a.key
	}

	return ""
}

func (a *Attribute) SetKey(v string) {
	if a != nil {
		a.key = v
	}
}

func (a *Attribute) GetValue() string {
	if a != nil {
		return a.val
	}

	return ""
}

func (a *Attribute) SetValue(v string) {
	if a != nil {
		a.val = v
	}
}

func (c *Container) GetVersion() *refs.Version {
	if c != nil {
		return c.version
	}

	return nil
}

func (c *Container) SetVersion(v *refs.Version) {
	if c != nil {
		c.version = v
	}
}

func (c *Container) GetOwnerID() *refs.OwnerID {
	if c != nil {
		return c.ownerID
	}

	return nil
}

func (c *Container) SetOwnerID(v *refs.OwnerID) {
	if c != nil {
		c.ownerID = v
	}
}

func (c *Container) GetNonce() []byte {
	if c != nil {
		return c.nonce
	}

	return nil
}

func (c *Container) SetNonce(v []byte) {
	if c != nil {
		c.nonce = v
	}
}

func (c *Container) GetBasicACL() uint32 {
	if c != nil {
		return c.basicACL
	}

	return 0
}

func (c *Container) SetBasicACL(v uint32) {
	if c != nil {
		c.basicACL = v
	}
}

func (c *Container) GetAttributes() []*Attribute {
	if c != nil {
		return c.attr
	}

	return nil
}

func (c *Container) SetAttributes(v []*Attribute) {
	if c != nil {
		c.attr = v
	}
}

func (c *Container) GetPlacementPolicy() *netmap.PlacementPolicy {
	if c != nil {
		return c.policy
	}

	return nil
}

func (c *Container) SetPlacementPolicy(v *netmap.PlacementPolicy) {
	if c != nil {
		c.policy = v
	}
}

func (r *PutRequestBody) GetContainer() *Container {
	if r != nil {
		return r.cnr
	}

	return nil
}

func (r *PutRequestBody) SetContainer(v *Container) {
	if r != nil {
		r.cnr = v
	}
}

func (r *PutRequestBody) GetSignature() *refs.Signature {
	if r != nil {
		return r.sig
	}

	return nil
}

func (r *PutRequestBody) SetSignature(v *refs.Signature) {
	if r != nil {
		r.sig = v
	}
}

func (r *PutRequest) GetBody() *PutRequestBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *PutRequest) SetBody(v *PutRequestBody) {
	if r != nil {
		r.body = v
	}
}

func (r *PutRequest) GetMetaHeader() *service.RequestMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *PutRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *PutRequest) GetVerificationHeader() *service.RequestVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *PutRequest) SetVerificationHeader(v *service.RequestVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (r *PutResponseBody) GetContainerID() *refs.ContainerID {
	if r != nil {
		return r.cid
	}

	return nil
}

func (r *PutResponseBody) SetContainerID(v *refs.ContainerID) {
	if r != nil {
		r.cid = v
	}
}

func (r *PutResponse) GetBody() *PutResponseBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *PutResponse) SetBody(v *PutResponseBody) {
	if r != nil {
		r.body = v
	}
}

func (r *PutResponse) GetMetaHeader() *service.ResponseMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *PutResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *PutResponse) GetVerificationHeader() *service.ResponseVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *PutResponse) SetVerificationHeader(v *service.ResponseVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (r *GetRequestBody) GetContainerID() *refs.ContainerID {
	if r != nil {
		return r.cid
	}

	return nil
}

func (r *GetRequestBody) SetContainerID(v *refs.ContainerID) {
	if r != nil {
		r.cid = v
	}
}

func (r *GetRequest) GetBody() *GetRequestBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *GetRequest) SetBody(v *GetRequestBody) {
	if r != nil {
		r.body = v
	}
}

func (r *GetRequest) GetMetaHeader() *service.RequestMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *GetRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *GetRequest) GetVerificationHeader() *service.RequestVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *GetRequest) SetVerificationHeader(v *service.RequestVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (r *GetResponseBody) GetContainer() *Container {
	if r != nil {
		return r.cnr
	}

	return nil
}

func (r *GetResponseBody) SetContainer(v *Container) {
	if r != nil {
		r.cnr = v
	}
}

func (r *GetResponse) GetBody() *GetResponseBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *GetResponse) SetBody(v *GetResponseBody) {
	if r != nil {
		r.body = v
	}
}

func (r *GetResponse) GetMetaHeader() *service.ResponseMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *GetResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *GetResponse) GetVerificationHeader() *service.ResponseVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *GetResponse) SetVerificationHeader(v *service.ResponseVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (r *DeleteRequestBody) GetContainerID() *refs.ContainerID {
	if r != nil {
		return r.cid
	}

	return nil
}

func (r *DeleteRequestBody) SetContainerID(v *refs.ContainerID) {
	if r != nil {
		r.cid = v
	}
}

func (r *DeleteRequestBody) GetSignature() *refs.Signature {
	if r != nil {
		return r.sig
	}

	return nil
}

func (r *DeleteRequestBody) SetSignature(v *refs.Signature) {
	if r != nil {
		r.sig = v
	}
}

func (r *DeleteRequest) GetBody() *DeleteRequestBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *DeleteRequest) SetBody(v *DeleteRequestBody) {
	if r != nil {
		r.body = v
	}
}

func (r *DeleteRequest) GetMetaHeader() *service.RequestMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *DeleteRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *DeleteRequest) GetVerificationHeader() *service.RequestVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *DeleteRequest) SetVerificationHeader(v *service.RequestVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (r *DeleteResponse) GetBody() *DeleteResponseBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *DeleteResponse) SetBody(v *DeleteResponseBody) {
	if r != nil {
		r.body = v
	}
}

func (r *DeleteResponse) GetMetaHeader() *service.ResponseMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *DeleteResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *DeleteResponse) GetVerificationHeader() *service.ResponseVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *DeleteResponse) SetVerificationHeader(v *service.ResponseVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (r *ListRequestBody) GetOwnerID() *refs.OwnerID {
	if r != nil {
		return r.ownerID
	}

	return nil
}

func (r *ListRequestBody) SetOwnerID(v *refs.OwnerID) {
	if r != nil {
		r.ownerID = v
	}
}

func (r *ListRequest) GetBody() *ListRequestBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *ListRequest) SetBody(v *ListRequestBody) {
	if r != nil {
		r.body = v
	}
}

func (r *ListRequest) GetMetaHeader() *service.RequestMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *ListRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *ListRequest) GetVerificationHeader() *service.RequestVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *ListRequest) SetVerificationHeader(v *service.RequestVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (r *ListResponseBody) GetContainerIDs() []*refs.ContainerID {
	if r != nil {
		return r.cidList
	}

	return nil
}

func (r *ListResponseBody) SetContainerIDs(v []*refs.ContainerID) {
	if r != nil {
		r.cidList = v
	}
}

func (r *ListResponse) GetBody() *ListResponseBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *ListResponse) SetBody(v *ListResponseBody) {
	if r != nil {
		r.body = v
	}
}

func (r *ListResponse) GetMetaHeader() *service.ResponseMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *ListResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *ListResponse) GetVerificationHeader() *service.ResponseVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *ListResponse) SetVerificationHeader(v *service.ResponseVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (r *SetExtendedACLRequestBody) GetEACL() *acl.Table {
	if r != nil {
		return r.eacl
	}

	return nil
}

func (r *SetExtendedACLRequestBody) SetEACL(v *acl.Table) {
	if r != nil {
		r.eacl = v
	}
}

func (r *SetExtendedACLRequestBody) GetSignature() *refs.Signature {
	if r != nil {
		return r.sig
	}

	return nil
}

func (r *SetExtendedACLRequestBody) SetSignature(v *refs.Signature) {
	if r != nil {
		r.sig = v
	}
}

func (r *SetExtendedACLRequest) GetBody() *SetExtendedACLRequestBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *SetExtendedACLRequest) SetBody(v *SetExtendedACLRequestBody) {
	if r != nil {
		r.body = v
	}
}

func (r *SetExtendedACLRequest) GetMetaHeader() *service.RequestMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *SetExtendedACLRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *SetExtendedACLRequest) GetVerificationHeader() *service.RequestVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *SetExtendedACLRequest) SetVerificationHeader(v *service.RequestVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (r *SetExtendedACLResponse) GetBody() *SetExtendedACLResponseBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *SetExtendedACLResponse) SetBody(v *SetExtendedACLResponseBody) {
	if r != nil {
		r.body = v
	}
}

func (r *SetExtendedACLResponse) GetMetaHeader() *service.ResponseMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *SetExtendedACLResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *SetExtendedACLResponse) GetVerificationHeader() *service.ResponseVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *SetExtendedACLResponse) SetVerificationHeader(v *service.ResponseVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (r *GetExtendedACLRequestBody) GetContainerID() *refs.ContainerID {
	if r != nil {
		return r.cid
	}

	return nil
}

func (r *GetExtendedACLRequestBody) SetContainerID(v *refs.ContainerID) {
	if r != nil {
		r.cid = v
	}
}

func (r *GetExtendedACLRequest) GetBody() *GetExtendedACLRequestBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *GetExtendedACLRequest) SetBody(v *GetExtendedACLRequestBody) {
	if r != nil {
		r.body = v
	}
}

func (r *GetExtendedACLRequest) GetMetaHeader() *service.RequestMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *GetExtendedACLRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *GetExtendedACLRequest) GetVerificationHeader() *service.RequestVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *GetExtendedACLRequest) SetVerificationHeader(v *service.RequestVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (r *GetExtendedACLResponseBody) GetEACL() *acl.Table {
	if r != nil {
		return r.eacl
	}

	return nil
}

func (r *GetExtendedACLResponseBody) SetEACL(v *acl.Table) {
	if r != nil {
		r.eacl = v
	}
}

func (r *GetExtendedACLResponseBody) GetSignature() *refs.Signature {
	if r != nil {
		return r.sig
	}

	return nil
}

func (r *GetExtendedACLResponseBody) SetSignature(v *refs.Signature) {
	if r != nil {
		r.sig = v
	}
}

func (r *GetExtendedACLResponse) GetBody() *GetExtendedACLResponseBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *GetExtendedACLResponse) SetBody(v *GetExtendedACLResponseBody) {
	if r != nil {
		r.body = v
	}
}

func (r *GetExtendedACLResponse) GetMetaHeader() *service.ResponseMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *GetExtendedACLResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *GetExtendedACLResponse) GetVerificationHeader() *service.ResponseVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *GetExtendedACLResponse) SetVerificationHeader(v *service.ResponseVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}
