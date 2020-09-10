package pkg

import (
	"fmt"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

// Version represents v2-compatible version.
type Version refs.Version

const sdkMjr, sdkMnr = 2, 0

// NewVersionFromV2 wraps v2 Version message to Version.
func NewVersionFromV2(v *refs.Version) *Version {
	return (*Version)(v)
}

// NewVersion creates and initializes blank Version.
//
// Works similar as NewVersionFromV2(new(Version)).
func NewVersion() *Version {
	return NewVersionFromV2(new(refs.Version))
}

// SDKVersion returns Version instance that
// initialized to current SDK revision number.
func SDKVersion() *Version {
	v := NewVersion()
	v.SetMajor(sdkMjr)
	v.SetMinor(sdkMnr)

	return v
}

// GetMajor returns major number of the revision.
func (v *Version) GetMajor() uint32 {
	return (*refs.Version)(v).
		GetMajor()
}

// SetMajor sets major number of the revision.
func (v *Version) SetMajor(val uint32) {
	(*refs.Version)(v).
		SetMajor(val)
}

// GetMinor returns minor number of the revision.
func (v *Version) GetMinor() uint32 {
	return (*refs.Version)(v).
		GetMinor()
}

// SetMinor sets minor number of the revision.
func (v *Version) SetMinor(val uint32) {
	(*refs.Version)(v).
		SetMinor(val)
}

// ToV2 converts Version to v2 Version message.
func (v *Version) ToV2() *refs.Version {
	return (*refs.Version)(v)
}

func (v *Version) String() string {
	return fmt.Sprintf("v%d.%d", v.GetMajor(), v.GetMinor())
}
