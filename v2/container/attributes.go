package container

// SysAttributePrefix is a prefix of key to system attribute.
const SysAttributePrefix = "__NEOFS__"

const (
	// SysAttributeSubnet is a string ID of container's storage subnet.
	SysAttributeSubnet = SysAttributePrefix + "SUBNET"

	// SysAttributeName is a string of human-friendly container name registered as the domain in NNS contract.
	SysAttributeName = SysAttributePrefix + "NAME"

	// SysAttributeZone is a string of zone for container name.
	SysAttributeZone = SysAttributePrefix + "ZONE"
)

// SysAttributeZoneDefault is a default value for SysAttributeZone attribute.
const SysAttributeZoneDefault = "container"
