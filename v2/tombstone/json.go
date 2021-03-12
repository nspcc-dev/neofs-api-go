package tombstone

import (
	"github.com/nspcc-dev/neofs-api-go/rpc/message"
	tombstone "github.com/nspcc-dev/neofs-api-go/v2/tombstone/grpc"
)

func (s *Tombstone) MarshalJSON() ([]byte, error) {
	return message.MarshalJSON(s)
}

func (s *Tombstone) UnmarshalJSON(data []byte) error {
	return message.UnmarshalJSON(s, data, new(tombstone.Tombstone))
}
