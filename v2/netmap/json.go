package netmap

import (
	"errors"

	netmap "github.com/nspcc-dev/neofs-api-go/v2/netmap/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	errEmptyInput = errors.New("empty input")
)

func NodeInfoToJSON(n *NodeInfo) ([]byte, error) {
	if n == nil {
		return nil, errEmptyInput
	}

	msg := NodeInfoToGRPCMessage(n)

	return protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(msg)
}

func NodeInfoFromJSON(data []byte) (*NodeInfo, error) {
	if len(data) == 0 {
		return nil, errEmptyInput
	}

	msg := new(netmap.NodeInfo)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return nil, err
	}

	return NodeInfoFromGRPCMessage(msg), nil
}
