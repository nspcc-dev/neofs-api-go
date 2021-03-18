package proto

type Stream interface {
	WriteUvarint(v uint64) (int, error)
	BytesMarshal(field int, v []byte) (int, error)
	StringMarshal(field int, v string) (int, error)
	BoolMarshal(field int, v bool) (int, error)
	UInt64Marshal(field int, v uint64) (int, error)
	Int64Marshal(field int, v int64) (int, error)
	UInt32Marshal(field int, v uint32) (int, error)
	Int32Marshal(field int, v int32) (int, error)
	EnumMarshal(field int, v int32) (int, error)
	RepeatedBytesMarshal(field int, v [][]byte) (int, error)
	RepeatedStringMarshal(field int, v []string) (int, error)
	RepeatedStringSize(field int, v []string) int
	RepeatedUInt64Marshal(field int, v []uint64) (int, error)
	RepeatedInt64Marshal(field int, v []int64) (int, error)
	RepeatedUInt32Marshal(field int, v []uint32) (int, error)
	RepeatedInt32Marshal(field int, v []int32) (int, error)
	NestedStructureMarshal(field int64, v StreamMarshaler) (int, error)
	Fixed64Marshal(field int, v uint64) (int, error)
}
