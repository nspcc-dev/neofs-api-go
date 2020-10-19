package container

import (
	"github.com/golang/protobuf/jsonpb"
	container "github.com/nspcc-dev/neofs-api-go/v2/container/grpc"
)

func ContainerToJSON(c *Container) []byte {
	if c == nil {
		return nil
	}

	msg := ContainerToGRPCMessage(c)
	m := jsonpb.Marshaler{}

	s, err := m.MarshalToString(msg)
	if err != nil {
		return nil
	}

	return []byte(s)
}

func ContainerFromJSON(data []byte) *Container {
	if len(data) == 0 {
		return nil
	}

	msg := new(container.Container)

	if err := jsonpb.UnmarshalString(string(data), msg); err != nil {
		return nil
	}

	return ContainerFromGRPCMessage(msg)
}
