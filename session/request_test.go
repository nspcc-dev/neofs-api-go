package session

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateRequestGettersSetters(t *testing.T) {
	t.Run("owner ID", func(t *testing.T) {
		id := OwnerID{1, 2, 3}
		m := new(CreateRequest)

		m.SetOwnerID(id)

		require.Equal(t, id, m.GetOwnerID())
	})

	t.Run("lifetime", func(t *testing.T) {
		e1, e2 := uint64(3), uint64(4)
		m := new(CreateRequest)

		m.SetCreationEpoch(e1)
		m.SetExpirationEpoch(e2)

		require.Equal(t, e1, m.CreationEpoch())
		require.Equal(t, e2, m.ExpirationEpoch())
	})
}

func TestCreateRequest_SignedData(t *testing.T) {
	var (
		id = OwnerID{1, 2, 3}
		e1 = uint64(1)
		e2 = uint64(2)
	)

	// create new message
	m := new(CreateRequest)

	// fill the fields
	m.SetOwnerID(id)
	m.SetCreationEpoch(e1)
	m.SetExpirationEpoch(e2)

	// calculate initial signed data
	d, err := m.SignedData()
	require.NoError(t, err)

	items := []struct {
		change func()
		reset  func()
	}{
		{ // OwnerID
			change: func() {
				id2 := id
				id2[0]++
				m.SetOwnerID(id2)
			},
			reset: func() {
				m.SetOwnerID(id)
			},
		},
		{ // CreationEpoch
			change: func() {
				m.SetCreationEpoch(e1 + 1)
			},
			reset: func() {
				m.SetCreationEpoch(e1)
			},
		},
		{ // ExpirationEpoch
			change: func() {
				m.SetExpirationEpoch(e2 + 1)
			},
			reset: func() {
				m.SetExpirationEpoch(e2)
			},
		},
	}

	for _, item := range items {
		item.change()

		d2, err := m.SignedData()
		require.NoError(t, err)

		require.NotEqual(t, d, d2)

		item.reset()
	}
}
