package client

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"io"

	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	"github.com/nspcc-dev/neofs-api-go/pkg/object"
	signer "github.com/nspcc-dev/neofs-api-go/util/signature"
	"github.com/nspcc-dev/neofs-api-go/v2/client"
	v2object "github.com/nspcc-dev/neofs-api-go/v2/object"
	v2refs "github.com/nspcc-dev/neofs-api-go/v2/refs"
	v2session "github.com/nspcc-dev/neofs-api-go/v2/session"
	"github.com/nspcc-dev/neofs-api-go/v2/signature"
	"github.com/pkg/errors"
)

type PutObjectParams struct {
	obj *object.Object

	r io.Reader
}

type DeleteObjectParams struct {
	addr *object.Address
}

type GetObjectParams struct {
	addr *object.Address

	w io.Writer
}

type ObjectHeaderParams struct {
	addr *object.Address

	short bool
}

type RangeDataParams struct {
	addr *object.Address

	r *object.Range

	w io.Writer
}

type RangeChecksumParams struct {
	typ checksumType

	addr *object.Address

	rs []*object.Range

	salt []byte
}

type SearchObjectParams struct {
	cid *container.ID

	filters object.SearchFilters
}

type putObjectV2Writer struct {
	key *ecdsa.PrivateKey

	chunkPart *v2object.PutObjectPartChunk

	req *v2object.PutRequest

	stream v2object.PutObjectStreamer
}

type checksumType int

const (
	_ checksumType = iota
	checksumSHA256
	checksumTZ
)

const chunkSize = 3 * (1 << 20)

const TZSize = 64

const searchQueryVersion uint32 = 1

func rangesToV2(rs []*object.Range) []*v2object.Range {
	r2 := make([]*v2object.Range, 0, len(rs))

	for i := range rs {
		r2 = append(r2, rs[i].ToV2())
	}

	return r2
}

func (t checksumType) toV2() v2refs.ChecksumType {
	switch t {
	case checksumSHA256:
		return v2refs.SHA256
	case checksumTZ:
		return v2refs.TillichZemor
	default:
		panic(fmt.Sprintf("invalid checksum type %d", t))
	}
}

func (w *putObjectV2Writer) Write(p []byte) (int, error) {
	w.chunkPart.SetChunk(p)

	w.req.SetVerificationHeader(nil)

	if err := signature.SignServiceMessage(w.key, w.req); err != nil {
		return 0, errors.Wrap(err, "could not sign chunk request message")
	}

	if err := w.stream.Send(w.req); err != nil {
		return 0, errors.Wrap(err, "could not send chunk request message")
	}

	return len(p), nil
}

func (p *PutObjectParams) WithObject(v *object.Object) *PutObjectParams {
	if p != nil {
		p.obj = v
	}

	return p
}

func (p *PutObjectParams) WithPayloadReader(v io.Reader) *PutObjectParams {
	if p != nil {
		p.r = v
	}

	return p
}

func (c *Client) PutObject(ctx context.Context, p *PutObjectParams, opts ...CallOption) (*object.ID, error) {
	// check remote node version
	switch c.remoteNode.Version.Major() {
	case 2:
		return c.putObjectV2(ctx, p, opts...)
	default:
		return nil, unsupportedProtocolErr
	}
}

