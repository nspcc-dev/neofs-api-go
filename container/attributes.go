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

	// SysAttributeHomomorphicHashing is a container's homomorphic hashing state.
	SysAttributeHomomorphicHashing = SysAttributePrefix + "DISABLE_HOMOMORPHIC_HASHING"
)

// SysAttributeZoneDefault is a default value for SysAttributeZone attribute.
const SysAttributeZoneDefault = "container"

const (
	disabledHomomorphicHashingValue = "true"
)

// HomomorphicHashingState returns container's homomorphic
// hashing state:
// 	* true if hashing is enabled;
// 	* false if hashing is disabled.
//
// All container's attributes must be unique, otherwise behavior
// is undefined.
//
// See also SetHomomorphicHashingState.
func (c Container) HomomorphicHashingState() bool {
	for i := range c.attr {
		if c.attr[i].GetKey() == SysAttributeHomomorphicHashing {
			return c.attr[i].GetValue() != disabledHomomorphicHashingValue
		}
	}

	return true
}

// SetHomomorphicHashingState sets homomorphic hashing state for
// container.
//
// All container's attributes must be unique, otherwise behavior
// is undefined.
//
// See also HomomorphicHashingState.
func (c *Container) SetHomomorphicHashingState(enable bool) {
	for i := range c.attr {
		if c.attr[i].GetKey() == SysAttributeHomomorphicHashing {
			if enable {
				// approach without allocation/waste
				// coping works since the attributes
				// order is not important
				c.attr[i] = c.attr[len(c.attr)-1]
				c.attr = c.attr[:len(c.attr)-1]
			} else {
				c.attr[i].SetValue(disabledHomomorphicHashingValue)
			}

			return
		}
	}

	if !enable {
		attr := Attribute{}
		attr.SetKey(SysAttributeHomomorphicHashing)
		attr.SetValue(disabledHomomorphicHashingValue)

		c.attr = append(c.attr, attr)
	}
}
