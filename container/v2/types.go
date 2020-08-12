package v2

import (
	netmap "github.com/nspcc-dev/neofs-api-go/netmap/v2"
	refs "github.com/nspcc-dev/neofs-api-go/refs/v2"
)

// SetKey sets key to the container attribute.
func (m *Container_Attribute) SetKey(v string) {
	if m != nil {
		m.Key = v
	}
}

// SetValue sets value of the container attribute.
func (m *Container_Attribute) SetValue(v string) {
	if m != nil {
		m.Value = v
	}
}

// SetOwnerId sets identifier of the container owner,
func (m *Container) SetOwnerId(v *refs.OwnerID) {
	if m != nil {
		m.OwnerId = v
	}
}

// SetNonce sets nonce of the container structure.
func (m *Container) SetNonce(v []byte) {
	if m != nil {
		m.Nonce = v
	}
}

// SetBasicAcl sets basic ACL of the container.
func (m *Container) SetBasicAcl(v uint32) {
	if m != nil {
		m.BasicAcl = v
	}
}

// SetAttributes sets list of the container attributes.
func (m *Container) SetAttributes(v []*Container_Attribute) {
	if m != nil {
		m.Attributes = v
	}
}

// SetRules sets placement rules of the container.
func (m *Container) SetRules(v *netmap.PlacementRule) {
	if m != nil {
		m.Rules = v
	}
}
