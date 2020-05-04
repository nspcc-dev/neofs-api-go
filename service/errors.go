package service

import "github.com/nspcc-dev/neofs-api-go/internal"

// ErrNilToken is returned by functions that expect a non-nil token argument, but received nil.
const ErrNilToken = internal.Error("token is nil")

// ErrInvalidTTL means that the TTL value does not satisfy a specific criterion.
const ErrInvalidTTL = internal.Error("invalid TTL value")

// ErrInvalidPublicKeyBytes means that the public key could not be unmarshaled.
const ErrInvalidPublicKeyBytes = internal.Error("cannot load public key")

// ErrCannotFindOwner is raised when signatures empty in GetOwner.
const ErrCannotFindOwner = internal.Error("cannot find owner public key")

// ErrWrongOwner is raised when passed OwnerID not equal to present PublicKey
const ErrWrongOwner = internal.Error("wrong owner")
