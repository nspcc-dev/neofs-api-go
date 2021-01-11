package session

// ReservedXHeaderPrefix is a prefix of keys to "well-known" X-headers.
const ReservedXHeaderPrefix = "__NEOFS__"

const (
	// XHeaderNetmapEpoch is a key to the reserved X-header that specifies netmap epoch
	// to use for object placement calculation. If set to '0' or not set, the current
	// epoch only will be used.
	XHeaderNetmapEpoch = ReservedXHeaderPrefix + "NETMAP_EPOCH"

	// XHeaderNetmapLookupDepth is a key to the reserved X-header that limits
	// how many past epochs back the node will can lookup. If set to '0' or not
	// set, the current epoch only will be used.
	XHeaderNetmapLookupDepth = ReservedXHeaderPrefix + "NETMAP_LOOKUP_DEPTH"
)
