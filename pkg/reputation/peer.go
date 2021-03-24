package reputation

// PeerID represents peer ID compatible with NeoFS API v2.
type PeerID []byte

// NewPeerID creates and returns blank PeerID.
func NewPeerID() *PeerID {
	return PeerIDFromV2(nil)
}

// PeerIDFromV2 converts bytes slice to PeerID.
func PeerIDFromV2(data []byte) *PeerID {
	return (*PeerID)(&data)
}

// SetBytes sets bytes of peer ID.
func (x *PeerID) SetBytes(v []byte) {
	*x = v
}

// Bytes returns bytes of peer ID.
func (x PeerID) Bytes() []byte {
	return x
}

// ToV2 converts PeerID to byte slice.
func (x PeerID) ToV2() []byte {
	return x
}
