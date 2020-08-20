package accounting

// SetValue sets value of the decimal number.
func (m *Decimal) SetValue(v int64) {
	if m != nil {
		m.Value = v
	}
}

// SetPrecision sets precision of the decimal number.
func (m *Decimal) SetPrecision(v uint32) {
	if m != nil {
		m.Precision = v
	}
}
