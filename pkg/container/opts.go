package container

import (
	"github.com/google/uuid"
	"github.com/nspcc-dev/neofs-api-go/pkg/acl"
	"github.com/nspcc-dev/neofs-api-go/pkg/netmap"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
)

type (
	NewOption interface {
		apply(*containerOptions)
	}

	containerOptions struct {
		acl        uint32
		policy     *netmap.PlacementPolicy
		attributes Attributes
		owner      *owner.ID
		nonce      uuid.UUID
	}
)

func defaultContainerOptions() containerOptions {
	rand, err := uuid.NewRandom()
	if err != nil {
		panic("can't create new random " + err.Error())
	}

	return containerOptions{
		acl:   acl.PrivateBasicRule,
		nonce: rand,
	}
}

type funcContainerOption struct {
	f func(*containerOptions)
}

func (fco *funcContainerOption) apply(co *containerOptions) {
	fco.f(co)
}

func newFuncContainerOption(f func(option *containerOptions)) *funcContainerOption {
	return &funcContainerOption{
		f: f,
	}
}

func WithPublicBasicACL() NewOption {
	return newFuncContainerOption(func(option *containerOptions) {
		option.acl = acl.PublicBasicRule
	})
}

func WithReadOnlyBasicACL() NewOption {
	return newFuncContainerOption(func(option *containerOptions) {
		option.acl = acl.ReadOnlyBasicRule
	})
}

func WithCustomBasicACL(acl uint32) NewOption {
	return newFuncContainerOption(func(option *containerOptions) {
		option.acl = acl
	})
}

func WithNonce(nonce uuid.UUID) NewOption {
	return newFuncContainerOption(func(option *containerOptions) {
		option.nonce = nonce
	})
}

func WithOwnerID(id *owner.ID) NewOption {
	return newFuncContainerOption(func(option *containerOptions) {
		option.owner = id
	})
}

func WithNEO3Wallet(w *owner.NEO3Wallet) NewOption {
	return newFuncContainerOption(func(option *containerOptions) {
		if option.owner == nil {
			option.owner = new(owner.ID)
		}

		option.owner.SetNeo3Wallet(w)
	})
}

func WithPolicy(policy *netmap.PlacementPolicy) NewOption {
	return newFuncContainerOption(func(option *containerOptions) {
		option.policy = policy
	})
}

func WithAttribute(key, value string) NewOption {
	return newFuncContainerOption(func(option *containerOptions) {
		attr := NewAttribute()
		attr.SetKey(key)
		attr.SetValue(value)

		option.attributes = append(option.attributes, attr)
	})
}
