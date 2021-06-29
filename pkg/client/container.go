package client

import (
	"context"
	"errors"
	"fmt"

	neofsecdsa "github.com/nspcc-dev/neofs-api-go/crypto/ecdsa"
	neofsrfc6979 "github.com/nspcc-dev/neofs-api-go/crypto/rfc6979"
	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/acl/eacl"
	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	cid "github.com/nspcc-dev/neofs-api-go/pkg/container/id"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	"github.com/nspcc-dev/neofs-api-go/pkg/session"
	v2container "github.com/nspcc-dev/neofs-api-go/v2/container"
	apicrypto "github.com/nspcc-dev/neofs-api-go/v2/crypto"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	rpcapi "github.com/nspcc-dev/neofs-api-go/v2/rpc"
	v2signature "github.com/nspcc-dev/neofs-api-go/v2/signature"
)

// EACLWithSignature represents eACL table/signature pair.
type EACLWithSignature struct {
	table *eacl.Table
}

// EACL returns eACL table.
func (e EACLWithSignature) EACL() *eacl.Table {
	return e.table
}

// Signature returns table signature.
//
// Deprecated: use EACL().Signature() instead.
func (e EACLWithSignature) Signature() *pkg.Signature {
	return e.table.Signature()
}

func (x Client) PutContainer(ctx context.Context, cnr *container.Container, opts ...CallOption) (*cid.ID, error) {
	// apply all available options
	callOptions := defaultCallOptions()

	for i := range opts {
		opts[i](callOptions)
	}

	// set transport version
	cnr.SetVersion(pkg.SDKVersion())

	// if container owner is not set, then use client key as owner
	if cnr.OwnerID() == nil {
		w, err := owner.NEO3WalletFromECDSAPublicKey(callOptions.key.PublicKey)
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
	var (
		p   apicrypto.SignPrm
		sig = new(refs.Signature)
	)

	reqBody.SetSignature(sig)

	p.SetProtoMarshaler(v2signature.StableMarshalerCrypto(reqBody.GetContainer()))
	p.SetTargetSignature(sig)

	err := apicrypto.Sign(neofsrfc6979.Signer(callOptions.key), p)
	if err != nil {
		return nil, err
	}

	var req v2container.PutRequest
	req.SetBody(reqBody)

	meta := v2MetaHeaderFromOpts(callOptions)
	meta.SetSessionToken(cnr.SessionToken().ToV2())

	req.SetMetaHeader(meta)

	err = v2signature.SignServiceMessage(neofsecdsa.Signer(callOptions.key), &req)
	if err != nil {
		return nil, err
	}

	var prm rpcapi.PutContainerPrm

	prm.SetRequest(req)

	var res rpcapi.PutContainerRes

	err = rpcapi.PutContainer(ctx, x.c, prm, &res)
	if err != nil {
		return nil, err
	}

	resp := res.Response()

	err = v2signature.VerifyServiceMessage(&resp)
	if err != nil {
		return nil, fmt.Errorf("can't verify response message: %w", err)
	}

	return cid.NewFromV2(resp.GetBody().GetContainerID()), nil
}

// GetContainer receives container structure through NeoFS API call.
//
// Returns error if container structure is received but does not meet NeoFS API specification.
func (x Client) GetContainer(ctx context.Context, id *cid.ID, opts ...CallOption) (*container.Container, error) {
	// apply all available options
	callOptions := defaultCallOptions()

	for i := range opts {
		opts[i](callOptions)
	}

	reqBody := new(v2container.GetRequestBody)
	reqBody.SetContainerID(id.ToV2())

	var req v2container.GetRequest
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err := v2signature.SignServiceMessage(neofsecdsa.Signer(callOptions.key), &req)
	if err != nil {
		return nil, err
	}

	var prm rpcapi.GetContainerPrm

	prm.SetRequest(req)

	var res rpcapi.GetContainerRes

	err = rpcapi.GetContainer(ctx, x.c, prm, &res)
	if err != nil {
		return nil, fmt.Errorf("transport error: %w", err)
	}

	resp := res.Response()

	err = v2signature.VerifyServiceMessage(&resp)
	if err != nil {
		return nil, fmt.Errorf("can't verify response message: %w", err)
	}

	body := resp.GetBody()

	cnr := container.NewContainerFromV2(body.GetContainer())

	cnr.SetSessionToken(
		session.NewTokenFromV2(body.GetSessionToken()),
	)

	cnr.SetSignature(
		pkg.NewSignatureFromV2(body.GetSignature()),
	)

	return cnr, nil
}

// GetVerifiedContainerStructure is a wrapper over Client.GetContainer method
// which checks if the structure of the resulting container matches its identifier.
//
// Returns an error if container does not match the identifier.
func GetVerifiedContainerStructure(ctx context.Context, c Client, id *cid.ID, opts ...CallOption) (*container.Container, error) {
	cnr, err := c.GetContainer(ctx, id, opts...)
	if err != nil {
		return nil, err
	}

	if !container.CalculateID(cnr).Equal(id) {
		return nil, errors.New("container structure does not match the identifier")
	}

	return cnr, nil
}

func (x Client) ListContainers(ctx context.Context, ownerID *owner.ID, opts ...CallOption) ([]*cid.ID, error) {
	// apply all available options
	callOptions := defaultCallOptions()

	for i := range opts {
		opts[i](callOptions)
	}

	if ownerID == nil {
		w, err := owner.NEO3WalletFromECDSAPublicKey(callOptions.key.PublicKey)
		if err != nil {
			return nil, err
		}

		ownerID = new(owner.ID)
		ownerID.SetNeo3Wallet(w)
	}

	reqBody := new(v2container.ListRequestBody)
	reqBody.SetOwnerID(ownerID.ToV2())

	var req v2container.ListRequest
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err := v2signature.SignServiceMessage(neofsecdsa.Signer(callOptions.key), &req)
	if err != nil {
		return nil, err
	}

	var prm rpcapi.ListContainerPrm

	prm.SetRequest(req)

	var res rpcapi.ListContainerRes

	err = rpcapi.ListContainers(ctx, x.c, prm, &res)
	if err != nil {
		return nil, fmt.Errorf("transport error: %w", err)
	}

	resp := res.Response()

	err = v2signature.VerifyServiceMessage(&resp)
	if err != nil {
		return nil, fmt.Errorf("can't verify response message: %w", err)
	}

	result := make([]*cid.ID, 0, len(resp.GetBody().GetContainerIDs()))
	for _, cidV2 := range resp.GetBody().GetContainerIDs() {
		result = append(result, cid.NewFromV2(cidV2))
	}

	return result, nil
}

func (x Client) DeleteContainer(ctx context.Context, id *cid.ID, opts ...CallOption) error {
	// apply all available options
	callOptions := defaultCallOptions()

	for i := range opts {
		opts[i](callOptions)
	}

	reqBody := new(v2container.DeleteRequestBody)
	reqBody.SetContainerID(id.ToV2())

	// sign container ID
	var (
		p   apicrypto.SignPrm
		sig = new(refs.Signature)
	)

	reqBody.SetSignature(sig)

	p.SetProtoMarshaler(v2signature.StableMarshalerCrypto(reqBody))
	p.SetTargetSignature(sig)

	err := apicrypto.Sign(neofsrfc6979.Signer(callOptions.key), p)
	if err != nil {
		return err
	}

	var req v2container.DeleteRequest
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err = v2signature.SignServiceMessage(neofsecdsa.Signer(callOptions.key), &req)
	if err != nil {
		return err
	}

	var prm rpcapi.DeleteContainerPrm

	prm.SetRequest(req)

	var res rpcapi.DeleteContainerRes

	err = rpcapi.DeleteContainer(ctx, x.c, prm, &res)
	if err != nil {
		return fmt.Errorf("transport error: %w", err)
	}

	resp := res.Response()

	if err := v2signature.VerifyServiceMessage(&resp); err != nil {
		return fmt.Errorf("can't verify response message: %w", err)
	}

	return nil
}

func (x Client) GetEACL(ctx context.Context, id *cid.ID, opts ...CallOption) (*EACLWithSignature, error) {
	// apply all available options
	callOptions := defaultCallOptions()

	for i := range opts {
		opts[i](callOptions)
	}

	reqBody := new(v2container.GetExtendedACLRequestBody)
	reqBody.SetContainerID(id.ToV2())

	var req v2container.GetExtendedACLRequest
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err := v2signature.SignServiceMessage(neofsecdsa.Signer(callOptions.key), &req)
	if err != nil {
		return nil, err
	}

	var prm rpcapi.GetEACLPrm

	prm.SetRequest(req)

	var res rpcapi.GetEACLRes

	err = rpcapi.GetEACL(ctx, x.c, prm, &res)
	if err != nil {
		return nil, fmt.Errorf("transport error: %w", err)
	}

	resp := res.Response()

	err = v2signature.VerifyServiceMessage(&resp)
	if err != nil {
		return nil, fmt.Errorf("can't verify response message: %w", err)
	}

	body := resp.GetBody()

	table := eacl.NewTableFromV2(body.GetEACL())

	table.SetSessionToken(
		session.NewTokenFromV2(body.GetSessionToken()),
	)

	table.SetSignature(
		pkg.NewSignatureFromV2(body.GetSignature()),
	)

	return &EACLWithSignature{
		table: table,
	}, nil
}

func (x Client) SetEACL(ctx context.Context, eacl *eacl.Table, opts ...CallOption) error {
	// apply all available options
	callOptions := defaultCallOptions()

	for i := range opts {
		opts[i](callOptions)
	}

	reqBody := new(v2container.SetExtendedACLRequestBody)
	reqBody.SetEACL(eacl.ToV2())
	reqBody.GetEACL().SetVersion(pkg.SDKVersion().ToV2())

	// sign eACL table
	var (
		p   apicrypto.SignPrm
		sig = new(refs.Signature)
	)

	reqBody.SetSignature(sig)

	p.SetProtoMarshaler(v2signature.StableMarshalerCrypto(reqBody.GetEACL()))
	p.SetTargetSignature(sig)

	err := apicrypto.Sign(neofsrfc6979.Signer(callOptions.key), p)
	if err != nil {
		return err
	}

	var req v2container.SetExtendedACLRequest
	req.SetBody(reqBody)

	meta := v2MetaHeaderFromOpts(callOptions)
	meta.SetSessionToken(eacl.SessionToken().ToV2())

	req.SetMetaHeader(meta)

	err = v2signature.SignServiceMessage(neofsecdsa.Signer(callOptions.key), &req)
	if err != nil {
		return err
	}

	var prm rpcapi.SetEACLPrm

	prm.SetRequest(req)

	var res rpcapi.SetEACLRes

	err = rpcapi.SetEACL(ctx, x.c, prm, &res)
	if err != nil {
		return fmt.Errorf("transport error: %w", err)
	}

	resp := res.Response()

	err = v2signature.VerifyServiceMessage(&resp)
	if err != nil {
		return fmt.Errorf("can't verify response message: %w", err)
	}

	return nil
}

// AnnounceContainerUsedSpace used by storage nodes to estimate their container
// sizes during lifetime. Use it only in storage node applications.
func (x Client) AnnounceContainerUsedSpace(
	ctx context.Context,
	announce []container.UsedSpaceAnnouncement,
	opts ...CallOption) error {
	callOptions := defaultCallOptions() // apply all available options

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

	var req v2container.AnnounceUsedSpaceRequest
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	// sign the request
	err := v2signature.SignServiceMessage(neofsecdsa.Signer(callOptions.key), &req)
	if err != nil {
		return err
	}

	var prm rpcapi.AnnounceUsedSpacePrm

	prm.SetRequest(req)

	var res rpcapi.AnnounceUsedSpaceRes

	err = rpcapi.AnnounceUsedSpace(ctx, x.c, prm, &res)
	if err != nil {
		return fmt.Errorf("transport error: %w", err)
	}

	resp := res.Response()

	err = v2signature.VerifyServiceMessage(&resp)
	if err != nil {
		return fmt.Errorf("can't verify response message: %w", err)
	}

	return nil
}
