package refs

import (
	refs "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func (a *Address) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		AddressToGRPCMessage(a),
	)
}

func (a *Address) UnmarshalJSON(data []byte) error {
	msg := new(refs.Address)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*a = *AddressFromGRPCMessage(msg)

	return nil
}

func (o *ObjectID) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		ObjectIDToGRPCMessage(o),
	)
}

func (o *ObjectID) UnmarshalJSON(data []byte) error {
	msg := new(refs.ObjectID)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*o = *ObjectIDFromGRPCMessage(msg)

	return nil
}

func (c *ContainerID) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		ContainerIDToGRPCMessage(c),
	)
}

func (c *ContainerID) UnmarshalJSON(data []byte) error {
	msg := new(refs.ContainerID)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*c = *ContainerIDFromGRPCMessage(msg)

	return nil
}

func (o *OwnerID) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		OwnerIDToGRPCMessage(o),
	)
}

func (o *OwnerID) UnmarshalJSON(data []byte) error {
	msg := new(refs.OwnerID)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*o = *OwnerIDFromGRPCMessage(msg)

	return nil
}

func (v *Version) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		VersionToGRPCMessage(v),
	)
}

func (v *Version) UnmarshalJSON(data []byte) error {
	msg := new(refs.Version)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*v = *VersionFromGRPCMessage(msg)

	return nil
}
