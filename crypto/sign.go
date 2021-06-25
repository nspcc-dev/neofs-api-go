package neofscrypto

// BytesHandler is a function over []byte.
type BytesHandler func([]byte)

// DataSource is a data read function.
type DataSource func() ([]byte, error)

type commonPrm struct {
	dataSrc DataSource

	hData BytesHandler
}

// SetDataSource sets signed data source (SRC).
//
// The parameter is required and should not be nil.
func (x *commonPrm) SetDataSource(f DataSource) {
	x.dataSrc = f
}

// WrapSignedDataHandler wraps handler of the signed data buffer.
// If some handler is already wrapped, it will be called before f.
//
// If provided, all wrapped handlers are called right after
// the signature is calculated.
func (x *commonPrm) WrapSignedDataHandler(f BytesHandler) {
	x.hData = f
}

// SignPrm groups the parameters of Sign operation.
//
// Required parameters:
//  - data source (SRC);
//  - signature handler (SH).
type SignPrm struct {
	commonPrm

	hSign BytesHandler
}

// SetSignatureHandler sets handler of the calculated signature (SH).
//
// The parameter is required and should not be nil.
func (x *SignPrm) SetSignatureHandler(f BytesHandler) {
	x.hSign = f
}

// Sign reads signed data from source (SRC), calculates signature (S) via Signer,
// and passes (S) into parameterized (SH).
//
// If (KH) is proved, signer's public key is marshaled and passed into it.
// If MarshalBinary returns an error, it is ignored w/o handler call.
//
// Panics if SRC is nil. Panics if SH is nil.
func Sign(signer Signer, prm SignPrm) error {
	data, err := prm.dataSrc()
	if err != nil {
		return err
	}

	sig, err := signer.Sign(data)
	if err != nil {
		return err
	}

	if prm.hData != nil {
		prm.hData(data)
	}

	prm.hSign(sig)

	return nil
}

// VerifyPrm groups the parameters of Verify operation.
//
// Required parameters:
//  - data source (SRC);
//  - signature (S).
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

// Verify verifies the signature (S) of the data from (SRC) via PublicKey.
// Return true if signature is correct.
//
// Panics if SRC is nil.
func Verify(key PublicKey, prm VerifyPrm) bool {
	data, err := prm.dataSrc()
	if err != nil {
		return false
	}

	valid := key.Verify(
		data,
		prm.sig,
	)

	if prm.hData != nil {
		prm.hData(data)
	}

	return valid
}
