package service

import (
	"github.com/nspcc-dev/neofs-api-go/refs"
)

// TokenID is a type alias of UUID ref.
type TokenID = refs.UUID

// OwnerID is a type alias of OwnerID ref.
type OwnerID = refs.OwnerID

// Address is a type alias of Address ref.
type Address = refs.Address

// AddressContainer is a type alias of refs.AddressContainer.
type AddressContainer = refs.AddressContainer

// OwnerIDContainer is a type alias of refs.OwnerIDContainer.
type OwnerIDContainer = refs.OwnerIDContainer
