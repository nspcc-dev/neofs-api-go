package acl

import (
	"strconv"

	"github.com/nspcc-dev/neofs-api-go/object"
	"github.com/nspcc-dev/neofs-api-go/service"
)

type objectHeaderSource struct {
	obj *object.Object
}

type typedHeader struct {
	n string
	v string
	t HeaderType
}

type extendedHeadersWrapper struct {
	hdrSrc service.ExtendedHeadersSource
}

type typedExtendedHeader struct {
	hdr service.ExtendedHeader
}

const (
	_ HeaderType = iota

	// HdrTypeRequest is a HeaderType for request header.
	HdrTypeRequest

	// HdrTypeObjSys is a HeaderType for system headers of object.
	HdrTypeObjSys

	// HdrTypeObjUsr is a HeaderType for user headers of object.
	HdrTypeObjUsr
)

const (
	// HdrObjSysNameID is a name of ID field in system header of object.
	HdrObjSysNameID = "ID"

	// HdrObjSysNameCID is a name of CID field in system header of object.
	HdrObjSysNameCID = "CID"

	// HdrObjSysNameOwnerID is a name of OwnerID field in system header of object.
	HdrObjSysNameOwnerID = "OWNER_ID"

	// HdrObjSysNameVersion is a name of Version field in system header of object.
	HdrObjSysNameVersion = "VERSION"

	// HdrObjSysNamePayloadLength is a name of PayloadLength field in system header of object.
	HdrObjSysNamePayloadLength = "PAYLOAD_LENGTH"

	// HdrObjSysNameCreatedUnix is a name of CreatedAt.UnitTime field in system header of object.
	HdrObjSysNameCreatedUnix = "CREATED_UNIX"

	// HdrObjSysNameCreatedEpoch is a name of CreatedAt.Epoch field in system header of object.
	HdrObjSysNameCreatedEpoch = "CREATED_EPOCH"

	// HdrObjSysLinkPrev is a name of previous link header in extended headers of object.
	HdrObjSysLinkPrev = "LINK_PREV"

	// HdrObjSysLinkNext is a name of next link header in extended headers of object.
	HdrObjSysLinkNext = "LINK_NEXT"

	// HdrObjSysLinkChild is a name of child link header in extended headers of object.
	HdrObjSysLinkChild = "LINK_CHILD"

	// HdrObjSysLinkPar is a name of parent link header in extended headers of object.
	HdrObjSysLinkPar = "LINK_PAR"

	// HdrObjSysLinkSG is a name of storage group link header in extended headers of object.
	HdrObjSysLinkSG = "LINK_SG"
)

func newTypedHeader(name, value string, typ HeaderType) TypedHeader {
	return &typedHeader{
		n: name,
		v: value,
		t: typ,
	}
}

// Name is a name field getter.
func (s typedHeader) Name() string {
	return s.n
}

// Value is a value field getter.
func (s typedHeader) Value() string {
	return s.v
}

// HeaderType is a type field getter.
func (s typedHeader) HeaderType() HeaderType {
	return s.t
}

// TypedHeaderSourceFromObject wraps passed object and returns TypedHeaderSource interface.
func TypedHeaderSourceFromObject(obj *object.Object) TypedHeaderSource {
	return &objectHeaderSource{
		obj: obj,
	}
}

