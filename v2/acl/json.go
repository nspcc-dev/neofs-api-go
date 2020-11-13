package acl

import (
	"errors"

	acl "github.com/nspcc-dev/neofs-api-go/v2/acl/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	errEmptyInput = errors.New("empty input")
)

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

func (f *HeaderFilter) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		HeaderFilterToGRPCMessage(f),
	)
}

func (f *HeaderFilter) UnmarshalJSON(data []byte) error {
	msg := new(acl.EACLRecord_Filter)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*f = *HeaderFilterFromGRPCMessage(msg)

	return nil
}

func (t *Target) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		TargetToGRPCMessage(t),
	)
}

func (t *Target) UnmarshalJSON(data []byte) error {
	msg := new(acl.EACLRecord_Target)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*t = *TargetInfoFromGRPCMessage(msg)

	return nil
}

func (r *Record) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		RecordToGRPCMessage(r),
	)
}

func (r *Record) UnmarshalJSON(data []byte) error {
	msg := new(acl.EACLRecord)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*r = *RecordFromGRPCMessage(msg)

	return nil
}

func (t *Table) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		TableToGRPCMessage(t),
	)
}

func (t *Table) UnmarshalJSON(data []byte) error {
	msg := new(acl.EACLTable)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*t = *TableFromGRPCMessage(msg)

	return nil
}

func (l *TokenLifetime) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		TokenLifetimeToGRPCMessage(l),
	)
}

func (l *TokenLifetime) UnmarshalJSON(data []byte) error {
	msg := new(acl.BearerToken_Body_TokenLifetime)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*l = *TokenLifetimeFromGRPCMessage(msg)

	return nil
}
