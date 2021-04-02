package reputation_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/reputation"
	reputationtest "github.com/nspcc-dev/neofs-api-go/pkg/reputation/test"
	reputationtestV2 "github.com/nspcc-dev/neofs-api-go/v2/reputation/test"
	"github.com/stretchr/testify/require"
)

func TestTrust(t *testing.T) {
	trust := reputation.NewTrust()

	id := reputationtest.GeneratePeerID()
	trust.SetPeer(id)
	require.Equal(t, id, trust.Peer())

	val := 1.5
	trust.SetValue(val)
	require.Equal(t, val, trust.Value())

	t.Run("binary encoding", func(t *testing.T) {
		trust := reputationtest.GenerateTrust()
		data, err := trust.Marshal()
		require.NoError(t, err)

		trust2 := reputation.NewTrust()
		require.NoError(t, trust2.Unmarshal(data))
		require.Equal(t, trust, trust2)
	})

	t.Run("JSON encoding", func(t *testing.T) {
		trust := reputationtest.GenerateTrust()
		data, err := trust.MarshalJSON()
		require.NoError(t, err)

		trust2 := reputation.NewTrust()
		require.NoError(t, trust2.UnmarshalJSON(data))
		require.Equal(t, trust, trust2)
	})
}

func TestGlobalTrust(t *testing.T) {
	t.Run("v2", func(t *testing.T) {
		gtV2 := reputationtestV2.GenerateGlobalTrust(false)

		gt := reputation.GlobalTrustFromV2(gtV2)

		require.Equal(t, gtV2, gt.ToV2())
	})

	t.Run("getters+setters", func(t *testing.T) {
		gt := reputation.NewGlobalTrust()

		require.Equal(t, pkg.SDKVersion(), gt.Version())
		require.Nil(t, gt.Manager())
		require.Nil(t, gt.Trust())

		version := pkg.NewVersion()
		version.SetMajor(13)
		version.SetMinor(31)
		gt.SetVersion(version)
		require.Equal(t, version, gt.Version())

		mngr := reputationtest.GeneratePeerID()
		gt.SetManager(mngr)
		require.Equal(t, mngr, gt.Manager())

		trust := reputationtest.GenerateTrust()
		gt.SetTrust(trust)
		require.Equal(t, trust, gt.Trust())
	})

	t.Run("sign+verify", func(t *testing.T) {
		gt := reputationtest.GenerateSignedGlobalTrust(t)

		err := gt.VerifySignature()
		require.NoError(t, err)
	})

	t.Run("encoding", func(t *testing.T) {
		t.Run("binary", func(t *testing.T) {
			gt := reputationtest.GenerateSignedGlobalTrust(t)

			data, err := gt.Marshal()
			require.NoError(t, err)

			gt2 := reputation.NewGlobalTrust()
			require.NoError(t, gt2.Unmarshal(data))
			require.Equal(t, gt, gt2)
		})

		t.Run("JSON", func(t *testing.T) {
			gt := reputationtest.GenerateSignedGlobalTrust(t)
			data, err := gt.MarshalJSON()
			require.NoError(t, err)

			gt2 := reputation.NewGlobalTrust()
			require.NoError(t, gt2.UnmarshalJSON(data))
			require.Equal(t, gt, gt2)
		})
	})
}
