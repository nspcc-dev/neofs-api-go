package pkg

import (
	"fmt"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/pkg/errors"
)

// Version represents v2-compatible version.
type Version refs.Version

const sdkMjr, sdkMnr = 2, 6

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

// Major returns major number of the revision.
func (v *Version) Major() uint32 {
	return (*refs.Version)(v).
		GetMajor()
}

// SetMajor sets major number of the revision.
func (v *Version) SetMajor(val uint32) {
	(*refs.Version)(v).
		SetMajor(val)
}

// Minor returns minor number of the revision.
func (v *Version) Minor() uint32 {
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
	return fmt.Sprintf("v%d.%d", v.Major(), v.Minor())
}

// IsSupportedVersion returns error if v is not supported by current SDK.
func IsSupportedVersion(v *Version) error {
	mjr, mnr := v.Major(), v.Minor()

	if mjr != 2 || mnr > sdkMnr {
		return errors.Errorf("unsupported version %d.%d", mjr, mnr)
	}

	return nil
}

// Marshal marshals Version into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (v *Version) Marshal(b ...[]byte) ([]byte, error) {
	var buf []byte
	if len(b) > 0 {
		buf = b[0]
	}

	return (*refs.Version)(v).
		StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of Version.
func (v *Version) Unmarshal(data []byte) error {
	return (*refs.Version)(v).
		Unmarshal(data)
}

// MarshalJSON encodes Version to protobuf JSON format.
func (v *Version) MarshalJSON() ([]byte, error) {
	return (*refs.Version)(v).
		MarshalJSON()
}

// UnmarshalJSON decodes Version from protobuf JSON format.
func (v *Version) UnmarshalJSON(data []byte) error {
	return (*refs.Version)(v).
		UnmarshalJSON(data)
}
