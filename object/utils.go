package object

import (
	"io"

	"code.cloudfoundry.org/bytefmt"
	"github.com/nspcc-dev/neofs-proto/service"
	"github.com/nspcc-dev/neofs-proto/session"
	"github.com/pkg/errors"
)

const maxGetPayloadSize = 3584 * 1024 // 3.5 MiB

func splitBytes(data []byte, maxSize int) (result [][]byte) {
	l := len(data)
	if l == 0 {
		return [][]byte{data}
	}
	for i := 0; i < l; i += maxSize {
		last := i + maxSize
		if last > l {
			last = l
		}
		result = append(result, data[i:last])
	}
	return
}

// SendPutRequest prepares object and sends it in chunks through protobuf stream.
func SendPutRequest(s Service_PutClient, obj *Object, epoch uint64, ttl uint32) (*PutResponse, error) {
	// TODO split must take into account size of the marshalled Object
	chunks := splitBytes(obj.Payload, maxGetPayloadSize)
	obj.Payload = chunks[0]
	if err := s.Send(MakePutRequestHeader(obj, epoch, ttl, nil)); err != nil {
		return nil, err
	}
	for i := range chunks[1:] {
		if err := s.Send(MakePutRequestChunk(chunks[i+1])); err != nil {
			return nil, err
		}
	}
	resp, err := s.CloseAndRecv()
	if err != nil && err != io.EOF {
		return nil, err
	}
	return resp, nil
}

// MakePutRequestHeader combines object, epoch, ttl and session token value
// into header of object put request.
func MakePutRequestHeader(obj *Object, epoch uint64, ttl uint32, token *session.Token) *PutRequest {
	return &PutRequest{
		RequestMetaHeader: service.RequestMetaHeader{TTL: ttl, Epoch: epoch},
		R: &PutRequest_Header{Header: &PutRequest_PutHeader{
			Object: obj,
			Token:  token,
		}},
	}
}

// MakePutRequestChunk splits data into chunks that will be transferred
// in the protobuf stream.
func MakePutRequestChunk(chunk []byte) *PutRequest {
	return &PutRequest{R: &PutRequest_Chunk{Chunk: chunk}}
}

func errMaxSizeExceeded(size uint64) error {
	return errors.Errorf("object payload size exceed: %s", bytefmt.ByteSize(size))
}

// ReceiveGetResponse receives object by chunks from the protobuf stream
// and combine it into single get response structure.
func ReceiveGetResponse(c Service_GetClient, maxSize uint64) (*GetResponse, error) {
	res, err := c.Recv()
	if err == io.EOF {
		return res, err
	} else if err != nil {
		return nil, err
	}

	obj := res.GetObject()
	if obj == nil {
		return nil, ErrHeaderExpected
	}

	if obj.SystemHeader.PayloadLength > maxSize {
		return nil, errMaxSizeExceeded(maxSize)
	}

	if res.NotFull() {
		payload := make([]byte, obj.SystemHeader.PayloadLength)
		offset := copy(payload, obj.Payload)

		var r *GetResponse
		for r, err = c.Recv(); err == nil; r, err = c.Recv() {
			offset += copy(payload[offset:], r.GetChunk())
		}
		if err != io.EOF {
			return nil, err
		}
		obj.Payload = payload
	}

	return res, nil
}
