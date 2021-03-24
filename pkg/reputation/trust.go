package reputation

import (
	"github.com/nspcc-dev/neofs-api-go/v2/reputation"
)

// Trust represents peer's trust compatible with NeoFS API v2.
type Trust reputation.Trust

// NewTrust creates and returns blank Trust.
func NewTrust() *Trust {
	return TrustFromV2(new(reputation.Trust))
}

// TrustFromV2 converts NeoFS API v2
// reputation.Trust message structure to Trust.
func TrustFromV2(t *reputation.Trust) *Trust {
	return (*Trust)(t)
}

// ToV2 converts Trust to NeoFS API v2
// reputation.Trust message structure.
func (x *Trust) ToV2() *reputation.Trust {
	return (*reputation.Trust)(x)
}

// TrustsToV2 converts slice of Trust's to slice of
// NeoFS API v2 reputation.Trust message structures.
func TrustsToV2(xs []*Trust) (res []*reputation.Trust) {
	if xs != nil {
		res = make([]*reputation.Trust, 0, len(xs))

		for i := range xs {
			res = append(res, xs[i].ToV2())
		}
	}

	return
}

// SetPeer sets trusted peer ID.
func (x *Trust) SetPeer(id *PeerID) {
	(*reputation.Trust)(x).
		SetPeer(id.ToV2())
}

// Peer returns trusted peer ID.
func (x *Trust) Peer() *PeerID {
	return PeerIDFromV2(
		(*reputation.Trust)(x).GetPeer(),
	)
}

// SetValue sets trust value.
func (x *Trust) SetValue(val float64) {
	(*reputation.Trust)(x).
		SetValue(val)
}

// Value returns trust value.
func (x *Trust) Value() float64 {
	return (*reputation.Trust)(x).
		GetValue()
}

// Marshal marshals Trust into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (x *Trust) Marshal(b ...[]byte) ([]byte, error) {
	var buf []byte
	if len(b) > 0 {
		buf = b[0]
	}

	return (*reputation.Trust)(x).StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of Trust.
func (x *Trust) Unmarshal(data []byte) error {
	return (*reputation.Trust)(x).
		Unmarshal(data)
}

// MarshalJSON encodes Trust to protobuf JSON format.
func (x *Trust) MarshalJSON() ([]byte, error) {
	return (*reputation.Trust)(x).
		MarshalJSON()
}

// UnmarshalJSON decodes Trust from protobuf JSON format.
func (x *Trust) UnmarshalJSON(data []byte) error {
	return (*reputation.Trust)(x).
		UnmarshalJSON(data)
}
