package container

import (
	"github.com/google/uuid"
	"github.com/nspcc-dev/neofs-api-go/pkg/acl"
	"github.com/nspcc-dev/neofs-api-go/pkg/refs"
)

type (
	NewOption interface {
		apply(*containerOptions)
	}

	attribute struct {
		key   string
		value string
	}

	containerOptions struct {
		acl        uint32
		policy     string
		attributes []attribute
		owner      *refs.NEO3Wallet
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

func WithOwner(owner refs.NEO3Wallet) NewOption {
	return newFuncContainerOption(func(option *containerOptions) {
		option.owner = &owner
	})
}

func WithPolicy(policy string) NewOption {
	return newFuncContainerOption(func(option *containerOptions) {
		// todo: make sanity check and store binary structure
		option.policy = policy
	})
}

func WithAttribute(key, value string) NewOption {
	return newFuncContainerOption(func(option *containerOptions) {
		attr := attribute{
			key:   key,
			value: value,
		}
		option.attributes = append(option.attributes, attr)
	})
}
