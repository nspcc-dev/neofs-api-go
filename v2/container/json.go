package container

import (
	container "github.com/nspcc-dev/neofs-api-go/v2/container/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func ContainerToJSON(c *Container) (data []byte) {
	if c == nil {
		return nil
	}

	msg := ContainerToGRPCMessage(c)

	data, err := protojson.Marshal(msg)
	if err != nil {
		return nil
	}

	return
}

func ContainerFromJSON(data []byte) *Container {
	if len(data) == 0 {
		return nil
	}

	msg := new(container.Container)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return nil
	}

	return ContainerFromGRPCMessage(msg)
}
