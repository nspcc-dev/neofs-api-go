package reputation

import (
	"github.com/nspcc-dev/neofs-api-go/rpc/message"
	reputation "github.com/nspcc-dev/neofs-api-go/v2/reputation/grpc"
)

func (x *Trust) MarshalJSON() ([]byte, error) {
	return message.MarshalJSON(x)
}

func (x *Trust) UnmarshalJSON(data []byte) error {
	return message.UnmarshalJSON(x, data, new(reputation.Trust))
}
