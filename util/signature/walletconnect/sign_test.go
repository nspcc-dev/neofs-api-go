package walletconnect

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSignMessage(t *testing.T) {
	p1, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)

	msg := []byte("NEO")
	result, err := SignMessage(p1, msg)
	require.NoError(t, err)
	require.Equal(t, elliptic.MarshalCompressed(elliptic.P256(), p1.PublicKey.X, p1.PublicKey.Y), result.PublicKey)
	require.Equal(t, saltSize, len(result.Salt))
	require.Equal(t, 64, len(result.Data))
	require.Equal(t, 4+1+16*2+3+2, len(result.Message))

	require.True(t, VerifyMessage(&p1.PublicKey, result))

	t.Run("invalid public key", func(t *testing.T) {
		p2, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		require.NoError(t, err)
		require.False(t, VerifyMessage(&p2.PublicKey, result))
	})
	t.Run("invalid signature", func(t *testing.T) {
		result := result
		result.Data[0] ^= 0xFF
		require.False(t, VerifyMessage(&p1.PublicKey, result))
	})
}

func TestSign(t *testing.T) {
	p1, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)

	msg := []byte("NEO")
	sign, err := Sign(p1, msg)
	require.NoError(t, err)
	require.True(t, Verify(&p1.PublicKey, msg, sign))

	t.Run("invalid public key", func(t *testing.T) {
		p2, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		require.NoError(t, err)
		require.False(t, Verify(&p2.PublicKey, msg, sign))
	})
	t.Run("invalid signature", func(t *testing.T) {
		sign[0] ^= 0xFF
		require.False(t, Verify(&p1.PublicKey, msg, sign))
	})
}

func TestVerifyNeonWallet(t *testing.T) {
	testCases := [...]struct {
		publicKey       string
		data            string
		salt            string
		messageHex      string
		messageOriginal string
	}{
		{ // Test values from this GIF https://github.com/CityOfZion/neon-wallet/pull/2390 .
			publicKey:       "02ce6228ba2cb2fc235be93aff9cd5fc0851702eb9791552f60db062f01e3d83f6",
			data:            "90ab1886ca0bece59b982d9ade8f5598065d651362fb9ce45ad66d0474b89c0b80913c8f0118a282acbdf200a429ba2d81bc52534a53ab41a2c6dfe2f0b4fb1b",
			salt:            "d41e348afccc2f3ee45cd9f5128b16dc",
			messageHex:      "010001f05c6434316533343861666363633266336565343563643966353132386231366463436172616c686f2c206d756c65712c206f2062616775697520656820697373756d65726d6f2074616978206c696761646f206e61206d697373e36f3f0000",
			messageOriginal: "436172616c686f2c206d756c65712c206f2062616775697520656820697373756d65726d6f2074616978206c696761646f206e61206d697373e36f3f",
		},
		{ // Test value from wallet connect integration test
			publicKey:       "03bd9108c0b49f657e9eee50d1399022bd1e436118e5b7529a1b7cd606652f578f",
			data:            "510caa8cb6db5dedf04d215a064208d64be7496916d890df59aee132db8f2b07532e06f7ea664c4a99e3bcb74b43a35eb9653891b5f8701d2aef9e7526703eaa",
			salt:            "2c5b189569e92cce12e1c640f23e83ba",
			messageHex:      "010001f02632633562313839353639653932636365313265316336343066323365383362613132333435360000",
			messageOriginal: "313233343536", // ascii string "123456"
		},
		{ // Test value from wallet connect integration test
			publicKey:       "03bd9108c0b49f657e9eee50d1399022bd1e436118e5b7529a1b7cd606652f578f",
			data:            "1e13f248962d8b3b60708b55ddf448d6d6a28c6b43887212a38b00bf6bab695e61261e54451c6e3d5f1f000e5534d166c7ca30f662a296d3a9aafa6d8c173c01",
			salt:            "58c86b2e74215b4f36b47d731236be3b",
			messageHex:      "010001f02035386338366232653734323135623466333662343764373331323336626533620000",
			messageOriginal: "", // empty string
		},
	}

	for _, testCase := range testCases {
		rawPub, err := hex.DecodeString(testCase.publicKey)
		require.NoError(t, err)
		data, err := hex.DecodeString(testCase.data)
		require.NoError(t, err)
		salt, err := hex.DecodeString(testCase.salt)
		require.NoError(t, err)
		msg, err := hex.DecodeString(testCase.messageHex)
		require.NoError(t, err)
		orig, err := hex.DecodeString(testCase.messageOriginal)
		require.NoError(t, err)

		require.Equal(t, msg, createMessageWithSalt(orig, salt))

		sm := SignedMessage{
			Data:      data,
			Message:   msg,
			PublicKey: rawPub,
			Salt:      salt,
		}
		require.True(t, VerifyMessage(nil, sm))
	}

}
