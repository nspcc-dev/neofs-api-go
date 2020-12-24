package container

import (
	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/v2/container"
	"github.com/pkg/errors"
)

// NewVerifiedFromV2 constructs Container from NeoFS API V2 Container message.
// Returns error if message does not meet NeoFS API V2 specification.
//
// Additionally checks if message carries supported version.
func NewVerifiedFromV2(cnrV2 *container.Container) (*Container, error) {
	cnr := NewContainerFromV2(cnrV2)

	// check version support
	if err := pkg.IsSupportedVersion(cnr.Version()); err != nil {
		return nil, err
	}

	// check nonce format
	if _, err := cnr.NonceUUID(); err != nil {
		return nil, errors.Wrap(err, "invalid nonce")
	}

	return cnr, nil
}
