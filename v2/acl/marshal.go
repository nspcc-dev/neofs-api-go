package acl

import (
	protoutil "github.com/nspcc-dev/neofs-api-go/util/proto"
	acl "github.com/nspcc-dev/neofs-api-go/v2/acl/grpc"
	"google.golang.org/protobuf/proto"
)

const (
	filterHeaderTypeField = 1
	filterMatchTypeField  = 2
	filterNameField       = 3
	filterValueField      = 4

	targetTypeField = 1
	targetKeysField = 2

	recordOperationField = 1
	recordActionField    = 2
	recordFiltersField   = 3
	recordTargetsField   = 4

	tableVersionField     = 1
	tableContainerIDField = 2
	tableRecordsField     = 3

	lifetimeExpirationField     = 1
	lifetimeNotValidBeforeField = 2
	lifetimeIssuedAtField       = 3

	bearerTokenBodyACLField      = 1
	bearerTokenBodyOwnerField    = 2
	bearerTokenBodyLifetimeField = 3

	bearerTokenBodyField      = 1
	bearerTokenSignatureField = 2
)

// StableMarshal marshals unified acl table structure in a protobuf
// compatible way without field order shuffle.
func (t *Table) StableMarshal(buf []byte) ([]byte, error) {
	if t == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, t.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = protoutil.NestedStructureMarshal(tableVersionField, buf[offset:], t.version)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = protoutil.NestedStructureMarshal(tableContainerIDField, buf[offset:], t.cid)
	if err != nil {
		return nil, err
	}

	offset += n

	for i := range t.records {
		n, err = protoutil.NestedStructureMarshal(tableRecordsField, buf[offset:], t.records[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	return buf, nil
}

// StableSize of acl table structure marshalled by StableMarshal function.
func (t *Table) StableSize() (size int) {
	if t == nil {
		return 0
	}

	size += protoutil.NestedStructureSize(tableVersionField, t.version)
	size += protoutil.NestedStructureSize(tableContainerIDField, t.cid)

	for i := range t.records {
		size += protoutil.NestedStructureSize(tableRecordsField, t.records[i])
	}

	return size
}

// StableMarshal marshals unified acl record structure in a protobuf
// compatible way without field order shuffle.
func (r *Record) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = protoutil.EnumMarshal(recordOperationField, buf[offset:], int32(r.op))
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = protoutil.EnumMarshal(recordActionField, buf[offset:], int32(r.action))
	if err != nil {
		return nil, err
	}

	offset += n

	for i := range r.filters {
		n, err = protoutil.NestedStructureMarshal(recordFiltersField, buf[offset:], r.filters[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	for i := range r.targets {
		n, err = protoutil.NestedStructureMarshal(recordTargetsField, buf[offset:], r.targets[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	return buf, nil
}

// StableSize of acl record structure marshalled by StableMarshal function.
func (r *Record) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += protoutil.EnumSize(recordOperationField, int32(r.op))
	size += protoutil.EnumSize(recordActionField, int32(r.action))

	for i := range r.filters {
		size += protoutil.NestedStructureSize(recordFiltersField, r.filters[i])
	}

	for i := range r.targets {
		size += protoutil.NestedStructureSize(recordTargetsField, r.targets[i])
	}

	return size
}

// StableMarshal marshals unified header filter structure in a protobuf
// compatible way without field order shuffle.
func (f *HeaderFilter) StableMarshal(buf []byte) ([]byte, error) {
	if f == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, f.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = protoutil.EnumMarshal(filterHeaderTypeField, buf[offset:], int32(f.hdrType))
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = protoutil.EnumMarshal(filterMatchTypeField, buf[offset:], int32(f.matchType))
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = protoutil.StringMarshal(filterNameField, buf[offset:], f.key)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = protoutil.StringMarshal(filterValueField, buf[offset:], f.value)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// StableSize of header filter structure marshalled by StableMarshal function.
func (f *HeaderFilter) StableSize() (size int) {
	if f == nil {
		return 0
	}

	size += protoutil.EnumSize(filterHeaderTypeField, int32(f.hdrType))
	size += protoutil.EnumSize(filterMatchTypeField, int32(f.matchType))
	size += protoutil.StringSize(filterNameField, f.key)
	size += protoutil.StringSize(filterValueField, f.value)

	return size
}

func (f *HeaderFilter) Unmarshal(data []byte) error {
	m := new(acl.EACLRecord_Filter)
	if err := proto.Unmarshal(data, m); err != nil {
		return err
	}

	*f = *HeaderFilterFromGRPCMessage(m)

	return nil
}

// StableMarshal marshals unified role info structure in a protobuf
// compatible way without field order shuffle.
func (t *Target) StableMarshal(buf []byte) ([]byte, error) {
	if t == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, t.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = protoutil.EnumMarshal(targetTypeField, buf[offset:], int32(t.role))
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = protoutil.RepeatedBytesMarshal(targetKeysField, buf[offset:], t.keys)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// StableSize of role info structure marshalled by StableMarshal function.
func (t *Target) StableSize() (size int) {
	if t == nil {
		return 0
	}

	size += protoutil.EnumSize(targetTypeField, int32(t.role))
	size += protoutil.RepeatedBytesSize(targetKeysField, t.keys)

	return size
}

func (t *Target) Unmarshal(data []byte) error {
	m := new(acl.EACLRecord_Target)
	if err := proto.Unmarshal(data, m); err != nil {
		return err
	}

	*t = *TargetInfoFromGRPCMessage(m)

	return nil
}

func (l *TokenLifetime) StableMarshal(buf []byte) ([]byte, error) {
	if l == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, l.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = protoutil.UInt64Marshal(lifetimeExpirationField, buf[offset:], l.exp)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = protoutil.UInt64Marshal(lifetimeNotValidBeforeField, buf[offset:], l.nbf)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = protoutil.UInt64Marshal(lifetimeIssuedAtField, buf[offset:], l.iat)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (l *TokenLifetime) StableSize() (size int) {
	if l == nil {
		return 0
	}

	size += protoutil.UInt64Size(lifetimeExpirationField, l.exp)
	size += protoutil.UInt64Size(lifetimeNotValidBeforeField, l.nbf)
	size += protoutil.UInt64Size(lifetimeIssuedAtField, l.iat)

	return size
}

func (bt *BearerTokenBody) StableMarshal(buf []byte) ([]byte, error) {
	if bt == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, bt.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = protoutil.NestedStructureMarshal(bearerTokenBodyACLField, buf[offset:], bt.eacl)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = protoutil.NestedStructureMarshal(bearerTokenBodyOwnerField, buf[offset:], bt.ownerID)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = protoutil.NestedStructureMarshal(bearerTokenBodyLifetimeField, buf[offset:], bt.lifetime)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (bt *BearerTokenBody) StableSize() (size int) {
	if bt == nil {
		return 0
	}

	size += protoutil.NestedStructureSize(bearerTokenBodyACLField, bt.eacl)
	size += protoutil.NestedStructureSize(bearerTokenBodyOwnerField, bt.ownerID)
	size += protoutil.NestedStructureSize(bearerTokenBodyLifetimeField, bt.lifetime)

	return size
}

func (bt *BearerToken) StableMarshal(buf []byte) ([]byte, error) {
	if bt == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, bt.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = protoutil.NestedStructureMarshal(bearerTokenBodyField, buf[offset:], bt.body)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = protoutil.NestedStructureMarshal(bearerTokenSignatureField, buf[offset:], bt.sig)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (bt *BearerToken) StableSize() (size int) {
	if bt == nil {
		return 0
	}

	size += protoutil.NestedStructureSize(bearerTokenBodyField, bt.body)
	size += protoutil.NestedStructureSize(bearerTokenSignatureField, bt.sig)

	return size
}
