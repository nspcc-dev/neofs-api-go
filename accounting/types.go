package accounting

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/binary"
	"reflect"

	"github.com/mr-tron/base58"
	"github.com/nspcc-dev/neofs-api-go/chain"
	"github.com/nspcc-dev/neofs-api-go/decimal"
	"github.com/nspcc-dev/neofs-api-go/internal"
	"github.com/nspcc-dev/neofs-api-go/refs"
	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/pkg/errors"
)

type (
	// Cheque structure that describes a user request for withdrawal of funds.
	Cheque struct {
		ID         ChequeID
		Owner      refs.OwnerID
		Amount     *decimal.Decimal
		Height     uint64
		Signatures []ChequeSignature
	}

	// BalanceReceiver interface that is used to retrieve user balance by address.
	BalanceReceiver interface {
		Balance(accountAddress string) (*Account, error)
	}

	// ChequeID is identifier of user request for withdrawal of funds.
	ChequeID string

	// CID type alias.
	CID = refs.CID

	// SGID type alias.
	SGID = refs.SGID

	// ChequeSignature contains public key and hash, and is used to verify signatures.
	ChequeSignature struct {
		Key  *ecdsa.PublicKey
		Hash []byte
	}
)

const (
	// ErrWrongSignature is raised when wrong signature is passed.
	ErrWrongSignature = internal.Error("wrong signature")

	// ErrWrongPublicKey is raised when wrong public key is passed.
	ErrWrongPublicKey = internal.Error("wrong public key")

	// ErrWrongChequeData is raised when passed bytes cannot not be parsed as valid Cheque.
	ErrWrongChequeData = internal.Error("wrong cheque data")

	// ErrInvalidLength is raised when passed bytes cannot not be parsed as valid ChequeID.
	ErrInvalidLength = internal.Error("invalid length")

	u16size = 2
	u64size = 8

	signaturesOffset = chain.AddressLength + refs.OwnerIDSize + u64size + u64size
)

// NewChequeID generates valid random ChequeID using crypto/rand.Reader.
func NewChequeID() (ChequeID, error) {
	d := make([]byte, chain.AddressLength)
	if _, err := rand.Read(d); err != nil {
		return "", err
	}

	id := base58.Encode(d)

	return ChequeID(id), nil
}

// String returns string representation of ChequeID.
func (b ChequeID) String() string { return string(b) }

// Empty returns true, if ChequeID is empty.
func (b ChequeID) Empty() bool { return len(b) == 0 }

// Valid validates ChequeID.
func (b ChequeID) Valid() bool {
	d, err := base58.Decode(string(b))
	return err == nil && len(d) == chain.AddressLength
}

// Bytes returns bytes representation of ChequeID.
func (b ChequeID) Bytes() []byte {
	d, err := base58.Decode(string(b))
	if err != nil {
		return make([]byte, chain.AddressLength)
	}
	return d
}

// Equal checks that current ChequeID is equal to passed ChequeID.
func (b ChequeID) Equal(b2 ChequeID) bool {
	return b.Valid() && b2.Valid() && string(b) == string(b2)
}

// Unmarshal tries to parse []byte into valid ChequeID.
func (b *ChequeID) Unmarshal(data []byte) error {
	*b = ChequeID(base58.Encode(data))
	if !b.Valid() {
		return ErrInvalidLength
	}
	return nil
}

// Size returns size (chain.AddressLength).
func (b ChequeID) Size() int {
	return chain.AddressLength
}

// MarshalTo tries to marshal ChequeID into passed bytes and returns
// count of copied bytes or error, if bytes len is not enough to contain ChequeID.
func (b ChequeID) MarshalTo(data []byte) (int, error) {
	if len(data) < chain.AddressLength {
		return 0, ErrInvalidLength
	}
	return copy(data, b.Bytes()), nil
}

// Equals checks that m and tx are valid and equal Tx values.
func (m Tx) Equals(tx Tx) bool {
	return m.From == tx.From &&
		m.To == tx.To &&
		m.Type == tx.Type &&
		m.Amount == tx.Amount
}

// Verify validates current Cheque and Signatures that are generated for current Cheque.
func (b Cheque) Verify() error {
	data := b.marshalBody()
	for i, sign := range b.Signatures {
		if err := crypto.VerifyRFC6979(sign.Key, data, sign.Hash); err != nil {
			return errors.Wrapf(ErrWrongSignature, "item #%d: %s", i, err.Error())
		}
	}

	return nil
}

// Sign is used to sign current Cheque and stores result inside b.Signatures.
func (b *Cheque) Sign(key *ecdsa.PrivateKey) error {
	hash, err := crypto.SignRFC6979(key, b.marshalBody())
	if err != nil {
		return err
	}

	b.Signatures = append(b.Signatures, ChequeSignature{
		Key:  &key.PublicKey,
		Hash: hash,
	})

	return nil
}

func (b *Cheque) marshalBody() []byte {
	buf := make([]byte, signaturesOffset)

	var offset int

	offset += copy(buf, b.ID.Bytes())
	offset += copy(buf[offset:], b.Owner.Bytes())

	binary.LittleEndian.PutUint64(buf[offset:], uint64(b.Amount.Value))
	offset += u64size

	binary.LittleEndian.PutUint64(buf[offset:], b.Height)

	return buf
}

