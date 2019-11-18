package chain

import (
	"crypto/ecdsa"
	"encoding/hex"
	"testing"

	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
)

func TestAddress(t *testing.T) {
	var (
		multiSigVerificationScript = "512103c02a93134f98d9c78ec54b1b1f97fc64cd81360f53a293f41e4ad54aac3c57172103fea219d4ccfd7641cebbb2439740bb4bd7c4730c1abd6ca1dc44386533816df952ae"
		multiSigAddress            = "ANbvKqa2SfgTUkq43NRUhCiyxPrpUPn7S3"

		normalVerificationScript = "2102a33413277a319cc6fd4c54a2feb9032eba668ec587f307e319dc48733087fa61ac"
		normalAddress            = "AcraNnCuPKnUYtPYyrACRCVJhLpvskbfhu"
	)

	t.Run("check multi-sig address", func(t *testing.T) {
		data, err := hex.DecodeString(multiSigVerificationScript)
		require.NoError(t, err)
		require.Equal(t, multiSigAddress, Address(data))
	})

	t.Run("check normal address", func(t *testing.T) {
		data, err := hex.DecodeString(normalVerificationScript)
		require.NoError(t, err)
		require.Equal(t, normalAddress, Address(data))
	})
}

func TestVerificationScript(t *testing.T) {
	t.Run("check normal", func(t *testing.T) {
		pkString := "02a33413277a319cc6fd4c54a2feb9032eba668ec587f307e319dc48733087fa61"

		pkBytes, err := hex.DecodeString(pkString)
		require.NoError(t, err)

		pk := crypto.UnmarshalPublicKey(pkBytes)

		expect, err := hex.DecodeString(
			"21" + pkString + // PUSHBYTES33
				"ac", // CHECKSIG
		)

		require.Equal(t, expect, VerificationScript(pk))
	})

	t.Run("check multisig", func(t *testing.T) {
		pk1String := "03c02a93134f98d9c78ec54b1b1f97fc64cd81360f53a293f41e4ad54aac3c5717"
		pk2String := "03fea219d4ccfd7641cebbb2439740bb4bd7c4730c1abd6ca1dc44386533816df9"

		pk1Bytes, err := hex.DecodeString(pk1String)
		require.NoError(t, err)

		pk1 := crypto.UnmarshalPublicKey(pk1Bytes)

		pk2Bytes, err := hex.DecodeString(pk2String)
		require.NoError(t, err)

		pk2 := crypto.UnmarshalPublicKey(pk2Bytes)

		expect, err := hex.DecodeString(
			"51" + // one address
				"21" + pk1String + // PUSHBYTES33
				"21" + pk2String + // PUSHBYTES33
				"52" + // 2 PublicKeys
				"ae", // CHECKMULTISIG
		)

		require.Equal(t, expect, VerificationScript(pk1, pk2))
	})
}

func TestKeysToAddress(t *testing.T) {
	t.Run("check normal", func(t *testing.T) {
		pkString := "02a33413277a319cc6fd4c54a2feb9032eba668ec587f307e319dc48733087fa61"

		pkBytes, err := hex.DecodeString(pkString)
		require.NoError(t, err)

		pk := crypto.UnmarshalPublicKey(pkBytes)

		expect := "AcraNnCuPKnUYtPYyrACRCVJhLpvskbfhu"

		actual := KeysToAddress(pk)
		require.Equal(t, expect, actual)
		require.NoError(t, IsAddress(actual))
	})

	t.Run("check multisig", func(t *testing.T) {
		pk1String := "03c02a93134f98d9c78ec54b1b1f97fc64cd81360f53a293f41e4ad54aac3c5717"
		pk2String := "03fea219d4ccfd7641cebbb2439740bb4bd7c4730c1abd6ca1dc44386533816df9"

		pk1Bytes, err := hex.DecodeString(pk1String)
		require.NoError(t, err)

		pk1 := crypto.UnmarshalPublicKey(pk1Bytes)

		pk2Bytes, err := hex.DecodeString(pk2String)
		require.NoError(t, err)

		pk2 := crypto.UnmarshalPublicKey(pk2Bytes)

		expect := "ANbvKqa2SfgTUkq43NRUhCiyxPrpUPn7S3"
		actual := KeysToAddress(pk1, pk2)
		require.Equal(t, expect, actual)
		require.NoError(t, IsAddress(actual))
	})
}

