package chain

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"

	"github.com/mr-tron/base58"
	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/nspcc-dev/neofs-proto/internal"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ripemd160"
)

// WalletAddress implements NEO address.
type WalletAddress [AddressLength]byte

const (
	// AddressLength contains size of address,
	// 0x17 byte (address version) + 20 bytes of ScriptHash + 4 bytes of checksum.
	AddressLength = 25

	// ScriptHashLength contains size of ScriptHash.
	ScriptHashLength = 20

	// ErrEmptyAddress is raised when empty Address is passed.
	ErrEmptyAddress = internal.Error("empty address")

	// ErrAddressLength is raised when passed address has wrong size.
	ErrAddressLength = internal.Error("wrong address length")
)

func checksum(sign []byte) []byte {
	hash := sha256.Sum256(sign)
	hash = sha256.Sum256(hash[:])
	return hash[:4]
}

// FetchPublicKeys tries to parse public keys from verification script.
func FetchPublicKeys(vs []byte) []*ecdsa.PublicKey {
	var (
		count  int
		offset int
		ln     = len(vs)
		result []*ecdsa.PublicKey
	)

	switch {
	case ln < 1: // wrong data size
		return nil
	case vs[ln-1] == 0xac: // last byte is CHECKSIG
		count = 1
	case vs[ln-1] == 0xae: // last byte is CHECKMULTISIG
		// 2nd byte from the end indicates about PK's count
		count = int(vs[ln-2] - 0x50)
		// ignores CHECKMULTISIG
		offset = 1
	default: // unknown type
		return nil
	}

	result = make([]*ecdsa.PublicKey, 0, count)
	for i := 0; i < count; i++ {
		// ignores PUSHBYTE33 and tries to parse
		from, to := offset+1, offset+1+crypto.PublicKeyCompressedSize

		// when passed VerificationScript has wrong size
		if len(vs) < to {
			return nil
		}

		key := crypto.UnmarshalPublicKey(vs[from:to])
		// when wrong public key is passed
		if key == nil {
			return nil
		}
		result = append(result, key)

		offset += 1 + crypto.PublicKeyCompressedSize
	}
	return result
}

// VerificationScript returns VerificationScript composed from public keys.
func VerificationScript(pubs ...*ecdsa.PublicKey) []byte {
	var (
		pre    []byte
		suf    []byte
		body   []byte
		offset int
		lnPK   = len(pubs)
		ln     = crypto.PublicKeyCompressedSize*lnPK + lnPK // 33 * count + count * 1 (PUSHBYTES33)
	)

	if len(pubs) > 1 {
		pre = []byte{0x51}                    // one address
		suf = []byte{byte(0x50 + lnPK), 0xae} // count of PK's + CHECKMULTISIG
	} else {
		suf = []byte{0xac} // CHECKSIG
	}

	ln += len(pre) + len(suf)

	body = make([]byte, ln)
	offset += copy(body, pre)

	for i := range pubs {
		body[offset] = 0x21
		offset++
		offset += copy(body[offset:], crypto.MarshalPublicKey(pubs[i]))
	}

	copy(body[offset:], suf)

	return body
}

// KeysToAddress return NEO address composed from public keys.
func KeysToAddress(pubs ...*ecdsa.PublicKey) string {
	if len(pubs) == 0 {
		return ""
	}
	return Address(VerificationScript(pubs...))
}

// Address returns NEO address based on passed VerificationScript.
func Address(verificationScript []byte) string {
	sign := [AddressLength]byte{0x17}
	hash := sha256.Sum256(verificationScript)
	ripe := ripemd160.New()
	ripe.Write(hash[:])
	copy(sign[1:], ripe.Sum(nil))
	copy(sign[21:], checksum(sign[:21]))
	return base58.Encode(sign[:])
}

// ReversedScriptHashToAddress parses script hash and returns valid NEO address.
func ReversedScriptHashToAddress(sc string) (addr string, err error) {
	var data []byte
	if data, err = DecodeScriptHash(sc); err != nil {
		return
	}
	sign := [AddressLength]byte{0x17}
	copy(sign[1:], data)
	copy(sign[1+ScriptHashLength:], checksum(sign[:1+ScriptHashLength]))
	return base58.Encode(sign[:]), nil
}

// IsAddress checks that passed NEO Address is valid.
func IsAddress(s string) error {
	if s == "" {
		return ErrEmptyAddress
	} else if addr, err := base58.Decode(s); err != nil {
		return errors.Wrap(err, "base58 decode")
	} else if ln := len(addr); ln != AddressLength {
		return errors.Wrapf(ErrAddressLength, "length %d != %d", AddressLength, ln)
	} else if sum := checksum(addr[:21]); !bytes.Equal(addr[21:], sum) {
		return errors.Errorf("wrong checksum %0x != %0x",
			addr[21:], sum)
	}

	return nil
}

// ReverseBytes returns reversed []byte of given.
func ReverseBytes(data []byte) []byte {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return data
}

// DecodeScriptHash parses script hash into slice of bytes.
func DecodeScriptHash(s string) ([]byte, error) {
	if s == "" {
		return nil, ErrEmptyAddress
	} else if addr, err := hex.DecodeString(s); err != nil {
		return nil, errors.Wrap(err, "hex decode")
	} else if ln := len(addr); ln != ScriptHashLength {
		return nil, errors.Wrapf(ErrAddressLength, "length %d != %d", ScriptHashLength, ln)
	} else {
		return addr, nil
	}
}
