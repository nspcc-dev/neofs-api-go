package accounting

import (
	"io/ioutil"
	"testing"

	"github.com/mr-tron/base58"
	"github.com/nspcc-dev/neofs-api-go/chain"
	"github.com/nspcc-dev/neofs-api-go/decimal"
	"github.com/nspcc-dev/neofs-api-go/refs"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
)

func TestCheque(t *testing.T) {
	t.Run("new/valid", func(t *testing.T) {
		id, err := NewChequeID()
		require.NoError(t, err)
		require.True(t, id.Valid())

		d := make([]byte, chain.AddressLength+1)

		// expected size + 1 byte
		str := base58.Encode(d)
		require.False(t, ChequeID(str).Valid())

		// expected size - 1 byte
		str = base58.Encode(d[:len(d)-2])
		require.False(t, ChequeID(str).Valid())

		// wrong encoding
		d = d[:len(d)-1] // normal size
		require.False(t, ChequeID(string(d)).Valid())
	})

	t.Run("marshal/unmarshal", func(t *testing.T) {
		b2 := new(Cheque)

		key1 := test.DecodeKey(0)
		key2 := test.DecodeKey(1)

		id, err := NewChequeID()
		require.NoError(t, err)

		owner, err := refs.NewOwnerID(&key1.PublicKey)
		require.NoError(t, err)

		b1 := &Cheque{
			ID:     id,
			Owner:  owner,
			Height: 100,
			Amount: decimal.NewGAS(100),
		}

		require.NoError(t, b1.Sign(key1))
		require.NoError(t, b1.Sign(key2))

		data, err := b1.MarshalBinary()
		require.NoError(t, err)

		require.Len(t, data, b1.Size())
		require.NoError(t, b2.UnmarshalBinary(data))
		require.Equal(t, b1, b2)

		require.NoError(t, b1.Verify())
		require.NoError(t, b2.Verify())
	})

	t.Run("example from SC", func(t *testing.T) {
		pathToCheque := "fixtures/cheque_data"
		expect, err := ioutil.ReadFile(pathToCheque)
		require.NoError(t, err)

		var cheque Cheque
		require.NoError(t, cheque.UnmarshalBinary(expect))

		actual, err := cheque.MarshalBinary()
		require.NoError(t, err)

		require.Equal(t, expect, actual)

		require.NoError(t, cheque.Verify())
		require.Equal(t, cheque.Height, uint64(7777))
		require.Equal(t, cheque.Amount, decimal.NewGAS(42))
	})
}

func TestBalanceRequest_SetOwnerID(t *testing.T) {
	ownerID := OwnerID{1, 2, 3}
	m := new(BalanceRequest)

	m.SetOwnerID(ownerID)

	require.Equal(t, ownerID, m.GetOwnerID())
}

func TestGetRequestGettersSetters(t *testing.T) {
	t.Run("id", func(t *testing.T) {
		id := ChequeID("test id")
		m := new(GetRequest)

		m.SetID(id)

		require.Equal(t, id, m.GetID())
	})

	t.Run("owner", func(t *testing.T) {
		id := OwnerID{1, 2, 3}
		m := new(GetRequest)

		m.SetOwnerID(id)

		require.Equal(t, id, m.GetOwnerID())
	})
}

func TestPutRequestGettersSetters(t *testing.T) {
	t.Run("owner", func(t *testing.T) {
		id := OwnerID{1, 2, 3}
		m := new(PutRequest)

		m.SetOwnerID(id)

		require.Equal(t, id, m.GetOwnerID())
	})

	t.Run("message ID", func(t *testing.T) {
		id, err := refs.NewUUID()
		require.NoError(t, err)

		m := new(PutRequest)
		m.SetMessageID(id)

		require.Equal(t, id, m.GetMessageID())
	})

	t.Run("amount", func(t *testing.T) {
		amount := decimal.New(1)
		m := new(PutRequest)

		m.SetAmount(amount)

		require.Equal(t, amount, m.GetAmount())
	})

	t.Run("height", func(t *testing.T) {
		h := uint64(3)
		m := new(PutRequest)

		m.SetHeight(h)

		require.Equal(t, h, m.GetHeight())
	})
}

func TestListRequestGettersSetters(t *testing.T) {
	ownerID := OwnerID{1, 2, 3}
	m := new(ListRequest)

	m.SetOwnerID(ownerID)

	require.Equal(t, ownerID, m.GetOwnerID())
}

func TestDeleteRequestGettersSetters(t *testing.T) {
	t.Run("id", func(t *testing.T) {
		id := ChequeID("test id")
		m := new(DeleteRequest)

		m.SetID(id)

		require.Equal(t, id, m.GetID())
	})

	t.Run("owner", func(t *testing.T) {
		id := OwnerID{1, 2, 3}
		m := new(DeleteRequest)

		m.SetOwnerID(id)

		require.Equal(t, id, m.GetOwnerID())
	})

	t.Run("message ID", func(t *testing.T) {
		id, err := refs.NewUUID()
		require.NoError(t, err)

		m := new(DeleteRequest)
		m.SetMessageID(id)

		require.Equal(t, id, m.GetMessageID())
	})
}
