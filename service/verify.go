package service

import (
	"crypto/ecdsa"

	"github.com/gogo/protobuf/proto"
	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/nspcc-dev/neofs-proto/internal"
	"github.com/pkg/errors"
)

type (
	// VerifiableRequest adds possibility to sign and verify request header
	VerifiableRequest interface {
		proto.Message
		Marshal() ([]byte, error)
		AddSignature(*RequestVerificationHeader_Signature)
		GetSignatures() []*RequestVerificationHeader_Signature
		SetSignatures([]*RequestVerificationHeader_Signature)
	}

	// MaintainableRequest adds possibility to set and get (+validate)
	// owner (client) public key from RequestVerificationHeader.
	MaintainableRequest interface {
		proto.Message
		GetOwner() (*ecdsa.PublicKey, error)
		SetOwner(*ecdsa.PublicKey, []byte)
		GetLastPeer() (*ecdsa.PublicKey, error)
	}
)

const (
	// ErrCannotLoadPublicKey is raised when cannot unmarshal public key from RequestVerificationHeader_Sign
	ErrCannotLoadPublicKey = internal.Error("cannot load public key")

	// ErrCannotFindOwner is raised when signatures empty in GetOwner
	ErrCannotFindOwner = internal.Error("cannot find owner public key")
)

// SetSignatures replaces signatures stored in RequestVerificationHeader
func (m *RequestVerificationHeader) SetSignatures(signatures []*RequestVerificationHeader_Signature) {
	m.Signatures = signatures
}

// AddSignature adds new Signature into RequestVerificationHeader
func (m *RequestVerificationHeader) AddSignature(sig *RequestVerificationHeader_Signature) {
	if sig == nil {
		return
	}
	m.Signatures = append(m.Signatures, sig)
}

// SetOwner adds origin (sign and public key) of owner (client) into first signature.
func (m *RequestVerificationHeader) SetOwner(pub *ecdsa.PublicKey, sign []byte) {
	if len(m.Signatures) == 0 || pub == nil {
		return
	}

	m.Signatures[0].Origin = &RequestVerificationHeader_Sign{
		Sign: sign,
		Peer: crypto.MarshalPublicKey(pub),
	}
}

// GetOwner tries to get owner (client) public key from signatures.
// If signatures contains not empty Origin, we should try to validate,
// that session key was signed by owner (client), otherwise return error.
func (m *RequestVerificationHeader) GetOwner() (*ecdsa.PublicKey, error) {
	if len(m.Signatures) == 0 {
		return nil, ErrCannotFindOwner
	}

	// if first signature contains origin, we should try to validate session key
	if m.Signatures[0].Origin != nil {
		owner := crypto.UnmarshalPublicKey(m.Signatures[0].Origin.Peer)
		if owner == nil {
			return nil, ErrCannotLoadPublicKey
		} else if err := crypto.Verify(owner, m.Signatures[0].Peer, m.Signatures[0].Origin.Sign); err != nil {
			return nil, errors.Wrap(err, "could not verify session token")
		}

		return owner, nil
	} else if key := crypto.UnmarshalPublicKey(m.Signatures[0].Peer); key != nil {
		return key, nil
	}

	return nil, ErrCannotLoadPublicKey
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

		return nil, ErrCannotLoadPublicKey
	}
}

func newSignature(key *ecdsa.PrivateKey, data []byte) (*RequestVerificationHeader_Signature, error) {
	sign, err := crypto.Sign(key, data)
	if err != nil {
		return nil, err
	}

	return &RequestVerificationHeader_Signature{
		RequestVerificationHeader_Sign: RequestVerificationHeader_Sign{
			Sign: sign,
			Peer: crypto.MarshalPublicKey(&key.PublicKey),
		},
	}, nil
}

// SignRequestHeader receives private key and request with RequestVerificationHeader,
// tries to marshal and sign request with passed PrivateKey, after that adds
// new signature to headers. If something went wrong, returns error.
func SignRequestHeader(key *ecdsa.PrivateKey, req VerifiableRequest) error {
	msg := proto.Clone(req).(VerifiableRequest)

	// ignore meta header
	if meta, ok := msg.(MetaHeader); ok {
		meta.ResetMeta()
	}

	data, err := msg.Marshal()
	if err != nil {
		return err
	}

	signature, err := newSignature(key, data)
	if err != nil {
		return err
	}

	req.AddSignature(signature)

	return nil
}

// VerifyRequestHeader receives request with RequestVerificationHeader,
// tries to marshal and verify each signature from request
// If something went wrong, returns error.
func VerifyRequestHeader(req VerifiableRequest) error {
	msg := proto.Clone(req).(VerifiableRequest)
	// ignore meta header
	if meta, ok := msg.(MetaHeader); ok {
		meta.ResetMeta()
	}

	signatures := msg.GetSignatures()

	for i := range signatures {
		msg.SetSignatures(signatures[:i])
		peer := signatures[i].GetPeer()
		sign := signatures[i].GetSign()

		key := crypto.UnmarshalPublicKey(peer)
		if key == nil {
			return errors.Wrapf(ErrCannotLoadPublicKey, "%d: %02x", i, peer)
		}

		if data, err := msg.Marshal(); err != nil {
			return errors.Wrapf(err, "%d: %02x", i, peer)
		} else if err := crypto.Verify(key, data, sign); err != nil {
			return errors.Wrapf(err, "%d: %02x", i, peer)
		}
	}

	return nil
}
