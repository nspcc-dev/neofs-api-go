package accounting

import (
	"github.com/nspcc-dev/neofs-api-go/v2/accounting"
)

// Decimal represents v2-compatible decimal number.
type Decimal accounting.Decimal

// NewDecimal creates, initializes and returns blank Decimal instance.
func NewDecimal() *Decimal {
	return NewDecimalFromV2(new(accounting.Decimal))
}

// NewDecimalFromV2 converts v2 Decimal to Decimal.
func NewDecimalFromV2(d *accounting.Decimal) *Decimal {
	return (*Decimal)(d)
}

// Value returns value of the decimal number.
func (d *Decimal) Value() int64 {
	return (*accounting.Decimal)(d).
		GetValue()
}

// SetValue sets value of the decimal number.
func (d *Decimal) SetValue(v int64) {
	(*accounting.Decimal)(d).
		SetValue(v)
}

// Precision returns precision of the decimal number.
func (d *Decimal) Precision() uint32 {
	return (*accounting.Decimal)(d).
		GetPrecision()
}

// SetPrecision sets precision of the decimal number.
func (d *Decimal) SetPrecision(p uint32) {
	(*accounting.Decimal)(d).
		SetPrecision(p)
}
