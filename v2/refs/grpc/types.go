package refs

// SetValue sets container identifier in a binary format.
func (m *ContainerID) SetValue(v []byte) {
	if m != nil {
		m.Value = v
	}
}

// SetValue sets object identifier in a binary format.
func (m *ObjectID) SetValue(v []byte) {
	if m != nil {
		m.Value = v
	}
}

// SetValue sets owner identifier in a binary format.
func (m *OwnerID) SetValue(v []byte) {
	if m != nil {
		m.Value = v
	}
}

// SetContainerId sets container identifier of the address.
func (m *Address) SetContainerId(v *ContainerID) {
	if m != nil {
		m.ContainerId = v
	}
}

// SetObjectId sets object identifier of the address.
func (m *Address) SetObjectId(v *ObjectID) {
	if m != nil {
		m.ObjectId = v
	}
}

// SetChecksumType in generic checksum structure.
func (m *Checksum) SetChecksumType(v ChecksumType) {
	if m != nil {
		m.Type = v
	}
}

// SetChecksumSum in generic checksum structure.
func (m *Checksum) SetSum(v []byte) {
	if m != nil {
		m.Sum = v
	}
}