func (c *Client) putObjectV2(ctx context.Context, p *PutObjectParams, opts ...CallOption) (*object.ID, error) {
	// create V2 Object client
	cli, err := v2ObjectClient(c.remoteNode.Protocol, c.opts)
	if err != nil {
		return nil, errors.Wrap(err, "could not create Object V2 client")
	}

	stream, err := cli.Put(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "could not open Put object stream")
	}

	callOpts := c.defaultCallOptions()

	for i := range opts {
		if opts[i] != nil {
			opts[i].apply(&callOpts)
		}
	}

	// create request
	req := new(v2object.PutRequest)

	// initialize request body
	body := new(v2object.PutRequestBody)
	req.SetBody(body)

	v2Addr := new(v2refs.Address)
	v2Addr.SetObjectID(p.obj.GetID().ToV2())
	v2Addr.SetContainerID(p.obj.GetContainerID().ToV2())

	// set meta header
	meta := v2MetaHeaderFromOpts(callOpts)
	if err = c.attachV2SessionToken(callOpts, meta, v2SessionReqInfo{
		addr: v2Addr,
		verb: v2session.ObjectVerbPut,
	}); err != nil {
		return nil, errors.Wrap(err, "could not sign session token")
	}
	req.SetMetaHeader(meta)

	// initialize init part
	initPart := new(v2object.PutObjectPartInit)
	body.SetObjectPart(initPart)

	obj := p.obj.ToV2()

	// set init part fields
	initPart.SetObjectID(obj.GetObjectID())
	initPart.SetSignature(obj.GetSignature())
	initPart.SetHeader(obj.GetHeader())

	// sign the request
	if err := signature.SignServiceMessage(c.key, req); err != nil {
		return nil, errors.Wrapf(err, "could not sign %T", req)
	}

	// send init part
	if err := stream.Send(req); err != nil {
		return nil, errors.Wrapf(err, "could not send %T", req)
	}

	// create payload bytes reader
	var rPayload io.Reader = bytes.NewReader(obj.GetPayload())
	if p.r != nil {
		rPayload = io.MultiReader(rPayload, p.r)
	}

	// create v2 payload stream writer
	chunkPart := new(v2object.PutObjectPartChunk)
	body.SetObjectPart(chunkPart)

	w := &putObjectV2Writer{
		key:       c.key,
		chunkPart: chunkPart,
		req:       req,
		stream:    stream,
	}

	// copy payload from reader to stream writer
	_, err = io.CopyBuffer(w, rPayload, make([]byte, chunkSize))
	if err != nil && !errors.Is(errors.Cause(err), io.EOF) {
		return nil, errors.Wrap(err, "could not send payload bytes to Put object stream")
	}

	// close object stream and receive response from remote node
	resp, err := stream.CloseAndRecv()
	if err != nil {
		return nil, errors.Wrapf(err, "could not close %T", stream)
	}

	// verify response structure
	if err := signature.VerifyServiceMessage(resp); err != nil {
		return nil, errors.Wrapf(err, "could not verify %T", resp)
	}

	// convert object identifier
	id := object.NewIDFromV2(resp.GetBody().GetObjectID())

	return id, nil
}

func (p *DeleteObjectParams) WithAddress(v *object.Address) *DeleteObjectParams {
	if p != nil {
		p.addr = v
	}

	return p
}

func (c *Client) DeleteObject(ctx context.Context, p *DeleteObjectParams, opts ...CallOption) error {
	// check remote node version
	switch c.remoteNode.Version.Major() {
	case 2:
		return c.deleteObjectV2(ctx, p, opts...)
	default:
		return unsupportedProtocolErr
	}
}

func (c *Client) deleteObjectV2(ctx context.Context, p *DeleteObjectParams, opts ...CallOption) error {
	// create V2 Object client
	cli, err := v2ObjectClient(c.remoteNode.Protocol, c.opts)
	if err != nil {
		return errors.Wrap(err, "could not create Object V2 client")
	}

	callOpts := c.defaultCallOptions()

	for i := range opts {
		if opts[i] != nil {
			opts[i].apply(&callOpts)
		}
	}

	// create request
	req := new(v2object.DeleteRequest)

	// initialize request body
	body := new(v2object.DeleteRequestBody)
	req.SetBody(body)

	// set meta header
	meta := v2MetaHeaderFromOpts(callOpts)
	if err = c.attachV2SessionToken(callOpts, meta, v2SessionReqInfo{
		addr: p.addr.ToV2(),
		verb: v2session.ObjectVerbDelete,
	}); err != nil {
		return errors.Wrap(err, "could not sign session token")
	}
	req.SetMetaHeader(meta)

	// fill body fields
	body.SetAddress(p.addr.ToV2())

	// sign the request
	if err := signature.SignServiceMessage(c.key, req); err != nil {
		return errors.Wrapf(err, "could not sign %T", req)
	}

	// send request
	resp, err := cli.Delete(ctx, req)
	if err != nil {
		return errors.Wrapf(err, "could not send %T", req)
	}

	// verify response structure
	if err := signature.VerifyServiceMessage(resp); err != nil {
		return errors.Wrapf(err, "could not verify %T", resp)
	}

	return nil
}

