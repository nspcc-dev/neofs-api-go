package session

import (
	"sync"
)

type mapTokenStore struct {
	*sync.RWMutex

	tokens map[TokenID]PrivateToken
}

// NewMapTokenStore creates new PrivateTokenStore instance.
//
// The elements of the instance are stored in the map.
func NewMapTokenStore() PrivateTokenStore {
	return &mapTokenStore{
		RWMutex: new(sync.RWMutex),
		tokens:  make(map[TokenID]PrivateToken),
	}
}

// Store adds passed token to the map.
//
// Resulting error is always nil.
func (s *mapTokenStore) Store(id TokenID, token PrivateToken) error {
	s.Lock()
	s.tokens[id] = token
	s.Unlock()

	return nil
}

// Fetch returns the map element corresponding to the given key.
//
// Returns ErrPrivateTokenNotFound is there is no element in map.
func (s *mapTokenStore) Fetch(id TokenID) (PrivateToken, error) {
	s.RLock()
	defer s.RUnlock()

	t, ok := s.tokens[id]
	if !ok {
		return nil, ErrPrivateTokenNotFound
	}

	return t, nil
}

// RemoveExpired removes all the map elements that are expired in the passed epoch.
//
// Resulting error is always nil.
func (s *mapTokenStore) RemoveExpired(epoch uint64) error {
	s.Lock()

	for key, token := range s.tokens {
		if token.Expired(epoch) {
			delete(s.tokens, key)
		}
	}

	s.Unlock()

	return nil
}
