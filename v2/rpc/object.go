package rpc

import (
	"context"

	protoclient "github.com/nspcc-dev/neofs-api-go/rpc/client"
	"github.com/nspcc-dev/neofs-api-go/rpc/common"
	"github.com/nspcc-dev/neofs-api-go/v2/object"
)

const serviceObject = serviceNamePrefix + "object.ObjectService"

const (
	rpcObjectPut    = "Put"
	rpcObjectGet    = "Get"
	rpcObjectSearch = "Search"
	rpcObjectRange  = "GetRange"
	rpcObjectHash   = "GetRangeHash"
	rpcObjectHead   = "Head"
	rpcObjectDelete = "Delete"
)

// GetObjectPrm groups the parameters of PutObject call.
type PutObjectPrm struct{}

// PutObjectRes groups the results of PutObject call.
type PutObjectRes struct {
	stream PutObjectStream

	resp object.PutResponse
}

// Response returns client stream as PutObjectStream.
func (x PutObjectRes) Stream() PutObjectStream {
	return x.stream
}

// Response returns the server response.
func (x PutObjectRes) Response() object.PutResponse {
	return x.resp
}

// PutObjectStream is a wrapper over client.MessageWriterCloser
// which provides method to write object.PutRequest messages.
type PutObjectStream struct {
	m protoclient.MessageWriterCloser
}

// Write writes object.GetResponse into the stream.
//
// Write should not be called after CloseSend.
func (x PutObjectStream) Write(req object.PutRequest) error {
	return x.m.WriteMessage(&req)
}

// CloseSend calls CloseSend on underlying MessageWriterCloser.
//
// All Write calls should be done before CloseSend.
func (x PutObjectStream) CloseSend() error {
	return x.m.CloseSend()
}

// PutObject executes ObjectService.Put RPC.
//
// Client connection should be established.
//
// If there is no error, client can stream messages using res.Stream().
// When all messages are sent, stream should be closed via CloseSend.
// After the stream is closed w/o an error, res.Response() contains the server response.
//
// Context and res must not be nil.
func PutObject(ctx context.Context, cli protoclient.Client, _ PutObjectPrm, res *PutObjectRes) error {
	var info common.CallMethodInfo

	setMethodInfo(&info, serviceObject, rpcObjectPut, true, false)

	var csPrm protoclient.OpenClientStreamPrm

	csPrm.SetCallMethodInfo(info)
	csPrm.SetResponse(&res.resp)

	var csRes protoclient.OpenClientStreamRes

	err := protoclient.OpenClientStream(ctx, cli, csPrm, &csRes)
	if err != nil {
		return err
	}

	res.stream.m = csRes.Messager()

	return nil
}

// GetObjectPrm groups the parameters of GetObject call.
type GetObjectPrm struct {
	req object.GetRequest
}

// SetRequest sets request to send to the server.
//
// Required parameter.
func (x *GetObjectPrm) SetRequest(req object.GetRequest) {
	x.req = req
}

// GetObjectRes groups the results of GetObject call.
type GetObjectRes struct {
	r GetObjectStream
}

// m returns server stream as GetObjectStream.
func (x GetObjectRes) Stream() GetObjectStream {
	return x.r
}

// GetObjectStream a wrapper over client.MessageReaderCloser
// which provides method to read object.GetResponse messages.
type GetObjectStream struct {
	m protoclient.MessageReaderCloser
}

// Read reads next stream message into the object.GetResponse.
//
// Returns io.EOF of streaming is finished.
func (x *GetObjectStream) Read(resp *object.GetResponse) error {
	return x.m.ReadMessage(resp)
}

func (x *GetObjectStream) set(m protoclient.MessageReaderCloser) {
	x.m = m
}

// GetObject executes ObjectService.Get RPC.
//
// Client connection should be established.
//
// If there is no error, server stream is available in res.Stream().
//
// Context and res must not be nil.
func GetObject(ctx context.Context, cli protoclient.Client, prm GetObjectPrm, res *GetObjectRes) error {
	return openServerStream(ctx, cli, &prm.req, serviceObject, rpcObjectGet, res.r.set)
}

// SearchObjectsPrm groups the parameters of SearchObjects call.
type SearchObjectsPrm struct {
	req object.SearchRequest
}

// SetRequest sets request to send to the server.
//
// Required parameter.
func (x *SearchObjectsPrm) SetRequest(req object.SearchRequest) {
	x.req = req
}

// SearchObjectsRes groups the results of SearchObjects call.
type SearchObjectsRes struct {
	r SearchObjectsStream
}

// Stream returns server stream as SearchObjectsStream.
func (x SearchObjectsRes) Stream() SearchObjectsStream {
	return x.r
}

// SearchObjectsStream a wrapper over client.MessageReaderCloser
// which provides method to read object.SearchResponse messages.
type SearchObjectsStream struct {
	m protoclient.MessageReaderCloser
}

// Read reads next stream message into the object.SearchResponse.
//
// Returns io.EOF of streaming is finished.
func (x *SearchObjectsStream) Read(resp *object.SearchResponse) error {
	return x.m.ReadMessage(resp)
}

