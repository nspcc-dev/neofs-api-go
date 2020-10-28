package acl

import (
	"errors"

	acl "github.com/nspcc-dev/neofs-api-go/v2/acl/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	errEmptyInput = errors.New("empty input")
)

func RecordToJSON(r *Record) ([]byte, error) {
	if r == nil {
		return nil, errEmptyInput
	}

	msg := RecordToGRPCMessage(r)

	return protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(msg)
}

func RecordFromJSON(data []byte) (*Record, error) {
	if len(data) == 0 {
		return nil, errEmptyInput
	}

	msg := new(acl.EACLRecord)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return nil, err
	}

	return RecordFromGRPCMessage(msg), nil
}

func TableToJSON(t *Table) ([]byte, error) {
	if t == nil {
		return nil, errEmptyInput
	}

	msg := TableToGRPCMessage(t)

	return protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(msg)
}

func TableFromJSON(data []byte) (*Table, error) {
	if len(data) == 0 {
		return nil, errEmptyInput
	}

	msg := new(acl.EACLTable)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return nil, err
	}

	return TableFromGRPCMessage(msg), nil
}

func BearerTokenToJSON(t *BearerToken) ([]byte, error) {
	if t == nil {
		return nil, errEmptyInput
	}

	msg := BearerTokenToGRPCMessage(t)

	return protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(msg)
}

func BearerTokenFromJSON(data []byte) (*BearerToken, error) {
	if len(data) == 0 {
		return nil, errEmptyInput
	}

	msg := new(acl.BearerToken)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return nil, err
	}

	return BearerTokenFromGRPCMessage(msg), nil
}
