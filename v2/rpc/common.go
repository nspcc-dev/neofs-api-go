package rpc

import (
	"context"

	protoclient "github.com/nspcc-dev/neofs-api-go/rpc/client"
	"github.com/nspcc-dev/neofs-api-go/rpc/common"
	"github.com/nspcc-dev/neofs-api-go/rpc/message"
)

const serviceNamePrefix = "neo.fs.v2."

func setMethodInfo(i *common.CallMethodInfo, svc, mtd string, clientStream, serverStream bool) {
	i.SetServiceName(svc)
	i.SetMethodName(mtd)

	if clientStream {
		i.SetClientStream()
	}

	if serverStream {
		i.SetServerStream()
	}
}

func sendUnaryRPC(ctx context.Context, cli protoclient.Client, req, resp message.Message, svc, mtd string) error {
	var info common.CallMethodInfo

	setMethodInfo(&info, svc, mtd, false, false)

	var uPrm protoclient.SendUnaryPrm

	uPrm.SetCallMethodInfo(info)
	uPrm.SetMessages(req, resp)

	return protoclient.SendUnary(ctx, cli, uPrm, nil)
}

func openServerStream(ctx context.Context, cli protoclient.Client, req message.Message, svc, mtd string, set func(protoclient.MessageReaderCloser)) error {
	var info common.CallMethodInfo

	setMethodInfo(&info, svc, mtd, false, true)

	var ssPrm protoclient.OpenServerStreamPrm

	ssPrm.SetCallMethodInfo(info)
	ssPrm.SetRequest(req)

	var ssRes protoclient.OpenServerStreamRes

	err := protoclient.OpenServerStream(ctx, cli, ssPrm, &ssRes)
	if err != nil {
		return err
	}

	set(ssRes.Messager())

	return nil
}
