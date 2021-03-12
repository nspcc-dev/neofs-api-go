package audit

import (
	"github.com/nspcc-dev/neofs-api-go/rpc/message"
	audit "github.com/nspcc-dev/neofs-api-go/v2/audit/grpc"
)

func (a *DataAuditResult) MarshalJSON() ([]byte, error) {
	return message.MarshalJSON(a)
}

func (a *DataAuditResult) UnmarshalJSON(data []byte) error {
	return message.UnmarshalJSON(a, data, new(audit.DataAuditResult))
}
