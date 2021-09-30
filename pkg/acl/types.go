package acl

const (
	// PublicBasicRule is a basic ACL value for public-read-write container.
	PublicBasicRule = 0x1FBFBFFF

	// PrivateBasicRule is a basic ACL value for private container.
	PrivateBasicRule = 0x1C8C8CCC

	// ReadOnlyBasicRule is a basic ACL value for public-read container.
	ReadOnlyBasicRule = 0x1FBF8CFF

	// PublicAppendRule is a basic ACL value for public-append container.
	PublicAppendRule = 0x1FBF9FFF
)
