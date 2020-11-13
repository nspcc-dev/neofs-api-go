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

func (a *Attribute) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		AttributeToGRPCMessage(a),
	)
}

func (a *Attribute) UnmarshalJSON(data []byte) error {
	msg := new(object.Header_Attribute)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*a = *AttributeFromGRPCMessage(msg)

	return nil
}

func (h *SplitHeader) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		SplitHeaderToGRPCMessage(h),
	)
}

func (h *SplitHeader) UnmarshalJSON(data []byte) error {
	msg := new(object.Header_Split)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*h = *SplitHeaderFromGRPCMessage(msg)

	return nil
}
