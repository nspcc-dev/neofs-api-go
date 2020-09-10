package client

import (
	"context"

	"github.com/nspcc-dev/neofs-api-go/pkg/acl/eacl"
	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	"github.com/nspcc-dev/neofs-api-go/util/signature"
	"github.com/nspcc-dev/neofs-api-go/v2/client"
	v2container "github.com/nspcc-dev/neofs-api-go/v2/container"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	v2signature "github.com/nspcc-dev/neofs-api-go/v2/signature"
	"github.com/pkg/errors"
)

type delContainerSignWrapper struct {
	body *v2container.DeleteRequestBody
}

func (c delContainerSignWrapper) ReadSignedData(bytes []byte) ([]byte, error) {
	return c.body.GetContainerID().GetValue(), nil
}

func (c delContainerSignWrapper) SignedDataSize() int {
	return len(c.body.GetContainerID().GetValue())
}

func (c Client) PutContainer(ctx context.Context, cnr *container.Container, opts ...CallOption) (*container.ID, error) {
	switch c.remoteNode.Version.GetMajor() {
	case 2:
		return c.putContainerV2(ctx, cnr, opts...)
	default:
		return nil, unsupportedProtocolErr
	}
}

func (c Client) GetContainer(ctx context.Context, id *container.ID, opts ...CallOption) (*container.Container, error) {
	switch c.remoteNode.Version.GetMajor() {
	case 2:
		return c.getContainerV2(ctx, id, opts...)
	default:
		return nil, unsupportedProtocolErr
	}
}

func (c Client) ListContainers(ctx context.Context, owner *owner.ID, opts ...CallOption) ([]*container.ID, error) {
	switch c.remoteNode.Version.GetMajor() {
	case 2:
		return c.listContainerV2(ctx, owner, opts...)
	default:
		return nil, unsupportedProtocolErr
	}
}

func (c Client) ListSelfContainers(ctx context.Context, opts ...CallOption) ([]*container.ID, error) {
	w, err := owner.NEO3WalletFromPublicKey(&c.key.PublicKey)
	if err != nil {
		return nil, err
	}

	ownerID := new(owner.ID)
	ownerID.SetNeo3Wallet(w)

	return c.ListContainers(ctx, ownerID, opts...)
}

func (c Client) DeleteContainer(ctx context.Context, id *container.ID, opts ...CallOption) error {
	switch c.remoteNode.Version.GetMajor() {
	case 2:
		return c.delContainerV2(ctx, id, opts...)
	default:
		return unsupportedProtocolErr
	}
}

func (c Client) GetEACL(ctx context.Context, id *container.ID, opts ...CallOption) (*eacl.Table, error) {
	switch c.remoteNode.Version.GetMajor() {
	case 2:
		return c.getEACLV2(ctx, id, opts...)
	default:
		return nil, unsupportedProtocolErr
	}
}

func (c Client) SetEACL(ctx context.Context, eacl *eacl.Table, opts ...CallOption) error {
	switch c.remoteNode.Version.GetMajor() {
	case 2:
		return c.setEACLV2(ctx, eacl, opts...)
	default:
		return unsupportedProtocolErr
	}
}

func (c Client) putContainerV2(ctx context.Context, cnr *container.Container, opts ...CallOption) (*container.ID, error) {
	// apply all available options
	callOptions := defaultCallOptions()
	for i := range opts {
		opts[i].apply(&callOptions)
	}

	// set transport version
	cnr.SetVersion(c.remoteNode.Version.ToV2())

	// if container owner is not set, then use client key as owner
	if cnr.GetOwnerID() == nil {
		w, err := owner.NEO3WalletFromPublicKey(&c.key.PublicKey)
		if err != nil {
			return nil, err
		}

		ownerID := new(owner.ID)
		ownerID.SetNeo3Wallet(w)

		cnr.SetOwnerID(ownerID.ToV2())
	}

	reqBody := new(v2container.PutRequestBody)
	reqBody.SetContainer(&cnr.Container)

	// sign container
	signWrapper := v2signature.StableMarshalerWrapper{SM: reqBody.GetContainer()}
	err := signature.SignDataWithHandler(c.key, signWrapper, func(key []byte, sig []byte) {
		containerSignature := new(refs.Signature)
		containerSignature.SetKey(key)
		containerSignature.SetSign(sig)
		reqBody.SetSignature(containerSignature)
	}, signature.SignWithRFC6979())
	if err != nil {
		return nil, err
	}

	req := new(v2container.PutRequest)
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err = v2signature.SignServiceMessage(c.key, req)
	if err != nil {
		return nil, err
	}

	switch c.remoteNode.Protocol {
	case GRPC:
		cli, err := v2ContainerClientFromOptions(c.opts)
		if err != nil {
			return nil, errors.Wrap(err, "can't create grpc client")
		}

		resp, err := cli.Put(ctx, req)
		if err != nil {
			return nil, errors.Wrap(err, "transport error")
		}

		err = v2signature.VerifyServiceMessage(resp)
		if err != nil {
			return nil, errors.Wrap(err, "can't verify response message")
		}

		return container.NewIDFromV2(resp.GetBody().GetContainerID()), nil
	default:
		return nil, unsupportedProtocolErr
	}
}

