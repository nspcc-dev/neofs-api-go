package storagegroup

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/mr-tron/base58"
	"github.com/nspcc-dev/neofs-api-go/hash"
	"github.com/nspcc-dev/neofs-api-go/refs"
)

type (
	// Hash is alias of hash.Hash for proto definition.
	Hash = hash.Hash

	// Provider is an interface that defines storage group instance.
	// There was different storage group implementation. Right now it
	// is implemented as extended header in the object.
	Provider interface {
		// Group returns list of object ids of the storage group.
		// This list **should be** sorted.
		Group() []refs.ObjectID

		IDInfo() *IdentificationInfo
		Zones() []ZoneInfo
	}

	// ZoneInfo provides validation information of storage group.
	ZoneInfo struct {
		hash.Hash
		Size uint64
	}

	// IdentificationInfo provides meta information about storage group.
	IdentificationInfo struct {
		CID     refs.CID
		SGID    refs.SGID
		OwnerID refs.OwnerID
	}

	// IDList is a slice of object ids, that can be sorted.
	IDList []refs.ObjectID
)

var _ proto.Message = (*StorageGroup)(nil)

// String returns string representation of StorageGroup.
func (m *StorageGroup) String() string {
	b := new(strings.Builder)
	b.WriteString("<SG")
	b.WriteString(" VDS=" + strconv.FormatUint(m.ValidationDataSize, 10))
	data := base58.Encode(m.ValidationHash.Bytes())
	b.WriteString(" Hash=" + data[:3] + "..." + data[len(data)-3:])
	if m.Lifetime != nil {
		b.WriteString(" Lifetime=(" + m.Lifetime.Unit.String() + " " + strconv.FormatInt(m.Lifetime.Value, 10) + ")")
	}
	b.WriteByte('>')
	return b.String()
}

// Empty checks if storage group has some data for validation.
func (m StorageGroup) Empty() bool {
	return m.ValidationDataSize == 0 && m.ValidationHash.Equal(hash.Hash{})
}

// Len returns amount of object ids in IDList.
func (s IDList) Len() int { return len(s) }

// Less returns byte comparision between IDList[i] and IDList[j].
func (s IDList) Less(i, j int) bool { return bytes.Compare(s[i].Bytes(), s[j].Bytes()) == -1 }

// Swap swaps element with i and j index in IDList.
func (s IDList) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// CalculateSize combines length of all zones in storage group.
func CalculateSize(sg Provider) (size uint64) {
	zoneList := sg.Zones()
	for i := range zoneList {
		size += zoneList[i].Size
	}
	return
}

// CalculateHash returns homomorphic sum of hashes
// fromm all zones in storage group.
func CalculateHash(sg Provider) (hash.Hash, error) {
	var (
		zones  = sg.Zones()
		hashes = make([]hash.Hash, len(zones))
	)
	for i := range zones {
		hashes[i] = zones[i].Hash
	}
	return hash.Concat(hashes)
}