func (p *GetObjectParams) WithAddress(v *object.Address) *GetObjectParams {
	if p != nil {
		p.addr = v
	}

	return p
}

func (p *GetObjectParams) WithPayloadWriter(w io.Writer) *GetObjectParams {
	if p != nil {
		p.w = w
	}

	return p
}

func (c *Client) GetObject(ctx context.Context, p *GetObjectParams, opts ...CallOption) (*object.Object, error) {
	// check remote node version
	switch c.remoteNode.Version.Major() {
	case 2:
		return c.getObjectV2(ctx, p, opts...)
	default:
		return nil, unsupportedProtocolErr
	}
}

func (c *Client) getObjectV2(ctx context.Context, p *GetObjectParams, opts ...CallOption) (*object.Object, error) {
	// create V2 Object client
	cli, err := v2ObjectClient(c.remoteNode.Protocol, c.opts)
	if err != nil {
		return nil, errors.Wrap(err, "could not create Object V2 client")
	}

	callOpts := c.defaultCallOptions()

	for i := range opts {
		if opts[i] != nil {
			opts[i].apply(&callOpts)
		}
	}

	// create request
	req := new(v2object.GetRequest)

	// initialize request body
	body := new(v2object.GetRequestBody)
	req.SetBody(body)

	// set meta header
	meta := v2MetaHeaderFromOpts(callOpts)
	if err = c.attachV2SessionToken(callOpts, meta, v2SessionReqInfo{
		addr: p.addr.ToV2(),
		verb: v2session.ObjectVerbGet,
	}); err != nil {
		return nil, errors.Wrap(err, "could not sign session token")
	}
	req.SetMetaHeader(meta)

	// fill body fields
	body.SetAddress(p.addr.ToV2())

	// sign the request
	if err := signature.SignServiceMessage(c.key, req); err != nil {
		return nil, errors.Wrapf(err, "could not sign %T", req)
	}

	// create Get object stream
	stream, err := cli.Get(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "could not create Get object stream")
	}

	var (
		payload []byte
		obj     = new(v2object.Object)
	)

	for {
		// receive message from server stream
		resp, err := stream.Recv()
		if err != nil {
			if errors.Is(errors.Cause(err), io.EOF) {
				break
			}

			return nil, errors.Wrap(err, "could not receive Get response")
		}

		// verify response structure
		if err := signature.VerifyServiceMessage(resp); err != nil {
			return nil, errors.Wrapf(err, "could not verify %T", resp)
		}

		switch v := resp.GetBody().GetObjectPart().(type) {
		case nil:
			return nil, errors.New("received nil object part")
		case *v2object.GetObjectPartInit:
			obj.SetObjectID(v.GetObjectID())
			obj.SetSignature(v.GetSignature())

			hdr := v.GetHeader()
			obj.SetHeader(hdr)

			if p.w == nil {
				payload = make([]byte, 0, hdr.GetPayloadLength())
			}
		case *v2object.GetObjectPartChunk:
			if p.w != nil {
				if _, err := p.w.Write(v.GetChunk()); err != nil {
					return nil, errors.Wrap(err, "could not write payload chunk")
				}
			} else {
				payload = append(payload, v.GetChunk()...)
			}
		default:
			panic(fmt.Sprintf("unexpected Get object part type %T", v))
		}
	}

	obj.SetPayload(payload)

	// convert the object
	return object.NewFromV2(obj), nil
}

func (p *ObjectHeaderParams) WithAddress(v *object.Address) *ObjectHeaderParams {
	if p != nil {
		p.addr = v
	}

	return p
}

func (p *ObjectHeaderParams) WithAllFields() *ObjectHeaderParams {
	if p != nil {
		p.short = false
	}

	return p
}

func (p *ObjectHeaderParams) WithMainFields() *ObjectHeaderParams {
	if p != nil {
		p.short = true
	}

	return p
}

