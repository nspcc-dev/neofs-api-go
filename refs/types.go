package refs

type OwnerID struct {
	val []byte
}

type ContainerID struct {
	val []byte
}

type ObjectID struct {
	val []byte
}

type Address struct {
	cid *ContainerID

	oid *ObjectID
}

type Checksum struct {
	typ ChecksumType

	sum []byte
}

type ChecksumType uint32

type Signature struct {
	key, sign []byte
}

type Version struct {
	major, minor uint32
}

const (
	UnknownChecksum ChecksumType = iota
	TillichZemor
	SHA256
)

func (o *OwnerID) GetValue() []byte {
	if o != nil {
		return o.val
	}

	return nil
}

func (o *OwnerID) SetValue(v []byte) {
	if o != nil {
		o.val = v
	}
}

func (c *ContainerID) GetValue() []byte {
	if c != nil {
		return c.val
	}

	return nil
}

func (c *ContainerID) SetValue(v []byte) {
	if c != nil {
		c.val = v
	}
}

func (o *ObjectID) GetValue() []byte {
	if o != nil {
		return o.val
	}

	return nil
}

func (o *ObjectID) SetValue(v []byte) {
	if o != nil {
		o.val = v
	}
}

func (a *Address) GetContainerID() *ContainerID {
	if a != nil {
		return a.cid
	}

	return nil
}

func (a *Address) SetContainerID(v *ContainerID) {
	if a != nil {
		a.cid = v
	}
}

func (a *Address) GetObjectID() *ObjectID {
	if a != nil {
		return a.oid
	}

	return nil
}

func (a *Address) SetObjectID(v *ObjectID) {
	if a != nil {
		a.oid = v
	}
}

func (c *Checksum) GetType() ChecksumType {
	if c != nil {
		return c.typ
	}

	return UnknownChecksum
}

func (c *Checksum) SetType(v ChecksumType) {
	if c != nil {
		c.typ = v
	}
}

func (c *Checksum) GetSum() []byte {
	if c != nil {
		return c.sum
	}

	return nil
}

func (c *Checksum) SetSum(v []byte) {
	if c != nil {
		c.sum = v
	}
}

func (s *Signature) GetKey() []byte {
	if s != nil {
		return s.key
	}

	return nil
}

func (s *Signature) SetKey(v []byte) {
	if s != nil {
		s.key = v
	}
}

func (s *Signature) GetSign() []byte {
	if s != nil {
		return s.sign
	}

	return nil
}

func (s *Signature) SetSign(v []byte) {
	if s != nil {
		s.sign = v
	}
}

func (v *Version) GetMajor() uint32 {
	if v != nil {
		return v.major
	}

	return 0
}

func (v *Version) SetMajor(val uint32) {
	if v != nil {
		v.major = val
	}
}

func (v *Version) GetMinor() uint32 {
	if v != nil {
		return v.minor
	}

	return 0
}

func (v *Version) SetMinor(val uint32) {
	if v != nil {
		v.minor = val
	}
}
