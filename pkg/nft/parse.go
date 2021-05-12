package nft

import (
	"errors"

	v2container "github.com/nspcc-dev/neofs-api-go/v2/container"
)

var (
	ErrorMissingID      = errors.New("missing NFT-ID attribute")
	ErrorMissingAddress = errors.New("missing NFT-Address attribute")
)

func extractFromV2Container(attrs []*v2container.Attribute) (*Attributes, error) {
	res := new(Attributes)

	for _, attr := range attrs {
		switch attr.GetKey() {
		case v2container.SysAttributeNFTAddress:
			res.address = attr.GetValue()
		case v2container.SysAttributeNFTChain:
			chain, err := ParseNFTChain(attr.GetValue())
			if err != nil {
				return nil, err
			}
			res.chain = chain
		case v2container.SysAttributeNFTID:
			res.id = attr.GetValue()
		case v2container.SysAttributeNFTOptions:
			res.options = attr.GetValue()
		case v2container.SysAttributeNFTType:
			nftType, err := ParseNFTType(attr.GetValue())
			if err != nil {
				return nil, err
			}
			res.nftType = nftType
		default:
		}
	}

	// NFT-ID and NFT-Address are required attributes
	if len(res.id) == 0 {
		return nil, ErrorMissingID
	} else if len(res.address) == 0 {
		return nil, ErrorMissingAddress
	}

	return res, nil
}
