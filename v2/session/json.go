package session

import (
	session "github.com/nspcc-dev/neofs-api-go/v2/session/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func (c *ObjectSessionContext) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		ObjectSessionContextToGRPCMessage(c),
	)
}

func (c *ObjectSessionContext) UnmarshalJSON(data []byte) error {
	msg := new(session.ObjectSessionContext)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*c = *ObjectSessionContextFromGRPCMessage(msg)

	return nil
}
