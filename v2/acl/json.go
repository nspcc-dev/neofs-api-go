package acl

import (
	acl "github.com/nspcc-dev/neofs-api-go/v2/acl/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func RecordToJSON(r *Record) (data []byte) {
	if r == nil {
		return nil
	}

	msg := RecordToGRPCMessage(r)

	data, err := protojson.Marshal(msg)
	if err != nil {
		return nil
	}

	return
}

func RecordFromJSON(data []byte) *Record {
	if len(data) == 0 {
		return nil
	}

	msg := new(acl.EACLRecord)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return nil
	}

	return RecordFromGRPCMessage(msg)
}

func TableToJSON(t *Table) (data []byte) {
	if t == nil {
		return nil
	}

	msg := TableToGRPCMessage(t)

	data, err := protojson.Marshal(msg)
	if err != nil {
		return nil
	}

	return
}

func TableFromJSON(data []byte) *Table {
	if len(data) == 0 {
		return nil
	}

	msg := new(acl.EACLTable)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return nil
	}

	return TableFromGRPCMessage(msg)
}

func BearerTokenToJSON(t *BearerToken) (data []byte) {
	if t == nil {
		return nil
	}

	msg := BearerTokenToGRPCMessage(t)

	data, err := protojson.Marshal(msg)
	if err != nil {
		return nil
	}

	return
}

func BearerTokenFromJSON(data []byte) *BearerToken {
	if len(data) == 0 {
		return nil
	}

	msg := new(acl.BearerToken)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return nil
	}

	return BearerTokenFromGRPCMessage(msg)
}
