package token

import (
	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	"github.com/nspcc-dev/neofs-api-go/v2/session"
)

type SessionToken session.SessionToken

func NewSessionTokenFromV2(tV2 *session.SessionToken) *SessionToken {
	return (*SessionToken)(tV2)
}

func NewSessionToken() *SessionToken {
	return NewSessionTokenFromV2(new(session.SessionToken))
}

func (t *SessionToken) ToV2() *session.SessionToken {
	return (*session.SessionToken)(t)
}

func (t *SessionToken) setBodyField(setter func(*session.SessionTokenBody)) {
	token := (*session.SessionToken)(t)
	body := token.GetBody()

	if body == nil {
		body = new(session.SessionTokenBody)
		token.SetBody(body)
	}

	setter(body)
}

func (t *SessionToken) ID() []byte {
	return (*session.SessionToken)(t).
		GetBody().
		GetID()
}

func (t *SessionToken) SetID(v []byte) {
	t.setBodyField(func(body *session.SessionTokenBody) {
		body.SetID(v)
	})
}

func (t *SessionToken) OwnerID() *owner.ID {
	return owner.NewIDFromV2(
		(*session.SessionToken)(t).
			GetBody().
			GetOwnerID(),
	)
}

func (t *SessionToken) SetOwnerID(v *owner.ID) {
	t.setBodyField(func(body *session.SessionTokenBody) {
		body.SetOwnerID(v.ToV2())
	})
}

func (t *SessionToken) SessionKey() []byte {
	return (*session.SessionToken)(t).
		GetBody().
		GetSessionKey()
}

func (t *SessionToken) SetSessionKey(v []byte) {
	t.setBodyField(func(body *session.SessionTokenBody) {
		body.SetSessionKey(v)
	})
}

func (t *SessionToken) Signature() *pkg.Signature {
	return pkg.NewSignatureFromV2(
		(*session.SessionToken)(t).
			GetSignature(),
	)
}

// Marshal marshals SessionToken into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (t *SessionToken) Marshal(bs ...[]byte) ([]byte, error) {
	var buf []byte
	if len(bs) > 0 {
		buf = bs[0]
	}

	return (*session.SessionToken)(t).
		StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of SessionToken.
func (t *SessionToken) Unmarshal(data []byte) error {
	tV2 := new(session.SessionToken)
	if err := tV2.Unmarshal(data); err != nil {
		return err
	}

	*t = *NewSessionTokenFromV2(tV2)

	return nil
}

// MarshalJSON encodes SessionToken to protobuf JSON format.
func (t *SessionToken) MarshalJSON() ([]byte, error) {
	return (*session.SessionToken)(t).
		MarshalJSON()
}

// UnmarshalJSON decodes SessionToken from protobuf JSON format.
func (t *SessionToken) UnmarshalJSON(data []byte) error {
	tV2 := new(session.SessionToken)
	if err := tV2.UnmarshalJSON(data); err != nil {
		return err
	}

	*t = *NewSessionTokenFromV2(tV2)

	return nil
}
