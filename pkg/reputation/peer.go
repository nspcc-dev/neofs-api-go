package reputation

import (
	"bytes"

	"github.com/mr-tron/base58"
	"github.com/nspcc-dev/neofs-api-go/v2/reputation"
)

// PeerID represents peer ID compatible with NeoFS API v2.
type PeerID reputation.PeerID

// NewPeerID creates and returns blank PeerID.
//
// Defaults:
//  - publicKey: nil.
func NewPeerID() *PeerID {
	return PeerIDFromV2(new(reputation.PeerID))
}

// PeerIDFromV2 converts NeoFS API v2 reputation.PeerID message to PeerID.
//
// Nil reputation.PeerID converts to nil.
func PeerIDFromV2(id *reputation.PeerID) *PeerID {
	return (*PeerID)(id)
}

// SetPublicKey sets peer ID as a compressed public key.
func (x *PeerID) SetPublicKey(v []byte) {
	(*reputation.PeerID)(x).
		SetPublicKey(v)
}

// ToV2 converts PeerID to NeoFS API v2 reputation.PeerID message.
//
// Nil PeerID converts to nil.
func (x *PeerID) ToV2() *reputation.PeerID {
	return (*reputation.PeerID)(x)
}

// Equal returns true if identifiers are identical.
func (x *PeerID) Equal(x2 *PeerID) bool {
	return bytes.Equal(
		(*reputation.PeerID)(x).GetPublicKey(),
		(*reputation.PeerID)(x2).GetPublicKey(),
	)
}

// Parse parses PeerID from base58 string.
func (x *PeerID) Parse(s string) error {
	data, err := base58.Decode(s)
	if err != nil {
		return err
	}

	(*reputation.PeerID)(x).SetPublicKey(data)

	return nil
}

// String returns base58 string representation of PeerID.
func (x *PeerID) String() string {
	return base58.Encode(
		(*reputation.PeerID)(x).
			GetPublicKey(),
	)
}

// Marshal marshals PeerID into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (x *PeerID) Marshal(b ...[]byte) ([]byte, error) {
	var buf []byte
	if len(b) > 0 {
		buf = b[0]
	}

	return (*reputation.PeerID)(x).
		StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of PeerID.
func (x *PeerID) Unmarshal(data []byte) error {
	return (*reputation.PeerID)(x).
		Unmarshal(data)
}

// MarshalJSON encodes PeerID to protobuf JSON format.
func (x *PeerID) MarshalJSON() ([]byte, error) {
	return (*reputation.PeerID)(x).
		MarshalJSON()
}

// UnmarshalJSON decodes PeerID from protobuf JSON format.
func (x *PeerID) UnmarshalJSON(data []byte) error {
	return (*reputation.PeerID)(x).
		UnmarshalJSON(data)
}
