package object

import (
	"crypto/rand"
	"crypto/sha256"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	"github.com/nspcc-dev/neofs-api-go/v2/object"
	"github.com/stretchr/testify/require"
)

func randID(t *testing.T) *ID {
	id := NewID()
	id.SetSHA256(randSHA256Checksum(t))

	return id
}

func randSHA256Checksum(t *testing.T) (cs [sha256.Size]byte) {
	_, err := rand.Read(cs[:])
	require.NoError(t, err)

	return
}

func randTZChecksum(t *testing.T) (cs [64]byte) {
	_, err := rand.Read(cs[:])
	require.NoError(t, err)

	return
}

func TestRawObject_SetID(t *testing.T) {
	obj := NewRaw()

	id := randID(t)

	obj.SetID(id)

	require.Equal(t, id, obj.GetID())
}

func TestRawObject_SetSignature(t *testing.T) {
	obj := NewRaw()

	sig := pkg.NewSignature()
	sig.SetKey([]byte{1, 2, 3})
	sig.SetSign([]byte{4, 5, 6})

	obj.SetSignature(sig)

	require.Equal(t, sig, obj.GetSignature())
}

func TestRawObject_SetPayload(t *testing.T) {
	obj := NewRaw()

	payload := make([]byte, 10)
	_, _ = rand.Read(payload)

	obj.SetPayload(payload)

	require.Equal(t, payload, obj.GetPayload())
}

func TestRawObject_SetVersion(t *testing.T) {
	obj := NewRaw()

	ver := pkg.NewVersion()
	ver.SetMajor(1)
	ver.SetMinor(2)

	obj.SetVersion(ver)

	require.Equal(t, ver, obj.GetVersion())
}

func TestRawObject_SetPayloadSize(t *testing.T) {
	obj := NewRaw()

	sz := uint64(133)
	obj.SetPayloadSize(sz)

	require.Equal(t, sz, obj.GetPayloadSize())
}

func TestRawObject_SetContainerID(t *testing.T) {
	obj := NewRaw()

	checksum := randSHA256Checksum(t)

	cid := container.NewID()
	cid.SetSHA256(checksum)

	obj.SetContainerID(cid)

	require.Equal(t, cid, obj.GetContainerID())
}

func TestRawObject_SetOwnerID(t *testing.T) {
	obj := NewRaw()

	w := new(owner.NEO3Wallet)
	_, _ = rand.Read(w.Bytes())

	ownerID := owner.NewID()
	ownerID.SetNeo3Wallet(w)

	obj.SetOwnerID(ownerID)

	require.Equal(t, ownerID, obj.GetOwnerID())
}

func TestRawObject_SetCreationEpoch(t *testing.T) {
	obj := NewRaw()

	creat := uint64(228)
	obj.setCreationEpoch(creat)

	require.Equal(t, creat, obj.GetCreationEpoch())
}

func TestRawObject_SetPayloadChecksum(t *testing.T) {
	obj := NewRaw()

	cs := pkg.NewChecksum()
	cs.SetSHA256(randSHA256Checksum(t))

	obj.SetPayloadChecksum(cs)

	require.Equal(t, cs, obj.GetPayloadChecksum())
}

func TestRawObject_SetPayloadHomomorphicHash(t *testing.T) {
	obj := NewRaw()

	cs := pkg.NewChecksum()
	cs.SetTillichZemor(randTZChecksum(t))

	obj.SetPayloadHomomorphicHash(cs)

	require.Equal(t, cs, obj.GetPayloadHomomorphicHash())
}

func TestRawObject_SetAttributes(t *testing.T) {
	obj := NewRaw()

	a1 := NewAttribute()
	a1.SetKey("key1")
	a1.SetValue("val1")

	a2 := NewAttribute()
	a2.SetKey("key2")
	a2.SetValue("val2")

	obj.SetAttributes(a1, a2)

	require.Equal(t, []*Attribute{a1, a2}, obj.GetAttributes())
}

func TestRawObject_SetPreviousID(t *testing.T) {
	obj := NewRaw()

	prev := randID(t)

	obj.SetPreviousID(prev)

	require.Equal(t, prev, obj.GetPreviousID())
}

func TestRawObject_SetChildren(t *testing.T) {
	obj := NewRaw()

	id1 := randID(t)
	id2 := randID(t)

	obj.SetChildren(id1, id2)

	require.Equal(t, []*ID{id1, id2}, obj.GetChildren())
}

func TestRawObject_SetParent(t *testing.T) {
	obj := NewRaw()

	par := NewRaw()
	par.SetID(randID(t))
	parObj := par.Object()

	obj.SetParent(parObj)

	require.Equal(t, parObj, obj.GetParent())
}

func TestRawObject_ToV2(t *testing.T) {
	objV2 := new(object.Object)
	objV2.SetPayload([]byte{1, 2, 3})

	obj := NewRawFromV2(objV2)

	require.Equal(t, objV2, obj.ToV2())
}