func (c *Client) GetObjectHeader(ctx context.Context, p *ObjectHeaderParams, opts ...CallOption) (*object.Object, error) {
	// check remote node version
	switch c.remoteNode.Version.Major() {
	case 2:
		return c.getObjectHeaderV2(ctx, p, opts...)
	default:
		return nil, unsupportedProtocolErr
	}
}

func (c *Client) getObjectHeaderV2(ctx context.Context, p *ObjectHeaderParams, opts ...CallOption) (*object.Object, error) {
	// create V2 Object client
	cli, err := v2ObjectClient(c.remoteNode.Protocol, c.opts)
	if err != nil {
		return nil, errors.Wrap(err, "could not create Object V2 client")
	}

	callOpts := c.defaultCallOptions()

	for i := range opts {
		if opts[i] != nil {
			opts[i].apply(&callOpts)
		}
	}

	// create request
	req := new(v2object.HeadRequest)

	// initialize request body
	body := new(v2object.HeadRequestBody)
	req.SetBody(body)

	// set meta header
	meta := v2MetaHeaderFromOpts(callOpts)
	if err = c.attachV2SessionToken(callOpts, meta, v2SessionReqInfo{
		addr: p.addr.ToV2(),
		verb: v2session.ObjectVerbHead,
	}); err != nil {
		return nil, errors.Wrap(err, "could not sign session token")
	}
	req.SetMetaHeader(meta)

	// fill body fields
	body.SetAddress(p.addr.ToV2())
	body.SetMainOnly(p.short)

	// sign the request
	if err := signature.SignServiceMessage(c.key, req); err != nil {
		return nil, errors.Wrapf(err, "could not sign %T", req)
	}

	// send Head request
	resp, err := cli.Head(ctx, req)
	if err != nil {
		return nil, errors.Wrapf(err, "could not send %T", req)
	}

	// verify response structure
	if err := signature.VerifyServiceMessage(resp); err != nil {
		return nil, errors.Wrapf(err, "could not verify %T", resp)
	}

	var hdr *v2object.Header

	switch v := resp.GetBody().GetHeaderPart().(type) {
	case nil:
		return nil, errors.New("received nil object header part")
	case *v2object.GetHeaderPartShort:
		if !p.short {
			return nil, errors.Errorf("wrong header part type: expected %T, received %T",
				(*v2object.GetHeaderPartFull)(nil), (*v2object.GetHeaderPartShort)(nil),
			)
		}

		h := v.GetShortHeader()

		hdr = new(v2object.Header)
		hdr.SetPayloadLength(h.GetPayloadLength())
		hdr.SetVersion(h.GetVersion())
		hdr.SetOwnerID(h.GetOwnerID())
		hdr.SetObjectType(h.GetObjectType())
		hdr.SetCreationEpoch(h.GetCreationEpoch())
	case *v2object.GetHeaderPartFull:
		if p.short {
			return nil, errors.Errorf("wrong header part type: expected %T, received %T",
				(*v2object.GetHeaderPartShort)(nil), (*v2object.GetHeaderPartFull)(nil),
			)
		}

		hdrWithSig := v.GetHeaderWithSignature()
		if hdrWithSig == nil {
			return nil, errors.New("got nil instead of header with signature")
		}
		hdr = hdrWithSig.GetHeader()
		// todo: check signature there
	default:
		panic(fmt.Sprintf("unexpected Head object type %T", v))
	}

	obj := new(v2object.Object)
	obj.SetHeader(hdr)

	raw := object.NewRawFromV2(obj)
	raw.SetID(p.addr.ObjectID())

	// convert the object
	return raw.Object(), nil
}

func (p *RangeDataParams) WithAddress(v *object.Address) *RangeDataParams {
	if p != nil {
		p.addr = v
	}

	return p
}

func (p *RangeDataParams) WithRange(v *object.Range) *RangeDataParams {
	if p != nil {
		p.r = v
	}

	return p
}

func (p *RangeDataParams) WithDataWriter(v io.Writer) *RangeDataParams {
	if p != nil {
		p.w = v
	}

	return p
}

func (c *Client) ObjectPayloadRangeData(ctx context.Context, p *RangeDataParams, opts ...CallOption) ([]byte, error) {
	// check remote node version
	switch c.remoteNode.Version.Major() {
	case 2:
		return c.objectPayloadRangeV2(ctx, p, opts...)
	default:
		return nil, unsupportedProtocolErr
	}
}

