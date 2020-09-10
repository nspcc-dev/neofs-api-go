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
