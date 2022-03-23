package accounting

// SetValue sets value of the decimal number.
func (m *Decimal) SetValue(v int64) {
	m.Value = v
}

// SetPrecision sets precision of the decimal number.
func (m *Decimal) SetPrecision(v uint32) {
	m.Precision = v
}
