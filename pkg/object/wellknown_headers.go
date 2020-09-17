package object

// This file contains well-known header names for eACL filters and search
// request filters. Both of them encode header type as string. There are
// constant strings for well-known object headers.

const (
	// ReservedHeaderNamePrefix used in filter names to specify well known
	// headers such as container id, object id, owner id, etc.
	// All names without this prefix used to lookup through user defined headers
	// in object or x-headers in request.
	ReservedHeaderNamePrefix = "_"

	// HdrSysNameID is a name of ID field in system header of object.
	HdrSysNameID = ReservedHeaderNamePrefix + "ID"

	// HdrSysNameCID is a name of cid field in system header of object.
	HdrSysNameCID = ReservedHeaderNamePrefix + "CID"

	// HdrSysNameOwnerID is a name of OwnerID field in system header of object.
	HdrSysNameOwnerID = ReservedHeaderNamePrefix + "OWNER_ID"

	// HdrSysNameVersion is a name of version field in system header of object.
	HdrSysNameVersion = ReservedHeaderNamePrefix + "VERSION"

	// HdrSysNamePayloadLength is a name of PayloadLength field in system header of object.
	HdrSysNamePayloadLength = ReservedHeaderNamePrefix + "PAYLOAD_LENGTH"

	// HdrSysNameCreatedEpoch is a name of CreatedAt.Epoch field in system header of object.
	HdrSysNameCreatedEpoch = ReservedHeaderNamePrefix + "CREATED_EPOCH"
)
