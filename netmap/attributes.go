package netmap

import (
	"errors"
	"fmt"
	"strings"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

// prefix of keys to subnet attributes.
const attrSubnetPrefix = "__NEOFS__SUBNET_"

const (
	// subnet attribute's value denoting subnet entry
	attrSubnetValEntry = "True"

	// subnet attribute's value denoting subnet exit
	attrSubnetValExit = "False"
)

// NodeSubnetInfo groups information about subnet which can be written to NodeInfo.
//
// Zero value represents entry to zero subnet.
type NodeSubnetInfo struct {
	exit bool

	id *refs.SubnetID
}

// Enters returns true iff node enters the subnet.
func (x NodeSubnetInfo) Enters() bool {
	return !x.exit
}

// SetEntryFlag sets the subnet entry flag.
func (x *NodeSubnetInfo) SetEntryFlag(enters bool) {
	x.exit = !enters
}

// ID returns identifier of the subnet.
func (x NodeSubnetInfo) ID() *refs.SubnetID {
	return x.id
}

// SetID sets identifier of the subnet.
func (x *NodeSubnetInfo) SetID(id *refs.SubnetID) {
	x.id = id
}

func subnetAttributeKey(id *refs.SubnetID) string {
	txt, _ := id.MarshalText() // never returns an error

	return attrSubnetPrefix + string(txt)
}

// WriteSubnetInfo writes NodeSubnetInfo to NodeInfo via attributes. NodeInfo must not be nil.
//
// Does not add (removes existing) attribute if node:
//   * exists non-zero subnet;
//   * enters zero subnet.
//
// Attribute key is calculated from ID using format `__NEOFS__SUBNET_%s`.
// Attribute Value is:
//   * `True` if node enters the subnet;
//   * `False`, otherwise.
func WriteSubnetInfo(node *NodeInfo, info NodeSubnetInfo) {
	attrs := node.GetAttributes()

	id := info.ID()
	enters := info.Enters()

	// calculate attribute key
	key := subnetAttributeKey(id)

	if refs.IsZeroSubnet(id) == enters {
		for i := range attrs {
			if attrs[i].GetKey() == key {
				attrs = append(attrs[:i], attrs[i+1:]...)
			}
		}
	} else {
		var val string

		if enters {
			val = attrSubnetValEntry
		} else {
			val = attrSubnetValExit
		}

		presented := false

		for i := range attrs {
			if attrs[i].GetKey() == key {
				attrs[i].SetValue(val)
				presented = true
			}
		}

		if !presented {
			var attr Attribute

			attr.SetKey(key)
			attr.SetValue(val)

			attrs = append(attrs, &attr)
		}
	}

	node.SetAttributes(attrs)
}

// ErrRemoveSubnet is returned when a node needs to leave the subnet.
var ErrRemoveSubnet = errors.New("remove subnet")

// IterateSubnets iterates over all subnets the node belongs to and passes the IDs to f.
// Handler must not be nil.
//
// If f returns ErrRemoveSubnet, then removes subnet entry. Breaks on any other non-nil error and returns it.
//
// Returns an error if any subnet attribute has wrong format.
func IterateSubnets(node *NodeInfo, f func(refs.SubnetID) error) error {
	attrs := node.GetAttributes()

	var (
		err     error
		id      refs.SubnetID
		metZero bool // if zero subnet's attribute was met in for-loop
	)

	for i := 0; i < len(attrs); i++ { // range must not be used because of attrs mutation in body
		key := attrs[i].GetKey()

		// cut subnet ID string
		idTxt := strings.TrimPrefix(key, attrSubnetPrefix)
		if idTxt == key {
			// not a subnet attribute
			continue
		}

		// check value
		switch val := attrs[i].GetValue(); val {
		default:
			return fmt.Errorf("invalid attribute value: %s", val)
		case attrSubnetValExit:
			// node is outside the subnet
			continue
		case attrSubnetValEntry:
			// required to avoid default case
		}

		// decode subnet ID
		if err = id.UnmarshalText([]byte(idTxt)); err != nil {
			return fmt.Errorf("invalid ID text: %w", err)
		}

		// pass ID to the handler
		err = f(id)

		isRemoveErr := errors.Is(err, ErrRemoveSubnet)

		if err != nil && !isRemoveErr {
			return err
		}

		if !metZero { // in order to not reset if has been already set
			metZero = refs.IsZeroSubnet(&id)

			if !isRemoveErr {
				// no handler's error and non-zero subnet
				continue
			} else if metZero {
				// removal error and zero subnet.
				// we don't remove attribute of zero subnet because it means entry
				attrs[i].SetValue(attrSubnetValExit)

				continue
			}
		}

		if isRemoveErr {
			// removal error and non-zero subnet.
			// we can set False or remove attribute, latter is more memory/network efficient.
			attrs = append(attrs[:i], attrs[i+1:]...)
			i--
		}
	}

	if !metZero {
		// missing attribute of zero subnet equivalent to entry
		refs.MakeZeroSubnet(&id)

		err = f(id)
		if errors.Is(err, ErrRemoveSubnet) {
			// zero subnet should be clearly removed with False value
			var attr Attribute

			attr.SetKey(subnetAttributeKey(&id))
			attr.SetValue(attrSubnetValExit)

			attrs = append(attrs, &attr)
		}
	}

	node.SetAttributes(attrs)

	return nil
}
