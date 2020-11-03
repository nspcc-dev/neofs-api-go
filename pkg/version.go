package pkg

import (
	"fmt"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/pkg/errors"
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

// IsSupportedVersion returns error if v is not supported by current SDK.
func IsSupportedVersion(v *Version) error {
	switch mjr := v.GetMajor(); mjr {
	case 2:
		switch mnr := v.GetMinor(); mnr {
		case 0:
			return nil
		}
	}

	return errors.Errorf("unsupported version %d.%d",
		v.GetMajor(),
		v.GetMinor(),
	)
}
