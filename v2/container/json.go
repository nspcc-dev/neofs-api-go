package container

import (
	container "github.com/nspcc-dev/neofs-api-go/v2/container/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func (a *Attribute) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		AttributeToGRPCMessage(a),
	)
}

func (a *Attribute) UnmarshalJSON(data []byte) error {
	msg := new(container.Container_Attribute)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*a = *AttributeFromGRPCMessage(msg)

	return nil
}

func (c *Container) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		ContainerToGRPCMessage(c),
	)
}

func (c *Container) UnmarshalJSON(data []byte) error {
	msg := new(container.Container)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*c = *ContainerFromGRPCMessage(msg)

	return nil
}
