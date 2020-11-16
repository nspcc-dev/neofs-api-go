package pkg

import (
	"github.com/nspcc-dev/neofs-api-go/v2/session"
)

// XHeader represents v2-compatible XHeader.
type XHeader session.XHeader

// NewXHeaderFromV2 wraps v2 XHeader message to XHeader.
func NewXHeaderFromV2(v *session.XHeader) *XHeader {
	return (*XHeader)(v)
}

// NewXHeader creates, initializes and returns blank XHeader instance.
func NewXHeader() *XHeader {
	return NewXHeaderFromV2(new(session.XHeader))
}

// ToV2 converts XHeader to v2 XHeader message.
func (x *XHeader) ToV2() *session.XHeader {
	return (*session.XHeader)(x)
}

// Key returns key to X-Header.
func (x *XHeader) Key() string {
	return (*session.XHeader)(x).
		GetKey()
}

// SetKey sets key to X-Header.
func (x *XHeader) SetKey(k string) {
	(*session.XHeader)(x).
		SetKey(k)
}

// Value returns value of X-Header.
func (x *XHeader) Value() string {
	return (*session.XHeader)(x).
		GetValue()
}

// SetValue sets value of X-Header.
func (x *XHeader) SetValue(k string) {
	(*session.XHeader)(x).
		SetValue(k)
}
