package bootstrap

import (
	"github.com/nspcc-dev/neofs-proto/service"
)

// NodeType type alias.
type NodeType = service.NodeRole

// SetTTL sets ttl to Request to satisfy TTLRequest interface.
func (m *Request) SetTTL(v uint32) { m.TTL = v }
