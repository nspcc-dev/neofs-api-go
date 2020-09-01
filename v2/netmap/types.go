package netmap

type PlacementPolicy struct {
	// TODO: fill me
}

// Attribute of storage node.
type Attribute struct {
	key   string
	value string
}

// NodeInfo of storage node.
type NodeInfo struct {
	publicKey  []byte
	address    string
	attributes []*Attribute
}

// NodeState of storage node.
type NodeState uint32

const (
	Unspecified NodeState = iota
	Online
	Offline
)

func (a *Attribute) GetKey() string {
	if a != nil {
		return a.key
	}

	return ""
}

func (a *Attribute) SetKey(v string) {
	if a != nil {
		a.key = v
	}
}

func (a *Attribute) GetValue() string {
	if a != nil {
		return a.value
	}

	return ""
}

func (a *Attribute) SetValue(v string) {
	if a != nil {
		a.value = v
	}
}

func (ni *NodeInfo) GetPublicKey() []byte {
	if ni != nil {
		return ni.publicKey
	}

	return nil
}

func (ni *NodeInfo) SetPublicKey(v []byte) {
	if ni != nil {
		ni.publicKey = v
	}
}

func (ni *NodeInfo) GetAddress() string {
	if ni != nil {
		return ni.address
	}

	return ""
}

func (ni *NodeInfo) SetAddress(v string) {
	if ni != nil {
		ni.address = v
	}
}

func (ni *NodeInfo) GetAttributes() []*Attribute {
	if ni != nil {
		return ni.attributes
	}

	return nil
}

func (ni *NodeInfo) SetAttributes(v []*Attribute) {
	if ni != nil {
		ni.attributes = v
	}
}
