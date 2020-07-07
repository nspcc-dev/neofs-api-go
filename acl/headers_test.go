package acl

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/object"
	"github.com/stretchr/testify/require"
)

func TestNewTypedObjectExtendedHeader(t *testing.T) {
	var res TypedHeader

	hdr := object.Header{}

	// nil value
	require.Nil(t, newTypedObjectExtendedHeader(hdr))

	// UserHeader
	{
		key := "key"
		val := "val"
		hdr.Value = &object.Header_UserHeader{
			UserHeader: &object.UserHeader{
				Key:   key,
				Value: val,
			},
		}

		res = newTypedObjectExtendedHeader(hdr)
		require.Equal(t, HdrTypeObjUsr, res.HeaderType())
		require.Equal(t, key, res.Name())
		require.Equal(t, val, res.Value())
	}

	{ // Link
		link := new(object.Link)
		link.ID = object.ID{1, 2, 3}

		hdr.Value = &object.Header_Link{
			Link: link,
		}

		check := func(lt object.Link_Type, name string) {
			link.Type = lt

			res = newTypedObjectExtendedHeader(hdr)

			require.Equal(t, HdrTypeObjSys, res.HeaderType())
			require.Equal(t, name, res.Name())
			require.Equal(t, link.ID.String(), res.Value())
		}

		check(object.Link_Previous, HdrObjSysLinkPrev)
		check(object.Link_Next, HdrObjSysLinkNext)
		check(object.Link_Parent, HdrObjSysLinkPar)
		check(object.Link_Child, HdrObjSysLinkChild)
		check(object.Link_StorageGroup, HdrObjSysLinkSG)
	}
}