func TestFetchPublicKeys(t *testing.T) {
	var (
		multiSigVerificationScript = "512103c02a93134f98d9c78ec54b1b1f97fc64cd81360f53a293f41e4ad54aac3c57172103fea219d4ccfd7641cebbb2439740bb4bd7c4730c1abd6ca1dc44386533816df952ae"
		normalVerificationScript   = "2102a33413277a319cc6fd4c54a2feb9032eba668ec587f307e319dc48733087fa61ac"

		pk1String = "03c02a93134f98d9c78ec54b1b1f97fc64cd81360f53a293f41e4ad54aac3c5717"
		pk2String = "03fea219d4ccfd7641cebbb2439740bb4bd7c4730c1abd6ca1dc44386533816df9"
		pk3String = "02a33413277a319cc6fd4c54a2feb9032eba668ec587f307e319dc48733087fa61"
	)

	t.Run("shouls not fail", func(t *testing.T) {
		wrongVS, err := hex.DecodeString(multiSigVerificationScript)
		require.NoError(t, err)

		wrongVS[len(wrongVS)-1] = 0x1

		wrongPK, err := hex.DecodeString(multiSigVerificationScript)
		require.NoError(t, err)
		wrongPK[2] = 0x1

		var testCases = []struct {
			name  string
			value []byte
		}{
			{name: "empty VerificationScript"},
			{
				name:  "wrong size VerificationScript",
				value: []byte{0x1},
			},
			{
				name:  "wrong VerificationScript type",
				value: wrongVS,
			},
			{
				name:  "wrong public key in VerificationScript",
				value: wrongPK,
			},
		}

		for i := range testCases {
			tt := testCases[i]
			t.Run(tt.name, func(t *testing.T) {
				var keys []*ecdsa.PublicKey
				require.NotPanics(t, func() {
					keys = FetchPublicKeys(tt.value)
				})
				require.Nil(t, keys)
			})
		}
	})

	t.Run("check multi-sig address", func(t *testing.T) {
		data, err := hex.DecodeString(multiSigVerificationScript)
		require.NoError(t, err)

		pk1Bytes, err := hex.DecodeString(pk1String)
		require.NoError(t, err)

		pk2Bytes, err := hex.DecodeString(pk2String)
		require.NoError(t, err)

		pk1 := crypto.UnmarshalPublicKey(pk1Bytes)
		pk2 := crypto.UnmarshalPublicKey(pk2Bytes)

		keys := FetchPublicKeys(data)
		require.Len(t, keys, 2)
		require.Equal(t, keys[0], pk1)
		require.Equal(t, keys[1], pk2)
	})

	t.Run("check normal address", func(t *testing.T) {
		data, err := hex.DecodeString(normalVerificationScript)
		require.NoError(t, err)

		pkBytes, err := hex.DecodeString(pk3String)
		require.NoError(t, err)

		pk := crypto.UnmarshalPublicKey(pkBytes)

		keys := FetchPublicKeys(data)
		require.Len(t, keys, 1)
		require.Equal(t, keys[0], pk)
	})

	t.Run("generate 10 keys VerificationScript and try parse it", func(t *testing.T) {
		var (
			count  = 10
			expect = make([]*ecdsa.PublicKey, 0, count)
		)

		for i := 0; i < count; i++ {
			key := test.DecodeKey(i)
			expect = append(expect, &key.PublicKey)
		}

		vs := VerificationScript(expect...)

		actual := FetchPublicKeys(vs)
		require.Equal(t, expect, actual)
	})
}

func TestReversedScriptHashToAddress(t *testing.T) {
	var testCases = []struct {
		name   string
		value  string
		expect string
	}{
		{
			name:   "first",
			expect: "APfiG5imQgn8dzTTfaDfqHnxo3QDUkF69A",
			value:  "5696acd07f0927fd5f01946828638c9e2c90c5dc",
		},

		{
			name:   "second",
			expect: "AK2nJJpJr6o664CWJKi1QRXjqeic2zRp8y",
			value:  "23ba2703c53263e8d6e522dc32203339dcd8eee9",
		},
	}

	for i := range testCases {
		tt := testCases[i]
		t.Run(tt.name, func(t *testing.T) {
			actual, err := ReversedScriptHashToAddress(tt.value)
			require.NoError(t, err)
			require.Equal(t, tt.expect, actual)
			require.NoError(t, IsAddress(actual))
		})
	}
}

func TestReverseBytes(t *testing.T) {
	var testCases = []struct {
		name   string
		value  []byte
		expect []byte
	}{
		{name: "empty"},
		{
			name:   "single byte",
			expect: []byte{0x1},
			value:  []byte{0x1},
		},

		{
			name:   "two bytes",
			expect: []byte{0x2, 0x1},
			value:  []byte{0x1, 0x2},
		},

		{
			name:   "three bytes",
			expect: []byte{0x3, 0x2, 0x1},
			value:  []byte{0x1, 0x2, 0x3},
		},

		{
			name:   "five bytes",
			expect: []byte{0x5, 0x4, 0x3, 0x2, 0x1},
			value:  []byte{0x1, 0x2, 0x3, 0x4, 0x5},
		},

		{
			name:   "eight bytes",
			expect: []byte{0x8, 0x7, 0x6, 0x5, 0x4, 0x3, 0x2, 0x1},
			value:  []byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8},
		},
	}

	for i := range testCases {
		tt := testCases[i]
		t.Run(tt.name, func(t *testing.T) {
			actual := ReverseBytes(tt.value)
			require.Equal(t, tt.expect, actual)
		})
	}
}
