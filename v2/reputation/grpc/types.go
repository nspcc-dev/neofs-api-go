package reputation

// SetPeer sets trusted peer's ID.
func (x *Trust) SetPeer(v []byte) {
	if x != nil {
		x.Peer = v
	}
}

// SetValue sets trust value.
func (x *Trust) SetValue(v float64) {
	if x != nil {
		x.Value = v
	}
}
