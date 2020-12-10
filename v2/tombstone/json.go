package tombstone

import (
	tombstone "github.com/nspcc-dev/neofs-api-go/v2/tombstone/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func (s *Tombstone) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		TombstoneToGRPCMessage(s),
	)
}

func (s *Tombstone) UnmarshalJSON(data []byte) error {
	msg := new(tombstone.Tombstone)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*s = *TombstoneFromGRPCMessage(msg)

	return nil
}