func (b *Cheque) unmarshalBody(buf []byte) error {
	var offset int

	if len(buf) < signaturesOffset {
		return ErrWrongChequeData
	}

	{ // unmarshal UUID
		if err := b.ID.Unmarshal(buf[offset : offset+chain.AddressLength]); err != nil {
			return err
		}
		offset += chain.AddressLength
	}

	{ // unmarshal OwnerID
		if err := b.Owner.Unmarshal(buf[offset : offset+refs.OwnerIDSize]); err != nil {
			return err
		}
		offset += refs.OwnerIDSize
	}

	{ // unmarshal amount
		amount := int64(binary.LittleEndian.Uint64(buf[offset:]))
		b.Amount = decimal.New(amount)
		offset += u64size
	}

	{ // unmarshal height
		b.Height = binary.LittleEndian.Uint64(buf[offset:])
		offset += u64size
	}

	return nil
}

// MarshalBinary is used to marshal Cheque into bytes.
func (b Cheque) MarshalBinary() ([]byte, error) {
	var (
		count  = len(b.Signatures)
		buf    = make([]byte, b.Size())
		offset = copy(buf, b.marshalBody())
	)

	binary.LittleEndian.PutUint16(buf[offset:], uint16(count))
	offset += u16size

	for _, sign := range b.Signatures {
		key := crypto.MarshalPublicKey(sign.Key)
		offset += copy(buf[offset:], key)
		offset += copy(buf[offset:], sign.Hash)
	}

	return buf, nil
}

// Size returns size of Cheque (count of bytes needs to store it).
func (b Cheque) Size() int {
	return signaturesOffset + u16size +
		len(b.Signatures)*(crypto.PublicKeyCompressedSize+crypto.RFC6979SignatureSize)
}

// UnmarshalBinary tries to parse []byte into valid Cheque.
func (b *Cheque) UnmarshalBinary(buf []byte) error {
	if err := b.unmarshalBody(buf); err != nil {
		return err
	}

	body := buf[:signaturesOffset]

	count := int64(binary.LittleEndian.Uint16(buf[signaturesOffset:]))
	offset := signaturesOffset + u16size

	if ln := count * int64(crypto.PublicKeyCompressedSize+crypto.RFC6979SignatureSize); ln > int64(len(buf[offset:])) {
		return ErrWrongChequeData
	}

	for i := int64(0); i < count; i++ {
		sign := ChequeSignature{
			Key:  crypto.UnmarshalPublicKey(buf[offset : offset+crypto.PublicKeyCompressedSize]),
			Hash: make([]byte, crypto.RFC6979SignatureSize),
		}

		offset += crypto.PublicKeyCompressedSize
		if sign.Key == nil {
			return errors.Wrapf(ErrWrongPublicKey, "item #%d", i)
		}

		offset += copy(sign.Hash, buf[offset:offset+crypto.RFC6979SignatureSize])
		if err := crypto.VerifyRFC6979(sign.Key, body, sign.Hash); err != nil {
			return errors.Wrapf(ErrWrongSignature, "item #%d: %s (offset=%d, len=%d)", i, err.Error(), offset, len(sign.Hash))
		}

		b.Signatures = append(b.Signatures, sign)
	}

	return nil
}

// ErrNotEnoughFunds generates error using address and amounts.
func ErrNotEnoughFunds(addr string, needed, residue *decimal.Decimal) error {
	return errors.Errorf("not enough funds (requested=%s, residue=%s, addr=%s", needed, residue, addr)
}

func (m *Account) hasLockAcc(addr string) bool {
	for i := range m.LockAccounts {
		if m.LockAccounts[i].Address == addr {
			return true
		}
	}
	return false
}

// ValidateLock checks that account can be locked.
func (m *Account) ValidateLock() error {
	switch {
	case m.Address == "":
		return ErrEmptyAddress
	case m.ParentAddress == "":
		return ErrEmptyParentAddress
	case m.LockTarget == nil:
		return ErrEmptyLockTarget
	}

	switch v := m.LockTarget.Target.(type) {
	case *LockTarget_WithdrawTarget:
		if v.WithdrawTarget.Cheque != m.Address {
			return errors.Errorf("wrong cheque ID: expected %s, has %s", m.Address, v.WithdrawTarget.Cheque)
		}
	case *LockTarget_ContainerCreateTarget:
		switch {
		case v.ContainerCreateTarget.CID.Empty():
			return ErrEmptyContainerID
		}
	}
	return nil
}

// CanLock checks possibility to lock funds.
func (m *Account) CanLock(lockAcc *Account) error {
	switch {
	case m.ActiveFunds.LT(lockAcc.ActiveFunds):
		return ErrNotEnoughFunds(lockAcc.ParentAddress, lockAcc.ActiveFunds, m.ActiveFunds)
	case m.hasLockAcc(lockAcc.Address):
		return errors.Errorf("could not lock account(%s) funds: duplicating lock(%s)", m.Address, lockAcc.Address)
	default:
		return nil
	}
}

// LockForWithdraw checks that account contains locked funds by passed ChequeID.
func (m *Account) LockForWithdraw(chequeID string) bool {
	switch v := m.LockTarget.Target.(type) {
	case *LockTarget_WithdrawTarget:
		return v.WithdrawTarget.Cheque == chequeID
	}
	return false
}

// LockForContainerCreate checks that account contains locked funds for container creation.
func (m *Account) LockForContainerCreate(cid refs.CID) bool {
	switch v := m.LockTarget.Target.(type) {
	case *LockTarget_ContainerCreateTarget:
		return v.ContainerCreateTarget.CID.Equal(cid)
	}
	return false
}

// Equal checks that current Settlement is equal to passed Settlement.
func (m *Settlement) Equal(s *Settlement) bool {
	if s == nil || m.Epoch != s.Epoch || len(m.Transactions) != len(s.Transactions) {
		return false
	}
	return len(m.Transactions) == 0 || reflect.DeepEqual(m.Transactions, s.Transactions)
}
