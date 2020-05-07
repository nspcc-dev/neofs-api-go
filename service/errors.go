package service

import "github.com/nspcc-dev/neofs-api-go/internal"

// ErrNilToken is returned by functions that expect
// a non-nil token argument, but received nil.
const ErrNilToken = internal.Error("token is nil")

// ErrInvalidTTL means that the TTL value does not
// satisfy a specific criterion.
const ErrInvalidTTL = internal.Error("invalid TTL value")

// ErrInvalidPublicKeyBytes means that the public key could not be unmarshaled.
const ErrInvalidPublicKeyBytes = internal.Error("cannot load public key")

// ErrCannotFindOwner is raised when signatures empty in GetOwner.
const ErrCannotFindOwner = internal.Error("cannot find owner public key")

// ErrWrongOwner is raised when passed OwnerID
// not equal to present PublicKey
const ErrWrongOwner = internal.Error("wrong owner")

// ErrNilSignedDataSource returned by functions that expect a non-nil
// SignedDataSource, but received nil.
const ErrNilSignedDataSource = internal.Error("signed data source is nil")

// ErrNilSignatureKeySource is returned by functions that expect a non-nil
// SignatureKeySource, but received nil.
const ErrNilSignatureKeySource = internal.Error("empty key-signature source")

// ErrEmptyDataWithSignature is returned by functions that expect
// a non-nil DataWithSignature, but received nil.
const ErrEmptyDataWithSignature = internal.Error("empty data with signature")

// ErrNegativeLength is returned by functions that received
// negative length for slice allocation.
const ErrNegativeLength = internal.Error("negative slice length")

// ErrNilDataWithTokenSignAccumulator is returned by functions that expect
// a non-nil DataWithTokenSignAccumulator, but received nil.
const ErrNilDataWithTokenSignAccumulator = internal.Error("signed data with token is nil")

// ErrNilSignatureKeySourceWithToken is returned by functions that expect
// a non-nil SignatureKeySourceWithToken, but received nil.
const ErrNilSignatureKeySourceWithToken = internal.Error("key-signature source with token is nil")
