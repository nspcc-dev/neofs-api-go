package session

import (
	session "github.com/nspcc-dev/neofs-api-go/v2/session/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func (c *ObjectSessionContext) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		ObjectSessionContextToGRPCMessage(c),
	)
}

func (c *ObjectSessionContext) UnmarshalJSON(data []byte) error {
	msg := new(session.ObjectSessionContext)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*c = *ObjectSessionContextFromGRPCMessage(msg)

	return nil
}

func (l *TokenLifetime) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		TokenLifetimeToGRPCMessage(l),
	)
}

func (l *TokenLifetime) UnmarshalJSON(data []byte) error {
	msg := new(session.SessionToken_Body_TokenLifetime)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*l = *TokenLifetimeFromGRPCMessage(msg)

	return nil
}

func (t *SessionTokenBody) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		SessionTokenBodyToGRPCMessage(t),
	)
}

func (t *SessionTokenBody) UnmarshalJSON(data []byte) error {
	msg := new(session.SessionToken_Body)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*t = *SessionTokenBodyFromGRPCMessage(msg)

	return nil
}
