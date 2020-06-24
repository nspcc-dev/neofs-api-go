package service

import (
	"crypto/ecdsa"
	"io"

	"github.com/nspcc-dev/neofs-api-go/internal"
	crypto "github.com/nspcc-dev/neofs-crypto"
)

type signKeyPairsWrapper struct {
	items []SignKeyPair
}

// GetSessionToken returns SessionToken interface of Token field.
//
// If token field value is nil, nil returns.
func (m RequestVerificationHeader) GetSessionToken() SessionToken {
	if t := m.GetToken(); t != nil {
		return t
	}

	return nil
}

// AddSignKey adds new element to Signatures field.
//
// Sets Sign field to passed sign. Set Peer field to marshaled passed key.
func (m *RequestVerificationHeader) AddSignKey(sign []byte, key *ecdsa.PublicKey) {
	m.SetSignatures(
		append(
			m.GetSignatures(),
			&RequestVerificationHeader_Signature{
				Sign: sign,
				Peer: crypto.MarshalPublicKey(key),
			},
		),
	)
}

// GetSignKeyPairs returns the elements of Signatures field as SignKeyPair slice.
func (m RequestVerificationHeader) GetSignKeyPairs() []SignKeyPair {
	var (
		signs = m.GetSignatures()
		res   = make([]SignKeyPair, len(signs))
	)

	for i := range signs {
		res[i] = signs[i]
	}

	return res
}

// GetSignature returns the result of a Sign field getter.
func (m RequestVerificationHeader_Signature) GetSignature() []byte {
	return m.GetSign()
}

// GetPublicKey unmarshals and returns the result of a Peer field getter.
func (m RequestVerificationHeader_Signature) GetPublicKey() *ecdsa.PublicKey {
	return crypto.UnmarshalPublicKey(m.GetPeer())
}

// SetSignatures replaces signatures stored in RequestVerificationHeader.
func (m *RequestVerificationHeader) SetSignatures(signatures []*RequestVerificationHeader_Signature) {
	m.Signatures = signatures
}

// SetToken is a Token field setter.
func (m *RequestVerificationHeader) SetToken(token *Token) {
	m.Token = token
}

// SetBearer is a Bearer field setter.
func (m *RequestVerificationHeader) SetBearer(v *BearerTokenMsg) {
	m.Bearer = v
}

// testCustomField for test usage only.
type testCustomField [8]uint32

var _ internal.Custom = (*testCustomField)(nil)

// Reset skip, it's for test usage only.
func (t testCustomField) Reset() {}

// ProtoMessage skip, it's for test usage only.
func (t testCustomField) ProtoMessage() {}

// Size skip, it's for test usage only.
func (t testCustomField) Size() int { return 32 }

// String skip, it's for test usage only.
func (t testCustomField) String() string { return "" }

// Bytes skip, it's for test usage only.
func (t testCustomField) Bytes() []byte { return nil }

// Unmarshal skip, it's for test usage only.
func (t testCustomField) Unmarshal(data []byte) error { return nil }

// Empty skip, it's for test usage only.
func (t testCustomField) Empty() bool { return false }

// UnmarshalTo skip, it's for test usage only.
func (t testCustomField) MarshalTo(data []byte) (int, error) { return 0, nil }

// Marshal skip, it's for test usage only.
func (t testCustomField) Marshal() ([]byte, error) { return nil, nil }

// GetBearerToken wraps Bearer field and return BearerToken interface.
//
// If Bearer field value is nil, nil returns.
func (m RequestVerificationHeader) GetBearerToken() BearerToken {
	if t := m.GetBearer(); t != nil {
		return t
	}

	return nil
}

// SignKeyPairsSignedData wraps passed SignKeyPair slice and returns SignedDataSource interface.
func SignKeyPairsSignedData(v ...SignKeyPair) SignedDataSource {
	return &signKeyPairsWrapper{
		items: v,
	}
}

// SignedData returns signed SignKeyPair slice in a binary representation.
func (s signKeyPairsWrapper) SignedData() ([]byte, error) {
	return SignedDataFromReader(s)
}

// SignedDataSize returns the length of signed SignKeyPair slice.
func (s signKeyPairsWrapper) SignedDataSize() (sz int) {
	for i := range s.items {
		// add key length
		sz += len(
			crypto.MarshalPublicKey(s.items[i].GetPublicKey()),
		)
	}

	return
}

// ReadSignedData copies a binary representation of the signed SignKeyPair slice to passed buffer.
//
// If buffer length is less than required, io.ErrUnexpectedEOF returns.
func (s signKeyPairsWrapper) ReadSignedData(p []byte) (int, error) {
	sz := s.SignedDataSize()
	if len(p) < sz {
		return 0, io.ErrUnexpectedEOF
	}

	off := 0

	for i := range s.items {
		// copy public key bytes
		off += copy(p[off:], crypto.MarshalPublicKey(
			s.items[i].GetPublicKey(),
		))
	}

	return off, nil
}
