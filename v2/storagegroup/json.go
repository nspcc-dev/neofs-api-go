package storagegroup

import (
	storagegroup "github.com/nspcc-dev/neofs-api-go/v2/storagegroup/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func (s *StorageGroup) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		StorageGroupToGRPCMessage(s),
	)
}

func (s *StorageGroup) UnmarshalJSON(data []byte) error {
	msg := new(storagegroup.StorageGroup)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*s = *StorageGroupFromGRPCMessage(msg)

	return nil
}
