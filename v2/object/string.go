package object

const (
	typeRegularString      = "Regular"
	typeTombstoneString    = "Tombstone"
	typeStorageGroupString = "StorageGroup"
)

func (t Type) String() string {
	switch t {
	default:
		return typeRegularString
	case TypeTombstone:
		return typeTombstoneString
	case TypeStorageGroup:
		return typeStorageGroupString
	}
}

func TypeFromString(s string) Type {
	switch s {
	default:
		return TypeRegular
	case typeTombstoneString:
		return TypeTombstone
	case typeStorageGroupString:
		return TypeStorageGroup
	}
}
