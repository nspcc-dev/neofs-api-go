package container

// SysAttributePrefix is a prefix of key to system attribute.
const SysAttributePrefix = "__NEOFS__"

const (
	// SysAttributeSubnet is a string ID of container's storage subnet.
	SysAttributeSubnet = SysAttributePrefix + "SUBNET"
)

// NFT related attributes.
const (
	// SysAttributeNFTID is a string of token ID associated with the container.
	SysAttributeNFTID = SysAttributePrefix + "NFT-ID"
	// SysAttributeNFTAddress is a address of the contract that produced NFT,
	// encoded in LittleEndian.
	SysAttributeNFTAddress = SysAttributePrefix + "NFT-Address"
	// SysAttributeNFTType is a type of NFT. NeoFS supports only `NEP-11` NFT.
	SysAttributeNFTType = SysAttributePrefix + "NFT-Type"
	// SysAttributeNFTChain is a NFT blockchain. NeoFS supports only `N3`
	// main net.
	SysAttributeNFTChain = SysAttributePrefix + "NFT-Chain"
	// SysAttributeNFTOptions is an additional data for later use.
	SysAttributeNFTOptions = SysAttributePrefix + "NFT-Options"
)
