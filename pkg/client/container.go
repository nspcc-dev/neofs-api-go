package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/acl/eacl"
	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	"github.com/nspcc-dev/neofs-api-go/rpc/client"
	"github.com/nspcc-dev/neofs-api-go/util/signature"
	v2container "github.com/nspcc-dev/neofs-api-go/v2/container"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	rpcapi "github.com/nspcc-dev/neofs-api-go/v2/rpc"
	v2signature "github.com/nspcc-dev/neofs-api-go/v2/signature"
)

// Container contains methods related to container and ACL.
type Container interface {
	// PutContainer creates new container in the NeoFS network.
	PutContainer(context.Context, *container.Container, ...CallOption) (*container.ID, error)

	// GetContainer returns container by ID.
	GetContainer(context.Context, *container.ID, ...CallOption) (*container.Container, error)

	// ListContainers return container list with the provided owner.
	ListContainers(context.Context, *owner.ID, ...CallOption) ([]*container.ID, error)

	// DeleteContainer removes container from NeoFS network.
	DeleteContainer(context.Context, *container.ID, ...CallOption) error

	// GetEACL returns extended ACL for a given container.
	GetEACL(context.Context, *container.ID, ...CallOption) (*EACLWithSignature, error)

	// SetEACL sets extended ACL.
	SetEACL(context.Context, *eacl.Table, ...CallOption) error

	// AnnounceContainerUsedSpace announces amount of space which is taken by stored objects.
	AnnounceContainerUsedSpace(context.Context, []container.UsedSpaceAnnouncement, ...CallOption) error
}

type delContainerSignWrapper struct {
	body *v2container.DeleteRequestBody
}

// EACLWithSignature represents eACL table/signature pair.
type EACLWithSignature struct {
	table *eacl.Table

	sig *pkg.Signature
}

func (c delContainerSignWrapper) ReadSignedData(bytes []byte) ([]byte, error) {
	return c.body.GetContainerID().GetValue(), nil
}

func (c delContainerSignWrapper) SignedDataSize() int {
	return len(c.body.GetContainerID().GetValue())
}

// EACL returns eACL table.
func (e EACLWithSignature) EACL() *eacl.Table {
	return e.table
}

// Signature returns table signature.
func (e EACLWithSignature) Signature() *pkg.Signature {
	return e.sig
}

