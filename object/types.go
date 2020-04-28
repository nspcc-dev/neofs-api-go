package object

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"reflect"

	"github.com/gogo/protobuf/proto"
	"github.com/nspcc-dev/neofs-api-go/internal"
	"github.com/nspcc-dev/neofs-api-go/refs"
	"github.com/pkg/errors"
)

type (
	// Pred defines a predicate function that can check if passed header
	// satisfies predicate condition. It is used to find headers of
	// specific type.
	Pred = func(*Header) bool

	// Address is a type alias of object Address.
	Address = refs.Address

	// PositionReader defines object reader that returns slice of bytes
	// for specified object and data range.
	PositionReader interface {
		PRead(ctx context.Context, addr refs.Address, rng Range) ([]byte, error)
	}

	// RequestType of the object service requests.
	RequestType int

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
	// TokenHdr is a token header type.
	TokenHdr
	// HomoHashHdr is a homomorphic hash header type.
	HomoHashHdr
	// PayloadChecksumHdr is a payload checksum header type.
	PayloadChecksumHdr
	// IntegrityHdr is a integrity header type.
	IntegrityHdr
	// StorageGroupHdr is a storage group header type.
	StorageGroupHdr
	// PublicKeyHdr is a public key header type.
	PublicKeyHdr
)

const (
	_ RequestType = iota
	// RequestPut is a type for object put request.
	RequestPut
	// RequestGet is a type for object get request.
	RequestGet
	// RequestHead is a type for object head request.
	RequestHead
	// RequestSearch is a type for object search request.
	RequestSearch
	// RequestRange is a type for object range request.
	RequestRange
	// RequestRangeHash is a type for object hash range request.
	RequestRangeHash
	// RequestDelete is a type for object delete request.
	RequestDelete
)

var (
	_ internal.Custom = (*Object)(nil)

	emptyObject = new(Object).Bytes()
)

// String returns printable name of the request type.
func (s RequestType) String() string {
	switch s {
	case RequestPut:
		return "PUT"
	case RequestGet:
		return "GET"
	case RequestHead:
		return "HEAD"
	case RequestSearch:
		return "SEARCH"
	case RequestRange:
		return "RANGE"
	case RequestRangeHash:
		return "RANGE_HASH"
	case RequestDelete:
		return "DELETE"
	default:
		return "UNKNOWN"
	}
}

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
	case *Header_Token:
		_, ok = m.Value.(*Header_Token)
	case *Header_HomoHash:
		_, ok = m.Value.(*Header_HomoHash)
	case *Header_PayloadChecksum:
		_, ok = m.Value.(*Header_PayloadChecksum)
	case *Header_Integrity:
		_, ok = m.Value.(*Header_Integrity)
	case *Header_StorageGroup:
		_, ok = m.Value.(*Header_StorageGroup)
	case *Header_PublicKey:
		_, ok = m.Value.(*Header_PublicKey)
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
	case TokenHdr:
		return func(h *Header) bool { _, ok := h.Value.(*Header_Token); return ok }
	case HomoHashHdr:
		return func(h *Header) bool { _, ok := h.Value.(*Header_HomoHash); return ok }
	case PayloadChecksumHdr:
		return func(h *Header) bool { _, ok := h.Value.(*Header_PayloadChecksum); return ok }
	case IntegrityHdr:
		return func(h *Header) bool { _, ok := h.Value.(*Header_Integrity); return ok }
	case StorageGroupHdr:
		return func(h *Header) bool { _, ok := h.Value.(*Header_StorageGroup); return ok }
	case PublicKeyHdr:
		return func(h *Header) bool { _, ok := h.Value.(*Header_PublicKey); return ok }
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

func (m CreationPoint) String() string {
	return fmt.Sprintf(`{UnixTime=%d Epoch=%d}`, m.UnixTime, m.Epoch)
}

// Stringify converts object into string format.
func Stringify(dst io.Writer, obj *Object) error {
	// put empty line
	if _, err := fmt.Fprintln(dst); err != nil {
		return err
	}

	// put object line
	if _, err := fmt.Fprintln(dst, "Object:"); err != nil {
		return err
	}

	// put system headers
	if _, err := fmt.Fprintln(dst, "\tSystemHeader:"); err != nil {
		return err
	}

	sysHeaders := []string{"ID", "CID", "OwnerID", "Version", "PayloadLength", "CreatedAt"}
	v := reflect.ValueOf(obj.SystemHeader)
	for _, key := range sysHeaders {
		if !v.FieldByName(key).IsValid() {
			return errors.Errorf("invalid system header key: %q", key)
		}

		val := v.FieldByName(key).Interface()
		if _, err := fmt.Fprintf(dst, "\t\t- %s=%v\n", key, val); err != nil {
			return err
		}
	}

	// put user headers
	if _, err := fmt.Fprintln(dst, "\tUserHeaders:"); err != nil {
		return err
	}

	for _, header := range obj.Headers {
		var (
			typ = reflect.ValueOf(header.Value)
			key string
			val interface{}
		)

		switch t := typ.Interface().(type) {
		case *Header_Link:
			key = "Link"
			val = fmt.Sprintf(`{Type=%s ID=%s}`, t.Link.Type, t.Link.ID)
		case *Header_Redirect:
			key = "Redirect"
			val = fmt.Sprintf(`{CID=%s OID=%s}`, t.Redirect.CID, t.Redirect.ObjectID)
		case *Header_UserHeader:
			key = "UserHeader"
			val = fmt.Sprintf(`{Key=%s Val=%s}`, t.UserHeader.Key, t.UserHeader.Value)
		case *Header_Transform:
			key = "Transform"
			val = t.Transform.Type.String()
		case *Header_Tombstone:
			key = "Tombstone"
			val = "MARKED"
		case *Header_Token:
			key = "Token"
			val = fmt.Sprintf("{"+
				"ID=%s OwnerID=%s Verb=%s Address=%s Created=%d ValidUntil=%d SessionKey=%02x Signature=%02x"+
				"}",
				t.Token.Token_Info.ID,
				t.Token.Token_Info.OwnerID,
				t.Token.Token_Info.Verb,
				t.Token.Token_Info.Address,
				t.Token.Token_Info.Created,
				t.Token.Token_Info.ValidUntil,
				t.Token.Token_Info.SessionKey,
				t.Token.Signature)
		case *Header_HomoHash:
			key = "HomoHash"
			val = t.HomoHash
		case *Header_PayloadChecksum:
			key = "PayloadChecksum"
			val = t.PayloadChecksum
		case *Header_Integrity:
			key = "Integrity"
			val = fmt.Sprintf(`{Checksum=%02x Signature=%02x}`,
				t.Integrity.HeadersChecksum,
				t.Integrity.ChecksumSignature)
		case *Header_StorageGroup:
			key = "StorageGroup"
			val = fmt.Sprintf(`{DataSize=%d Hash=%02x Lifetime={Unit=%s Value=%d}}`,
				t.StorageGroup.ValidationDataSize,
				t.StorageGroup.ValidationHash,
				t.StorageGroup.Lifetime.Unit,
				t.StorageGroup.Lifetime.Value)
		case *Header_PublicKey:
			key = "PublicKey"
			val = t.PublicKey.Value
		default:
			key = "Unknown"
			val = t
		}

		if _, err := fmt.Fprintf(dst, "\t\t- Type=%s\n\t\t  Value=%v\n", key, val); err != nil {
			return err
		}
	}

	// put payload
	if _, err := fmt.Fprintf(dst, "\tPayload: %#v\n", obj.Payload); err != nil {
		return err
	}

	return nil
}
