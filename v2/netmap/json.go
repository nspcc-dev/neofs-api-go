package netmap

import (
	netmap "github.com/nspcc-dev/neofs-api-go/v2/netmap/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func NodeInfoToJSON(n *NodeInfo) (data []byte) {
	if n == nil {
		return nil
	}

	msg := NodeInfoToGRPCMessage(n)

	data, err := protojson.Marshal(msg)
	if err != nil {
		return nil
	}

	return
}

func NodeInfoFromJSON(data []byte) *NodeInfo {
	if len(data) == 0 {
		return nil
	}

	msg := new(netmap.NodeInfo)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return nil
	}

	return NodeInfoFromGRPCMessage(msg)
}
