package object

// ReservedFilterPrefix is a prefix of key to object header value or property.
const ReservedFilterPrefix = "$Object:"

const (
	// FilterHeaderVersion is a filter key to "version" field of the object header.
	FilterHeaderVersion = ReservedFilterPrefix + "version"

	// FilterHeaderObjectID is a filter key to "object_id" field of the object.
	FilterHeaderObjectID = ReservedFilterPrefix + "objectID"

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

	// FilterHeaderParent is a filter key to "split.splitID" field of the object header.
	FilterHeaderSplitID = ReservedFilterPrefix + "split.splitID"
)

const (
	// FilterPropertyRoot is a filter key to check if regular object is on top of split hierarchy.
	FilterPropertyRoot = ReservedFilterPrefix + "ROOT"

	// FilterPropertyPhy is a filter key to check if an object physically stored on a node.
	FilterPropertyPhy = ReservedFilterPrefix + "PHY"
)

const (
	// BooleanPropertyValueTrue is a true value for boolean property filters.
	BooleanPropertyValueTrue = "true"

	// BooleanPropertyValueFalse is a false value for boolean property filters.
	BooleanPropertyValueFalse = ""
)
