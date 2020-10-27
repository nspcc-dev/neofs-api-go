package object

// ReservedFilterPrefix is a prefix of key to object header value or property.
const ReservedFilterPrefix = "$Object:"

const (
	// FilterHeaderVersion is a filter key to "version" field of the object header.
	FilterHeaderVersion = ReservedFilterPrefix + "version"

	// FilterHeaderContainerID is a filter key to "container_id" field of the object header.
	FilterHeaderContainerID = ReservedFilterPrefix + "containerID"

	// FilterHeaderOwnerID is a filter key to "owner_id" field of the object header.
	FilterHeaderOwnerID = ReservedFilterPrefix + "ownerID"

	// FilterHeaderCreationEpoch is a filter key to "creation_epoch" field of the object header.
	FilterHeaderCreationEpoch = ReservedFilterPrefix + "creationEpoch"

	// FilterHeaderPayloadLength is a filter key to "payload_length" field of the object header.
	FilterHeaderPayloadLength = ReservedFilterPrefix + "payloadLength"

	// FilterHeaderPayloadHash is a filter key to "payload_hash" field of the object header.
	FilterHeaderPayloadHash = ReservedFilterPrefix + "payloadHash"

	// FilterHeaderObjectType is a filter key to "object_type" field of the object header.
	FilterHeaderObjectType = ReservedFilterPrefix + "objectType"

	// FilterHeaderHomomorphicHash is a filter key to "homomorphic_hash" field of the object header.
	FilterHeaderHomomorphicHash = ReservedFilterPrefix + "homomorphicHash"

	// FilterHeaderParent is a filter key to "split.parent" field of the object header.
	FilterHeaderParent = ReservedFilterPrefix + "split.parent"
)

const (
	// FilterPropertyRoot is a filter key to check if an object is a top object in a split hierarchy.
	FilterPropertyRoot = ReservedFilterPrefix + "ROOT"

	// FilterPropertyLeaf is a filter key to check if an object is a leaf in a split hierarchy.
	FilterPropertyLeaf = ReservedFilterPrefix + "LEAF"

	// FilterPropertyChildfree is a filter key to check if an object has empty children list in `Split` header.
	FilterPropertyChildfree = ReservedFilterPrefix + "CHILDFREE"
)
