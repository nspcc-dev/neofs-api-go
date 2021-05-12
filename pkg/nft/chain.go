package nft

import (
	"fmt"
)

// Chain is an enum for supported chains with NFT.
type Chain uint32

// Enum of NFT chains.
const (
	// N3 is a default NFT chain.
	N3 Chain = iota
	// N3String is a string representation for N3 NFT chain.
	N3String = "N3"
)

func (n Chain) String() (name string) {
	switch n {
	case N3:
		name = "N3"
	}
	return
}

func ParseNFTChain(s string) (Chain, error) {
	switch s {
	case "", N3String:
		return N3, nil
	default:
		return 0, fmt.Errorf("unknown NFT chain [%s]", s)
	}
}
