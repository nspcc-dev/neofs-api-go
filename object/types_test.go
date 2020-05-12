package object

import (
	"bytes"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/refs"
	"github.com/nspcc-dev/neofs-api-go/service"
	"github.com/nspcc-dev/neofs-api-go/storagegroup"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
)

func TestStringify(t *testing.T) {
	res := `
Object:
	SystemHeader:
		- ID=7e0b9c6c-aabc-4985-949e-2680e577b48b
		- CID=11111111111111111111111111111111
		- OwnerID=ALYeYC41emF6MrmUMc4a8obEPdgFhq9ran
		- Version=1
		- PayloadLength=1
		- CreatedAt={UnixTime=1 Epoch=1}
	UserHeaders:
		- Type=Link
		  Value={Type=Child ID=7e0b9c6c-aabc-4985-949e-2680e577b48b}
		- Type=Redirect
		  Value={CID=11111111111111111111111111111111 OID=7e0b9c6c-aabc-4985-949e-2680e577b48b}
		- Type=UserHeader
		  Value={Key=test_key Val=test_value}
		- Type=Transform
		  Value=Split
		- Type=Tombstone
		  Value=MARKED
		- Type=Token
		  Value={ID=7e0b9c6c-aabc-4985-949e-2680e577b48b OwnerID=ALYeYC41emF6MrmUMc4a8obEPdgFhq9ran Verb=Search Address=11111111111111111111111111111111/7e0b9c6c-aabc-4985-949e-2680e577b48b Created=1 ValidUntil=2 SessionKey=010203040506 Signature=010203040506}
		- Type=HomoHash
		  Value=1111111111111111111111111111111111111111111111111111111111111111
		- Type=PayloadChecksum
		  Value=[1 2 3 4 5 6]
		- Type=Integrity
		  Value={Checksum=010203040506 Signature=010203040506}
		- Type=StorageGroup
		  Value={DataSize=5 Hash=31313131313131313131313131313131313131313131313131313131313131313131313131313131313131313131313131313131313131313131313131313131 Lifetime={Unit=UnixTime Value=555}}
		- Type=PublicKey
		  Value=[1 2 3 4 5 6]
	Payload: []byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7}
`

	key := test.DecodeKey(0)

	uid, err := refs.NewOwnerID(&key.PublicKey)
	require.NoError(t, err)

	var oid refs.UUID

	require.NoError(t, oid.Parse("7e0b9c6c-aabc-4985-949e-2680e577b48b"))

	obj := &Object{
		SystemHeader: SystemHeader{
			Version:       1,
			PayloadLength: 1,
			ID:            oid,
			OwnerID:       uid,
			CID:           CID{},
			CreatedAt: CreationPoint{
				UnixTime: 1,
				Epoch:    1,
			},
		},
		Payload: []byte{1, 2, 3, 4, 5, 6, 7},
	}

	// *Header_Link
	obj.Headers = append(obj.Headers, Header{
		Value: &Header_Link{
			Link: &Link{ID: oid, Type: Link_Child},
		},
	})

	// *Header_Redirect
	obj.Headers = append(obj.Headers, Header{
		Value: &Header_Redirect{
			Redirect: &Address{ObjectID: oid, CID: CID{}},
		},
	})

	// *Header_UserHeader
	obj.Headers = append(obj.Headers, Header{
		Value: &Header_UserHeader{
			UserHeader: &UserHeader{
				Key:   "test_key",
				Value: "test_value",
			},
		},
	})

	// *Header_Transform
	obj.Headers = append(obj.Headers, Header{
		Value: &Header_Transform{
			Transform: &Transform{
				Type: Transform_Split,
			},
		},
	})

	// *Header_Tombstone
	obj.Headers = append(obj.Headers, Header{
		Value: &Header_Tombstone{
			Tombstone: &Tombstone{},
		},
	})

	token := new(Token)
	token.SetID(oid)
	token.SetOwnerID(uid)
	token.SetVerb(service.Token_Info_Search)
	token.SetAddress(Address{ObjectID: oid, CID: refs.CID{}})
	token.SetCreationEpoch(1)
	token.SetExpirationEpoch(2)
	token.SetSessionKey([]byte{1, 2, 3, 4, 5, 6})
	token.SetSignature([]byte{1, 2, 3, 4, 5, 6})

	// *Header_Token
	obj.Headers = append(obj.Headers, Header{
		Value: &Header_Token{
			Token: token,
		},
	})

	// *Header_HomoHash
	obj.Headers = append(obj.Headers, Header{
		Value: &Header_HomoHash{
			HomoHash: Hash{},
		},
	})

	// *Header_PayloadChecksum
	obj.Headers = append(obj.Headers, Header{
		Value: &Header_PayloadChecksum{
			PayloadChecksum: []byte{1, 2, 3, 4, 5, 6},
		},
	})

	// *Header_Integrity
	obj.Headers = append(obj.Headers, Header{
		Value: &Header_Integrity{
			Integrity: &IntegrityHeader{
				HeadersChecksum:   []byte{1, 2, 3, 4, 5, 6},
				ChecksumSignature: []byte{1, 2, 3, 4, 5, 6},
			},
		},
	})

	// *Header_StorageGroup
	obj.Headers = append(obj.Headers, Header{
		Value: &Header_StorageGroup{
			StorageGroup: &storagegroup.StorageGroup{
				ValidationDataSize: 5,
				ValidationHash:     storagegroup.Hash{},
				Lifetime: &storagegroup.StorageGroup_Lifetime{
					Unit:  storagegroup.StorageGroup_Lifetime_UnixTime,
					Value: 555,
				},
			},
		},
	})

	// *Header_PublicKey
	obj.Headers = append(obj.Headers, Header{
		Value: &Header_PublicKey{
			PublicKey: &PublicKey{Value: []byte{1, 2, 3, 4, 5, 6}},
		},
	})

	buf := new(bytes.Buffer)

	require.NoError(t, Stringify(buf, obj))
	require.Equal(t, res, buf.String())
}

func TestObject_Copy(t *testing.T) {
	t.Run("token header", func(t *testing.T) {
		token := new(Token)
		token.SetID(service.TokenID{1, 2, 3})

		obj := new(Object)

		obj.AddHeader(&Header{
			Value: &Header_Token{
				Token: token,
			},
		})

		cp := obj.Copy()

		_, h := cp.LastHeader(HeaderType(TokenHdr))
		require.NotNil(t, h)
		require.Equal(t, token, h.GetValue().(*Header_Token).Token)
	})
}
