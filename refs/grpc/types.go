package refs

// SetValue sets container identifier in a binary format.
func (x *ContainerID) SetValue(v []byte) {
	if x != nil {
		x.Value = v
	}
}

// SetValue sets object identifier in a binary format.
func (x *ObjectID) SetValue(v []byte) {
	if x != nil {
		x.Value = v
	}
}

// SetValue sets owner identifier in a binary format.
func (x *OwnerID) SetValue(v []byte) {
	if x != nil {
		x.Value = v
	}
}

// SetContainerId sets container identifier of the address.
func (x *Address) SetContainerId(v *ContainerID) {
	if x != nil {
		x.ContainerId = v
	}
}

// SetObjectId sets object identifier of the address.
func (x *Address) SetObjectId(v *ObjectID) {
	if x != nil {
		x.ObjectId = v
	}
}

// SetChecksumType in generic checksum structure.
func (x *Checksum) SetChecksumType(v ChecksumType) {
	if x != nil {
		x.Type = v
	}
}

// SetSum in generic checksum structure.
func (x *Checksum) SetSum(v []byte) {
	if x != nil {
		x.Sum = v
	}
}

// SetMajor sets major version number.
func (x *Version) SetMajor(v uint32) {
	if x != nil {
		x.Major = v
	}
}

// SetMinor sets minor version number.
func (x *Version) SetMinor(v uint32) {
	if x != nil {
		x.Minor = v
	}
}

// SetKey sets public key in a binary format.
func (x *Signature) SetKey(v []byte) {
	if x != nil {
		x.Key = v
	}
}

// SetSign sets signature.
func (x *Signature) SetSign(v []byte) {
	if x != nil {
		x.Sign = v
	}
}

// SetScheme sets signature scheme.
func (x *Signature) SetScheme(s SignatureScheme) {
	if x != nil {
		x.Scheme = s
	}
}

// SetKey sets public key in a binary format.
func (x *SignatureRFC6979) SetKey(v []byte) {
	if x != nil {
		x.Key = v
	}
}

// SetSign sets signature.
func (x *SignatureRFC6979) SetSign(v []byte) {
	if x != nil {
		x.Sign = v
	}
}

// FromString parses SignatureScheme from a string representation,
// It is a reverse action to String().
//
// Returns true if s was parsed successfully.
func (x *SignatureScheme) FromString(s string) bool {
	i, ok := SignatureScheme_value[s]
	if ok {
		*x = SignatureScheme(i)
	}

	return ok
}

// FromString parses ChecksumType from a string representation,
// It is a reverse action to String().
//
// Returns true if s was parsed successfully.
func (x *ChecksumType) FromString(s string) bool {
	i, ok := ChecksumType_value[s]
	if ok {
		*x = ChecksumType(i)
	}

	return ok
}

// SetValue sets subnet identifier in a base-10 integer format.
func (x *SubnetID) SetValue(v uint32) {
	if x != nil {
		x.Value = v
	}
}
