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

func (h *Header) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		HeaderToGRPCMessage(h),
	)
}

func (h *Header) UnmarshalJSON(data []byte) error {
	msg := new(object.Header)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*h = *HeaderFromGRPCMessage(msg)

	return nil
}

func (h *HeaderWithSignature) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		HeaderWithSignatureToGRPCMessage(h),
	)
}

func (h *HeaderWithSignature) UnmarshalJSON(data []byte) error {
	msg := new(object.HeaderWithSignature)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*h = *HeaderWithSignatureFromGRPCMessage(msg)

	return nil
}

func (o *Object) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		ObjectToGRPCMessage(o),
	)
}

func (o *Object) UnmarshalJSON(data []byte) error {
	msg := new(object.Object)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*o = *ObjectFromGRPCMessage(msg)

	return nil
}

func (f *SearchFilter) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		SearchFilterToGRPCMessage(f),
	)
}

func (f *SearchFilter) UnmarshalJSON(data []byte) error {
	msg := new(object.SearchRequest_Body_Filter)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*f = *SearchFilterFromGRPCMessage(msg)

	return nil
}
