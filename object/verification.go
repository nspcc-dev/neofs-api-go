package object

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"

	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/pkg/errors"
)

func (m Object) headersData(check bool) ([]byte, error) {
	var bytebuf = new(bytes.Buffer)

	// fixme: we must marshal fields one by one without protobuf marshaling
	//        protobuf marshaling does not guarantee the same result

	if sysheader, err := m.SystemHeader.Marshal(); err != nil {
		return nil, err
	} else if _, err := bytebuf.Write(sysheader); err != nil {
		return nil, err
	}

	n, _ := m.LastHeader(HeaderType(IntegrityHdr))
	for i := range m.Headers {
		if check && i == n {
			// ignore last integrity header in order to check headers data
			continue
		}

		if header, err := m.Headers[i].Marshal(); err != nil {
			return nil, err
		} else if _, err := bytebuf.Write(header); err != nil {
			return nil, err
		}
	}
	return bytebuf.Bytes(), nil
}

func (m Object) headersChecksum(check bool) ([]byte, error) {
	data, err := m.headersData(check)
	if err != nil {
		return nil, err
	}
	checksum := sha256.Sum256(data)
	return checksum[:], nil
}

// PayloadChecksum calculates sha256 checksum of object payload.
func (m Object) PayloadChecksum() []byte {
	checksum := sha256.Sum256(m.Payload)
	return checksum[:]
}

func (m Object) verifySignature(key []byte, ih *IntegrityHeader) error {
	pk := crypto.UnmarshalPublicKey(key)
	if crypto.Verify(pk, ih.HeadersChecksum, ih.ChecksumSignature) == nil {
		return nil
	}
	return ErrVerifySignature
}

// Verify performs local integrity check by finding verification header and
// integrity header. If header integrity is passed, function verifies
// checksum of the object payload.
// todo: move this verification logic into separate library
func (m Object) Verify() error {
	var (
		err      error
		checksum []byte
		pubkey   []byte
	)
	ind, ih := m.LastHeader(HeaderType(IntegrityHdr))
	if ih == nil || ind != len(m.Headers)-1 {
		return ErrHeaderNotFound
	}
	integrity := ih.Value.(*Header_Integrity).Integrity

	// Prepare structures
	_, vh := m.LastHeader(HeaderType(TokenHdr))
	if vh == nil {
		_, pkh := m.LastHeader(HeaderType(PublicKeyHdr))
		if pkh == nil {
			return ErrHeaderNotFound
		}
		pubkey = pkh.Value.(*Header_PublicKey).PublicKey.Value
	} else {
		pubkey = vh.Value.(*Header_Token).Token.SessionKey
	}

	// Verify signature
	err = m.verifySignature(pubkey, integrity)
	if err != nil {
		return errors.Wrapf(err, "public key: %x", pubkey)
	}

	// Verify checksum of header
	checksum, err = m.headersChecksum(true)
	if err != nil {
		return err
	}
	if !bytes.Equal(integrity.HeadersChecksum, checksum) {
		return ErrVerifyHeader
	}

	// Verify checksum of payload
	if m.SystemHeader.PayloadLength > 0 && !m.IsLinking() {
		checksum = m.PayloadChecksum()

		_, ph := m.LastHeader(HeaderType(PayloadChecksumHdr))
		if ph == nil {
			return ErrHeaderNotFound
		}
		if !bytes.Equal(ph.Value.(*Header_PayloadChecksum).PayloadChecksum, checksum) {
			return ErrVerifyPayload
		}
	}
	return nil
}

// CreateIntegrityHeader returns signed integrity header for the object
func CreateIntegrityHeader(obj *Object, key *ecdsa.PrivateKey) (*Header, error) {
	headerChecksum, err := obj.headersChecksum(false)
	if err != nil {
		return nil, err
	}
	headerChecksumSignature, err := crypto.Sign(key, headerChecksum)
	if err != nil {
		return nil, err
	}

	return &Header{Value: &Header_Integrity{
		Integrity: &IntegrityHeader{
			HeadersChecksum:   headerChecksum,
			ChecksumSignature: headerChecksumSignature,
		},
	}}, nil
}

// Sign creates new integrity header and adds it to the end of the list of
// extended headers.
func (m *Object) Sign(key *ecdsa.PrivateKey) error {
	ih, err := CreateIntegrityHeader(m, key)
	if err != nil {
		return err
	}
	m.AddHeader(ih)
	return nil
}
