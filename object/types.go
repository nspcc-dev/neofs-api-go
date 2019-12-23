package object

import (
	"bytes"
	"context"

	"github.com/gogo/protobuf/proto"
	"github.com/nspcc-dev/neofs-proto/internal"
	"github.com/nspcc-dev/neofs-proto/refs"
	"github.com/nspcc-dev/neofs-proto/session"
)

type (
	// Pred defines a predicate function that can check if passed header
	// satisfies predicate condition. It is used to find headers of
	// specific type.
	Pred = func(*Header) bool

	// Address is a type alias of object Address.
	Address = refs.Address

	// VerificationHeader is a type alias of session's verification header.
	VerificationHeader = session.VerificationHeader

	// PositionReader defines object reader that returns slice of bytes
	// for specified object and data range.
	PositionReader interface {
		PRead(ctx context.Context, addr refs.Address, rng Range) ([]byte, error)
	}

	headerType int
)

const (
	// ErrVerifyPayload is raised when payload checksum cannot be verified.
	ErrVerifyPayload = internal.Error("can't verify payload")

	// ErrVerifyHeader is raised when object integrity cannot be verified.
	ErrVerifyHeader = internal.Error("can't verify header")

	// ErrHeaderNotFound is raised when requested header not found.
	ErrHeaderNotFound = internal.Error("header not found")

	// ErrVerifySignature is raised when signature cannot be verified.
	ErrVerifySignature = internal.Error("can't verify signature")

	// ErrRangeOutOfBounds is raised when range is out of object payload bounds.
	ErrRangeOutOfBounds = internal.Error("payload range is out of bounds")

	// ErrInvalidRangeOrder is raised when payload range order is out of order.
	ErrInvalidRangeOrder = internal.Error("invalid payload range list order")

	// ErrEmptyPayloadRange is raised when payload range size is non-positive.
	ErrEmptyPayloadRange = internal.Error("empty payload range")

	// ErrIncompleteRangeCoverage is raised when checksum header range list covers object payload incompletely.
	ErrIncompleteRangeCoverage = internal.Error("incomplete range coverage of payload")
)

const (
	_ headerType = iota
	// LinkHdr is a link header type.
	LinkHdr
	// RedirectHdr is a redirect header type.
	RedirectHdr
	// UserHdr is a user defined header type.
	UserHdr
	// TransformHdr is a transformation header type.
	TransformHdr
	// TombstoneHdr is a tombstone header type.
	TombstoneHdr
	// VerifyHdr is a verification header type.
	VerifyHdr
	// HomoHashHdr is a homomorphic hash header type.
	HomoHashHdr
	// PayloadChecksumHdr is a payload checksum header type.
	PayloadChecksumHdr
	// IntegrityHdr is a integrity header type.
	IntegrityHdr
	// StorageGroupHdr is a storage group header type.
	StorageGroupHdr
)

var (
	_ internal.Custom = (*Object)(nil)

	emptyObject = new(Object).Bytes()
)

// Bytes returns marshaled object in a binary format.
func (m Object) Bytes() []byte { data, _ := m.Marshal(); return data }

// Empty checks if object does not contain any information.
func (m Object) Empty() bool { return bytes.Equal(m.Bytes(), emptyObject) }

// LastHeader returns last header of the specified type. Type must be
// specified as a Pred function.
func (m Object) LastHeader(f Pred) (int, *Header) {
	for i := len(m.Headers) - 1; i >= 0; i-- {
		if f != nil && f(&m.Headers[i]) {
			return i, &m.Headers[i]
		}
	}
	return -1, nil
}

// AddHeader adds passed header to the end of extended header list.
func (m *Object) AddHeader(h *Header) {
	m.Headers = append(m.Headers, *h)
}

// SetPayload sets payload field and payload length in the system header.
func (m *Object) SetPayload(payload []byte) {
	m.Payload = payload
	m.SystemHeader.PayloadLength = uint64(len(payload))
}

