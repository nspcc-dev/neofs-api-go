package container_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	containerV2 "github.com/nspcc-dev/neofs-api-go/v2/container"
	"github.com/stretchr/testify/require"
)

func TestNewVerifiedFromV2(t *testing.T) {
	cnrV2 := new(containerV2.Container)

	errAssert := func() {
		_, err := container.NewVerifiedFromV2(cnrV2)
		require.Error(t, err)
	}

	// set unsupported version
	v := pkg.SDKVersion()
	v.SetMajor(0)
	require.Error(t, pkg.IsSupportedVersion(v))
	cnrV2.SetVersion(v.ToV2())

	errAssert()

	// set supported version
	v.SetMajor(2)
	require.NoError(t, pkg.IsSupportedVersion(v))
	cnrV2.SetVersion(v.ToV2())

	errAssert()

	// set invalid nonce
	nonce := []byte{1, 2, 3}
	cnrV2.SetNonce(nonce)

	errAssert()

	// set valid nonce
	uid := uuid.New()
	data, _ := uid.MarshalBinary()
	cnrV2.SetNonce(data)

	_, err := container.NewVerifiedFromV2(cnrV2)
	require.NoError(t, err)
}
