package token

import (
	"github.com/google/uuid"
)

type SessionToken struct {
	id     uuid.UUID
	pubKey []byte
}

func CreateSessionToken(id, pub []byte) (*SessionToken, error) {
	var tokenID uuid.UUID

	err := tokenID.UnmarshalBinary(id)
	if err != nil {
		return nil, err
	}

	key := make([]byte, len(pub))
	copy(key[:], pub)

	return &SessionToken{
		id:     tokenID,
		pubKey: key,
	}, nil
}

func (s SessionToken) SessionKey() []byte {
	return s.pubKey
}

func (s SessionToken) ID() []byte {
	return s.id[:]
}
