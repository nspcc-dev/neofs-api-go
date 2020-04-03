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

	types := []RequestType{
		RequestPut,
		RequestGet,
		RequestHead,
		RequestSearch,
		RequestDelete,
		RequestRange,
		RequestRangeHash,
	}

	for i := range cases {
		v := cases[i]

		t.Run(fmt.Sprintf("%T", v), func(t *testing.T) {
			require.NotPanics(t, func() { v.CID() })
			require.Equal(t, types[i], v.Type())
		})
	}
}
