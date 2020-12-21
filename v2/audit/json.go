package audit

import (
	audit "github.com/nspcc-dev/neofs-api-go/v2/audit/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func (a *DataAuditResult) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		DataAuditResultToGRPCMessage(a),
	)
}

func (a *DataAuditResult) UnmarshalJSON(data []byte) error {
	msg := new(audit.DataAuditResult)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*a = *DataAuditResultFromGRPCMessage(msg)

	return nil
}