func (c Client) getContainerV2(ctx context.Context, id *container.ID, opts ...CallOption) (*container.Container, error) {
	// apply all available options
	callOptions := defaultCallOptions()
	for i := range opts {
		opts[i].apply(&callOptions)
	}

	reqBody := new(v2container.GetRequestBody)
	reqBody.SetContainerID(id.ToV2())

	req := new(v2container.GetRequest)
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err := v2signature.SignServiceMessage(c.key, req)
	if err != nil {
		return nil, err
	}

	switch c.remoteNode.Protocol {
	case GRPC:
		cli, err := v2ContainerClientFromOptions(c.opts)
		if err != nil {
			return nil, errors.Wrap(err, "can't create grpc client")
		}

		resp, err := cli.Get(ctx, req)
		if err != nil {
			return nil, errors.Wrap(err, "transport error")
		}

		err = v2signature.VerifyServiceMessage(resp)
		if err != nil {
			return nil, errors.Wrap(err, "can't verify response message")
		}

		return &container.Container{
			Container: *resp.GetBody().GetContainer(),
		}, nil
	default:
		return nil, unsupportedProtocolErr
	}
}

func (c Client) listContainerV2(ctx context.Context, owner *owner.ID, opts ...CallOption) ([]*container.ID, error) {
	// apply all available options
	callOptions := defaultCallOptions()
	for i := range opts {
		opts[i].apply(&callOptions)
	}

	reqBody := new(v2container.ListRequestBody)
	reqBody.SetOwnerID(owner.ToV2())

	req := new(v2container.ListRequest)
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err := v2signature.SignServiceMessage(c.key, req)
	if err != nil {
		return nil, err
	}

	switch c.remoteNode.Protocol {
	case GRPC:
		cli, err := v2ContainerClientFromOptions(c.opts)
		if err != nil {
			return nil, errors.Wrap(err, "can't create grpc client")
		}

		resp, err := cli.List(ctx, req)
		if err != nil {
			return nil, errors.Wrap(err, "transport error")
		}

		err = v2signature.VerifyServiceMessage(resp)
		if err != nil {
			return nil, errors.Wrap(err, "can't verify response message")
		}

		result := make([]*container.ID, 0, len(resp.GetBody().GetContainerIDs()))
		for _, cidV2 := range resp.GetBody().GetContainerIDs() {
			result = append(result, container.NewIDFromV2(cidV2))
		}

		return result, nil
	default:
		return nil, unsupportedProtocolErr
	}
}

func (c Client) delContainerV2(ctx context.Context, id *container.ID, opts ...CallOption) error {
	// apply all available options
	callOptions := defaultCallOptions()
	for i := range opts {
		opts[i].apply(&callOptions)
	}

	reqBody := new(v2container.DeleteRequestBody)
	reqBody.SetContainerID(id.ToV2())

	// sign container
	err := signature.SignDataWithHandler(c.key,
		delContainerSignWrapper{
			body: reqBody,
		},
		func(key []byte, sig []byte) {
			containerSignature := new(refs.Signature)
			containerSignature.SetKey(key)
			containerSignature.SetSign(sig)
			reqBody.SetSignature(containerSignature)
		},
		signature.SignWithRFC6979())
	if err != nil {
		return err
	}

	req := new(v2container.DeleteRequest)
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err = v2signature.SignServiceMessage(c.key, req)
	if err != nil {
		return err
	}

	switch c.remoteNode.Protocol {
	case GRPC:
		cli, err := v2ContainerClientFromOptions(c.opts)
		if err != nil {
			return errors.Wrap(err, "can't create grpc client")
		}

		resp, err := cli.Delete(ctx, req)
		if err != nil {
			return errors.Wrap(err, "transport error")
		}

		err = v2signature.VerifyServiceMessage(resp)
		if err != nil {
			return errors.Wrap(err, "can't verify response message")
		}

		return nil
	default:
		return unsupportedProtocolErr
	}
}

