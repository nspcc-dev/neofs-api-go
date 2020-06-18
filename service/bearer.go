package service

import (
	"crypto/ecdsa"
	"io"

	"github.com/nspcc-dev/neofs-api-go/refs"
	crypto "github.com/nspcc-dev/neofs-crypto"
)

type signedBearerToken struct {
	BearerToken
}

const fixedBearerTokenDataSize = 0 +
	refs.OwnerIDSize +
	8

// NewSignedBearerToken wraps passed BearerToken in a component suitable for signing.
//
// Result can be used in AddSignatureWithKey function.
func NewSignedBearerToken(token BearerToken) DataWithSignKeyAccumulator {
	return &signedBearerToken{
		BearerToken: token,
	}
}

// NewVerifiedBearerToken wraps passed SessionToken in a component suitable for signature verification.
//
// Result can be used in VerifySignatureWithKey function.
func NewVerifiedBearerToken(token BearerToken) DataWithSignature {
	return &signedBearerToken{
		BearerToken: token,
	}
}

// AddSignKey calls a Signature field setter and an OwnerKey field setter with corresponding arguments.
func (s signedBearerToken) AddSignKey(sig []byte, key *ecdsa.PublicKey) {
	if s.BearerToken != nil {
		s.SetSignature(sig)

		s.SetOwnerKey(
			crypto.MarshalPublicKey(key),
		)
	}
}

// SignedData returns token information in a binary representation.
func (s signedBearerToken) SignedData() ([]byte, error) {
	return SignedDataFromReader(s)
}

// SignedDataSize returns the length of signed token information slice.
func (s signedBearerToken) SignedDataSize() int {
	return bearerTokenInfoSize(s.BearerToken)
}

// ReadSignedData copies a binary representation of the token information to passed buffer.
//
// If buffer length is less than required, io.ErrUnexpectedEOF returns.
func (s signedBearerToken) ReadSignedData(p []byte) (int, error) {
	sz := s.SignedDataSize()
	if len(p) < sz {
		return 0, io.ErrUnexpectedEOF
	}

	copyBearerTokenSignedData(p, s.BearerToken)

	return sz, nil
}

func bearerTokenInfoSize(v ACLRulesSource) int {
	if v == nil {
		return 0
	}
	return fixedBearerTokenDataSize + len(v.GetACLRules())
}

// Fills passed buffer with signing token information bytes.
// Does not check buffer length, it is understood that enough space is allocated in it.
//
// If passed BearerTokenInfo, buffer remains unchanged.
func copyBearerTokenSignedData(buf []byte, token BearerTokenInfo) {
	if token == nil {
		return
	}

	var off int

	off += copy(buf[off:], token.GetACLRules())

	off += copy(buf[off:], token.GetOwnerID().Bytes())

	tokenEndianness.PutUint64(buf[off:], token.ExpirationEpoch())
	off += 8
}
