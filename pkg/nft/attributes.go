package nft

import (
	"github.com/nspcc-dev/neo-go/pkg/util"
	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	v2container "github.com/nspcc-dev/neofs-api-go/v2/container"
)

// Attributes is a structure with all NFT related attributes.
type Attributes struct {
	id      string
	address string
	nftType Type
	chain   Chain
	options string
}

func (n Attributes) ID() string {
	return n.id
}

func (n Attributes) Address() string {
	return n.address
}

func (n Attributes) NFTType() Type {
	return n.nftType
}

func (n Attributes) Chain() Chain {
	return n.chain
}

func (n Attributes) Options() string {
	return n.options
}

func (n Attributes) ToContainerAttributes() container.Attributes {
	res := make(container.Attributes, 0, 5)

	attr := container.NewAttribute()
	attr.SetKey(v2container.SysAttributeNFTID)
	attr.SetValue(n.id)

	res = append(res, attr)

	attr = container.NewAttribute()
	attr.SetKey(v2container.SysAttributeNFTAddress)
	attr.SetValue(n.address)

	res = append(res, attr)

	if n.nftType != NEP11 { // NEP11 is a default type and can be omitted
		attr = container.NewAttribute()
		attr.SetKey(v2container.SysAttributeNFTType)
		attr.SetValue(n.nftType.String())

		res = append(res, attr)
	}

	if n.chain != N3 { // N3 is a default chain and can be omitted
		attr = container.NewAttribute()
		attr.SetKey(v2container.SysAttributeNFTChain)
		attr.SetValue(n.chain.String())

		res = append(res, attr)
	}

	if len(n.options) != 0 {
		attr = container.NewAttribute()
		attr.SetKey(v2container.SysAttributeNFTOptions)
		attr.SetValue(n.options)

		res = append(res, attr)
	}

	return res
}

func (n Attributes) AsN3NFTAttributes() (*N3NFTAttributes, error) {
	u160, err := util.Uint160DecodeStringLE(n.address)
	if err != nil {
		return nil, err
	}

	return &N3NFTAttributes{
		Attributes:  n,
		u160Address: u160,
	}, nil
}

func ExtractFromContainer(attrs container.Attributes) (*Attributes, error) {
	return extractFromV2Container(attrs.ToV2())
}

type (
	// N3NFTAttributes is a wrapper around Attributes that provides
	// parsed ID and Address values.
	N3NFTAttributes struct {
		Attributes
		u160Address util.Uint160
	}

	N3NFTAttributesParam struct {
		ID      []byte
		Address util.Uint160
	}
)

func (n N3NFTAttributes) ID() []byte {
	return []byte(n.id)
}

func (n N3NFTAttributes) Address() util.Uint160 {
	return n.u160Address
}

func NewN3NFTAttributes(param N3NFTAttributesParam, opts ...Option) *N3NFTAttributes {
	defaultCfg := &nftOpt{
		nftType: NEP11,
	}

	for _, opt := range opts {
		opt(defaultCfg)
	}

	nftAttributes := Attributes{
		id:      string(param.ID),
		address: param.Address.StringLE(),
		chain:   N3,
		nftType: defaultCfg.nftType,
		options: defaultCfg.options,
	}

	return &N3NFTAttributes{
		Attributes:  nftAttributes,
		u160Address: param.Address,
	}
}

type (
	nftOpt struct {
		options string
		nftType Type
	}

	Option func(cfg *nftOpt)
)

func WithNFTOptions(options string) Option {
	return func(cfg *nftOpt) {
		cfg.options = options
	}
}

func WithNFTType(t Type) Option {
	return func(cfg *nftOpt) {
		cfg.nftType = t
	}
}
