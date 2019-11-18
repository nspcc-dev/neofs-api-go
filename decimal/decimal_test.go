package decimal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecimal_Parse(t *testing.T) {
	tests := []struct {
		value  float64
		name   string
		expect *Decimal
	}{
		{name: "empty", expect: &Decimal{Precision: GASPrecision}},

		{
			value:  100,
			name:   "100 GAS",
			expect: &Decimal{Value: 1e10, Precision: GASPrecision},
		},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expect, ParseFloat(tt.value))
		})
	}
}

func TestDecimal_ParseWithPrecision(t *testing.T) {
	type args struct {
		v float64
		p int
	}
	tests := []struct {
		args   args
		name   string
		expect *Decimal
	}{
		{name: "empty", expect: &Decimal{}},

		{
			name:   "empty precision",
			expect: &Decimal{Value: 0, Precision: 0},
		},

		{
			name:   "100 GAS",
			args:   args{100, GASPrecision},
			expect: &Decimal{Value: 1e10, Precision: GASPrecision},
		},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expect,
				ParseFloatWithPrecision(tt.args.v, tt.args.p))
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name   string
		val    int64
		expect *Decimal
	}{
		{name: "empty", expect: &Decimal{Value: 0, Precision: GASPrecision}},
		{name: "100 GAS", val: 1e10, expect: &Decimal{Value: 1e10, Precision: GASPrecision}},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			require.Equalf(t, tt.expect, New(tt.val), tt.name)
		})
	}
}

func TestNewGAS(t *testing.T) {
	tests := []struct {
		name   string
		val    int64
		expect *Decimal
	}{
		{name: "empty", expect: &Decimal{Value: 0, Precision: GASPrecision}},
		{name: "100 GAS", val: 100, expect: &Decimal{Value: 1e10, Precision: GASPrecision}},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			require.Equalf(t, tt.expect, NewGAS(tt.val), tt.name)
		})
	}
}
func TestNewWithPrecision(t *testing.T) {
	tests := []struct {
		name   string
		val    int64
		pre    uint32
		expect *Decimal
	}{
		{name: "empty", expect: &Decimal{}},
		{name: "100 GAS", val: 1e10, pre: GASPrecision, expect: &Decimal{Value: 1e10, Precision: GASPrecision}},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			require.Equalf(t, tt.expect, NewWithPrecision(tt.val, tt.pre), tt.name)
		})
	}
}

func TestDecimal_Neg(t *testing.T) {
	tests := []struct {
		name   string
		val    int64
		expect *Decimal
	}{
		{name: "empty", expect: &Decimal{Value: 0, Precision: GASPrecision}},
		{name: "100 GAS", val: 1e10, expect: &Decimal{Value: -1e10, Precision: GASPrecision}},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			require.NotPanicsf(t, func() {
				require.Equalf(t, tt.expect, New(tt.val).Neg(), tt.name)
			}, tt.name)
		})
	}
}

func TestDecimal_String(t *testing.T) {
	tests := []struct {
		name   string
		expect string
		value  *Decimal
	}{
		{name: "empty", expect: "0", value: &Decimal{}},
		{name: "100 GAS", expect: "100", value: &Decimal{Value: 1e10, Precision: GASPrecision}},
		{name: "-100 GAS", expect: "-100", value: &Decimal{Value: -1e10, Precision: GASPrecision}},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			require.Equalf(t, tt.expect, tt.value.String(), tt.name)
		})
	}
}

const SomethingElsePrecision = 5

func TestDecimal_Add(t *testing.T) {
	tests := []struct {
		name   string
		expect *Decimal
		values [2]*Decimal
	}{
		{name: "empty", expect: &Decimal{}, values: [2]*Decimal{{}, {}}},
		{
			name:   "5 GAS + 2 GAS",
			expect: &Decimal{Value: 7e8, Precision: GASPrecision},
			values: [2]*Decimal{
				{Value: 2e8, Precision: GASPrecision},
				{Value: 5e8, Precision: GASPrecision},
			},
		},
		{
			name:   "1e2 + 1e3",
			expect: &Decimal{Value: 1.1e3, Precision: 3},
			values: [2]*Decimal{
				{Value: 1e2, Precision: 2},
				{Value: 1e3, Precision: 3},
			},
		},
		{
			name:   "5 GAS + 10 SomethingElse",
			expect: &Decimal{Value: 5.01e8, Precision: GASPrecision},
			values: [2]*Decimal{
				{Value: 5e8, Precision: GASPrecision},
				{Value: 1e6, Precision: SomethingElsePrecision},
			},
		},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			require.NotPanicsf(t, func() {
				{ // A + B
					one := tt.values[0]
					two := tt.values[1]
					require.Equalf(t, tt.expect, one.Add(two), tt.name)
					t.Log(one.Add(two))
				}

				{ // B + A
					one := tt.values[0]
					two := tt.values[1]
					require.Equalf(t, tt.expect, two.Add(one), tt.name)
					t.Log(two.Add(one))
				}
			}, tt.name)
		})
	}
}

func TestDecimal_Copy(t *testing.T) {
	tests := []struct {
		name   string
		expect *Decimal
		value  *Decimal
	}{
		{name: "zero", expect: Zero},
		{
			name:   "5 GAS",
			expect: &Decimal{Value: 5e8, Precision: GASPrecision},
		},
		{
			name:   "100 GAS",
			expect: &Decimal{Value: 1e10, Precision: GASPrecision},
		},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			require.NotPanicsf(t, func() {
				require.Equal(t, tt.expect, tt.expect.Copy())
			}, tt.name)
		})
	}
}