func (c Client) getEACLV2(ctx context.Context, id *container.ID, opts ...CallOption) (*eacl.Table, error) {
	// apply all available options
	callOptions := defaultCallOptions()
	for i := range opts {
		opts[i].apply(&callOptions)
	}

	reqBody := new(v2container.GetExtendedACLRequestBody)
	reqBody.SetContainerID(id.ToV2())

	req := new(v2container.GetExtendedACLRequest)
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err := v2signature.SignServiceMessage(c.key, req)
	if err != nil {
		return nil, err
	}

	switch c.remoteNode.Protocol {
	case GRPC:
		cli, err := v2ContainerClientFromOptions(c.opts)
		if err != nil {
			return nil, errors.Wrap(err, "can't create grpc client")
		}

		resp, err := cli.GetExtendedACL(ctx, req)
		if err != nil {
			return nil, errors.Wrap(err, "transport error")
		}

		err = v2signature.VerifyServiceMessage(resp)
		if err != nil {
			return nil, errors.Wrap(err, "can't verify response message")
		}

		body := resp.GetBody()
		if body == nil {
			return nil, errors.New("response body is nil")
		}

		result := eacl.NewTableFromV2(body.GetEACL())

		// todo: check signature

		return result, nil
	default:
		return nil, unsupportedProtocolErr
	}
}

func (c Client) setEACLV2(ctx context.Context, eacl *eacl.Table, opts ...CallOption) error {
	// apply all available options
	callOptions := defaultCallOptions()
	for i := range opts {
		opts[i].apply(&callOptions)
	}

	reqBody := new(v2container.SetExtendedACLRequestBody)
	reqBody.SetEACL(eacl.ToV2())

	signWrapper := v2signature.StableMarshalerWrapper{SM: reqBody.GetEACL()}
	err := signature.SignDataWithHandler(c.key, signWrapper, func(key []byte, sig []byte) {
		eaclSignature := new(refs.Signature)
		eaclSignature.SetKey(key)
		eaclSignature.SetSign(sig)
		reqBody.SetSignature(eaclSignature)
	}, signature.SignWithRFC6979())
	if err != nil {
		return err
	}

	req := new(v2container.SetExtendedACLRequest)
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err = v2signature.SignServiceMessage(c.key, req)
	if err != nil {
		return err
	}

	switch c.remoteNode.Protocol {
	case GRPC:
		cli, err := v2ContainerClientFromOptions(c.opts)
		if err != nil {
			return errors.Wrap(err, "can't create grpc client")
		}

		resp, err := cli.SetExtendedACL(ctx, req)
		if err != nil {
			return errors.Wrap(err, "transport error")
		}

		err = v2signature.VerifyServiceMessage(resp)
		if err != nil {
			return errors.Wrap(err, "can't verify response message")
		}

		return nil
	default:
		return unsupportedProtocolErr
	}
}

func v2ContainerClientFromOptions(opts *clientOptions) (cli *v2container.Client, err error) {
	switch {
	case opts.grpcOpts.v2ContainerClient != nil:
		// return value from client cache
		return opts.grpcOpts.v2ContainerClient, nil

	case opts.grpcOpts.conn != nil:
		cli, err = v2container.NewClient(v2container.WithGlobalOpts(
			client.WithGRPCConn(opts.grpcOpts.conn)),
		)

	case opts.addr != "":
		cli, err = v2container.NewClient(v2container.WithGlobalOpts(
			client.WithNetworkAddress(opts.addr)),
		)

	default:
		return nil, errors.New("lack of sdk client options to create accounting client")
	}

	// check if client correct and save in cache
	if err != nil {
		return nil, err
	}

	opts.grpcOpts.v2ContainerClient = cli

	return cli, nil
}
