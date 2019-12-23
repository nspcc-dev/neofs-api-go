package object

import (
	"crypto/rand"
	"crypto/sha256"
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
	obj.SetHeader(&Header{Value: &Header_PayloadChecksum{
		PayloadChecksum: &PayloadChecksum{
			ChecksumList: []PayloadChecksum_RangeChecksum{
				{
					Range: Range{
						Offset: 0,
						Length: uint64(len(payload)),
					},
					Value: []byte("incorrect checksum"),
				},
			},
		},
	}})

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

	checksum, err := obj.PayloadRangeChecksum(Range{
		Offset: 0,
		Length: uint64(len(obj.Payload)),
	})
	require.NoError(t, err)

	obj.SetHeader(&Header{Value: &Header_PayloadChecksum{
		PayloadChecksum: &PayloadChecksum{
			ChecksumList: []PayloadChecksum_RangeChecksum{
				{
					Range: Range{
						Offset: 0,
						Length: uint64(len(payload)),
					},
					Value: checksum,
				},
			},
		},
	}})
	require.NoError(t, obj.Sign(sessionkey))

	t.Run("correct", func(t *testing.T) {
		err = obj.Verify()
		require.NoError(t, err)
	})
}

func TestObject_PayloadRangeChecksum(t *testing.T) {
	payloadSize := uint64(10)

	obj := &Object{Payload: testData(t, payloadSize)}

	t.Run("out-of-bounds range", func(t *testing.T) {
		_, err := obj.PayloadRangeChecksum(Range{
			Offset: 0,
			Length: payloadSize + 1,
		})
		require.EqualError(t, err, ErrRangeOutOfBounds.Error())
	})

	t.Run("correct range", func(t *testing.T) {
		r := Range{Offset: 0, Length: payloadSize / 2}

		checksum, err := obj.PayloadRangeChecksum(r)
		require.NoError(t, err)

		calculatedChecksum := sha256.Sum256(obj.Payload[r.Offset : r.Offset+r.Length])

		require.Equal(t, checksum, calculatedChecksum[:])
	})
}

func TestObject_VerifyPayload(t *testing.T) {
	payloadSize := uint64(10)

	obj := &Object{Payload: testData(t, payloadSize)}

	t.Run("missing checksum header", func(t *testing.T) {
		require.EqualError(t, obj.VerifyPayload(), ErrHeaderNotFound.Error())
	})

	checksumList := []PayloadChecksum_RangeChecksum{{}}

	payloadChecksum := &PayloadChecksum{ChecksumList: checksumList}

	obj.SetHeader(&Header{Value: &Header_PayloadChecksum{PayloadChecksum: payloadChecksum}})

	t.Run("invalid order", func(t *testing.T) {
		payloadChecksum.ChecksumList[0].Range.Offset = 1

		require.EqualError(t, obj.VerifyPayload(), ErrInvalidRangeOrder.Error())
	})

	payloadChecksum.ChecksumList[0].Range.Offset = 0

	t.Run("empty range", func(t *testing.T) {
		require.EqualError(t, obj.VerifyPayload(), ErrEmptyPayloadRange.Error())
	})

	t.Run("range out of bounds", func(t *testing.T) {
		payloadChecksum.ChecksumList[0].Range.Length = payloadSize + 1

		require.EqualError(t, obj.VerifyPayload(), ErrRangeOutOfBounds.Error())
	})

	payloadChecksum.ChecksumList[0].Range.Length = payloadSize / 2

	t.Run("incorrect value", func(t *testing.T) {
		require.EqualError(t, obj.VerifyPayload(), ErrVerifyPayload.Error())
	})

	cs, err := obj.PayloadRangeChecksum(payloadChecksum.ChecksumList[0].Range)
	require.NoError(t, err)

	payloadChecksum.ChecksumList[0].Value = cs

	t.Run("incomplete coverage", func(t *testing.T) {
		require.EqualError(t, obj.VerifyPayload(), ErrIncompleteRangeCoverage.Error())
	})

	tailRange := Range{Offset: payloadSize - payloadChecksum.ChecksumList[0].Range.Length}
	tailRange.Length = payloadSize - tailRange.Offset

	cs, err = obj.PayloadRangeChecksum(tailRange)
	require.NoError(t, err)

	payloadChecksum.ChecksumList = append(payloadChecksum.ChecksumList, PayloadChecksum_RangeChecksum{
		Range: tailRange,
		Value: cs,
	})

	t.Run("correct", func(t *testing.T) {
		require.NoError(t, obj.VerifyPayload())
	})
}

func testData(t *testing.T, size uint64) []byte {
	data := make([]byte, size)
	_, err := rand.Read(data)
	require.NoError(t, err)
	return data
}
