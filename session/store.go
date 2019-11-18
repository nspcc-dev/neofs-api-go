package session

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"sync"

	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/nspcc-dev/neofs-proto/refs"
)

type simpleStore struct {
	*sync.RWMutex

	tokens map[TokenID]*PToken
}

// TODO get curve from neofs-crypto
func defaultCurve() elliptic.Curve {
	return elliptic.P256()
}

// NewSimpleStore creates simple token storage
func NewSimpleStore() TokenStore {
	return &simpleStore{
		RWMutex: new(sync.RWMutex),
		tokens:  make(map[TokenID]*PToken),
	}
}

// New returns new token with specified parameters.
func (s *simpleStore) New(p TokenParams) *PToken {
	tid, err := refs.NewUUID()
	if err != nil {
		return nil
	}

	key, err := ecdsa.GenerateKey(defaultCurve(), rand.Reader)
	if err != nil {
		return nil
	}

	if p.FirstEpoch > p.LastEpoch || p.OwnerID.Empty() {
		return nil
	}

	t := &PToken{
		mtx: new(sync.Mutex),
		Token: Token{
			ID:         tid,
			Header:     VerificationHeader{PublicKey: crypto.MarshalPublicKey(&key.PublicKey)},
			FirstEpoch: p.FirstEpoch,
			LastEpoch:  p.LastEpoch,
			ObjectID:   p.ObjectID,
			OwnerID:    p.OwnerID,
		},
		PrivateKey: key,
	}

	s.Lock()
	s.tokens[t.ID] = t
	s.Unlock()

	return t
}

// Fetch tries to fetch a token with specified id.
func (s *simpleStore) Fetch(id TokenID) *PToken {
	s.RLock()
	defer s.RUnlock()

	return s.tokens[id]
}

// Remove removes token with id from store.
func (s *simpleStore) Remove(id TokenID) {
	s.Lock()
	delete(s.tokens, id)
	s.Unlock()
}
