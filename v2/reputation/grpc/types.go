package reputation

import (
	refs "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
)

// SetPublicKey sets binary public key of ID.
func (x *PeerID) SetPublicKey(v []byte) {
	if x != nil {
		x.PublicKey = v
	}
}

// SetPeer sets trusted peer's ID.
func (x *Trust) SetPeer(v *PeerID) {
	if x != nil {
		x.Peer = v
	}
}

// SetValue sets trust value.
func (x *Trust) SetValue(v float64) {
	if x != nil {
		x.Value = v
	}
}

// SetTrustingPeer sets trusting peer ID.
func (x *PeerToPeerTrust) SetTrustingPeer(v *PeerID) {
	if x != nil {
		x.TrustingPeer = v
	}
}

// SetTrust sets trust value of trusting peer to the trusted one.
func (x *PeerToPeerTrust) SetTrust(v *Trust) {
	if x != nil {
		x.Trust = v
	}
}

// SetManager sets manager ID.
func (x *GlobalTrust_Body) SetManager(v *PeerID) {
	if x != nil {
		x.Manager = v
	}
}

// SetTrust sets global trust value.
func (x *GlobalTrust_Body) SetTrust(v *Trust) {
	if x != nil {
		x.Trust = v
	}
}

// SetVersion sets message format version.
func (x *GlobalTrust) SetVersion(v *refs.Version) {
	if x != nil {
		x.Version = v
	}
}

// SetBody sets message body.
func (x *GlobalTrust) SetBody(v *GlobalTrust_Body) {
	if x != nil {
		x.Body = v
	}
}

// SetSignature sets body signature.
func (x *GlobalTrust) SetSignature(v *refs.Signature) {
	if x != nil {
		x.Signature = v
	}
}
