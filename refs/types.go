// This package contains basic structures implemented in Go, such as
//
// CID - container id
// OwnerID - owner id
// ObjectID - object id
// SGID - storage group id
// Address - contains object id and container id
// UUID - a 128 bit (16 byte) Universal Unique Identifier as defined in RFC 4122

package refs

import (
	"crypto/sha256"

	"github.com/google/uuid"
	"github.com/nspcc-dev/neofs-api-go/chain"
	"github.com/nspcc-dev/neofs-api-go/internal"
)

type (
	// CID is implementation of ContainerID.
	CID [CIDSize]byte

	// UUID wrapper over github.com/google/uuid.UUID.
	UUID uuid.UUID

	// SGID is type alias of UUID.
	SGID = UUID

	// ObjectID is type alias of UUID.
	ObjectID = UUID

	// MessageID is type alias of UUID.
	MessageID = UUID

	// OwnerID is wrapper over neofs-proto/chain.WalletAddress.
	OwnerID chain.WalletAddress
)

// OwnerIDSource is an interface of the container of an OwnerID value with read access.
type OwnerIDSource interface {
	GetOwnerID() OwnerID
}

// OwnerIDContainer is an interface of the container of an OwnerID value.
type OwnerIDContainer interface {
	OwnerIDSource
	SetOwnerID(OwnerID)
}

// AddressContainer is an interface of the container of object address value.
type AddressContainer interface {
	GetAddress() Address
	SetAddress(Address)
}

const (
	// UUIDSize contains size of UUID.
	UUIDSize = 16

	// SGIDSize contains size of SGID.
	SGIDSize = UUIDSize

	// CIDSize contains size of CID.
	CIDSize = sha256.Size

	// OwnerIDSize contains size of OwnerID.
	OwnerIDSize = chain.AddressLength

	// ErrWrongDataSize is raised when passed bytes into Unmarshal have wrong size.
	ErrWrongDataSize = internal.Error("wrong data size")

	// ErrEmptyOwner is raised when empty OwnerID is passed into container.New.
	ErrEmptyOwner = internal.Error("owner cant be empty")

	// ErrEmptyCapacity is raised when empty Capacity is passed container.New.
	ErrEmptyCapacity = internal.Error("capacity cant be empty")

	// ErrEmptyContainer is raised when it CID method is called for an empty container.
	ErrEmptyContainer = internal.Error("cannot return ID for empty container")
)

var (
	emptyCID   = (CID{}).Bytes()
	emptyUUID  = (UUID{}).Bytes()
	emptyOwner = (OwnerID{}).Bytes()

	_ internal.Custom = (*CID)(nil)
	_ internal.Custom = (*SGID)(nil)
	_ internal.Custom = (*UUID)(nil)
	_ internal.Custom = (*OwnerID)(nil)
	_ internal.Custom = (*ObjectID)(nil)
	_ internal.Custom = (*MessageID)(nil)

	// NewSGID method alias.
	NewSGID = NewUUID

	// NewObjectID method alias.
	NewObjectID = NewUUID

	// NewMessageID method alias.
	NewMessageID = NewUUID
)

// NewUUID returns a Random (Version 4) UUID.
//
// The strength of the UUIDs is based on the strength of the crypto/rand
// package.
//
// A note about uniqueness derived from the UUID Wikipedia entry:
//
//  Randomly generated UUIDs have 122 random bits.  One's annual risk of being
//  hit by a meteorite is estimated to be one chance in 17 billion, that
//  means the probability is about 0.00000000006 (6 × 10−11),
//  equivalent to the odds of creating a few tens of trillions of UUIDs in a
//  year and having one duplicate.
func NewUUID() (UUID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return UUID{}, err
	}
	return UUID(id), nil
}