func TestDecimal_Zero(t *testing.T) {
	tests := []struct {
		name   string
		expect bool
		value  *Decimal
	}{
		{name: "zero", expect: true, value: Zero},
		{
			name:   "5 GAS",
			expect: false,
			value:  &Decimal{Value: 5e8, Precision: GASPrecision},
		},
		{
			name:   "100 GAS",
			expect: false,
			value:  &Decimal{Value: 1e10, Precision: GASPrecision},
		},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			require.NotPanicsf(t, func() {
				require.Truef(t, tt.expect == tt.value.Zero(), tt.name)
			}, tt.name)
		})
	}
}

func TestDecimal_Equal(t *testing.T) {
	tests := []struct {
		name   string
		expect bool
		values [2]*Decimal
	}{
		{name: "zero == zero", expect: true, values: [2]*Decimal{Zero, Zero}},
		{
			name:   "5 GAS != 2 GAS",
			expect: false,
			values: [2]*Decimal{
				{Value: 5e8, Precision: GASPrecision},
				{Value: 2e8, Precision: GASPrecision},
			},
		},
		{
			name:   "100 GAS == 100 GAS",
			expect: true,
			values: [2]*Decimal{
				{Value: 1e10, Precision: GASPrecision},
				{Value: 1e10, Precision: GASPrecision},
			},
		},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			require.NotPanicsf(t, func() {
				require.Truef(t, tt.expect == (tt.values[0].Equal(tt.values[1])), tt.name)
			}, tt.name)
		})
	}
}

func TestDecimal_GT(t *testing.T) {
	tests := []struct {
		name   string
		expect bool
		values [2]*Decimal
	}{
		{name: "two zeros", expect: false, values: [2]*Decimal{Zero, Zero}},
		{
			name:   "5 GAS > 2 GAS",
			expect: true,
			values: [2]*Decimal{
				{Value: 5e8, Precision: GASPrecision},
				{Value: 2e8, Precision: GASPrecision},
			},
		},
		{
			name:   "100 GAS !> 100 GAS",
			expect: false,
			values: [2]*Decimal{
				{Value: 1e10, Precision: GASPrecision},
				{Value: 1e10, Precision: GASPrecision},
			},
		},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			require.NotPanicsf(t, func() {
				require.Truef(t, tt.expect == (tt.values[0].GT(tt.values[1])), tt.name)
			}, tt.name)
		})
	}
}

func TestDecimal_GTE(t *testing.T) {
	tests := []struct {
		name   string
		expect bool
		values [2]*Decimal
	}{
		{name: "two zeros", expect: true, values: [2]*Decimal{Zero, Zero}},
		{
			name:   "5 GAS >= 2 GAS",
			expect: true,
			values: [2]*Decimal{
				{Value: 5e8, Precision: GASPrecision},
				{Value: 2e8, Precision: GASPrecision},
			},
		},
		{
			name:   "1 GAS !>= 100 GAS",
			expect: false,
			values: [2]*Decimal{
				{Value: 1e8, Precision: GASPrecision},
				{Value: 1e10, Precision: GASPrecision},
			},
		},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			require.NotPanicsf(t, func() {
				require.Truef(t, tt.expect == (tt.values[0].GTE(tt.values[1])), tt.name)
			}, tt.name)
		})
	}
}

func TestDecimal_LT(t *testing.T) {
	tests := []struct {
		name   string
		expect bool
		values [2]*Decimal
	}{
		{name: "two zeros", expect: false, values: [2]*Decimal{Zero, Zero}},
		{
			name:   "5 GAS !< 2 GAS",
			expect: false,
			values: [2]*Decimal{
				{Value: 5e8, Precision: GASPrecision},
				{Value: 2e8, Precision: GASPrecision},
			},
		},
		{
			name:   "1 GAS < 100 GAS",
			expect: true,
			values: [2]*Decimal{
				{Value: 1e8, Precision: GASPrecision},
				{Value: 1e10, Precision: GASPrecision},
			},
		},
		{
			name:   "100 GAS !< 100 GAS",
			expect: false,
			values: [2]*Decimal{
				{Value: 1e10, Precision: GASPrecision},
				{Value: 1e10, Precision: GASPrecision},
			},
		},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			require.NotPanicsf(t, func() {
				require.Truef(t, tt.expect == (tt.values[0].LT(tt.values[1])), tt.name)
			}, tt.name)
		})
	}
}

func TestDecimal_LTE(t *testing.T) {
	tests := []struct {
		name   string
		expect bool
		values [2]*Decimal
	}{
		{name: "two zeros", expect: true, values: [2]*Decimal{Zero, Zero}},
		{
			name:   "5 GAS <= 2 GAS",
			expect: false,
			values: [2]*Decimal{
				{Value: 5e8, Precision: GASPrecision},
				{Value: 2e8, Precision: GASPrecision},
			},
		},
		{
			name:   "1 GAS <= 100 GAS",
			expect: true,
			values: [2]*Decimal{
				{Value: 1e8, Precision: GASPrecision},
				{Value: 1e10, Precision: GASPrecision},
			},
		},
		{
			name:   "100 GAS !<= 1 GAS",
			expect: false,
			values: [2]*Decimal{
				{Value: 1e10, Precision: GASPrecision},
				{Value: 1e8, Precision: GASPrecision},
			},
		},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			require.NotPanicsf(t, func() {
				require.Truef(t, tt.expect == (tt.values[0].LTE(tt.values[1])), tt.name)
			}, tt.name)
		})
	}
}
