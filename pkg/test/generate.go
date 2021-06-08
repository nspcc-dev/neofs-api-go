package refstest

import (
	"crypto/sha256"
	"math/rand"

	"github.com/nspcc-dev/neofs-api-go/pkg"
)

// Checksum returns random pkg.Checksum.
func Checksum() *pkg.Checksum {
	var cs [sha256.Size]byte

	rand.Read(cs[:])

	x := pkg.NewChecksum()

	x.SetSHA256(cs)

	return x
}

// Signature returns random pkg.Signature.
func Signature() *pkg.Signature {
	x := pkg.NewSignature()

	x.SetKey([]byte("key"))
	x.SetSign([]byte("sign"))

	return x
}

// Version returns random pkg.Version.
func Version() *pkg.Version {
	x := pkg.NewVersion()

	x.SetMajor(2)
	x.SetMinor(1)

	return x
}

// XHeader returns random pkg.XHeader.
func XHeader() *pkg.XHeader {
	x := pkg.NewXHeader()

	x.SetKey("key")
	x.SetValue("value")

	return x
}
