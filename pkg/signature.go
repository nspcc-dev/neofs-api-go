package pkg

import (
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

// Signature represents v2-compatible signature.
type Signature refs.Signature

// NewSignatureFromV2 wraps v2 Signature message to Signature.
//
// Nil refs.Signature converts to nil.
func NewSignatureFromV2(sV2 *refs.Signature) *Signature {
	return (*Signature)(sV2)
}

// NewSignature creates and initializes blank Signature.
//
// Works similar as NewSignatureFromV2(new(Signature)).
//
// Defaults:
//  - key: nil;
//  - signature: nil.
func NewSignature() *Signature {
	return NewSignatureFromV2(new(refs.Signature))
}

// Key sets binary public key.
func (s *Signature) Key() []byte {
	return (*refs.Signature)(s).GetKey()
}

// SetKey returns binary public key.
func (s *Signature) SetKey(v []byte) {
	(*refs.Signature)(s).SetKey(v)
}

// Sign return signature value.
func (s *Signature) Sign() []byte {
	return (*refs.Signature)(s).GetSign()
}

// SetSign sets signature value.
func (s *Signature) SetSign(v []byte) {
	(*refs.Signature)(s).SetSign(v)
}

// ToV2 converts Signature to v2 Signature message.
//
// Nil Signature converts to nil.
func (s *Signature) ToV2() *refs.Signature {
	return (*refs.Signature)(s)
}

// Marshal marshals Signature into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (s *Signature) Marshal(b ...[]byte) ([]byte, error) {
	var buf []byte
	if len(b) > 0 {
		buf = b[0]
	}

	return (*refs.Signature)(s).
		StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of Signature.
func (s *Signature) Unmarshal(data []byte) error {
	return (*refs.Signature)(s).
		Unmarshal(data)
}

// MarshalJSON encodes Signature to protobuf JSON format.
func (s *Signature) MarshalJSON() ([]byte, error) {
	return (*refs.Signature)(s).
		MarshalJSON()
}

// UnmarshalJSON decodes Signature from protobuf JSON format.
func (s *Signature) UnmarshalJSON(data []byte) error {
	return (*refs.Signature)(s).
		UnmarshalJSON(data)
}
