package acl

// ObjectFilterPrefix is a prefix of key to object header value or property.
const ObjectFilterPrefix = "$Object:"

const (
	// FilterObjectVersion is a filter key to "version" field of the object header.
	FilterObjectVersion = ObjectFilterPrefix + "version"

	// FilterObjectID is a filter key to "object_id" field of the object.
	FilterObjectID = ObjectFilterPrefix + "objectID"

	// FilterObjectContainerID is a filter key to "container_id" field of the object header.
	FilterObjectContainerID = ObjectFilterPrefix + "containerID"

	// FilterObjectOwnerID is a filter key to "owner_id" field of the object header.
	FilterObjectOwnerID = ObjectFilterPrefix + "ownerID"

	// FilterObjectCreationEpoch is a filter key to "creation_epoch" field of the object header.
	FilterObjectCreationEpoch = ObjectFilterPrefix + "creationEpoch"

	// FilterObjectPayloadLength is a filter key to "payload_length" field of the object header.
	FilterObjectPayloadLength = ObjectFilterPrefix + "payloadLength"

	// FilterObjectPayloadHash is a filter key to "payload_hash" field of the object header.
	FilterObjectPayloadHash = ObjectFilterPrefix + "payloadHash"

	// FilterObjectType is a filter key to "object_type" field of the object header.
	FilterObjectType = ObjectFilterPrefix + "objectType"

	// FilterObjectHomomorphicHash is a filter key to "homomorphic_hash" field of the object header.
	FilterObjectHomomorphicHash = ObjectFilterPrefix + "homomorphicHash"

	// FilterObjectParent is a filter key to "split.parent" field of the object header.
	FilterObjectParent = ObjectFilterPrefix + "split.parent"
)
