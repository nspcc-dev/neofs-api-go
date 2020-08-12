package signature

import (
	"crypto/ecdsa"

	crypto "github.com/nspcc-dev/neofs-crypto"
)

type DataSource interface {
	ReadSignedData([]byte) ([]byte, error)
	SignedDataSize() int
}

type DataWithSignature interface {
	DataSource
	GetSignatureWithKey() (key, sig []byte)
	SetSignatureWithKey(key, sig []byte)
}

type SignOption func(*cfg)

type KeySignatureHandler func(key []byte, sig []byte)

type KeySignatureSource func() (key, sig []byte)

func DataSignature(key *ecdsa.PrivateKey, src DataSource, opts ...SignOption) ([]byte, error) {
	if key == nil {
		return nil, crypto.ErrEmptyPrivateKey
	}

	data, err := dataForSignature(src)
	if err != nil {
		return nil, err
	}
	defer bytesPool.Put(data)

	cfg := defaultCfg()

	for i := range opts {
		opts[i](cfg)
	}

	return cfg.signFunc(key, data)
}

func SignDataWithHandler(key *ecdsa.PrivateKey, src DataSource, handler KeySignatureHandler, opts ...SignOption) error {
	sig, err := DataSignature(key, src, opts...)
	if err != nil {
		return err
	}

	handler(crypto.MarshalPublicKey(&key.PublicKey), sig)

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

	key, sig := sigSrc()

	return cfg.verifyFunc(
		crypto.UnmarshalPublicKey(key),
		data,
		sig,
	)
}

func SignData(key *ecdsa.PrivateKey, v DataWithSignature, opts ...SignOption) error {
	return SignDataWithHandler(key, v, v.SetSignatureWithKey, opts...)
}

func VerifyData(src DataWithSignature, opts ...SignOption) error {
	return VerifyDataWithSource(src, src.GetSignatureWithKey, opts...)
}
