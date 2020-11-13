package netmap

import (
	netmap "github.com/nspcc-dev/neofs-api-go/v2/netmap/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func (p *PlacementPolicy) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		PlacementPolicyToGRPCMessage(p),
	)
}

func (p *PlacementPolicy) UnmarshalJSON(data []byte) error {
	msg := new(netmap.PlacementPolicy)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*p = *PlacementPolicyFromGRPCMessage(msg)

	return nil
}

func (f *Filter) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		FilterToGRPCMessage(f),
	)
}

func (f *Filter) UnmarshalJSON(data []byte) error {
	msg := new(netmap.Filter)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*f = *FilterFromGRPCMessage(msg)

	return nil
}

func (s *Selector) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		SelectorToGRPCMessage(s),
	)
}

func (s *Selector) UnmarshalJSON(data []byte) error {
	msg := new(netmap.Selector)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*s = *SelectorFromGRPCMessage(msg)

	return nil
}

func (r *Replica) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		ReplicaToGRPCMessage(r),
	)
}

func (r *Replica) UnmarshalJSON(data []byte) error {
	msg := new(netmap.Replica)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*r = *ReplicaFromGRPCMessage(msg)

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
	msg := new(netmap.NodeInfo_Attribute)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*a = *AttributeFromGRPCMessage(msg)

	return nil
}

func (ni *NodeInfo) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		NodeInfoToGRPCMessage(ni),
	)
}

func (ni *NodeInfo) UnmarshalJSON(data []byte) error {
	msg := new(netmap.NodeInfo)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*ni = *NodeInfoFromGRPCMessage(msg)

	return nil
}
