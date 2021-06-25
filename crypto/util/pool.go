package cryptoutil

import (
	"sync"
)

var bytesPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 5<<20)
	},
}

// Buffer gets buffer with length sz from sync.Pool
// or allocates new one via make.
//
// Buffer can be returned to pool via ReleaseBuffer.
func Buffer(sz int) []byte {
	buf := bytesPool.Get().([]byte)

	if sz <= cap(buf) {
		buf = buf[:sz]
	} else {
		buf = make([]byte, sz)
	}

	return buf
}

// ReleaseBuffer returns the buffer to the sync.Pool.
func ReleaseBuffer(buf []byte) {
	bytesPool.Put(buf)
}
