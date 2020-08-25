package pkg

import (
	"fmt"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

type (
	Version struct {
		Major uint32
		Minor uint32
	}
)

var SDKVersion = Version{2, 0}

func (v Version) String() string {
	return fmt.Sprintf("v%d.%d", v.Major, v.Minor)
}

func (v Version) ToV2Version() *refs.Version {
	result := new(refs.Version)
	result.SetMajor(v.Major)
	result.SetMinor(v.Minor)

	return result
}
