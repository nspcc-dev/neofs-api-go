package object

import (
	"github.com/nspcc-dev/neofs-api-go/hash"
	"github.com/nspcc-dev/neofs-api-go/internal"
	"github.com/nspcc-dev/neofs-api-go/refs"
	"github.com/nspcc-dev/neofs-api-go/service"
	"github.com/nspcc-dev/neofs-api-go/session"
)

type (
	// ID is a type alias of object id.
	ID = refs.ObjectID

	// CID is a type alias of container id.
	CID = refs.CID

	// SGID is a type alias of storage group id.
	SGID = refs.SGID

	// OwnerID is a type alias of owner id.
	OwnerID = refs.OwnerID

	// Hash is a type alias of Homomorphic hash.
	Hash = hash.Hash

	// Token is a type alias of session token.
	Token = session.Token

	// Request defines object rpc requests.
	// All object operations must have TTL, Epoch, Type, Container ID and
	// permission of usage previous network map.
	Request interface {
		service.MetaHeader

		CID() CID
		Type() RequestType
		AllowPreviousNetMap() bool
	}
)

const (
	// starts enum for amount of bytes.
	_ int64 = 1 << (10 * iota)

	// UnitsKB defines amount of bytes in one kilobyte.
	UnitsKB

	// UnitsMB defines amount of bytes in one megabyte.
	UnitsMB

	// UnitsGB defines amount of bytes in one gigabyte.
	UnitsGB

	// UnitsTB defines amount of bytes in one terabyte.
	UnitsTB
)

const (
	// ErrNotFound is raised when object is not found in the system.
	ErrNotFound = internal.Error("could not find object")

	// ErrHeaderExpected is raised when first message in protobuf stream does not contain user header.
	ErrHeaderExpected = internal.Error("expected header as a first message in stream")

	// KeyStorageGroup is a key for a search object by storage group id.
	KeyStorageGroup = "STORAGE_GROUP"

	// KeyNoChildren is a key for searching object that have no children links.
	KeyNoChildren = "LEAF"

	// KeyParent is a key for searching object by id of parent object.
	KeyParent = "PARENT"

	// KeyHasParent is a key for searching object that have parent link.
	KeyHasParent = "HAS_PAR"

	// KeyTombstone is a key for searching object that have tombstone header.
	KeyTombstone = "TOMBSTONE"

	// KeyChild is a key for searching object by id of child link.
	KeyChild = "CHILD"

	// KeyPrev is a key for searching object by id of previous link.
	KeyPrev = "PREV"

	// KeyNext is a key for searching object by id of next link.
	KeyNext = "NEXT"

	// KeyID is a key for searching object by object id.
	KeyID = "ID"

	// KeyCID is a key for searching object by container id.
	KeyCID = "CID"

	// KeyOwnerID is a key for searching object by owner id.
	KeyOwnerID = "OWNERID"

	// KeyRootObject is a key for searching object that are zero-object or do
	// not have any children.
	KeyRootObject = "ROOT_OBJECT"
)

func checkIsNotFull(v interface{}) bool {
	var obj *Object

	switch t := v.(type) {
	case *GetResponse:
		obj = t.GetObject()
	case *PutRequest:
		if h := t.GetHeader(); h != nil {
			obj = h.Object
		}
	default:
		panic("unknown type")
	}

	return obj == nil || obj.SystemHeader.PayloadLength != uint64(len(obj.Payload)) && !obj.IsLinking()
}

// NotFull checks if protobuf stream provided whole object for get operation.
func (m *GetResponse) NotFull() bool { return checkIsNotFull(m) }

// NotFull checks if protobuf stream provided whole object for put operation.
func (m *PutRequest) NotFull() bool { return checkIsNotFull(m) }

// CID returns container id value from object put request.
func (m *PutRequest) CID() CID {
	if header := m.GetHeader(); header != nil {
		return header.Object.SystemHeader.CID
	}
	return refs.CID{}
}

// CID returns container id value from object get request.
func (m *GetRequest) CID() CID { return m.Address.CID }

// CID returns container id value from object head request.
func (m *HeadRequest) CID() CID { return m.Address.CID }

// CID returns container id value from object search request.
func (m *SearchRequest) CID() CID { return m.ContainerID }

// CID returns container id value from object delete request.
func (m *DeleteRequest) CID() CID { return m.Address.CID }

// CID returns container id value from object get range request.
func (m *GetRangeRequest) CID() CID { return m.Address.CID }

// CID returns container id value from object get range hash request.
func (m *GetRangeHashRequest) CID() CID { return m.Address.CID }

// AllowPreviousNetMap returns permission to use previous network map in object put request.
func (m *PutRequest) AllowPreviousNetMap() bool { return false }

// AllowPreviousNetMap returns permission to use previous network map in object get request.
func (m *GetRequest) AllowPreviousNetMap() bool { return true }

// AllowPreviousNetMap returns permission to use previous network map in object head request.
func (m *HeadRequest) AllowPreviousNetMap() bool { return true }

// AllowPreviousNetMap returns permission to use previous network map in object search request.
func (m *SearchRequest) AllowPreviousNetMap() bool { return true }

// AllowPreviousNetMap returns permission to use previous network map in object delete request.
func (m *DeleteRequest) AllowPreviousNetMap() bool { return false }

// AllowPreviousNetMap returns permission to use previous network map in object get range request.
func (m *GetRangeRequest) AllowPreviousNetMap() bool { return false }

// AllowPreviousNetMap returns permission to use previous network map in object get range hash request.
func (m *GetRangeHashRequest) AllowPreviousNetMap() bool { return false }

// Type returns type of the object put request.
func (m *PutRequest) Type() RequestType { return RequestPut }

// Type returns type of the object get request.
func (m *GetRequest) Type() RequestType { return RequestGet }

// Type returns type of the object head request.
func (m *HeadRequest) Type() RequestType { return RequestHead }

// Type returns type of the object search request.
func (m *SearchRequest) Type() RequestType { return RequestSearch }

// Type returns type of the object delete request.
func (m *DeleteRequest) Type() RequestType { return RequestDelete }

// Type returns type of the object get range request.
func (m *GetRangeRequest) Type() RequestType { return RequestRange }

// Type returns type of the object get range hash request.
func (m *GetRangeHashRequest) Type() RequestType { return RequestRangeHash }