// SetHeader replaces existing extended header or adds new one to the end of
// extended header list.
func (m *Object) SetHeader(h *Header) {
	// looking for the header of that type
	for i := range m.Headers {
		if m.Headers[i].typeOf(h.Value) {
			// if we found one - set it with new value and return
			m.Headers[i] = *h
			return
		}
	}
	// if we did not find one - add this header
	m.AddHeader(h)
}

func (m Header) typeOf(t isHeader_Value) (ok bool) {
	switch t.(type) {
	case *Header_Link:
		_, ok = m.Value.(*Header_Link)
	case *Header_Redirect:
		_, ok = m.Value.(*Header_Redirect)
	case *Header_UserHeader:
		_, ok = m.Value.(*Header_UserHeader)
	case *Header_Transform:
		_, ok = m.Value.(*Header_Transform)
	case *Header_Tombstone:
		_, ok = m.Value.(*Header_Tombstone)
	case *Header_Verify:
		_, ok = m.Value.(*Header_Verify)
	case *Header_HomoHash:
		_, ok = m.Value.(*Header_HomoHash)
	case *Header_PayloadChecksum:
		_, ok = m.Value.(*Header_PayloadChecksum)
	case *Header_Integrity:
		_, ok = m.Value.(*Header_Integrity)
	case *Header_StorageGroup:
		_, ok = m.Value.(*Header_StorageGroup)
	}
	return
}

// HeaderType returns predicate that check if extended header is a header
// of specified type.
func HeaderType(t headerType) Pred {
	switch t {
	case LinkHdr:
		return func(h *Header) bool { _, ok := h.Value.(*Header_Link); return ok }
	case RedirectHdr:
		return func(h *Header) bool { _, ok := h.Value.(*Header_Redirect); return ok }
	case UserHdr:
		return func(h *Header) bool { _, ok := h.Value.(*Header_UserHeader); return ok }
	case TransformHdr:
		return func(h *Header) bool { _, ok := h.Value.(*Header_Transform); return ok }
	case TombstoneHdr:
		return func(h *Header) bool { _, ok := h.Value.(*Header_Tombstone); return ok }
	case VerifyHdr:
		return func(h *Header) bool { _, ok := h.Value.(*Header_Verify); return ok }
	case HomoHashHdr:
		return func(h *Header) bool { _, ok := h.Value.(*Header_HomoHash); return ok }
	case PayloadChecksumHdr:
		return func(h *Header) bool { _, ok := h.Value.(*Header_PayloadChecksum); return ok }
	case IntegrityHdr:
		return func(h *Header) bool { _, ok := h.Value.(*Header_Integrity); return ok }
	case StorageGroupHdr:
		return func(h *Header) bool { _, ok := h.Value.(*Header_StorageGroup); return ok }
	default:
		return nil
	}
}

// Copy creates full copy of the object.
func (m *Object) Copy() (obj *Object) {
	obj = new(Object)
	m.CopyTo(obj)
	return
}

// CopyTo creates fills passed object with the data from the current object.
// This function creates copies on every available data slice.
func (m *Object) CopyTo(o *Object) {
	o.SystemHeader = m.SystemHeader
	o.Headers = make([]Header, len(m.Headers))
	o.Payload = make([]byte, len(m.Payload))

	for i := range m.Headers {
		switch v := m.Headers[i].Value.(type) {
		case *Header_Link:
			link := *v.Link
			o.Headers[i] = Header{
				Value: &Header_Link{
					Link: &link,
				},
			}
		case *Header_HomoHash:
			o.Headers[i] = Header{
				Value: &Header_HomoHash{
					HomoHash: v.HomoHash,
				},
			}
		default:
			o.Headers[i] = *proto.Clone(&m.Headers[i]).(*Header)
		}
	}

	copy(o.Payload, m.Payload)
}

// Address returns object's address.
func (m Object) Address() *refs.Address {
	return &refs.Address{
		ObjectID: m.SystemHeader.ID,
		CID:      m.SystemHeader.CID,
	}
}