func (x *SearchObjectsStream) set(m protoclient.MessageReaderCloser) {
	x.m = m
}

// SearchObjects executes ObjectService.Search RPC.
//
// Client connection should be established.
//
// If there is no error, server stream is available in res.Stream().
//
// Context and res must not be nil.
func SearchObjects(ctx context.Context, cli protoclient.Client, prm SearchObjectsPrm, res *SearchObjectsRes) error {
	return openServerStream(ctx, cli, &prm.req, serviceObject, rpcObjectSearch, res.r.set)
}

// GetObjectRangePrm groups the parameters of GetObjectRange call.
type GetObjectRangePrm struct {
	req object.GetRangeRequest
}

// SetRequest sets request to send to the server.
//
// Required parameter.
func (x *GetObjectRangePrm) SetRequest(req object.GetRangeRequest) {
	x.req = req
}

// GetObjectRangeRes groups the results of GetObjectRange call.
type GetObjectRangeRes struct {
	r GetObjectRangeStream
}

// Stream returns server stream as GetObjectRangeStream.
func (x GetObjectRangeRes) Stream() GetObjectRangeStream {
	return x.r
}

// GetObjectRangeStream a wrapper over client.MessageReaderCloser
// which provides method to read object.GetRangeResponse messages.
type GetObjectRangeStream struct {
	m protoclient.MessageReaderCloser
}

// Read reads next stream message into the object.GetRangeResponse.
//
// Returns io.EOF of streaming is finished.
func (x *GetObjectRangeStream) Read(resp *object.GetRangeResponse) error {
	return x.m.ReadMessage(resp)
}

func (x *GetObjectRangeStream) set(m protoclient.MessageReaderCloser) {
	x.m = m
}

// GetObjectRange executes ObjectService.GetRange RPC.
//
// Client connection should be established.
//
// If there is no error, server stream is available in res.Stream().
//
// Context and res must not be nil.
func GetObjectRange(ctx context.Context, cli protoclient.Client, prm GetObjectRangePrm, res *GetObjectRangeRes) error {
	return openServerStream(ctx, cli, &prm.req, serviceObject, rpcObjectRange, res.r.set)
}

// HeadObjectPrm groups the parameters of HeadObject call.
type HeadObjectPrm struct {
	req object.HeadRequest
}

// SetRequest sets request to send to the server.
//
// Required parameter.
func (x *HeadObjectPrm) SetRequest(req object.HeadRequest) {
	x.req = req
}

// HeadObjectRes groups the results of HeadObject call.
type HeadObjectRes struct {
	resp object.HeadResponse
}

// Response returns the server response.
func (x *HeadObjectRes) Response() object.HeadResponse {
	return x.resp
}

// HeadObject executes ObjectService.Head RPC.
//
// Client connection should be established.
//
// If there is no error, res.Response() contains the server response.
//
// Context and res must not be nil.
func HeadObject(ctx context.Context, cli protoclient.Client, prm HeadObjectPrm, res *HeadObjectRes) error {
	return sendUnaryRPC(ctx, cli, &prm.req, &res.resp, serviceObject, rpcObjectHead)
}

// DeleteObjectPrm groups the parameters of DeleteObjectRange call.
type DeleteObjectPrm struct {
	req object.DeleteRequest
}

// SetRequest sets request to send to the server.
//
// Required parameter.
func (x *DeleteObjectPrm) SetRequest(req object.DeleteRequest) {
	x.req = req
}

// DeleteObjectRes groups the results of DeleteObject call.
type DeleteObjectRes struct {
	resp object.DeleteResponse
}

// Response returns the server response.
func (x *DeleteObjectRes) Response() object.DeleteResponse {
	return x.resp
}

// DeleteObject executes ObjectService.Delete RPC.
//
// Client connection should be established.
//
// If there is no error, res.Response() contains the server response.
//
// Context and res must not be nil.
func DeleteObject(ctx context.Context, cli protoclient.Client, prm DeleteObjectPrm, res *DeleteObjectRes) error {
	return sendUnaryRPC(ctx, cli, &prm.req, &res.resp, serviceObject, rpcObjectDelete)
}

// HashObjectRangePrm groups the parameters of HashObjectRange call.
type HashObjectRangePrm struct {
	req object.GetRangeHashRequest
}

// SetRequest sets request to send to the server.
//
// Required parameter.
func (x *HashObjectRangePrm) SetRequest(req object.GetRangeHashRequest) {
	x.req = req
}

// HashObjectRangeRes groups the results of HashObjectRange call.
type HashObjectRangeRes struct {
	resp object.GetRangeHashResponse
}

// Response returns the server response.
func (x *HashObjectRangeRes) Response() object.GetRangeHashResponse {
	return x.resp
}

// HashObjectRange executes ObjectService.GetRangeHash RPC.
//
// Client connection should be established.
//
// If there is no error, res.Response() contains the server response.
//
// Context and res must not be nil.
func HashObjectRange(ctx context.Context, cli protoclient.Client, prm HashObjectRangePrm, res *HashObjectRangeRes) error {
	return sendUnaryRPC(ctx, cli, &prm.req, &res.resp, serviceObject, rpcObjectHash)
}
