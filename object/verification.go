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

// PayloadRangeChecksum calculates sha256 checksum of object payload range.
func (m Object) PayloadRangeChecksum(r Range) ([]byte, error) {
	if r.Offset+r.Length > uint64(len(m.Payload)) {
		return nil, ErrRangeOutOfBounds
	}
	checksum := sha256.Sum256(m.Payload[r.Offset : r.Offset+r.Length])
	return checksum[:], nil
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
func (m Object) Verify() error {
	var (
		err      error
		checksum []byte
	)
	// Prepare structures
	_, vh := m.LastHeader(HeaderType(VerifyHdr))
	if vh == nil {
		return ErrHeaderNotFound
	}
	verify := vh.Value.(*Header_Verify).Verify

	_, ih := m.LastHeader(HeaderType(IntegrityHdr))
	if ih == nil {
		return ErrHeaderNotFound
	}
	integrity := ih.Value.(*Header_Integrity).Integrity

	// Verify signature
	err = m.verifySignature(verify.PublicKey, integrity)
	if err != nil {
		return errors.Wrapf(err, "public key: %x", verify.PublicKey)
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
		return m.VerifyPayload()
	}
	return nil
}

// Sign creates new integrity header and adds it to the end of the list of
// extended headers.
func (m *Object) Sign(key *ecdsa.PrivateKey) error {
	headerChecksum, err := m.headersChecksum(false)
	if err != nil {
		return err
	}
	headerChecksumSignature, err := crypto.Sign(key, headerChecksum)
	if err != nil {
		return err
	}
	m.AddHeader(&Header{Value: &Header_Integrity{
		Integrity: &IntegrityHeader{
			HeadersChecksum:   headerChecksum,
			ChecksumSignature: headerChecksumSignature,
		},
	}})
	return nil
}

// VerifyPayload checks the correctness of the payload checksum header structurally and by value.
func (m *Object) VerifyPayload() error {
	_, ph := m.LastHeader(HeaderType(PayloadChecksumHdr))
	if ph == nil {
		return ErrHeaderNotFound
	}

	p := uint64(0)

	for _, rangeChecksum := range ph.Value.(*Header_PayloadChecksum).PayloadChecksum.ChecksumList {
		if rangeChecksum.Range.Offset != p {
			return ErrInvalidRangeOrder
		} else if rangeChecksum.Range.Length <= 0 {
			return ErrEmptyPayloadRange
		} else if calculatedChecksum, err := m.PayloadRangeChecksum(rangeChecksum.Range); err != nil {
			return err
		} else if !bytes.Equal(calculatedChecksum, rangeChecksum.Value) {
			return ErrVerifyPayload
		}
		p += rangeChecksum.Range.Length
	}

	if p != uint64(len(m.Payload)) {
		return ErrIncompleteRangeCoverage
	}

	return nil
}
