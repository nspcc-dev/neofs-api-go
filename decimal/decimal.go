package decimal

import (
	"math"
	"strconv"
	"strings"
)

// GASPrecision contains precision for NEO Gas token.
const GASPrecision = 8

// Zero is empty Decimal value.
var Zero = &Decimal{}

// New returns new Decimal (in satoshi).
func New(v int64) *Decimal {
	return NewWithPrecision(v, GASPrecision)
}

// NewGAS returns new Decimal * 1e8 (in GAS).
func NewGAS(v int64) *Decimal {
	v *= int64(math.Pow10(GASPrecision))
	return NewWithPrecision(v, GASPrecision)
}

// NewWithPrecision returns new Decimal with custom precision.
func NewWithPrecision(v int64, p uint32) *Decimal {
	return &Decimal{Value: v, Precision: p}
}

// ParseFloat return new Decimal parsed from float64 * 1e8 (in GAS).
func ParseFloat(v float64) *Decimal {
	return new(Decimal).Parse(v, GASPrecision)
}

// ParseFloatWithPrecision returns new Decimal parsed from float64 * 1^p.
func ParseFloatWithPrecision(v float64, p int) *Decimal {
	return new(Decimal).Parse(v, p)
}

// Copy returns copy of current Decimal.
func (m *Decimal) Copy() *Decimal { return &Decimal{Value: m.Value, Precision: m.Precision} }

// Parse returns parsed Decimal from float64 * 1^p.
func (m *Decimal) Parse(v float64, p int) *Decimal {
	m.Value = int64(v * math.Pow10(p))
	m.Precision = uint32(p)
	return m
}

// String returns string representation of Decimal.
func (m Decimal) String() string {
	buf := new(strings.Builder)
	val := m.Value
	dec := int64(math.Pow10(int(m.Precision)))
	if val < 0 {
		buf.WriteRune('-')
		val = -val
	}
	str := strconv.FormatInt(val/dec, 10)
	buf.WriteString(str)
	val %= dec
	if val > 0 {
		buf.WriteRune('.')
		str = strconv.FormatInt(val, 10)
		for i := len(str); i < int(m.Precision); i++ {
			buf.WriteRune('0')
		}
		buf.WriteString(strings.TrimRight(str, "0"))
	}
	return buf.String()
}

// Add returns d + m.
func (m Decimal) Add(d *Decimal) *Decimal {
	precision := m.Precision
	if precision < d.Precision {
		precision = d.Precision
	}
	return &Decimal{
		Value:     m.Value + d.Value,
		Precision: precision,
	}
}

// Zero checks that Decimal is empty.
func (m Decimal) Zero() bool { return m.Value == 0 }

// Equal checks that current Decimal is equal to passed Decimal.
func (m Decimal) Equal(v *Decimal) bool { return m.Value == v.Value && m.Precision == v.Precision }

// GT checks that m > v.
func (m Decimal) GT(v *Decimal) bool { return m.Value > v.Value }

// GTE checks that m >= v.
func (m Decimal) GTE(v *Decimal) bool { return m.Value >= v.Value }

// LT checks that m < v.
func (m Decimal) LT(v *Decimal) bool { return m.Value < v.Value }

// LTE checks that m <= v.
func (m Decimal) LTE(v *Decimal) bool { return m.Value <= v.Value }

// Neg returns negative representation of current Decimal (m * -1).
func (m Decimal) Neg() *Decimal {
	return &Decimal{
		Value:     m.Value * -1,
		Precision: m.Precision,
	}
}