func (c *Client) objectPayloadRangeV2(ctx context.Context, p *RangeDataParams, opts ...CallOption) ([]byte, error) {
	// create V2 Object client
	cli, err := v2ObjectClient(c.remoteNode.Protocol, c.opts)
	if err != nil {
		return nil, errors.Wrap(err, "could not create Object V2 client")
	}

	callOpts := c.defaultCallOptions()

	for i := range opts {
		if opts[i] != nil {
			opts[i].apply(&callOpts)
		}
	}

	// create request
	req := new(v2object.GetRangeRequest)

	// initialize request body
	body := new(v2object.GetRangeRequestBody)
	req.SetBody(body)

	// set meta header
	meta := v2MetaHeaderFromOpts(callOpts)
	if err = c.attachV2SessionToken(callOpts, meta, v2SessionReqInfo{
		addr: p.addr.ToV2(),
		verb: v2session.ObjectVerbRange,
	}); err != nil {
		return nil, errors.Wrap(err, "could not sign session token")
	}
	req.SetMetaHeader(meta)

	// fill body fields
	body.SetAddress(p.addr.ToV2())
	body.SetRange(p.r.ToV2())

	// sign the request
	if err := signature.SignServiceMessage(c.key, req); err != nil {
		return nil, errors.Wrapf(err, "could not sign %T", req)
	}

	// create Get payload range stream
	stream, err := cli.GetRange(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "could not create Get payload range stream")
	}

	var payload []byte
	if p.w != nil {
		payload = make([]byte, p.r.GetLength())
	}

	for {
		// receive message from server stream
		resp, err := stream.Recv()
		if err != nil {
			if errors.Is(errors.Cause(err), io.EOF) {
				break
			}

			return nil, errors.Wrap(err, "could not receive Get payload range response")
		}

		// verify response structure
		if err := signature.VerifyServiceMessage(resp); err != nil {
			return nil, errors.Wrapf(err, "could not verify %T", resp)
		}

		chunk := resp.GetBody().GetChunk()

		if p.w != nil {
			if _, err := p.w.Write(chunk); err != nil {
				return nil, errors.Wrap(err, "could not write payload chunk")
			}
		} else {
			payload = append(payload, chunk...)
		}
	}

	return payload, nil
}

func (p *RangeChecksumParams) WithAddress(v *object.Address) *RangeChecksumParams {
	if p != nil {
		p.addr = v
	}

	return p
}

func (p *RangeChecksumParams) WithRangeList(rs ...*object.Range) *RangeChecksumParams {
	if p != nil {
		p.rs = rs
	}

	return p
}

func (p *RangeChecksumParams) WithSalt(v []byte) *RangeChecksumParams {
	if p != nil {
		p.salt = v
	}

	return p
}

func (p *RangeChecksumParams) withChecksumType(t checksumType) *RangeChecksumParams {
	if p != nil {
		p.typ = t
	}

	return p
}

func (c *Client) ObjectPayloadRangeSHA256(ctx context.Context, p *RangeChecksumParams, opts ...CallOption) ([][sha256.Size]byte, error) {
	res, err := c.objectPayloadRangeHash(ctx, p.withChecksumType(checksumSHA256), opts...)
	if err != nil {
		return nil, err
	}

	return res.([][sha256.Size]byte), nil
}

func (c *Client) ObjectPayloadRangeTZ(ctx context.Context, p *RangeChecksumParams, opts ...CallOption) ([][TZSize]byte, error) {
	res, err := c.objectPayloadRangeHash(ctx, p.withChecksumType(checksumTZ), opts...)
	if err != nil {
		return nil, err
	}

	return res.([][TZSize]byte), nil
}

func (c *Client) objectPayloadRangeHash(ctx context.Context, p *RangeChecksumParams, opts ...CallOption) (interface{}, error) {
	// check remote node version
	switch c.remoteNode.Version.Major() {
	case 2:
		return c.objectPayloadRangeHashV2(ctx, p, opts...)
	default:
		return nil, unsupportedProtocolErr
	}
}

