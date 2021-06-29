package rpc

import (
	"context"

	protoclient "github.com/nspcc-dev/neofs-api-go/rpc/client"
	"github.com/nspcc-dev/neofs-api-go/v2/container"
)

const serviceContainer = serviceNamePrefix + "container.ContainerService"

const (
	rpcContainerPut       = "Put"
	rpcContainerGet       = "Get"
	rpcContainerDel       = "Delete"
	rpcContainerList      = "List"
	rpcContainerSetEACL   = "SetExtendedACL"
	rpcContainerGetEACL   = "GetExtendedACL"
	rpcContainerUsedSpace = "AnnounceUsedSpace"
)

// PutContainerPrm groups the parameters of PutContainer call.
type PutContainerPrm struct {
	req container.PutRequest
}

// SetRequest sets request to send to the server.
//
// Required parameter.
func (x *PutContainerPrm) SetRequest(req container.PutRequest) {
	x.req = req
}

// PutContainerRes groups the results of PutContainer call.
type PutContainerRes struct {
	resp container.PutResponse
}

// Response returns the server response.
func (x *PutContainerRes) Response() container.PutResponse {
	return x.resp
}

// PutContainer executes ContainerService.Put RPC.
//
// Client connection should be established.
//
// If there is no error, res.Response() contains the server response.
//
// Context and res must not be nil.
func PutContainer(ctx context.Context, cli protoclient.Client, prm PutContainerPrm, res *PutContainerRes) error {
	return sendUnaryRPC(ctx, cli, &prm.req, &res.resp, serviceContainer, rpcContainerPut)
}

// GetContainerPrm groups the parameters of GetContainer call.
type GetContainerPrm struct {
	req container.GetRequest
}

// SetRequest sets request to send to the server.
//
// Required parameter.
func (x *GetContainerPrm) SetRequest(req container.GetRequest) {
	x.req = req
}

// GetContainerRes groups the results of GetContainer call.
type GetContainerRes struct {
	resp container.GetResponse
}

// Response returns the server response.
func (x *GetContainerRes) Response() container.GetResponse {
	return x.resp
}

// GetContainer executes ContainerService.Get RPC.
//
// Client connection should be established.
//
// If there is no error, res.Response() contains the server response.
//
// Context and res must not be nil.
func GetContainer(ctx context.Context, cli protoclient.Client, prm GetContainerPrm, res *GetContainerRes) error {
	return sendUnaryRPC(ctx, cli, &prm.req, &res.resp, serviceContainer, rpcContainerGet)
}

// DeleteContainerPrm groups the parameters of DeleteContainer call.
type DeleteContainerPrm struct {
	req container.DeleteRequest
}

// SetRequest sets request to send to the server.
//
// Required parameter.
func (x *DeleteContainerPrm) SetRequest(req container.DeleteRequest) {
	x.req = req
}

// DeleteContainerRes groups the results of DeleteContainer call.
type DeleteContainerRes struct {
	resp container.DeleteResponse
}

// Response returns the server response.
func (x *DeleteContainerRes) Response() container.DeleteResponse {
	return x.resp
}

// DeleteContainer executes ContainerService.Delete RPC.
//
// Client connection should be established.
//
// If there is no error, res.Response() contains the server response.
//
// Context and res must not be nil.
func DeleteContainer(ctx context.Context, cli protoclient.Client, prm DeleteContainerPrm, res *DeleteContainerRes) error {
	return sendUnaryRPC(ctx, cli, &prm.req, &res.resp, serviceContainer, rpcContainerDel)
}

// ListContainerPrm groups the parameters of ListContainers call.
type ListContainerPrm struct {
	req container.ListRequest
}

// SetRequest sets request to send to the server.
//
// Required parameter.
func (x *ListContainerPrm) SetRequest(req container.ListRequest) {
	x.req = req
}

