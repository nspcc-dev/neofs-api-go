package netmap

import (
	"github.com/golang/protobuf/jsonpb"
	netmap "github.com/nspcc-dev/neofs-api-go/v2/netmap/grpc"
)

func NodeInfoToJSON(n *NodeInfo) []byte {
	if n == nil {
		return nil
	}

	msg := NodeInfoToGRPCMessage(n)
	m := jsonpb.Marshaler{}

	s, err := m.MarshalToString(msg)
	if err != nil {
		return nil
	}

	return []byte(s)
}

func NodeInfoFromJSON(data []byte) *NodeInfo {
	if len(data) == 0 {
		return nil
	}

	msg := new(netmap.NodeInfo)

	if err := jsonpb.UnmarshalString(string(data), msg); err != nil {
		return nil
	}

	return NodeInfoFromGRPCMessage(msg)
}
