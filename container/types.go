package container

import (
	"bytes"

	"github.com/google/uuid"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/nspcc-dev/neofs-proto/internal"
	"github.com/nspcc-dev/neofs-proto/refs"
	"github.com/nspcc-dev/netmap"
	"github.com/pkg/errors"
)

// AccessMode is a container access mode type.
type AccessMode uint32

const (
	// AccessModeRead is a read access mode.
	AccessModeRead AccessMode = 1 << iota
	// AccessModeWrite is a write access mode.
	AccessModeWrite
)

// AccessModeReadWrite is a read/write container access mode.
const AccessModeReadWrite = AccessModeRead | AccessModeWrite

var (
	_ internal.Custom = (*Container)(nil)

	emptySalt  = (UUID{}).Bytes()
	emptyOwner = (OwnerID{}).Bytes()
)

// New creates new user container based on capacity, OwnerID and PlacementRules.
func New(cap uint64, owner OwnerID, rules netmap.PlacementRule) (*Container, error) {
	if bytes.Equal(owner[:], emptyOwner) {
		return nil, refs.ErrEmptyOwner
	} else if cap == 0 {
		return nil, refs.ErrEmptyCapacity
	}

	salt, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "could not create salt")
	}

	return &Container{
		OwnerID:  owner,
		Salt:     UUID(salt),
		Capacity: cap,
		Rules:    rules,
	}, nil
}

// Bytes returns bytes representation of Container.
func (m *Container) Bytes() []byte {
	data, err := m.Marshal()
	if err != nil {
		return nil
	}

	return data
}

// ID returns generated ContainerID based on Container (data).
func (m *Container) ID() (CID, error) {
	if m.Empty() {
		return CID{}, refs.ErrEmptyContainer
	}
	data, err := m.Marshal()
	if err != nil {
		return CID{}, err
	}

	return refs.CIDForBytes(data), nil
}

// Empty checks that container is empty.
func (m *Container) Empty() bool {
	return m.Capacity == 0 || bytes.Equal(m.Salt.Bytes(), emptySalt) || bytes.Equal(m.OwnerID.Bytes(), emptyOwner)
}

// -- Test container definition -- //

// NewTestContainer returns test container.
// WARNING: DON'T USE THIS OUTSIDE TESTS.
func NewTestContainer() (*Container, error) {
	key := test.DecodeKey(0)
	owner, err := refs.NewOwnerID(&key.PublicKey)
	if err != nil {
		return nil, err
	}
	return New(100, owner, netmap.PlacementRule{
		ReplFactor: 2,
		SFGroups: []netmap.SFGroup{
			{
				Selectors: []netmap.Select{
					{Key: "Country", Count: 1},
					{Key: netmap.NodesBucket, Count: 2},
				},
				Filters: []netmap.Filter{
					{Key: "Country", F: netmap.FilterIn("USA")},
				},
			},
		},
	})
}
