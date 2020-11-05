package client_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/nspcc-dev/neofs-api-go/pkg/client"
	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	"github.com/nspcc-dev/neofs-api-go/pkg/netmap"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestExample(t *testing.T) {
	t.Skip()
	target := "s01.localtest.nspcc.ru:50501"
	key := test.DecodeKey(-1)

	// create client from address
	cli, err := client.New(key, client.WithAddress(target))
	require.NoError(t, err)

	// ask for balance
	resp, err := cli.GetSelfBalance(context.Background())
	require.NoError(t, err)

	fmt.Println(resp.GetValue(), resp.GetPrecision())

	// create client from grpc connection
	conn, err := grpc.DialContext(context.Background(), target, grpc.WithBlock(), grpc.WithInsecure())
	require.NoError(t, err)

	cli, err = client.New(key, client.WithGRPCConnection(conn))
	require.NoError(t, err)

	replica := new(netmap.Replica)
	replica.SetCount(2)
	replica.SetSelector("*")

	policy := new(netmap.PlacementPolicy)
	policy.SetContainerBackupFactor(2)
	policy.SetReplicas(replica)

	// this container has random nonce and it does not set owner id
	cnr := container.New(
		container.WithAttribute("CreatedAt", time.Now().String()),
		container.WithPolicy(policy),
		container.WithReadOnlyBasicACL(),
	)
	require.NoError(t, err)

	// here container will have owner id from client key, and it will be signed
	containerID, err := cli.PutContainer(context.Background(), cnr, client.WithTTL(10))
	require.NoError(t, err)

	fmt.Println(containerID)

	list, err := cli.ListSelfContainers(context.Background())
	require.NoError(t, err)

	for i := range list {
		fmt.Println("found container:", list[i])
	}
}
