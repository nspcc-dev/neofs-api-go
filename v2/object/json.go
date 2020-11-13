package object

import (
	object "github.com/nspcc-dev/neofs-api-go/v2/object/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func (h *ShortHeader) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		ShortHeaderToGRPCMessage(h),
	)
}

func (h *ShortHeader) UnmarshalJSON(data []byte) error {
	msg := new(object.ShortHeader)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*h = *ShortHeaderFromGRPCMessage(msg)

	return nil
}