func (c *Client) objectPayloadRangeHashV2(ctx context.Context, p *RangeChecksumParams, opts ...CallOption) (interface{}, error) {
	// create V2 Object client
	cli, err := v2ObjectClient(c.remoteNode.Protocol, c.opts)
	if err != nil {
		return nil, errors.Wrap(err, "could not create Object V2 client")
	}

	callOpts := c.defaultCallOptions()

	for i := range opts {
		if opts[i] != nil {
			opts[i].apply(&callOpts)
		}
	}

	// create request
	req := new(v2object.GetRangeHashRequest)

	// initialize request body
	body := new(v2object.GetRangeHashRequestBody)
	req.SetBody(body)

	// set meta header
	meta := v2MetaHeaderFromOpts(callOpts)
	if err = c.attachV2SessionToken(callOpts, meta, v2SessionReqInfo{
		addr: p.addr.ToV2(),
		verb: v2session.ObjectVerbRangeHash,
	}); err != nil {
		return nil, errors.Wrap(err, "could not sign session token")
	}
	req.SetMetaHeader(meta)

	// fill body fields
	body.SetAddress(p.addr.ToV2())
	body.SetSalt(p.salt)

	typV2 := p.typ.toV2()
	body.SetType(typV2)

	rsV2 := rangesToV2(p.rs)
	body.SetRanges(rsV2)

	// sign the request
	if err := signature.SignServiceMessage(c.key, req); err != nil {
		return nil, errors.Wrapf(err, "could not sign %T", req)
	}

	// send request
	resp, err := cli.GetRangeHash(ctx, req)
	if err != nil {
		return nil, errors.Wrapf(err, "could not send %T", req)
	}

	// verify response structure
	if err := signature.VerifyServiceMessage(resp); err != nil {
		return nil, errors.Wrapf(err, "could not verify %T", resp)
	}

	respBody := resp.GetBody()
	respType := respBody.GetType()
	respHashes := respBody.GetHashList()

	if t := p.typ.toV2(); respType != t {
		return nil, errors.Errorf("invalid checksum type: expected %v, received %v", t, respType)
	} else if reqLn, respLn := len(rsV2), len(respHashes); reqLn != respLn {
		return nil, errors.Errorf("wrong checksum number: expected %d, received %d", reqLn, respLn)
	}

	var res interface{}

	switch p.typ {
	case checksumSHA256:
		r := make([][sha256.Size]byte, 0, len(respHashes))

		for i := range respHashes {
			if ln := len(respHashes[i]); ln != sha256.Size {
				return nil, errors.Errorf("invalid checksum length: expected %d, received %d", sha256.Size, ln)
			}

			cs := [sha256.Size]byte{}
			copy(cs[:], respHashes[i])

			r = append(r, cs)
		}

		res = r
	case checksumTZ:
		r := make([][TZSize]byte, 0, len(respHashes))

		for i := range respHashes {
			if ln := len(respHashes[i]); ln != TZSize {
				return nil, errors.Errorf("invalid checksum length: expected %d, received %d", TZSize, ln)
			}

			cs := [TZSize]byte{}
			copy(cs[:], respHashes[i])

			r = append(r, cs)
		}

		res = r
	}

	return res, nil
}

func (p *SearchObjectParams) WithContainerID(v *container.ID) *SearchObjectParams {
	if p != nil {
		p.cid = v
	}

	return p
}

func (p *SearchObjectParams) WithSearchFilters(v object.SearchFilters) *SearchObjectParams {
	if p != nil {
		p.filters = v
	}

	return p
}

func (c *Client) SearchObject(ctx context.Context, p *SearchObjectParams, opts ...CallOption) ([]*object.ID, error) {
	// check remote node version
	switch c.remoteNode.Version.Major() {
	case 2:
		return c.searchObjectV2(ctx, p, opts...)
	default:
		return nil, unsupportedProtocolErr
	}
}

