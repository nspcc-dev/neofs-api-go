package netmap

import (
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
)

// NetworkInfo represents v2-compatible structure
// with information about NeoFS network.
type NetworkInfo netmap.NetworkInfo

// NewNetworkInfoFromV2 wraps v2 NetworkInfo message to NetworkInfo.
func NewNetworkInfoFromV2(iV2 *netmap.NetworkInfo) *NetworkInfo {
	return (*NetworkInfo)(iV2)
}

// NewNetworkInfo creates and initializes blank NetworkInfo.
func NewNetworkInfo() *NetworkInfo {
	return NewNetworkInfoFromV2(new(netmap.NetworkInfo))
}

// ToV2 converts NetworkInfo to v2 NetworkInfo.
func (i *NetworkInfo) ToV2() *netmap.NetworkInfo {
	return (*netmap.NetworkInfo)(i)
}

// CurrentEpoch returns current epoch of the NeoFS network.
func (i *NetworkInfo) CurrentEpoch() uint64 {
	return (*netmap.NetworkInfo)(i).
		GetCurrentEpoch()
}

// SetCurrentEpoch sets current epoch of the NeoFS network.
func (i *NetworkInfo) SetCurrentEpoch(epoch uint64) {
	(*netmap.NetworkInfo)(i).
		SetCurrentEpoch(epoch)
}

// MagicNumber returns magic number of the sidechain.
func (i *NetworkInfo) MagicNumber() uint64 {
	return (*netmap.NetworkInfo)(i).
		GetMagicNumber()
}

// SetMagicNumber sets magic number of the sidechain.
func (i *NetworkInfo) SetMagicNumber(epoch uint64) {
	(*netmap.NetworkInfo)(i).
		SetMagicNumber(epoch)
}

// Marshal marshals NetworkInfo into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (i *NetworkInfo) Marshal(b ...[]byte) ([]byte, error) {
	var buf []byte
	if len(b) > 0 {
		buf = b[0]
	}

	return (*netmap.NetworkInfo)(i).
		StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of NetworkInfo.
func (i *NetworkInfo) Unmarshal(data []byte) error {
	return (*netmap.NetworkInfo)(i).
		Unmarshal(data)
}

// MarshalJSON encodes NetworkInfo to protobuf JSON format.
func (i *NetworkInfo) MarshalJSON() ([]byte, error) {
	return (*netmap.NetworkInfo)(i).
		MarshalJSON()
}

// UnmarshalJSON decodes NetworkInfo from protobuf JSON format.
func (i *NetworkInfo) UnmarshalJSON(data []byte) error {
	return (*netmap.NetworkInfo)(i).
		UnmarshalJSON(data)
}
