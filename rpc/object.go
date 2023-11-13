package rpc

import (
	"context"

	"github.com/nspcc-dev/neofs-api-go/v2/object"
	"github.com/nspcc-dev/neofs-api-go/v2/rpc/client"
	"github.com/nspcc-dev/neofs-api-go/v2/rpc/common"
	"github.com/nspcc-dev/neofs-api-go/v2/rpc/message"
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

// PutRequestWriter is an object.PutRequest
// message streaming component.
type PutRequestWriter struct {
	wc client.MessageWriterCloser

	resp message.Message
}

// Write writes req to the stream.
func (w *PutRequestWriter) Write(req *object.PutRequest) error {
	return w.wc.WriteMessage(req)
}

// Close closes the stream.
func (w *PutRequestWriter) Close() error {
	return w.wc.Close()
}

// PutObject executes ObjectService.Put RPC.
func PutObject(
	cli *client.Client,
	resp *object.PutResponse,
	opts ...client.CallOption,
) (*PutRequestWriter, error) {
	wc, err := client.OpenClientStream(cli, common.CallMethodInfoClientStream(serviceObject, rpcObjectPut), resp, opts...)
	if err != nil {
		return nil, err
	}

	return &PutRequestWriter{
		wc:   wc,
		resp: resp,
	}, nil
}

// PutRequestBinaryWriter represents stream of binary-encoded request messages
// of the NeoFS API V2 ObjectService.Put RPC.
type PutRequestBinaryWriter struct {
	wc client.MessageWriterCloser
}

// Write writes next binary-encoded request message to the stream. Note that
// message sequence and format should comply to the protocol requirements.
func (w *PutRequestBinaryWriter) Write(msg []byte) error {
	return w.wc.WriteMessage(client.BinaryMessage(msg))
}

// Close closes the stream and decodes response message.
func (w *PutRequestBinaryWriter) Close() error {
	return w.wc.Close()
}

// PutObjectBinary opens and returns binary object stream using NeoFS API V2
// ObjectService.Put RPC. Object is transmitted by sequentially calling
// [PutRequestBinaryWriter.Write] method. When stream is completed,
// [PutRequestBinaryWriter.Close] must be called. If transmission succeeds,
// provided response is decoded from the received message.
func PutObjectBinary(ctx context.Context, cli *client.Client, resp *object.PutResponse, opts ...client.CallOption) (*PutRequestBinaryWriter, error) {
	wc, err := client.OpenClientStream(cli, common.CallMethodInfoClientStream(serviceObject, rpcObjectPut), resp,
		append(opts, client.WithContext(ctx), client.AllowBinarySendingOnly())...)
	if err != nil {
		return nil, err
	}

	return &PutRequestBinaryWriter{
		wc: wc,
	}, nil
}

// GetResponseReader is an object.GetResponse
// stream reader.
type GetResponseReader struct {
	r client.MessageReader
}

// Read reads response from the stream.
//
// Returns io.EOF of streaming is finished.
func (r *GetResponseReader) Read(resp *object.GetResponse) error {
	return r.r.ReadMessage(resp)
}

// GetObject executes ObjectService.Get RPC.
func GetObject(
	cli *client.Client,
	req *object.GetRequest,
	opts ...client.CallOption,
) (*GetResponseReader, error) {
	wc, err := client.OpenServerStream(cli, common.CallMethodInfoServerStream(serviceObject, rpcObjectGet), req, opts...)
	if err != nil {
		return nil, err
	}

	return &GetResponseReader{
		r: wc,
	}, nil
}

// GetResponseReader is an object.SearchResponse
// stream reader.
type SearchResponseReader struct {
	r client.MessageReader
}

// Read reads response from the stream.
//
// Returns io.EOF of streaming is finished.
func (r *SearchResponseReader) Read(resp *object.SearchResponse) error {
	return r.r.ReadMessage(resp)
}

// SearchObjects executes ObjectService.Search RPC.
func SearchObjects(
	cli *client.Client,
	req *object.SearchRequest,
	opts ...client.CallOption,
) (*SearchResponseReader, error) {
	wc, err := client.OpenServerStream(cli, common.CallMethodInfoServerStream(serviceObject, rpcObjectSearch), req, opts...)
	if err != nil {
		return nil, err
	}

	return &SearchResponseReader{
		r: wc,
	}, nil
}

// GetResponseReader is an object.GetRangeResponse
// stream reader.
type ObjectRangeResponseReader struct {
	r client.MessageReader
}

// Read reads response from the stream.
//
// Returns io.EOF of streaming is finished.
func (r *ObjectRangeResponseReader) Read(resp *object.GetRangeResponse) error {
	return r.r.ReadMessage(resp)
}

// GetObjectRange executes ObjectService.GetRange RPC.
func GetObjectRange(
	cli *client.Client,
	req *object.GetRangeRequest,
	opts ...client.CallOption,
) (*ObjectRangeResponseReader, error) {
	wc, err := client.OpenServerStream(cli, common.CallMethodInfoServerStream(serviceObject, rpcObjectRange), req, opts...)
	if err != nil {
		return nil, err
	}

	return &ObjectRangeResponseReader{
		r: wc,
	}, nil
}

// HeadObject executes ObjectService.Head RPC.
func HeadObject(
	cli *client.Client,
	req *object.HeadRequest,
	opts ...client.CallOption,
) (*object.HeadResponse, error) {
	resp := new(object.HeadResponse)

	err := client.SendUnary(cli, common.CallMethodInfoUnary(serviceObject, rpcObjectHead), req, resp, opts...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteObject executes ObjectService.Delete RPC.
func DeleteObject(
	cli *client.Client,
	req *object.DeleteRequest,
	opts ...client.CallOption,
) (*object.DeleteResponse, error) {
	resp := new(object.DeleteResponse)

	err := client.SendUnary(cli, common.CallMethodInfoUnary(serviceObject, rpcObjectDelete), req, resp, opts...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// HashObjectRange executes ObjectService.GetRangeHash RPC.
func HashObjectRange(
	cli *client.Client,
	req *object.GetRangeHashRequest,
	opts ...client.CallOption,
) (*object.GetRangeHashResponse, error) {
	resp := new(object.GetRangeHashResponse)

	err := client.SendUnary(cli, common.CallMethodInfoUnary(serviceObject, rpcObjectHash), req, resp, opts...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
