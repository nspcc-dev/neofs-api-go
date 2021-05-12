package nft_test

import (
	"testing"

	"github.com/nspcc-dev/neo-go/pkg/util"
	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	"github.com/nspcc-dev/neofs-api-go/pkg/nft"
	v2container "github.com/nspcc-dev/neofs-api-go/v2/container"
	"github.com/stretchr/testify/require"
)

const (
	address = "0172fc1ed977a6e37f261d1727f675ccb05360db"
	id      = "5qPQkgjk9orvP25G1ygCRytJH4FCUuMdZtLE8FGhDxz5"
	options = "sator arepo tenet opera rotas"
)

func TestExtractFromContainer(t *testing.T) {
	from := createAttributes(withID(id), withAddress(address), withOptions(options))

	nftAttrs, err := nft.ExtractFromContainer(from)
	require.NoError(t, err)
	require.Equal(t, nft.NEP11, nftAttrs.NFTType())
	require.Equal(t, nft.N3, nftAttrs.Chain())
	require.Equal(t, address, nftAttrs.Address())
	require.Equal(t, id, nftAttrs.ID())
	require.Equal(t, options, nftAttrs.Options())

	t.Run("missing ID", func(t *testing.T) {
		from = createAttributes(withAddress(address), withOptions(options))
		_, err = nft.ExtractFromContainer(from)
		require.EqualError(t, err, nft.ErrorMissingID.Error())
	})

	t.Run("missing address", func(t *testing.T) {
		from = createAttributes(withID(id), withOptions(options))
		_, err = nft.ExtractFromContainer(from)
		require.EqualError(t, err, nft.ErrorMissingAddress.Error())
	})

	t.Run("invalid chain", func(t *testing.T) {
		from = createAttributes(withID(id), withAddress(address), withChain("bad"))
		_, err = nft.ExtractFromContainer(from)
		require.Error(t, err)
	})

	t.Run("invalid NFT type", func(t *testing.T) {
		from = createAttributes(withID(id), withAddress(address), withType("bad"))
		_, err = nft.ExtractFromContainer(from)
		require.Error(t, err)
	})
}

func TestAsN3NFTAttribute(t *testing.T) {
	from := createAttributes(withID(id), withAddress(address), withOptions(options))
	u160, err := util.Uint160DecodeStringLE(address)
	require.NoError(t, err)

	nftAttrs, err := nft.ExtractFromContainer(from)
	require.NoError(t, err)

	n3Attrs, err := nftAttrs.AsN3NFTAttributes()
	require.NoError(t, err)

	require.Equal(t, nft.NEP11, n3Attrs.NFTType())
	require.Equal(t, nft.N3, n3Attrs.Chain())
	require.Equal(t, options, n3Attrs.Options())
	require.Equal(t, u160, n3Attrs.Address())
	require.Equal(t, []byte(id), n3Attrs.ID())

	t.Run("invalid address", func(t *testing.T) {
		from = createAttributes(withID(id), withAddress("invalid"))

		nftAttrs, err = nft.ExtractFromContainer(from)
		require.NoError(t, err)

		_, err = nftAttrs.AsN3NFTAttributes()
		require.Error(t, err)
	})
}

func TestToContainerAttributes(t *testing.T) {
	from := createAttributes(withID(id), withAddress(address), withOptions(options))

	nftAttrs, err := nft.ExtractFromContainer(from)
	require.NoError(t, err)

	n3Attrs, err := nftAttrs.AsN3NFTAttributes()
	require.NoError(t, err)

	to := n3Attrs.ToContainerAttributes()
	require.Equal(t, from, to)
}

func TestNewN3NFTAttributes(t *testing.T) {
	u160, err := util.Uint160DecodeStringLE(address)
	require.NoError(t, err)

	n3Attrs := nft.NewN3NFTAttributes(nft.N3NFTAttributesParam{
		ID:      []byte(id),
		Address: u160,
	}, nft.WithNFTType(nft.NEP11), nft.WithNFTOptions(options))

	to := createAttributes(withID(id), withAddress(address), withOptions(options))
	require.Equal(t, n3Attrs.ToContainerAttributes(), to)
}

type (
	attrOpt struct {
		attrs container.Attributes
	}

	attrOption func(c *attrOpt)
)

func createAttributes(opts ...attrOption) container.Attributes {
	cfg := &attrOpt{
		attrs: make(container.Attributes, 0, 5),
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return cfg.attrs
}

func withID(s string) attrOption {
	return func(c *attrOpt) {
		attr := container.NewAttribute()
		attr.SetKey(v2container.SysAttributeNFTID)
		attr.SetValue(s)
		c.attrs = append(c.attrs, attr)
	}
}

func withAddress(s string) attrOption {
	return func(c *attrOpt) {
		attr := container.NewAttribute()
		attr.SetKey(v2container.SysAttributeNFTAddress)
		attr.SetValue(s)
		c.attrs = append(c.attrs, attr)
	}
}

func withOptions(s string) attrOption {
	return func(c *attrOpt) {
		attr := container.NewAttribute()
		attr.SetKey(v2container.SysAttributeNFTOptions)
		attr.SetValue(s)
		c.attrs = append(c.attrs, attr)
	}
}

func withType(s string) attrOption {
	return func(c *attrOpt) {
		attr := container.NewAttribute()
		attr.SetKey(v2container.SysAttributeNFTType)
		attr.SetValue(s)
		c.attrs = append(c.attrs, attr)
	}
}

func withChain(s string) attrOption {
	return func(c *attrOpt) {
		attr := container.NewAttribute()
		attr.SetKey(v2container.SysAttributeNFTChain)
		attr.SetValue(s)
		c.attrs = append(c.attrs, attr)
	}
}
