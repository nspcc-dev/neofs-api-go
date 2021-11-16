package object

// SysAttributePrefix is a prefix of key to system attribute.
const SysAttributePrefix = "__NEOFS__"

const (
	// SysAttributeUploadID marks smaller parts of a split bigger object.
	SysAttributeUploadID = SysAttributePrefix + "UPLOAD_ID"

	// SysAttributeExpEpoch tells GC to delete object after that epoch.
	SysAttributeExpEpoch = SysAttributePrefix + "EXPIRATION_EPOCH"
)
