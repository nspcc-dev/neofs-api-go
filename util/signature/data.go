package signature

import (
	"crypto/ecdsa"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	crypto "github.com/nspcc-dev/neofs-crypto"
)

type DataSource interface {
	ReadSignedData([]byte) ([]byte, error)
	SignedDataSize() int
}

type DataWithSignature interface {
	DataSource
	GetSignature() *refs.Signature
	SetSignature(*refs.Signature)
}

type SignOption func(*cfg)

type KeySignatureHandler func(*refs.Signature)

type KeySignatureSource func() *refs.Signature

func SignDataWithHandler(key *ecdsa.PrivateKey, src DataSource, handler KeySignatureHandler, opts ...SignOption) error {
	if key == nil {
		return crypto.ErrEmptyPrivateKey
	}

	data, err := dataForSignature(src)
	if err != nil {
		return err
	}
	defer bytesPool.Put(data)

	cfg := defaultCfg()

	for i := range opts {
		opts[i](cfg)
	}

	sigData, err := sign(cfg, key, data)
	if err != nil {
		return err
	}

	sig := new(refs.Signature)
	if cfg.useScheme != cfg.defaultScheme {
		sig.SetScheme(cfg.useScheme)
	}
	sig.SetKey(crypto.MarshalPublicKey(&key.PublicKey))
	sig.SetSign(sigData)
	handler(sig)

	return nil
}

func VerifyDataWithSource(dataSrc DataSource, sigSrc KeySignatureSource, opts ...SignOption) error {
	data, err := dataForSignature(dataSrc)
	if err != nil {
		return err
	}
	defer bytesPool.Put(data)

	cfg := defaultCfg()

	for i := range opts {
		opts[i](cfg)
	}

	return verify(cfg, data, sigSrc())
}

func SignData(key *ecdsa.PrivateKey, v DataWithSignature, opts ...SignOption) error {
	return SignDataWithHandler(key, v, v.SetSignature, opts...)
}

func VerifyData(src DataWithSignature, opts ...SignOption) error {
	return VerifyDataWithSource(src, src.GetSignature, opts...)
}
