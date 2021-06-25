package apicrypto

import (
	neofscrypto "github.com/nspcc-dev/neofs-api-go/crypto"
	cryprotobuf "github.com/nspcc-dev/neofs-api-go/crypto/proto"
	cryptoutil "github.com/nspcc-dev/neofs-api-go/crypto/util"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

type commonPrm struct {
	protoMarshaler cryprotobuf.StableProtoMarshaler
}

// SignPrm groups the parameters of Sign operation.
type SignPrm struct {
	commonPrm

	sigMsg *refs.Signature
}

// SetProtoMarshaler sets cryprotobuf.StableProtoMarshaler (SPM) which will
// be used as signed data source.
//
// It is a required parameter and must not be nil.
func (x *commonPrm) SetProtoMarshaler(v cryprotobuf.StableProtoMarshaler) {
	x.protoMarshaler = v
}

// SetTargetSignature sets target signature message (TSM) to write
// calculated signature and marshaled signer's public key.
//
// It is a required parameter and must not be nil.
func (x *SignPrm) SetTargetSignature(sigMsg *refs.Signature) {
	x.sigMsg = sigMsg
}

// Sign signs (SPM) data with signer using neofscrypto.Sign and
// writes the signature and marshaled signer's public key to (TSM).
//
// Buffer for reading data is allocated using cryptoutil.Buffer
// and released using cryptoutil.ReleaseBuffer.
func Sign(signer neofscrypto.Signer, prm SignPrm) error {
	// it seems to be easier than neofscrypto.Sign
	binKey, err := signer.Public().MarshalBinary()
	if err != nil {
		return err
	}

	var sigPrm neofscrypto.SignPrm

	cryprotobuf.SetDataSource(&sigPrm, prm.protoMarshaler, cryptoutil.Buffer)
	sigPrm.SetSignatureHandler(prm.sigMsg.SetSign)
	sigPrm.WrapSignedDataHandler(cryptoutil.ReleaseBuffer)

	err = neofscrypto.Sign(signer, sigPrm)
	if err != nil {
		return err
	}

	prm.sigMsg.SetKey(binKey)

	return nil
}

// VerifyPrm groups the parameters of Verify operation.
type VerifyPrm struct {
	commonPrm

	sig []byte
}

// SetSignature sets signature to check (S).
//
// (S) should not be modified before the end of the Verify call.
func (x *VerifyPrm) SetSignature(s []byte) {
	x.sig = s
}

// Verify verifies the signature (S) of (SPM) using neofscrypto.Verify.
//
// Buffer for reading data is allocated using cryptoutil.Buffer
// and released using cryptoutil.ReleaseBuffer.
func Verify(key neofscrypto.PublicKey, prm VerifyPrm) bool {
	var verPrm neofscrypto.VerifyPrm

	cryprotobuf.SetDataSource(&verPrm, prm.protoMarshaler, cryptoutil.Buffer)
	verPrm.WrapSignedDataHandler(cryptoutil.ReleaseBuffer)
	verPrm.SetSignature(prm.sig)

	return neofscrypto.Verify(key, verPrm)
}
