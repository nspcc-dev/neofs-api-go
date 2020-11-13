package accounting

import (
	accounting "github.com/nspcc-dev/neofs-api-go/v2/accounting/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func (d *Decimal) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		DecimalToGRPCMessage(d),
	)
}

func (d *Decimal) UnmarshalJSON(data []byte) error {
	msg := new(accounting.Decimal)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*d = *DecimalFromGRPCMessage(msg)

	return nil
}
