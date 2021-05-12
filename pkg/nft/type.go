package nft

import (
	"fmt"
)

// Type is an enum for supported NFT standards.
type Type uint32

// Enum of NFT types.
const (
	// NEP11 is a default NFT type.
	NEP11 Type = iota
	// NEP11String is a string representation for NEP-11 NFT type.
	NEP11String = "NEP-11"
)

func (n Type) String() (name string) {
	switch n {
	case NEP11:
		name = NEP11String
	}
	return
}

func ParseNFTType(s string) (Type, error) {
	switch s {
	case "", NEP11String:
		return NEP11, nil
	default:
		return 0, fmt.Errorf("unknown NFT type [%s]", s)
	}
}
