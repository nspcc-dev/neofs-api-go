package refs

import (
	"crypto/sha256"
	"strings"

	"github.com/nspcc-dev/neofs-api-go/internal"
)

const (
	joinSeparator = "/"

	// ErrWrongAddress is raised when wrong address is passed to Address.Parse ParseAddress.
	ErrWrongAddress = internal.Error("wrong address")

	// ErrEmptyAddress is raised when empty address is passed to Address.Parse ParseAddress.
	ErrEmptyAddress = internal.Error("empty address")
)

// ParseAddress parses address from string representation into new Address.
func ParseAddress(str string) (*Address, error) {
	var addr Address
	return &addr, addr.Parse(str)
}

// Parse parses address from string representation into current Address.
func (m *Address) Parse(addr string) error {
	if m == nil {
		return ErrEmptyAddress
	}

	items := strings.Split(addr, joinSeparator)
	if len(items) != 2 {
		return ErrWrongAddress
	}

	if err := m.CID.Parse(items[0]); err != nil {
		return err
	} else if err := m.ObjectID.Parse(items[1]); err != nil {
		return err
	}

	return nil
}

// String returns string representation of Address.
func (m Address) String() string {
	return strings.Join([]string{m.CID.String(), m.ObjectID.String()}, joinSeparator)
}

// IsFull checks that ContainerID and ObjectID is not empty.
func (m Address) IsFull() bool {
	return !m.CID.Empty() && !m.ObjectID.Empty()
}

// Equal checks that current Address is equal to passed Address.
func (m Address) Equal(a2 *Address) bool {
	return m.CID.Equal(a2.CID) && m.ObjectID.Equal(a2.ObjectID)
}

// Hash returns []byte that used as a key for storage bucket.
func (m Address) Hash() ([]byte, error) {
	if !m.IsFull() {
		return nil, ErrEmptyAddress
	}
	h := sha256.Sum256(append(m.ObjectID.Bytes(), m.CID.Bytes()...))
	return h[:], nil
}
