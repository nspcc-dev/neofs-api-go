package object_test

import (
	"crypto/rand"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg/object"
	"github.com/stretchr/testify/require"
)

func TestSplitInfo(t *testing.T) {
	s := object.NewSplitInfo()
	splitID := object.NewSplitID()
	lastPart := generateID()
	link := generateID()

	s.SetSplitID(splitID)
	require.Equal(t, splitID, s.SplitID())

	s.SetLastPart(lastPart)
	require.Equal(t, lastPart, s.LastPart())

	s.SetLink(link)
	require.Equal(t, link, s.Link())

	t.Run("to and from v2", func(t *testing.T) {
		v2 := s.ToV2()
		newS := object.NewSplitInfoFromV2(v2)

		require.Equal(t, s, newS)
	})

	t.Run("marshal and unmarshal", func(t *testing.T) {
		data, err := s.Marshal()
		require.NoError(t, err)

		newS := object.NewSplitInfo()

		err = newS.Unmarshal(data)
		require.NoError(t, err)
		require.Equal(t, s, newS)
	})
}

func generateID() *object.ID {
	var buf [32]byte
	_, _ = rand.Read(buf[:])

	id := object.NewID()
	id.SetSHA256(buf)

	return id
}
