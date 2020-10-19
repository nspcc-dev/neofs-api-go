package acl

import (
	"github.com/golang/protobuf/jsonpb"
	acl "github.com/nspcc-dev/neofs-api-go/v2/acl/grpc"
)

func RecordToJSON(r *Record) []byte {
	if r == nil {
		return nil
	}

	msg := RecordToGRPCMessage(r)
	m := jsonpb.Marshaler{}

	s, err := m.MarshalToString(msg)
	if err != nil {
		return nil
	}

	return []byte(s)
}

func RecordFromJSON(data []byte) *Record {
	if len(data) == 0 {
		return nil
	}

	msg := new(acl.EACLRecord)

	if err := jsonpb.UnmarshalString(string(data), msg); err != nil {
		return nil
	}

	return RecordFromGRPCMessage(msg)
}

func TableToJSON(t *Table) (data []byte) {
	if t == nil {
		return nil
	}

	msg := TableToGRPCMessage(t)
	m := jsonpb.Marshaler{}

	s, err := m.MarshalToString(msg)
	if err != nil {
		return nil
	}

	return []byte(s)
}

func TableFromJSON(data []byte) *Table {
	if len(data) == 0 {
		return nil
	}

	msg := new(acl.EACLTable)

	if jsonpb.UnmarshalString(string(data), msg) != nil {
		return nil
	}

	return TableFromGRPCMessage(msg)
}

func BearerTokenToJSON(t *BearerToken) (data []byte) {
	if t == nil {
		return nil
	}

	msg := BearerTokenToGRPCMessage(t)
	m := jsonpb.Marshaler{}

	s, err := m.MarshalToString(msg)
	if err != nil {
		return nil
	}

	return []byte(s)
}

func BearerTokenFromJSON(data []byte) *BearerToken {
	if len(data) == 0 {
		return nil
	}

	msg := new(acl.BearerToken)

	if jsonpb.UnmarshalString(string(data), msg) != nil {
		return nil
	}

	return BearerTokenFromGRPCMessage(msg)
}
