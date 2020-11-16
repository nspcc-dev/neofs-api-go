package acl

const (
	// PublicBasicRule is a basic ACL value for public container.
	PublicBasicRule = 0x1FFFFFFF

	// PrivateBasicRule is a basic ACL value for private container.
	PrivateBasicRule = 0x18888888

	// ReadOnlyBasicRule is a basic ACL value for read-only container.
	ReadOnlyBasicRule = 0x1FFF88FF
)
