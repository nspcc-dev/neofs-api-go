package object

import (
	"testing"

	"github.com/google/uuid"
	"github.com/nspcc-dev/neofs-api-go/container"
	"github.com/nspcc-dev/neofs-api-go/refs"
	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
)

func TestObject_Verify(t *testing.T) {
	key := test.DecodeKey(0)
	sessionkey := test.DecodeKey(1)

	payload := make([]byte, 1024*1024)

	cnr, err := container.NewTestContainer()
	require.NoError(t, err)

	cid, err := cnr.ID()
	require.NoError(t, err)

	id, err := uuid.NewRandom()
	uid := refs.UUID(id)
	require.NoError(t, err)

	obj := &Object{
		SystemHeader: SystemHeader{
			ID:      uid,
			CID:     cid,
			OwnerID: refs.OwnerID([refs.OwnerIDSize]byte{}),
		},
		Headers: []Header{
			{
				Value: &Header_UserHeader{
					UserHeader: &UserHeader{
						Key:   "Profession",
						Value: "Developer",
					},
				},
			},
			{
				Value: &Header_UserHeader{
					UserHeader: &UserHeader{
						Key:   "Language",
						Value: "GO",
					},
				},
			},
		},
	}
	obj.SetPayload(payload)
	obj.SetHeader(&Header{Value: &Header_PayloadChecksum{[]byte("incorrect checksum")}})

	t.Run("error no integrity header and pubkey", func(t *testing.T) {
		err = obj.Verify()
		require.EqualError(t, err, ErrHeaderNotFound.Error())
	})

	badHeaderChecksum := []byte("incorrect checksum")
	signature, err := crypto.Sign(sessionkey, badHeaderChecksum)
	require.NoError(t, err)
	ih := &IntegrityHeader{
		HeadersChecksum:   badHeaderChecksum,
		ChecksumSignature: signature,
	}
	obj.SetHeader(&Header{Value: &Header_Integrity{ih}})

	t.Run("error no validation header", func(t *testing.T) {
		err = obj.Verify()
		require.EqualError(t, err, ErrHeaderNotFound.Error())
	})

	dataPK := crypto.MarshalPublicKey(&sessionkey.PublicKey)
	signature, err = crypto.Sign(key, dataPK)
	tok := new(Token)
	tok.SetSignature(signature)
	tok.SetSessionKey(dataPK)
	obj.AddHeader(&Header{Value: &Header_Token{Token: tok}})

	// validation header is not last
	t.Run("error validation header is not last", func(t *testing.T) {
		err = obj.Verify()
		require.EqualError(t, err, ErrHeaderNotFound.Error())
	})

	obj.Headers = obj.Headers[:len(obj.Headers)-2]
	obj.AddHeader(&Header{Value: &Header_Token{Token: tok}})
	obj.SetHeader(&Header{Value: &Header_Integrity{ih}})

	t.Run("error invalid header checksum", func(t *testing.T) {
		err = obj.Verify()
		require.EqualError(t, err, ErrVerifyHeader.Error())
	})

	obj.Headers = obj.Headers[:len(obj.Headers)-1]
	genIH, err := CreateIntegrityHeader(obj, sessionkey)
	require.NoError(t, err)
	obj.SetHeader(genIH)

	t.Run("error invalid payload checksum", func(t *testing.T) {
		err = obj.Verify()
		require.EqualError(t, err, ErrVerifyPayload.Error())
	})

	obj.SetHeader(&Header{Value: &Header_PayloadChecksum{obj.PayloadChecksum()}})

	obj.Headers = obj.Headers[:len(obj.Headers)-1]
	genIH, err = CreateIntegrityHeader(obj, sessionkey)
	require.NoError(t, err)
	obj.SetHeader(genIH)

	t.Run("correct with tok", func(t *testing.T) {
		err = obj.Verify()
		require.NoError(t, err)
	})

	pkh := Header{Value: &Header_PublicKey{&PublicKey{
		Value: crypto.MarshalPublicKey(&key.PublicKey),
	}}}
	// replace tok with pkh
	obj.Headers[len(obj.Headers)-2] = pkh
	// re-sign object
	obj.Sign(sessionkey)

	t.Run("incorrect with bad public key", func(t *testing.T) {
		err = obj.Verify()
		require.Error(t, err)
	})

	obj.SetHeader(&Header{Value: &Header_PublicKey{&PublicKey{
		Value: dataPK,
	}}})
	obj.Sign(sessionkey)

	t.Run("correct with good public key", func(t *testing.T) {
		err = obj.Verify()
		require.NoError(t, err)
	})

}
