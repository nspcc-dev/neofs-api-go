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

// SetMajor sets major version number.
func (m *Version) SetMajor(v uint32) {
	if m != nil {
		m.Major = v
	}
}

// SetMinor sets minor version number.
func (m *Version) SetMinor(v uint32) {
	if m != nil {
		m.Minor = v
	}
}

// SetKey sets public key in a binary format.
func (m *Signature) SetKey(v []byte) {
	if m != nil {
		m.Key = v
	}
}

// SetSign sets signature.
func (m *Signature) SetSign(v []byte) {
	if m != nil {
		m.Sign = v
	}
}
