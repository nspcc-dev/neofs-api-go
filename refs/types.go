package refs

import (
	"fmt"
	"strconv"
)

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

type SubnetID struct {
	value uint32
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

func (s *SubnetID) SetValue(id uint32) {
	if s != nil {
		s.value = id
	}
}

func (s *SubnetID) GetValue() uint32 {
	if s != nil {
		return s.value
	}
	return 0
}

// MarshalText encodes SubnetID into text format according to NeoFS API V2 protocol:
// value in base-10 integer string format.
//
// Implements encoding.TextMarshaler.
func (s *SubnetID) MarshalText() ([]byte, error) {
	num := s.GetValue() // NPE safe, returns zero on nil (zero subnet)

	return []byte(strconv.FormatUint(uint64(num), 10)), nil
}

// UnmarshalText decodes SubnetID from the text according to NeoFS API V2 protocol:
// should be base-10 integer string format with bitsize = 32.
//
// Returns strconv.ErrRange if integer overflows uint32.
//
// Must not be called on nil.
//
// Implements encoding.TextUnmarshaler.
func (s *SubnetID) UnmarshalText(txt []byte) error {
	num, err := strconv.ParseUint(string(txt), 10, 32)
	if err != nil {
		return fmt.Errorf("invalid numeric value: %w", err)
	}

	s.value = uint32(num)

	return nil
}

// IsZeroSubnet returns true iff the SubnetID refers to zero subnet.
func IsZeroSubnet(id *SubnetID) bool {
	return id.GetValue() == 0
}

// MakeZeroSubnet makes the SubnetID to refer to zero subnet.
func MakeZeroSubnet(id *SubnetID) {
	id.SetValue(0)
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
