package cryptoalgo

// SignatureAlgorithm represents enumeration of
// cryptographic signature algorithms.
type SignatureAlgorithm int

const (
	_ SignatureAlgorithm = iota - 1

	// ECDSA is a SignatureAlgorithm for Elliptic Curve Digital Signature Algorithm,
	// as defined in FIPS 186-3.
	ECDSA

	// RFC6979 is a SignatureAlgorithm for RFC 6979's deterministic DSA.
	RFC6979
)
