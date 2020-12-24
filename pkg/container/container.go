package container

import (
	"crypto/sha256"

	"github.com/google/uuid"
	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/netmap"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	"github.com/nspcc-dev/neofs-api-go/v2/container"
)

type Container struct {
	v2 container.Container
}

func New(opts ...NewOption) *Container {
	cnrOptions := defaultContainerOptions()

	for i := range opts {
		opts[i].apply(&cnrOptions)
	}

	cnr := new(Container)
	cnr.SetNonceUUID(cnrOptions.nonce)
	cnr.SetBasicACL(cnrOptions.acl)

	if cnrOptions.owner != nil {
		cnr.SetOwnerID(cnrOptions.owner)
	}

	if cnrOptions.policy != nil {
		cnr.SetPlacementPolicy(cnrOptions.policy)
	}

	cnr.SetAttributes(cnrOptions.attributes)

	return cnr
}

func (c *Container) ToV2() *container.Container {
	return &c.v2
}

// NewVerifiedFromV2 constructs Container from NeoFS API V2 Container message.
//
// Does not perform if message meets NeoFS API V2 specification. To do this
// use NewVerifiedFromV2 constructor.
func NewContainerFromV2(c *container.Container) *Container {
	cnr := new(Container)

	if c != nil {
		cnr.v2 = *c
	}

	return cnr
}

// CalculateID calculates container identifier
// based on its structure.
func CalculateID(c *Container) *ID {
	data, err := c.ToV2().StableMarshal(nil)
	if err != nil {
		panic(err)
	}

	id := NewID()
	id.SetSHA256(sha256.Sum256(data))

	return id
}

func (c *Container) Version() *pkg.Version {
	return pkg.NewVersionFromV2(c.v2.GetVersion())
}

func (c *Container) SetVersion(v *pkg.Version) {
	c.v2.SetVersion(v.ToV2())
}

func (c *Container) OwnerID() *owner.ID {
	return owner.NewIDFromV2(c.v2.GetOwnerID())
}

func (c *Container) SetOwnerID(v *owner.ID) {
	c.v2.SetOwnerID(v.ToV2())
}

// Nonce returns container nonce in a binary format.
//
// Returns nil if container nonce is not a valid UUID.
//
// Deprecated: use NonceUUID instead.
func (c *Container) Nonce() []byte {
	uid, err := c.NonceUUID()
	if err == nil {
		data, _ := uid.MarshalBinary()
		return data
	}

	return nil
}

// SetNonce sets container nonce in a binary format.
//
// If slice length is less than UUID size, than
// value is padded with a sequence of zeros.
// If slice length is more than UUID size, than
// value is cut.
//
// Deprecated: use SetNonceUUID instead.
func (c *Container) SetNonce(v []byte) {
	u := uuid.UUID{}
	copy(u[:], v)
	c.v2.SetNonce(u[:])
}

// Returns container nonce in UUID format.
//
// Returns error if container nonce is not a valid UUID.
func (c *Container) NonceUUID() (uuid.UUID, error) {
	return uuid.FromBytes(c.v2.GetNonce())
}

// SetNonceUUID sets container nonce as UUID.
func (c *Container) SetNonceUUID(v uuid.UUID) {
	data, _ := v.MarshalBinary()
	c.v2.SetNonce(data)
}

func (c *Container) BasicACL() uint32 {
	return c.v2.GetBasicACL()
}

func (c *Container) SetBasicACL(v uint32) {
	c.v2.SetBasicACL(v)
}

func (c *Container) Attributes() Attributes {
	return NewAttributesFromV2(c.v2.GetAttributes())
}

func (c *Container) SetAttributes(v Attributes) {
	c.v2.SetAttributes(v.ToV2())
}

func (c *Container) PlacementPolicy() *netmap.PlacementPolicy {
	return netmap.NewPlacementPolicyFromV2(c.v2.GetPlacementPolicy())
}

func (c *Container) SetPlacementPolicy(v *netmap.PlacementPolicy) {
	c.v2.SetPlacementPolicy(v.ToV2())
}

// Marshal marshals Container into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (c *Container) Marshal(b ...[]byte) ([]byte, error) {
	var buf []byte
	if len(b) > 0 {
		buf = b[0]
	}

	return c.v2.
		StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of Container.
func (c *Container) Unmarshal(data []byte) error {
	return c.v2.
		Unmarshal(data)
}

// MarshalJSON encodes Container to protobuf JSON format.
func (c *Container) MarshalJSON() ([]byte, error) {
	return c.v2.
		MarshalJSON()
}

// UnmarshalJSON decodes Container from protobuf JSON format.
func (c *Container) UnmarshalJSON(data []byte) error {
	return c.v2.
		UnmarshalJSON(data)
}
