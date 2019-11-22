package object

import (
	"io"

	"code.cloudfoundry.org/bytefmt"
	"github.com/nspcc-dev/neofs-proto/session"
	"github.com/pkg/errors"
)

// MakePutRequestHeader combines object and session token value
// into header of object put request.
func MakePutRequestHeader(obj *Object, token *session.Token) *PutRequest {
	return &PutRequest{
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
