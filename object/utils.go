package object

import (
	"io"
	"strconv"

	"github.com/pkg/errors"
)

// FilenameHeader is a user header key for names of files, stored by third
// party apps. We recommend to use this header to be compatible with neofs
// http gate, neofs minio gate and neofs-dropper application.
const FilenameHeader = "filename"

// ByteSize used to format bytes
type ByteSize uint64

// String represents ByteSize in string format
func (b ByteSize) String() string {
	var (
		dec  int64
		unit string
		num  = int64(b)
	)

	switch {
	case num > UnitsTB:
		unit = "TB"
		dec = UnitsTB
	case num > UnitsGB:
		unit = "GB"
		dec = UnitsGB
	case num > UnitsMB:
		unit = "MB"
		dec = UnitsMB
	case num > UnitsKB:
		unit = "KB"
		dec = UnitsKB
	default:
		dec = 1
	}

	return strconv.FormatFloat(float64(num)/float64(dec), 'g', 6, 64) + unit
}

// MakePutRequestHeader combines object and session token value
// into header of object put request.
func MakePutRequestHeader(obj *Object) *PutRequest {
	return &PutRequest{
		R: &PutRequest_Header{Header: &PutRequest_PutHeader{
			Object: obj,
		}},
	}
}

// MakePutRequestChunk splits data into chunks that will be transferred
// in the protobuf stream.
func MakePutRequestChunk(chunk []byte) *PutRequest {
	return &PutRequest{R: &PutRequest_Chunk{Chunk: chunk}}
}

func errMaxSizeExceeded(size uint64) error {
	return errors.Errorf("object payload size exceed: %s", ByteSize(size).String())
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