func (c *clientImpl) PutContainer(ctx context.Context, cnr *container.Container, opts ...CallOption) (*container.ID, error) {
	// apply all available options
	callOptions := c.defaultCallOptions()

	for i := range opts {
		opts[i](callOptions)
	}

	// set transport version
	cnr.SetVersion(pkg.SDKVersion())

	// if container owner is not set, then use client key as owner
	if cnr.OwnerID() == nil {
		w, err := owner.NEO3WalletFromPublicKey(&callOptions.key.PublicKey)
		if err != nil {
			return nil, err
		}

		ownerID := new(owner.ID)
		ownerID.SetNeo3Wallet(w)

		cnr.SetOwnerID(ownerID)
	}

	reqBody := new(v2container.PutRequestBody)
	reqBody.SetContainer(cnr.ToV2())

	// sign container
	signWrapper := v2signature.StableMarshalerWrapper{SM: reqBody.GetContainer()}

	err := signature.SignDataWithHandler(callOptions.key, signWrapper, func(key []byte, sig []byte) {
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

	meta := v2MetaHeaderFromOpts(callOptions)
	meta.SetSessionToken(cnr.SessionToken().ToV2())

	req.SetMetaHeader(meta)

	err = v2signature.SignServiceMessage(callOptions.key, req)
	if err != nil {
		return nil, err
	}

	resp, err := rpcapi.PutContainer(c.Raw(), req, client.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	err = v2signature.VerifyServiceMessage(resp)
	if err != nil {
		return nil, fmt.Errorf("can't verify response message: %w", err)
	}

	return container.NewIDFromV2(resp.GetBody().GetContainerID()), nil
}

// GetContainer receives container structure through NeoFS API call.
//
// Returns error if container structure is received but does not meet NeoFS API specification.
func (c *clientImpl) GetContainer(ctx context.Context, id *container.ID, opts ...CallOption) (*container.Container, error) {
	// apply all available options
	callOptions := c.defaultCallOptions()

	for i := range opts {
		opts[i](callOptions)
	}

	reqBody := new(v2container.GetRequestBody)
	reqBody.SetContainerID(id.ToV2())

	req := new(v2container.GetRequest)
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err := v2signature.SignServiceMessage(callOptions.key, req)
	if err != nil {
		return nil, err
	}

	resp, err := rpcapi.GetContainer(c.Raw(), req, client.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("transport error: %w", err)
	}

	err = v2signature.VerifyServiceMessage(resp)
	if err != nil {
		return nil, fmt.Errorf("can't verify response message: %w", err)
	}

	return container.NewVerifiedFromV2(resp.GetBody().GetContainer())
}

// GetVerifiedContainerStructure is a wrapper over Client.GetContainer method
// which checks if the structure of the resulting container matches its identifier.
//
// Returns an error if container does not match the identifier.
func GetVerifiedContainerStructure(ctx context.Context, c Client, id *container.ID, opts ...CallOption) (*container.Container, error) {
	cnr, err := c.GetContainer(ctx, id, opts...)
	if err != nil {
		return nil, err
	}

	if !container.CalculateID(cnr).Equal(id) {
		return nil, errors.New("container structure does not match the identifier")
	}

	return cnr, nil
}

func (c *clientImpl) ListContainers(ctx context.Context, ownerID *owner.ID, opts ...CallOption) ([]*container.ID, error) {
	// apply all available options
	callOptions := c.defaultCallOptions()

	for i := range opts {
		opts[i](callOptions)
	}

	if ownerID == nil {
		w, err := owner.NEO3WalletFromPublicKey(&callOptions.key.PublicKey)
		if err != nil {
			return nil, err
		}

		ownerID = new(owner.ID)
		ownerID.SetNeo3Wallet(w)
	}

	reqBody := new(v2container.ListRequestBody)
	reqBody.SetOwnerID(ownerID.ToV2())

	req := new(v2container.ListRequest)
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err := v2signature.SignServiceMessage(callOptions.key, req)
	if err != nil {
		return nil, err
	}

	resp, err := rpcapi.ListContainers(c.Raw(), req, client.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("transport error: %w", err)
	}

	err = v2signature.VerifyServiceMessage(resp)
	if err != nil {
		return nil, fmt.Errorf("can't verify response message: %w", err)
	}

	result := make([]*container.ID, 0, len(resp.GetBody().GetContainerIDs()))
	for _, cidV2 := range resp.GetBody().GetContainerIDs() {
		result = append(result, container.NewIDFromV2(cidV2))
	}

	return result, nil
}

func (c *clientImpl) DeleteContainer(ctx context.Context, id *container.ID, opts ...CallOption) error {
	// apply all available options
	callOptions := c.defaultCallOptions()

	for i := range opts {
		opts[i](callOptions)
	}

	reqBody := new(v2container.DeleteRequestBody)
	reqBody.SetContainerID(id.ToV2())

	// sign container
	err := signature.SignDataWithHandler(callOptions.key,
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

	err = v2signature.SignServiceMessage(callOptions.key, req)
	if err != nil {
		return err
	}

	resp, err := rpcapi.DeleteContainer(c.Raw(), req, client.WithContext(ctx))
	if err != nil {
		return fmt.Errorf("transport error: %w", err)
	}

	if err := v2signature.VerifyServiceMessage(resp); err != nil {
		return fmt.Errorf("can't verify response message: %w", err)
	}

	return nil
}

func (c *clientImpl) GetEACL(ctx context.Context, id *container.ID, opts ...CallOption) (*EACLWithSignature, error) {
	// apply all available options
	callOptions := c.defaultCallOptions()

	for i := range opts {
		opts[i](callOptions)
	}

	reqBody := new(v2container.GetExtendedACLRequestBody)
	reqBody.SetContainerID(id.ToV2())

	req := new(v2container.GetExtendedACLRequest)
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err := v2signature.SignServiceMessage(callOptions.key, req)
	if err != nil {
		return nil, err
	}

	resp, err := rpcapi.GetEACL(c.Raw(), req, client.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("transport error: %w", err)
	}

	err = v2signature.VerifyServiceMessage(resp)
	if err != nil {
		return nil, fmt.Errorf("can't verify response message: %w", err)
	}

	body := resp.GetBody()

	return &EACLWithSignature{
		table: eacl.NewTableFromV2(body.GetEACL()),
		sig:   pkg.NewSignatureFromV2(body.GetSignature()),
	}, nil
}

func (c *clientImpl) SetEACL(ctx context.Context, eacl *eacl.Table, opts ...CallOption) error {
	// apply all available options
	callOptions := c.defaultCallOptions()

	for i := range opts {
		opts[i](callOptions)
	}

	reqBody := new(v2container.SetExtendedACLRequestBody)
	reqBody.SetEACL(eacl.ToV2())
	reqBody.GetEACL().SetVersion(pkg.SDKVersion().ToV2())

	signWrapper := v2signature.StableMarshalerWrapper{SM: reqBody.GetEACL()}

	err := signature.SignDataWithHandler(callOptions.key, signWrapper, func(key []byte, sig []byte) {
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

	err = v2signature.SignServiceMessage(callOptions.key, req)
	if err != nil {
		return err
	}

	resp, err := rpcapi.SetEACL(c.Raw(), req, client.WithContext(ctx))
	if err != nil {
		return fmt.Errorf("transport error: %w", err)
	}

	err = v2signature.VerifyServiceMessage(resp)
	if err != nil {
		return fmt.Errorf("can't verify response message: %w", err)
	}

	return nil
}

// AnnounceContainerUsedSpace used by storage nodes to estimate their container
// sizes during lifetime. Use it only in storage node applications.
func (c *clientImpl) AnnounceContainerUsedSpace(
	ctx context.Context,
	announce []container.UsedSpaceAnnouncement,
	opts ...CallOption) error {
	callOptions := c.defaultCallOptions() // apply all available options

	for i := range opts {
		opts[i](callOptions)
	}

	// convert list of SDK announcement structures into NeoFS-API v2 list
	v2announce := make([]*v2container.UsedSpaceAnnouncement, 0, len(announce))
	for i := range announce {
		v2announce = append(v2announce, announce[i].ToV2())
	}

	// prepare body of the NeoFS-API v2 request and request itself
	reqBody := new(v2container.AnnounceUsedSpaceRequestBody)
	reqBody.SetAnnouncements(v2announce)

	req := new(v2container.AnnounceUsedSpaceRequest)
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	// sign the request
	err := v2signature.SignServiceMessage(callOptions.key, req)
	if err != nil {
		return err
	}

	resp, err := rpcapi.AnnounceUsedSpace(c.Raw(), req, client.WithContext(ctx))
	if err != nil {
		return fmt.Errorf("transport error: %w", err)
	}

	err = v2signature.VerifyServiceMessage(resp)
	if err != nil {
		return fmt.Errorf("can't verify response message: %w", err)
	}

	return nil
}
