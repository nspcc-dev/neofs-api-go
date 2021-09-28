package netmap

import (
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
)

// NetworkInfo represents v2-compatible structure
// with information about NeoFS network.
type NetworkInfo netmap.NetworkInfo

// NewNetworkInfoFromV2 wraps v2 NetworkInfo message to NetworkInfo.
//
// Nil netmap.NetworkInfo converts to nil.
func NewNetworkInfoFromV2(iV2 *netmap.NetworkInfo) *NetworkInfo {
	return (*NetworkInfo)(iV2)
}

// NewNetworkInfo creates and initializes blank NetworkInfo.
//
// Defaults:
//  - curEpoch: 0;
//  - magicNum: 0;
//  - msPerBlock: 0;
//  - network config: nil.
func NewNetworkInfo() *NetworkInfo {
	return NewNetworkInfoFromV2(new(netmap.NetworkInfo))
}

// ToV2 converts NetworkInfo to v2 NetworkInfo.
//
// Nil NetworkInfo converts to nil.
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

// MsPerBlock returns MillisecondsPerBlock network parameter.
func (i *NetworkInfo) MsPerBlock() int64 {
	return (*netmap.NetworkInfo)(i).
		GetMsPerBlock()
}

// SetMsPerBlock sets MillisecondsPerBlock network parameter.
func (i *NetworkInfo) SetMsPerBlock(v int64) {
	(*netmap.NetworkInfo)(i).
		SetMsPerBlock(v)
}

// NetworkConfig returns NeoFS network configuration.
func (i *NetworkInfo) NetworkConfig() *NetworkConfig {
	return NewNetworkConfigFromV2(
		(*netmap.NetworkInfo)(i).
			GetNetworkConfig(),
	)
}

// SetNetworkConfig sets NeoFS network configuration.
func (i *NetworkInfo) SetNetworkConfig(v *NetworkConfig) {
	(*netmap.NetworkInfo)(i).
		SetNetworkConfig(v.ToV2())
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

// NetworkParameter represents v2-compatible NeoFS network parameter.
type NetworkParameter netmap.NetworkParameter

// NewNetworkParameterFromV2 wraps v2 NetworkParameter message to NetworkParameter.
//
// Nil netmap.NetworkParameter converts to nil.
func NewNetworkParameterFromV2(pv2 *netmap.NetworkParameter) *NetworkParameter {
	return (*NetworkParameter)(pv2)
}

// NewNetworkParameter creates and initializes blank NetworkParameter.
//
// Defaults:
//  - key: nil;
//  - value: nil.
func NewNetworkParameter() *NetworkParameter {
	return NewNetworkParameterFromV2(new(netmap.NetworkParameter))
}

// ToV2 converts NetworkParameter to v2 NetworkParameter.
//
// Nil NetworkParameter converts to nil.
func (x *NetworkParameter) ToV2() *netmap.NetworkParameter {
	return (*netmap.NetworkParameter)(x)
}

// Key returns key to network parameter.
func (x *NetworkParameter) Key() []byte {
	return (*netmap.NetworkParameter)(x).
		GetKey()
}

// SetKey sets key to the network parameter.
func (x *NetworkParameter) SetKey(key []byte) {
	(*netmap.NetworkParameter)(x).
		SetKey(key)
}

// Value returns value of the network parameter.
func (x *NetworkParameter) Value() []byte {
	return (*netmap.NetworkParameter)(x).
		GetValue()
}

// SetValue sets value of the network parameter.
func (x *NetworkParameter) SetValue(val []byte) {
	(*netmap.NetworkParameter)(x).
		SetValue(val)
}

// NetworkConfig represents v2-compatible NeoFS network configuration.
type NetworkConfig netmap.NetworkConfig

// NewNetworkConfigFromV2 wraps v2 NetworkConfig message to NetworkConfig.
//
// Nil netmap.NetworkConfig converts to nil.
func NewNetworkConfigFromV2(cv2 *netmap.NetworkConfig) *NetworkConfig {
	return (*NetworkConfig)(cv2)
}

// NewNetworkConfig creates and initializes blank NetworkConfig.
//
// Defaults:
//  - parameters num: 0.
func NewNetworkConfig() *NetworkConfig {
	return NewNetworkConfigFromV2(new(netmap.NetworkConfig))
}

// ToV2 converts NetworkConfig to v2 NetworkConfig.
//
// Nil NetworkConfig converts to nil.
func (x *NetworkConfig) ToV2() *netmap.NetworkConfig {
	return (*netmap.NetworkConfig)(x)
}

// NumberOfParameters returns number of network parameters.
func (x *NetworkConfig) NumberOfParameters() int {
	return (*netmap.NetworkConfig)(x).NumberOfParameters()
}

// IterateAddresses iterates over network parameters.
// Breaks iteration on f's true return.
//
// Handler should not be nil.
func (x *NetworkConfig) IterateParameters(f func(*NetworkParameter) bool) {
	(*netmap.NetworkConfig)(x).
		IterateParameters(func(p *netmap.NetworkParameter) bool {
			return f(NewNetworkParameterFromV2(p))
		})
}

// Value returns value of the network parameter.
func (x *NetworkConfig) SetParameters(ps ...*NetworkParameter) {
	var psV2 []*netmap.NetworkParameter

	if ps != nil {
		ln := len(ps)

		psV2 = make([]*netmap.NetworkParameter, 0, ln)

		for i := 0; i < ln; i++ {
			psV2 = append(psV2, ps[i].ToV2())
		}
	}

	(*netmap.NetworkConfig)(x).
		SetParameters(psV2...)
}
