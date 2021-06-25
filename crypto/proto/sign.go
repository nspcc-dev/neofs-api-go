package cryprotobuf

import (
	neofscrypto "github.com/nspcc-dev/neofs-api-go/crypto"
)

// StableProtoMarshaler is the interface implemented by an object that can
// marshal itself into a stable protobuf binary form. Stable means the
// independence of the result from the number and place of calls.
type StableProtoMarshaler interface {
	// StableMarshal reads binary protobuf data into a buffer.
	//
	// Must return an error if the buffer size is not enough.
	StableMarshal([]byte) error

	// StableSize returns size of protobuf data.f
	//
	// Size must not be negative.
	StableSize() int
}

// BufferAllocator is a function which allocates
// byte slices of particular length.
type BufferAllocator func(int) []byte

// SetDataSource composes data source of neofscrypto.SignPrm from StableProtoMarshaler.
//
// To read the data to be signed, a buffer is allocated using BufferAllocator.
// If alloc if not provided (nil), buffer is allocated via make.
//
// StableProtoMarshaler is required and should not be nil.
func SetDataSource(prm interface {
	SetDataSource(neofscrypto.DataSource)
}, m StableProtoMarshaler, alloc BufferAllocator) {
	prm.SetDataSource(func() ([]byte, error) {
		var (
			sz = m.StableSize()

			buf []byte
		)

		if alloc != nil {
			buf = alloc(sz)
		} else {
			buf = make([]byte, sz)
		}

		return buf, m.StableMarshal(buf)
	})
}