func (c *Client) searchObjectV2(ctx context.Context, p *SearchObjectParams, opts ...CallOption) ([]*object.ID, error) {
	// create V2 Object client
	cli, err := v2ObjectClient(c.remoteNode.Protocol, c.opts)
	if err != nil {
		return nil, errors.Wrap(err, "could not create Object V2 client")
	}

	callOpts := c.defaultCallOptions()

	for i := range opts {
		if opts[i] != nil {
			opts[i].apply(&callOpts)
		}
	}

	// create request
	req := new(v2object.SearchRequest)

	// initialize request body
	body := new(v2object.SearchRequestBody)
	req.SetBody(body)

	v2Addr := new(v2refs.Address)
	v2Addr.SetContainerID(p.cid.ToV2())

	// set meta header
	meta := v2MetaHeaderFromOpts(callOpts)
	if err = c.attachV2SessionToken(callOpts, meta, v2SessionReqInfo{
		addr: v2Addr,
		verb: v2session.ObjectVerbSearch,
	}); err != nil {
		return nil, errors.Wrap(err, "could not sign session token")
	}
	req.SetMetaHeader(meta)

	// fill body fields
	body.SetContainerID(v2Addr.GetContainerID())
	body.SetVersion(searchQueryVersion)
	body.SetFilters(p.filters.ToV2())

	// sign the request
	if err := signature.SignServiceMessage(c.key, req); err != nil {
		return nil, errors.Wrapf(err, "could not sign %T", req)
	}

	// create search stream
	stream, err := cli.Search(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "could not create search stream")
	}

	var searchResult []*object.ID

	for {
		// receive message from server stream
		resp, err := stream.Recv()
		if err != nil {
			if errors.Is(errors.Cause(err), io.EOF) {
				break
			}

			return nil, errors.Wrap(err, "could not receive search response")
		}

		// verify response structure
		if err := signature.VerifyServiceMessage(resp); err != nil {
			return nil, errors.Wrapf(err, "could not verify %T", resp)
		}

		chunk := resp.GetBody().GetIDList()
		for i := range chunk {
			searchResult = append(searchResult, object.NewIDFromV2(chunk[i]))
		}
	}

	return searchResult, nil
}

func v2ObjectClient(proto TransportProtocol, opts *clientOptions) (*v2object.Client, error) {
	switch proto {
	case GRPC:
		var err error

		if opts.grpcOpts.objectClientV2 == nil {
			var optsV2 []v2object.Option

			if opts.grpcOpts.conn != nil {
				optsV2 = []v2object.Option{
					v2object.WithGlobalOpts(
						client.WithGRPCConn(opts.grpcOpts.conn),
					),
				}
			} else {
				optsV2 = []v2object.Option{
					v2object.WithGlobalOpts(
						client.WithNetworkAddress(opts.addr),
					),
				}
			}

			opts.grpcOpts.objectClientV2, err = v2object.NewClient(optsV2...)
		}

		return opts.grpcOpts.objectClientV2, err
	default:
		return nil, unsupportedProtocolErr
	}
}

func (c Client) attachV2SessionToken(opts callOptions, hdr *v2session.RequestMetaHeader, info v2SessionReqInfo) error {
	if opts.session == nil {
		return nil
	}

	// Do not resign already prepared session token
	if opts.session.Signature() != nil {
		hdr.SetSessionToken(opts.session.ToV2())
		return nil
	}

	token := new(v2session.SessionToken)
	token.SetBody(opts.session.ToV2().GetBody())

	opCtx := new(v2session.ObjectSessionContext)
	opCtx.SetAddress(info.addr)
	opCtx.SetVerb(info.verb)

	lt := new(v2session.TokenLifetime)
	lt.SetIat(info.iat)
	lt.SetNbf(info.nbf)
	lt.SetExp(info.exp)

	body := token.GetBody()
	body.SetSessionKey(opts.session.SessionKey())
	body.SetContext(opCtx)
	body.SetLifetime(lt)

	signWrapper := signature.StableMarshalerWrapper{SM: token.GetBody()}
	err := signer.SignDataWithHandler(c.key, signWrapper, func(key []byte, sig []byte) {
		sessionTokenSignature := new(v2refs.Signature)
		sessionTokenSignature.SetKey(key)
		sessionTokenSignature.SetSign(sig)
		token.SetSignature(sessionTokenSignature)
	})
	if err != nil {
		return err
	}

	hdr.SetSessionToken(token)

	return nil
}
