package object

import (
	"testing"

	"github.com/google/uuid"
	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/nspcc-dev/neofs-proto/container"
	"github.com/nspcc-dev/neofs-proto/refs"
	"github.com/nspcc-dev/neofs-proto/session"
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

	t.Run("error no integrity header", func(t *testing.T) {
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
	vh := &session.VerificationHeader{
		PublicKey:    dataPK,
		KeySignature: signature,
	}
	obj.SetVerificationHeader(vh)

	t.Run("error invalid header checksum", func(t *testing.T) {
		err = obj.Verify()
		require.EqualError(t, err, ErrVerifyHeader.Error())
	})

	require.NoError(t, obj.Sign(sessionkey))

	t.Run("error invalid payload checksum", func(t *testing.T) {
		err = obj.Verify()
		require.EqualError(t, err, ErrVerifyPayload.Error())
	})

	obj.SetHeader(&Header{Value: &Header_PayloadChecksum{obj.PayloadChecksum()}})
	require.NoError(t, obj.Sign(sessionkey))

	t.Run("correct", func(t *testing.T) {
		err = obj.Verify()
		require.NoError(t, err)
	})
}
