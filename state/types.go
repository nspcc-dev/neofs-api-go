package state

import (
	"encoding/binary"
)

// SetState is a State field setter.
func (m *ChangeStateRequest) SetState(st ChangeStateRequest_State) {
	m.State = st
}

// Size returns the size of the state binary representation.
func (ChangeStateRequest_State) Size() int {
	return 4
}

// Bytes returns the state binary representation.
func (x ChangeStateRequest_State) Bytes() []byte {
	data := make([]byte, x.Size())

	binary.BigEndian.PutUint32(data, uint32(x))

	return data
}
