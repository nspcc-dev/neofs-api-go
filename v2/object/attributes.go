package object

// SysAttrPrefix is a prefix of key to system attribute.
const SysAttrPrefix = "__NEOFS__"

const (
	// SysAttrUploadID marks smaller parts of a split bigger object.
	SysAttrUploadID = SysAttrPrefix + "UPLOAD_ID"

	// SysAttrExpEpoch tells GC to delete object after that epoch.
	SysAttrExpEpoch = SysAttrPrefix + "EXPIRATION_EPOCH"
)
