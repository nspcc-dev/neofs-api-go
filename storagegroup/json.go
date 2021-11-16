package storagegroup

import (
	"github.com/nspcc-dev/neofs-api-go/v2/rpc/message"
	storagegroup "github.com/nspcc-dev/neofs-api-go/v2/storagegroup/grpc"
)

func (s *StorageGroup) MarshalJSON() ([]byte, error) {
	return message.MarshalJSON(s)
}

func (s *StorageGroup) UnmarshalJSON(data []byte) error {
	return message.UnmarshalJSON(s, data, new(storagegroup.StorageGroup))
}
