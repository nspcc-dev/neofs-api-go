package object

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestByteSize_String(t *testing.T) {
	cases := []struct {
		name   string
		expect string
		actual ByteSize
	}{
		{
			name:   "0 bytes",
			expect: "0",
			actual: ByteSize(0),
		},
		{
			name:   "101 bytes",
			expect: "101",
			actual: ByteSize(101),
		},
		{
			name:   "112.84KB",
			expect: "112.84KB",
			actual: ByteSize(115548),
		},
		{
			name:   "80.44MB",
			expect: "80.44MB",
			actual: ByteSize(84347453),
		},
		{
			name:   "905.144GB",
			expect: "905.144GB",
			actual: ByteSize(971891061884),
		},
		{
			name:   "1.857TB",
			expect: "1.857TB",
			actual: ByteSize(2041793092780),
		},
	}

	for i := range cases {
		tt := cases[i]
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expect, tt.actual.String())
		})
	}
}
