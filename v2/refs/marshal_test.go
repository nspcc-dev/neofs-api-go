package refs_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	grpc "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
	"github.com/stretchr/testify/require"
)

func TestOwnerID_StableMarshal(t *testing.T) {
	ownerFrom := new(refs.OwnerID)
	ownerTransport := new(grpc.OwnerID)

	t.Run("non empty", func(t *testing.T) {
		ownerFrom.SetValue([]byte("Owner ID"))

		wire, err := ownerFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = ownerTransport.Unmarshal(wire)
		require.NoError(t, err)

		ownerTo := refs.OwnerIDFromGRPCMessage(ownerTransport)
		require.Equal(t, ownerFrom, ownerTo)
	})
}

func TestContainerID_StableMarshal(t *testing.T) {
	cnrFrom := new(refs.ContainerID)
	cnrTransport := new(grpc.ContainerID)

	t.Run("non empty", func(t *testing.T) {
		cnrFrom.SetValue([]byte("Container ID"))

		wire, err := cnrFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = cnrTransport.Unmarshal(wire)
		require.NoError(t, err)

		cnrTo := refs.ContainerIDFromGRPCMessage(cnrTransport)
		require.Equal(t, cnrFrom, cnrTo)
	})
}

func TestObjectID_StableMarshal(t *testing.T) {
	objectIDFrom := new(refs.ObjectID)
	objectIDTransport := new(grpc.ObjectID)

	t.Run("non empty", func(t *testing.T) {
		objectIDFrom.SetValue([]byte("Object ID"))

		wire, err := objectIDFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = objectIDTransport.Unmarshal(wire)
		require.NoError(t, err)

		objectIDTo := refs.ObjectIDFromGRPCMessage(objectIDTransport)
		require.Equal(t, objectIDFrom, objectIDTo)
	})
}

func TestAddress_StableMarshal(t *testing.T) {
	addressFrom := new(refs.Address)

	cnr := new(refs.ContainerID)
	cnr.SetValue([]byte("Container ID"))

	objectID := new(refs.ObjectID)
	objectID.SetValue([]byte("Object ID"))

	addressTransport := new(grpc.Address)

	t.Run("non empty", func(t *testing.T) {
		addressFrom.SetContainerID(cnr)
		addressFrom.SetObjectID(objectID)

		wire, err := addressFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = addressTransport.Unmarshal(wire)
		require.NoError(t, err)

		addressTo := refs.AddressFromGRPCMessage(addressTransport)
		require.Equal(t, addressFrom, addressTo)
	})
}

func TestChecksum_StableMarshal(t *testing.T) {
	checksumFrom := new(refs.Checksum)
	transport := new(grpc.Checksum)

	t.Run("non empty", func(t *testing.T) {
		checksumFrom.SetType(refs.TillichZemor)
		checksumFrom.SetSum([]byte("Homomorphic Hash"))

		wire, err := checksumFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		checksumTo := refs.ChecksumFromGRPCMessage(transport)
		require.Equal(t, checksumFrom, checksumTo)
	})
}

func TestSignature_StableMarshal(t *testing.T) {
	signatureFrom := generateSignature("Public Key", "Signature")
	transport := new(grpc.Signature)

	t.Run("non empty", func(t *testing.T) {
		wire, err := signatureFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		signatureTo := refs.SignatureFromGRPCMessage(transport)
		require.Equal(t, signatureFrom, signatureTo)
	})
}

func TestVersion_StableMarshal(t *testing.T) {
	versionFrom := generateVersion(2, 0)
	transport := new(grpc.Version)

	t.Run("non empty", func(t *testing.T) {
		wire, err := versionFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		versionTo := refs.VersionFromGRPCMessage(transport)
		require.Equal(t, versionFrom, versionTo)
	})
}

func generateSignature(k, v string) *refs.Signature {
	sig := new(refs.Signature)
	sig.SetKey([]byte(k))
	sig.SetSign([]byte(v))

	return sig
}

func generateVersion(maj, min uint32) *refs.Version {
	version := new(refs.Version)
	version.SetMajor(maj)
	version.SetMinor(min)

	return version
}
