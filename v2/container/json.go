package container

import (
	"errors"

	container "github.com/nspcc-dev/neofs-api-go/v2/container/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	errEmptyInput = errors.New("empty input")
)

func ContainerToJSON(c *Container) ([]byte, error) {
	if c == nil {
		return nil, errEmptyInput
	}

	msg := ContainerToGRPCMessage(c)

	return protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(msg)
}

func ContainerFromJSON(data []byte) (*Container, error) {
	if len(data) == 0 {
		return nil, errEmptyInput
	}

	msg := new(container.Container)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return nil, err
	}

	return ContainerFromGRPCMessage(msg), nil
}
