package container

import (
	"bytes"

	"github.com/google/uuid"
	"github.com/nspcc-dev/neofs-api-go/internal"
	"github.com/nspcc-dev/neofs-api-go/refs"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/nspcc-dev/netmap"
	"github.com/pkg/errors"
)

var (
	_ internal.Custom = (*Container)(nil)

	emptySalt  = (UUID{}).Bytes()
	emptyOwner = (OwnerID{}).Bytes()
)

// New creates new user container based on capacity, OwnerID, ACL and PlacementRules.
func New(cap uint64, owner OwnerID, acl uint32, rules netmap.PlacementRule) (*Container, error) {
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
		BasicACL: acl,
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
	return New(100, owner, 0xFFFFFFFF, netmap.PlacementRule{
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

// GetMessageID is a MessageID field getter.
func (m PutRequest) GetMessageID() MessageID {
	return m.MessageID
}

// SetMessageID is a MessageID field getter.
func (m *PutRequest) SetMessageID(id MessageID) {
	m.MessageID = id
}

// SetCapacity is a Capacity field setter.
func (m *PutRequest) SetCapacity(c uint64) {
	m.Capacity = c
}

// GetOwnerID is an OwnerID field getter.
func (m PutRequest) GetOwnerID() OwnerID {
	return m.OwnerID
}

// SetOwnerID is an OwnerID field setter.
func (m *PutRequest) SetOwnerID(owner OwnerID) {
	m.OwnerID = owner
}

// SetRules is a Rules field setter.
func (m *PutRequest) SetRules(rules netmap.PlacementRule) {
	m.Rules = rules
}

// SetBasicACL is a BasicACL field setter.
func (m *PutRequest) SetBasicACL(acl uint32) {
	m.BasicACL = acl
}

// GetCID is a CID field getter.
func (m DeleteRequest) GetCID() CID {
	return m.CID
}

// SetCID is a CID field setter.
func (m *DeleteRequest) SetCID(cid CID) {
	m.CID = cid
}

// GetCID is a CID field getter.
func (m GetRequest) GetCID() CID {
	return m.CID
}

// SetCID is a CID field setter.
func (m *GetRequest) SetCID(cid CID) {
	m.CID = cid
}

// GetOwnerID is an OwnerID field getter.
func (m ListRequest) GetOwnerID() OwnerID {
	return m.OwnerID
}

// SetOwnerID is an OwnerID field setter.
func (m *ListRequest) SetOwnerID(owner OwnerID) {
	m.OwnerID = owner
}

// CID is an ID field getter.
func (m ExtendedACLKey) CID() CID {
	return m.ID
}

// SetCID is an ID field setter.
func (m *ExtendedACLKey) SetCID(v CID) {
	m.ID = v
}

// SetEACL is an EACL field setter.
func (m *ExtendedACLValue) SetEACL(v []byte) {
	m.EACL = v
}

// SetSignature is a Signature field setter.
func (m *ExtendedACLValue) SetSignature(sig []byte) {
	m.Signature = sig
}
