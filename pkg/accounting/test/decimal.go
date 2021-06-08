package accountingtest

import (
	"github.com/nspcc-dev/neofs-api-go/pkg/accounting"
)

// Generate returns random accounting.Decimal.
func Generate() *accounting.Decimal {
	d := accounting.NewDecimal()
	d.SetValue(1)
	d.SetPrecision(2)

	return d
}
