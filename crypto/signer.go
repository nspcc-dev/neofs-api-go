package neofscrypto

import (
	"encoding"
)

// Signer is an interface for an opaque private key that can be used for signing operations.
// It is an adaptation of crypto.Signer to the needs of NeoFS.
type Signer interface {
	// Sign signs data with the private key.
	Sign(data []byte) ([]byte, error)

	// Public returns the public key corresponding to the Signer.
	Public() PublicKey
}

// PublicKey represents a public key using a particular algorithm.
type PublicKey interface {
	// Verify checks if signature calculated from
	// the data via corresponding Signer is correct.
	Verify(data, signature []byte) bool

	// Binary format for network transmission.
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}
