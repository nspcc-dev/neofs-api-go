package session

import (
	"crypto/ecdsa"
	"encoding/binary"
	"sync"

	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/nspcc-dev/neofs-proto/chain"
	"github.com/nspcc-dev/neofs-proto/internal"
	"github.com/nspcc-dev/neofs-proto/refs"
	"github.com/pkg/errors"
)

type (
	// ObjectID type alias.
	ObjectID = refs.ObjectID
	// OwnerID type alias.
	OwnerID = refs.OwnerID
	// TokenID type alias.
	TokenID = refs.UUID

	// PToken is a wrapper around Token that allows to sign data
	// and to do thread-safe manipulations.
	PToken struct {
		Token

		mtx        *sync.Mutex
		PrivateKey *ecdsa.PrivateKey
	}
)

const (
	// ErrWrongFirstEpoch is raised when passed Token contains wrong first epoch.
	// First epoch is an epoch since token is valid
	ErrWrongFirstEpoch = internal.Error("wrong first epoch")

	// ErrWrongLastEpoch is raised when passed Token contains wrong last epoch.
	// Last epoch is an epoch until token is valid
	ErrWrongLastEpoch = internal.Error("wrong last epoch")

	// ErrWrongOwner is raised when passed Token contains wrong OwnerID.
	ErrWrongOwner = internal.Error("wrong owner")

	// ErrEmptyPublicKey is raised when passed Token contains wrong public key.
	ErrEmptyPublicKey = internal.Error("empty public key")

	// ErrWrongObjectsCount is raised when passed Token contains wrong objects count.
	ErrWrongObjectsCount = internal.Error("wrong objects count")

	// ErrWrongObjects is raised when passed Token contains wrong object ids.
	ErrWrongObjects = internal.Error("wrong objects")

	// ErrInvalidSignature is raised when wrong signature is passed to VerificationHeader.VerifyData().
	ErrInvalidSignature = internal.Error("invalid signature")
)

// verificationData returns byte array to sign.
// Note: protobuf serialization is inconsistent as
// wire order is unspecified.
func (m *Token) verificationData() (data []byte) {
	var size int
	if l := len(m.ObjectID); l > 0 {
		size = m.ObjectID[0].Size()
		data = make([]byte, 16+l*size)
	} else {
		data = make([]byte, 16)
	}
	binary.BigEndian.PutUint64(data, m.FirstEpoch)
	binary.BigEndian.PutUint64(data[8:], m.LastEpoch)
	for i := range m.ObjectID {
		copy(data[16+i*size:], m.ObjectID[i].Bytes())
	}
	return
}

// IsSame checks if the passed token is valid and equal to current token
func (m *Token) IsSame(t *Token) error {
	switch {
	case m.FirstEpoch != t.FirstEpoch:
		return ErrWrongFirstEpoch
	case m.LastEpoch != t.LastEpoch:
		return ErrWrongLastEpoch
	case !m.OwnerID.Equal(t.OwnerID):
		return ErrWrongOwner
	case m.Header.PublicKey == nil:
		return ErrEmptyPublicKey
	case len(m.ObjectID) != len(t.ObjectID):
		return ErrWrongObjectsCount
	default:
		for i := range m.ObjectID {
			if !m.ObjectID[i].Equal(t.ObjectID[i]) {
				return errors.Wrapf(ErrWrongObjects, "expect %s, actual: %s", m.ObjectID[i], t.ObjectID[i])
			}
		}
	}
	return nil
}

// Sign tries to sign current Token data and stores signature inside it.
func (m *Token) Sign(key *ecdsa.PrivateKey) error {
	if err := m.Header.Sign(key); err != nil {
		return err
	}

	s, err := crypto.Sign(key, m.verificationData())
	if err != nil {
		return err
	}

	m.Signature = s
	return nil
}

// SetPublicKeys sets owner's public keys to the token
func (m *Token) SetPublicKeys(keys... *ecdsa.PublicKey) {
	m.PublicKeys = m.PublicKeys[:0]
	for i := range keys {
		m.PublicKeys = append(m.PublicKeys, crypto.MarshalPublicKey(keys[i]))
	}
}

// Verify checks if token is correct and signed.
func (m *Token) Verify(keys ...*ecdsa.PublicKey) bool {
	if m.FirstEpoch > m.LastEpoch {
		return false
	}
	ownerFromKeys := chain.KeysToAddress(keys...)
	if m.OwnerID.String() != ownerFromKeys {
		return false
	}

	for i := range keys {
		if m.Header.Verify(keys[i]) && crypto.Verify(keys[i], m.verificationData(), m.Signature) == nil {
			return true
		}
	}
	return false
}

// AddSignatures adds token signatures.
func (t *PToken) AddSignatures(signH, signT []byte) {
	t.mtx.Lock()

	t.Header.KeySignature = signH
	t.Signature = signT

	t.mtx.Unlock()
}

// SignData signs data with session private key.
func (t *PToken) SignData(data []byte) ([]byte, error) {
	return crypto.Sign(t.PrivateKey, data)
}

// VerifyData checks if signature of data by token is equal to sign.
func (m *VerificationHeader) VerifyData(data, sign []byte) error {
	if crypto.Verify(crypto.UnmarshalPublicKey(m.PublicKey), data, sign) != nil {
		return ErrInvalidSignature
	}
	return nil
}

// Verify checks if verification header was issued by id.
func (m *VerificationHeader) Verify(keys ...*ecdsa.PublicKey) bool {
	for i := range keys {
		if crypto.Verify(keys[i], m.PublicKey, m.KeySignature) == nil {
			return true
		}
	}
	return false
}

// UnmarshalPublicKeys returns unmarshal public keys from the token
func UnmarshalPublicKeys(t *Token) []*ecdsa.PublicKey {
	r := make([]*ecdsa.PublicKey, 0, len(t.PublicKeys))
	for i := range t.PublicKeys {
		r = append(r, crypto.UnmarshalPublicKey(t.PublicKeys[i]))
	}
	return r
}
