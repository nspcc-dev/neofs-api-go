package accounting

import (
	"github.com/nspcc-dev/neofs-api-go/rpc/message"
	accounting "github.com/nspcc-dev/neofs-api-go/v2/accounting/grpc"
)

func (d *Decimal) MarshalJSON() ([]byte, error) {
	return message.MarshalJSON(d)
}

func (d *Decimal) UnmarshalJSON(data []byte) error {
	return message.UnmarshalJSON(d, data, new(accounting.Decimal))
}