// ListContainerRes groups the results of ListContainers call.
type ListContainerRes struct {
	resp container.ListResponse
}

// Response returns the server response.
func (x *ListContainerRes) Response() container.ListResponse {
	return x.resp
}

// ListContainer executes ContainerService.List RPC.
//
// Client connection should be established.
//
// If there is no error, res.Response() contains the server response.
//
// Context and res must not be nil.
func ListContainers(ctx context.Context, cli protoclient.Client, prm ListContainerPrm, res *ListContainerRes) error {
	return sendUnaryRPC(ctx, cli, &prm.req, &res.resp, serviceContainer, rpcContainerList)
}

// SetEACLPrm groups the parameters of SetEACL call.
type SetEACLPrm struct {
	req container.SetExtendedACLRequest
}

// SetRequest sets request to send to the server.
//
// Required parameter.
func (x *SetEACLPrm) SetRequest(req container.SetExtendedACLRequest) {
	x.req = req
}

// SetEACLRes groups the results of SetEACL call.
type SetEACLRes struct {
	resp container.SetExtendedACLResponse
}

// Response returns the server response.
func (x *SetEACLRes) Response() container.SetExtendedACLResponse {
	return x.resp
}

// SetEACL executes ContainerService.SetExtendedACL RPC.
//
// Client connection should be established.
//
// If there is no error, res.Response() contains the server response.
//
// Context and res must not be nil.
func SetEACL(ctx context.Context, cli protoclient.Client, prm SetEACLPrm, res *SetEACLRes) error {
	return sendUnaryRPC(ctx, cli, &prm.req, &res.resp, serviceContainer, rpcContainerSetEACL)
}

// GetEACLPrm groups the parameters of GetEACL call.
type GetEACLPrm struct {
	req container.GetExtendedACLRequest
}

// SetRequest sets request to send to the server.
//
// Required parameter.
func (x *GetEACLPrm) SetRequest(req container.GetExtendedACLRequest) {
	x.req = req
}

// GetEACLRes groups the results of GetEACL call.
type GetEACLRes struct {
	resp container.GetExtendedACLResponse
}

// Response returns the server response.
func (x *GetEACLRes) Response() container.GetExtendedACLResponse {
	return x.resp
}

// GetEACL executes ContainerService.GetExtendedACL RPC.
//
// Client connection should be established.
//
// If there is no error, res.Response() contains the server response.
//
// Context and res must not be nil.
func GetEACL(ctx context.Context, cli protoclient.Client, prm GetEACLPrm, res *GetEACLRes) error {
	return sendUnaryRPC(ctx, cli, &prm.req, &res.resp, serviceContainer, rpcContainerGetEACL)
}

// AnnounceUsedSpacePrm groups the parameters of AnnounceUsedSpace call.
type AnnounceUsedSpacePrm struct {
	req container.AnnounceUsedSpaceRequest
}

// SetRequest sets request to send to the server.
//
// Required parameter.
func (x *AnnounceUsedSpacePrm) SetRequest(req container.AnnounceUsedSpaceRequest) {
	x.req = req
}

// AnnounceUsedSpaceRes groups the results of AnnounceUsedSpace call.
type AnnounceUsedSpaceRes struct {
	resp container.AnnounceUsedSpaceResponse
}

// Response returns the server response.
func (x *AnnounceUsedSpaceRes) Response() container.AnnounceUsedSpaceResponse {
	return x.resp
}

// AnnounceUsedSpace executes ContainerService.AnnounceUsedSpace RPC.
//
// Client connection should be established.
//
// If there is no error, res.Response() contains the server response.
//
// Context and res must not be nil.
func AnnounceUsedSpace(ctx context.Context, cli protoclient.Client, prm AnnounceUsedSpacePrm, res *AnnounceUsedSpaceRes) error {
	return sendUnaryRPC(ctx, cli, &prm.req, &res.resp, serviceContainer, rpcContainerUsedSpace)
}
