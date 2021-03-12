package refstest

import (
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

func GenerateVersion(empty bool) *refs.Version {
	m := new(refs.Version)

	if !empty {
		m.SetMajor(2)
		m.SetMinor(1)
	}

	return m
}

func GenerateOwnerID(empty bool) *refs.OwnerID {
	m := new(refs.OwnerID)

	if !empty {
		m.SetValue([]byte{1, 2, 3})
	}

	return m
}

func GenerateAddress(empty bool) *refs.Address {
	m := new(refs.Address)

	m.SetObjectID(GenerateObjectID(empty))
	m.SetContainerID(GenerateContainerID(empty))

	return m
}

func GenerateObjectID(empty bool) *refs.ObjectID {
	m := new(refs.ObjectID)

	if !empty {
		m.SetValue([]byte{1, 2, 3})
	}

	return m
}

func GenerateObjectIDs(empty bool) []*refs.ObjectID {
	ids := make([]*refs.ObjectID, 0)

	if !empty {
		ids = append(ids,
			GenerateObjectID(false),
			GenerateObjectID(false),
		)
	}

	return ids
}

func GenerateContainerID(empty bool) *refs.ContainerID {
	m := new(refs.ContainerID)

	if !empty {
		m.SetValue([]byte{1, 2, 3})
	}

	return m
}

func GenerateContainerIDs(empty bool) (res []*refs.ContainerID) {
	if !empty {
		res = append(res,
			GenerateContainerID(false),
			GenerateContainerID(false),
		)
	}

	return
}

func GenerateSignature(empty bool) *refs.Signature {
	m := new(refs.Signature)

	if !empty {
		m.SetKey([]byte{1})
		m.SetSign([]byte{2})
	}

	return m
}

func GenerateChecksum(empty bool) *refs.Checksum {
	m := new(refs.Checksum)

	if !empty {
		m.SetType(1)
		m.SetSum([]byte{1, 2, 3})
	}

	return m
}
