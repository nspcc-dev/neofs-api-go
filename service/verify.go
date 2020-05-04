package service

import (
	"crypto/ecdsa"
	"sync"

	"github.com/nspcc-dev/neofs-api-go/internal"
	"github.com/nspcc-dev/neofs-api-go/refs"
	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/pkg/errors"
)

type (
	// VerifiableRequest adds possibility to sign and verify request header.
	VerifiableRequest interface {
		Size() int
		MarshalTo([]byte) (int, error)
		AddSignature(*RequestVerificationHeader_Signature)
		GetSignatures() []*RequestVerificationHeader_Signature
		SetSignatures([]*RequestVerificationHeader_Signature)
	}

	// MaintainableRequest adds possibility to set and get (+validate)
	// owner (client) public key from RequestVerificationHeader.
	MaintainableRequest interface {
		GetOwner() (*ecdsa.PublicKey, error)
		SetOwner(*ecdsa.PublicKey, []byte)
		GetLastPeer() (*ecdsa.PublicKey, error)
	}

	// TokenHeader is an interface of the container of a Token pointer value
	TokenHeader interface {
		GetToken() *Token
		SetToken(*Token)
	}
)

// SetSignatures replaces signatures stored in RequestVerificationHeader.
func (m *RequestVerificationHeader) SetSignatures(signatures []*RequestVerificationHeader_Signature) {
	m.Signatures = signatures
}

// AddSignature adds new Signature into RequestVerificationHeader.
func (m *RequestVerificationHeader) AddSignature(sig *RequestVerificationHeader_Signature) {
	if sig == nil {
		return
	}
	m.Signatures = append(m.Signatures, sig)
}

// CheckOwner validates, that passed OwnerID is equal to present PublicKey of owner.
func (m *RequestVerificationHeader) CheckOwner(owner refs.OwnerID) error {
	if key, err := m.GetOwner(); err != nil {
		return err
	} else if user, err := refs.NewOwnerID(key); err != nil {
		return err
	} else if !user.Equal(owner) {
		return ErrWrongOwner
	}
	return nil
}

// GetOwner tries to get owner (client) public key from signatures.
// If signatures contains not empty Origin, we should try to validate,
// that session key was signed by owner (client), otherwise return error.
func (m *RequestVerificationHeader) GetOwner() (*ecdsa.PublicKey, error) {
	if len(m.Signatures) == 0 {
		return nil, ErrCannotFindOwner
	} else if key := crypto.UnmarshalPublicKey(m.Signatures[0].Peer); key != nil {
		return key, nil
	}

	return nil, ErrInvalidPublicKeyBytes
}

// GetLastPeer tries to get last peer public key from signatures.
// If signatures has zero length, returns ErrCannotFindOwner.
// If signatures has length equal to one, uses GetOwner.
// Otherwise tries to unmarshal last peer public key.
func (m *RequestVerificationHeader) GetLastPeer() (*ecdsa.PublicKey, error) {
	switch ln := len(m.Signatures); ln {
	case 0:
		return nil, ErrCannotFindOwner
	case 1:
		return m.GetOwner()
	default:
		if key := crypto.UnmarshalPublicKey(m.Signatures[ln-1].Peer); key != nil {
			return key, nil
		}

		return nil, ErrInvalidPublicKeyBytes
	}
}

// SetToken is a Token field setter.
func (m *RequestVerificationHeader) SetToken(token *Token) {
	m.Token = token
}

func newSignature(key *ecdsa.PrivateKey, data []byte) (*RequestVerificationHeader_Signature, error) {
	sign, err := crypto.Sign(key, data)
	if err != nil {
		return nil, err
	}

	return &RequestVerificationHeader_Signature{
		Sign: sign,
		Peer: crypto.MarshalPublicKey(&key.PublicKey),
	}, nil
}

var bytesPool = sync.Pool{New: func() interface{} {
	return make([]byte, 4.5*1024*1024) // 4.5MB
}}

// SignRequestHeader receives private key and request with RequestVerificationHeader,
// tries to marshal and sign request with passed PrivateKey, after that adds
// new signature to headers. If something went wrong, returns error.
func SignRequestHeader(key *ecdsa.PrivateKey, msg VerifiableRequest) error {
	// ignore meta header
	if meta, ok := msg.(MetaHeader); ok {
		h := meta.ResetMeta()

		defer func() {
			meta.RestoreMeta(h)
		}()
	}

	data := bytesPool.Get().([]byte)
	defer func() {
		bytesPool.Put(data)
	}()

	if size := msg.Size(); size <= cap(data) {
		data = data[:size]
	} else {
		data = make([]byte, size)
	}

	size, err := msg.MarshalTo(data)
	if err != nil {
		return err
	}

	signature, err := newSignature(key, data[:size])
	if err != nil {
		return err
	}

	msg.AddSignature(signature)

	return nil
}

// VerifyRequestHeader receives request with RequestVerificationHeader,
// tries to marshal and verify each signature from request.
// If something went wrong, returns error.
func VerifyRequestHeader(msg VerifiableRequest) error {
	// ignore meta header
	if meta, ok := msg.(MetaHeader); ok {
		h := meta.ResetMeta()

		defer func() {
			meta.RestoreMeta(h)
		}()
	}

	data := bytesPool.Get().([]byte)
	signatures := msg.GetSignatures()
	defer func() {
		bytesPool.Put(data)
		msg.SetSignatures(signatures)
	}()

	for i := range signatures {
		msg.SetSignatures(signatures[:i])
		peer := signatures[i].GetPeer()
		sign := signatures[i].GetSign()

		key := crypto.UnmarshalPublicKey(peer)
		if key == nil {
			return errors.Wrapf(ErrInvalidPublicKeyBytes, "%d: %02x", i, peer)
		}

		if size := msg.Size(); size <= cap(data) {
			data = data[:size]
		} else {
			data = make([]byte, size)
		}

		if size, err := msg.MarshalTo(data); err != nil {
			return errors.Wrapf(err, "%d: %02x", i, peer)
		} else if err := crypto.Verify(key, data[:size], sign); err != nil {
			return errors.Wrapf(err, "%d: %02x", i, peer)
		}
	}

	return nil
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
