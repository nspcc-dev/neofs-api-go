package object

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRequest(t *testing.T) {
	cases := []Request{
		&PutRequest{},
		&GetRequest{},
		&HeadRequest{},
		&SearchRequest{},
		&DeleteRequest{},
		&GetRangeRequest{},
		&GetRangeHashRequest{},
	}

	for i := range cases {
		v := cases[i]

		t.Run(fmt.Sprintf("%T", v), func(t *testing.T) {
			require.NotPanics(t, func() { v.CID() })
		})
	}
}
