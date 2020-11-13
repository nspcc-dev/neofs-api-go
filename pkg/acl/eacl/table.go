package eacl

import (
	"crypto/sha256"

	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	v2acl "github.com/nspcc-dev/neofs-api-go/v2/acl"
)

type (
	// Table is a group of EACL records for single container.
	Table struct {
		version pkg.Version
		cid     *container.ID
		records []Record
	}
)

func (t Table) CID() *container.ID {
	return t.cid
}

func (t *Table) SetCID(cid *container.ID) {
	t.cid = cid
}

func (t Table) Version() pkg.Version {
	return t.version
}

func (t *Table) SetVersion(version pkg.Version) {
	t.version = version
}

func (t Table) Records() []Record {
	return t.records
}

func (t *Table) AddRecord(r *Record) {
	if r != nil {
		t.records = append(t.records, *r)
	}
}

func (t *Table) ToV2() *v2acl.Table {
	v2 := new(v2acl.Table)

	if t.cid != nil {
		v2.SetContainerID(t.cid.ToV2())
	}

	records := make([]*v2acl.Record, 0, len(t.records))
	for _, record := range t.records {
		records = append(records, record.ToV2())
	}

	v2.SetVersion(t.version.ToV2())
	v2.SetRecords(records)

	return v2
}

func NewTable() *Table {
	t := new(Table)
	t.SetVersion(*pkg.SDKVersion())

	return t
}

func CreateTable(cid container.ID) *Table {
	t := NewTable()
	t.SetCID(&cid)

	return t
}

func NewTableFromV2(table *v2acl.Table) *Table {
	t := new(Table)

	if table == nil {
		return t
	}

	// set version
	if v := table.GetVersion(); v != nil {
		version := pkg.Version{}
		version.SetMajor(v.GetMajor())
		version.SetMinor(v.GetMinor())

		t.SetVersion(version)
	}

	// set container id
	if cid := table.GetContainerID(); cid != nil {
		if t.cid == nil {
			t.cid = new(container.ID)
		}

		var h [sha256.Size]byte
		copy(h[:], table.GetContainerID().GetValue())
		t.cid.SetSHA256(h)
	}

	// set eacl records
	v2records := table.GetRecords()
	t.records = make([]Record, 0, len(v2records))
	for i := range v2records {
		t.records = append(t.records, *NewRecordFromV2(v2records[i]))
	}

	return t
}

// Marshal marshals Table into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (t *Table) Marshal(b ...[]byte) ([]byte, error) {
	var buf []byte
	if len(b) > 0 {
		buf = b[0]
	}

	return t.ToV2().
		StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of Table.
func (t *Table) Unmarshal(data []byte) error {
	fV2 := new(v2acl.Table)
	if err := fV2.Unmarshal(data); err != nil {
		return err
	}

	*t = *NewTableFromV2(fV2)

	return nil
}

// MarshalJSON encodes Table to protobuf JSON format.
func (t *Table) MarshalJSON() ([]byte, error) {
	return t.ToV2().
		MarshalJSON()
}

// UnmarshalJSON decodes Table from protobuf JSON format.
func (t *Table) UnmarshalJSON(data []byte) error {
	tV2 := new(v2acl.Table)
	if err := tV2.UnmarshalJSON(data); err != nil {
		return err
	}

	*t = *NewTableFromV2(tV2)

	return nil
}