// HeaderOfType gathers object headers of passed type and returns Header list.
//
// If value of some header can not be calculated (e.g. nil extended header), it does not appear in list.
//
// Always returns true.
func (s objectHeaderSource) HeadersOfType(typ HeaderType) ([]Header, bool) {
	if s.obj == nil {
		return nil, true
	}

	var res []Header

	switch typ {
	case HdrTypeObjUsr:
		objHeaders := s.obj.GetHeaders()

		res = make([]Header, 0, len(objHeaders)) // 7 system header fields

		for _, extHdr := range objHeaders {
			if h := newTypedObjectExtendedHeader(extHdr); h != nil {
				res = append(res, h)
			}
		}
	case HdrTypeObjSys:
		res = make([]Header, 0, 7)

		sysHdr := s.obj.GetSystemHeader()

		// ID
		res = append(res, newTypedHeader(
			HdrObjSysNameID,
			sysHdr.ID.String(),
			HdrTypeObjSys),
		)

		// CID
		res = append(res, newTypedHeader(
			HdrObjSysNameCID,
			sysHdr.CID.String(),
			HdrTypeObjSys),
		)

		// OwnerID
		res = append(res, newTypedHeader(
			HdrObjSysNameOwnerID,
			sysHdr.OwnerID.String(),
			HdrTypeObjSys),
		)

		// Version
		res = append(res, newTypedHeader(
			HdrObjSysNameVersion,
			strconv.FormatUint(sysHdr.GetVersion(), 10),
			HdrTypeObjSys),
		)

		// PayloadLength
		res = append(res, newTypedHeader(
			HdrObjSysNamePayloadLength,
			strconv.FormatUint(sysHdr.GetPayloadLength(), 10),
			HdrTypeObjSys),
		)

		created := sysHdr.GetCreatedAt()

		// CreatedAt.UnitTime
		res = append(res, newTypedHeader(
			HdrObjSysNameCreatedUnix,
			strconv.FormatUint(uint64(created.GetUnixTime()), 10),
			HdrTypeObjSys),
		)

		// CreatedAt.Epoch
		res = append(res, newTypedHeader(
			HdrObjSysNameCreatedEpoch,
			strconv.FormatUint(created.GetEpoch(), 10),
			HdrTypeObjSys),
		)
	}

	return res, true
}

func newTypedObjectExtendedHeader(h object.Header) TypedHeader {
	val := h.GetValue()
	if val == nil {
		return nil
	}

	res := new(typedHeader)
	res.t = HdrTypeObjSys

	switch hdr := val.(type) {
	case *object.Header_UserHeader:
		if hdr.UserHeader == nil {
			return nil
		}

		res.t = HdrTypeObjUsr
		res.n = hdr.UserHeader.GetKey()
		res.v = hdr.UserHeader.GetValue()
	case *object.Header_Link:
		if hdr.Link == nil {
			return nil
		}

		switch hdr.Link.GetType() {
		case object.Link_Previous:
			res.n = HdrObjSysLinkPrev
		case object.Link_Next:
			res.n = HdrObjSysLinkNext
		case object.Link_Child:
			res.n = HdrObjSysLinkChild
		case object.Link_Parent:
			res.n = HdrObjSysLinkPar
		case object.Link_StorageGroup:
			res.n = HdrObjSysLinkSG
		default:
			return nil
		}

		res.v = hdr.Link.ID.String()
	default:
		return nil
	}

	return res
}

// TypedHeaderSourceFromExtendedHeaders wraps passed ExtendedHeadersSource and returns TypedHeaderSource interface.
func TypedHeaderSourceFromExtendedHeaders(hdrSrc service.ExtendedHeadersSource) TypedHeaderSource {
	return &extendedHeadersWrapper{
		hdrSrc: hdrSrc,
	}
}

// Name returns the result of Key method.
func (s typedExtendedHeader) Name() string {
	return s.hdr.Key()
}

// Value returns the result of Value method.
func (s typedExtendedHeader) Value() string {
	return s.hdr.Value()
}

// HeaderType always returns HdrTypeRequest.
func (s typedExtendedHeader) HeaderType() HeaderType {
	return HdrTypeRequest
}

// TypedHeaders gathers extended request headers and returns TypedHeader list.
//
// Nil headers are ignored.
//
// Always returns true.
func (s extendedHeadersWrapper) HeadersOfType(typ HeaderType) ([]Header, bool) {
	if s.hdrSrc == nil {
		return nil, true
	}

	var res []Header

	switch typ {
	case HdrTypeRequest:
		hs := s.hdrSrc.ExtendedHeaders()

		res = make([]Header, 0, len(hs))

		for i := range hs {
			if hs[i] == nil {
				continue
			}

			res = append(res, &typedExtendedHeader{
				hdr: hs[i],
			})
		}
	}

	return res, true
}
