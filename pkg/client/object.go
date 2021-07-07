package client

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"

	cid "github.com/nspcc-dev/neofs-api-go/pkg/container/id"
	"github.com/nspcc-dev/neofs-api-go/pkg/object"
	"github.com/nspcc-dev/neofs-api-go/rpc/client"
	signer "github.com/nspcc-dev/neofs-api-go/util/signature"
	v2object "github.com/nspcc-dev/neofs-api-go/v2/object"
	v2refs "github.com/nspcc-dev/neofs-api-go/v2/refs"
	rpcapi "github.com/nspcc-dev/neofs-api-go/v2/rpc"
	v2session "github.com/nspcc-dev/neofs-api-go/v2/session"
	"github.com/nspcc-dev/neofs-api-go/v2/signature"
)

// Object contains methods for working with objects.
type Object interface {
	// PutObject puts new object to NeoFS.
	PutObject(context.Context, *PutObjectParams, ...CallOption) (*object.ID, error)

	// DeleteObject deletes object to NeoFS.
	DeleteObject(context.Context, *DeleteObjectParams, ...CallOption) error

	// GetObject returns object stored in NeoFS.
	GetObject(context.Context, *GetObjectParams, ...CallOption) (*object.Object, error)

	// GetObjectHeader returns object header.
	GetObjectHeader(context.Context, *ObjectHeaderParams, ...CallOption) (*object.Object, error)

	// ObjectPayloadRangeData returns range of object payload.
	ObjectPayloadRangeData(context.Context, *RangeDataParams, ...CallOption) ([]byte, error)

	// ObjectPayloadRangeSHA256 returns sha-256 hashes of object sub-ranges from NeoFS.
	ObjectPayloadRangeSHA256(context.Context, *RangeChecksumParams, ...CallOption) ([][sha256.Size]byte, error)

	// ObjectPayloadRangeTZ returns homomorphic hashes of object sub-ranges from NeoFS.
	ObjectPayloadRangeTZ(context.Context, *RangeChecksumParams, ...CallOption) ([][TZSize]byte, error)

	// SearchObject searches for objects in NeoFS using provided parameters.
	SearchObject(context.Context, *SearchObjectParams, ...CallOption) ([]*object.ID, error)
}

type PutObjectParams struct {
	obj *object.Object

	r io.Reader
}

// ObjectAddressWriter is an interface of the
// component that writes the object address.
type ObjectAddressWriter interface {
	SetAddress(*object.Address)
}

type objectAddressWriter struct {
	addr *object.Address
}

type DeleteObjectParams struct {
	addr *object.Address

	tombTgt ObjectAddressWriter
}

type GetObjectParams struct {
	addr *object.Address

	raw bool

	w io.Writer

	readerHandler ReaderHandler
}

type ObjectHeaderParams struct {
	addr *object.Address

	raw bool

	short bool
}

type RangeDataParams struct {
	addr *object.Address

	raw bool

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
	cid *cid.ID

	filters object.SearchFilters
}

type putObjectV2Reader struct {
	r io.Reader
}

type putObjectV2Writer struct {
	key *ecdsa.PrivateKey

	chunkPart *v2object.PutObjectPartChunk

	req *v2object.PutRequest

	stream *rpcapi.PutRequestWriter
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

var errNilObjectPart = errors.New("received nil object part")

func (w *objectAddressWriter) SetAddress(addr *object.Address) {
	w.addr = addr
}

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

func (w *putObjectV2Reader) Read(p []byte) (int, error) {
	return w.r.Read(p)
}

func (w *putObjectV2Writer) Write(p []byte) (int, error) {
	w.chunkPart.SetChunk(p)

	w.req.SetVerificationHeader(nil)

	if err := signature.SignServiceMessage(w.key, w.req); err != nil {
		return 0, fmt.Errorf("could not sign chunk request message: %w", err)
	}

	if err := w.stream.Write(w.req); err != nil {
		return 0, fmt.Errorf("could not send chunk request message: %w", err)
	}

	return len(p), nil
}

func (p *PutObjectParams) WithObject(v *object.Object) *PutObjectParams {
	if p != nil {
		p.obj = v
	}

	return p
}

func (p *PutObjectParams) Object() *object.Object {
	if p != nil {
		return p.obj
	}

	return nil
}

func (p *PutObjectParams) WithPayloadReader(v io.Reader) *PutObjectParams {
	if p != nil {
		p.r = v
	}

	return p
}

func (p *PutObjectParams) PayloadReader() io.Reader {
	if p != nil {
		return p.r
	}

	return nil
}

func (c *clientImpl) PutObject(ctx context.Context, p *PutObjectParams, opts ...CallOption) (*object.ID, error) {
	callOpts := c.defaultCallOptions()

	for i := range opts {
		if opts[i] != nil {
			opts[i](callOpts)
		}
	}

	// create request
	req := new(v2object.PutRequest)

	// initialize request body
	body := new(v2object.PutRequestBody)
	req.SetBody(body)

	v2Addr := new(v2refs.Address)
	v2Addr.SetObjectID(p.obj.ID().ToV2())
	v2Addr.SetContainerID(p.obj.ContainerID().ToV2())

	// set meta header
	meta := v2MetaHeaderFromOpts(callOpts)

	if err := c.attachV2SessionToken(callOpts, meta, v2SessionReqInfo{
		addr: v2Addr,
		verb: v2session.ObjectVerbPut,
	}); err != nil {
		return nil, fmt.Errorf("could not attach session token: %w", err)
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
	if err := signature.SignServiceMessage(callOpts.key, req); err != nil {
		return nil, fmt.Errorf("signing the request failed: %w", err)
	}

	// open stream
	resp := new(v2object.PutResponse)

	stream, err := rpcapi.PutObject(c.Raw(), resp, client.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("stream opening failed: %w", err)
	}

	// send init part
	err = stream.Write(req)
	if err != nil {
		return nil, fmt.Errorf("sending the initial message to stream failed: %w", err)
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
		key:       callOpts.key,
		chunkPart: chunkPart,
		req:       req,
		stream:    stream,
	}

	r := &putObjectV2Reader{r: rPayload}

	// copy payload from reader to stream writer
	_, err = io.CopyBuffer(w, r, make([]byte, chunkSize))
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("payload streaming failed: %w", err)
	}

	// close object stream and receive response from remote node
	err = stream.Close()
	if err != nil {
		return nil, fmt.Errorf("closing the stream failed: %w", err)
	}

	// verify response structure
	if err := signature.VerifyServiceMessage(resp); err != nil {
		return nil, fmt.Errorf("response verification failed: %w", err)
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

func (p *DeleteObjectParams) Address() *object.Address {
	if p != nil {
		return p.addr
	}

	return nil
}

// WithTombstoneAddressTarget sets target component to write tombstone address.
func (p *DeleteObjectParams) WithTombstoneAddressTarget(v ObjectAddressWriter) *DeleteObjectParams {
	if p != nil {
		p.tombTgt = v
	}

	return p
}

// TombstoneAddressTarget returns target component to write tombstone address.
func (p *DeleteObjectParams) TombstoneAddressTarget() ObjectAddressWriter {
	if p != nil {
		return p.tombTgt
	}

	return nil
}

// DeleteObject is a wrapper over Client.DeleteObject method
// that provides the ability to receive tombstone address
// without setting a target in the parameters.
func DeleteObject(ctx context.Context, c Client, p *DeleteObjectParams, opts ...CallOption) (*object.Address, error) {
	w := new(objectAddressWriter)

	err := c.DeleteObject(ctx, p.WithTombstoneAddressTarget(w), opts...)
	if err != nil {
		return nil, err
	}

	return w.addr, nil
}

// DeleteObject removes object by address.
//
// If target of tombstone address is not set, the address is ignored.
func (c *clientImpl) DeleteObject(ctx context.Context, p *DeleteObjectParams, opts ...CallOption) error {
	callOpts := c.defaultCallOptions()

	for i := range opts {
		if opts[i] != nil {
			opts[i](callOpts)
		}
	}

	// create request
	req := new(v2object.DeleteRequest)

	// initialize request body
	body := new(v2object.DeleteRequestBody)
	req.SetBody(body)

	// set meta header
	meta := v2MetaHeaderFromOpts(callOpts)

	if err := c.attachV2SessionToken(callOpts, meta, v2SessionReqInfo{
		addr: p.addr.ToV2(),
		verb: v2session.ObjectVerbDelete,
	}); err != nil {
		return fmt.Errorf("could not attach session token: %w", err)
	}

	req.SetMetaHeader(meta)

	// fill body fields
	body.SetAddress(p.addr.ToV2())

	// sign the request
	if err := signature.SignServiceMessage(callOpts.key, req); err != nil {
		return fmt.Errorf("signing the request failed: %w", err)
	}

	// send request
	resp, err := rpcapi.DeleteObject(c.Raw(), req, client.WithContext(ctx))
	if err != nil {
		return fmt.Errorf("sending the request failed: %w", err)
	}

	// verify response structure
	if err := signature.VerifyServiceMessage(resp); err != nil {
		return fmt.Errorf("response verification failed: %w", err)
	}

	if p.tombTgt != nil {
		p.tombTgt.SetAddress(object.NewAddressFromV2(resp.GetBody().GetTombstone()))
	}

	return nil
}

func (p *GetObjectParams) WithAddress(v *object.Address) *GetObjectParams {
	if p != nil {
		p.addr = v
	}

	return p
}

func (p *GetObjectParams) Address() *object.Address {
	if p != nil {
		return p.addr
	}

	return nil
}

func (p *GetObjectParams) WithPayloadWriter(w io.Writer) *GetObjectParams {
	if p != nil {
		p.w = w
	}

	return p
}

func (p *GetObjectParams) PayloadWriter() io.Writer {
	if p != nil {
		return p.w
	}

	return nil
}

func (p *GetObjectParams) WithRawFlag(v bool) *GetObjectParams {
	if p != nil {
		p.raw = v
	}

	return p
}

func (p *GetObjectParams) RawFlag() bool {
	if p != nil {
		return p.raw
	}

	return false
}

// ReaderHandler is a function over io.Reader.
type ReaderHandler func(io.Reader)

// WithPayloadReaderHandler sets handler of the payload reader.
//
// If provided, payload reader is composed after receiving the header.
// In this case payload writer set via WithPayloadWriter is ignored.
//
// Handler should not be nil.
func (p *GetObjectParams) WithPayloadReaderHandler(f ReaderHandler) *GetObjectParams {
	if p != nil {
		p.readerHandler = f
	}

	return p
}

// wrapper over the Object Get stream that provides io.Reader.
type objectPayloadReader struct {
	stream interface {
		Read(*v2object.GetResponse) error
	}

	resp v2object.GetResponse

	tail []byte
}

func (x *objectPayloadReader) Read(p []byte) (read int, err error) {
	// read remaining tail
	read = copy(p, x.tail)

	x.tail = x.tail[read:]

	if len(p)-read == 0 {
		return
	}

	// receive message from server stream
	err = x.stream.Read(&x.resp)
	if err != nil {
		if errors.Is(err, io.EOF) {
			err = io.EOF
			return
		}

		err = fmt.Errorf("reading the response failed: %w", err)
		return
	}

	// get chunk part message
	part := x.resp.GetBody().GetObjectPart()

	chunkPart, ok := part.(*v2object.GetObjectPartChunk)
	if !ok {
		err = errWrongMessageSeq
		return
	}

	// verify response structure
	if err = signature.VerifyServiceMessage(&x.resp); err != nil {
		err = fmt.Errorf("response verification failed: %w", err)
		return
	}

	// read new chunk
	chunk := chunkPart.GetChunk()

	tailOffset := copy(p[read:], chunk)

	read += tailOffset

	// save the tail
	x.tail = append(x.tail, chunk[tailOffset:]...)

	return
}

var errWrongMessageSeq = errors.New("incorrect message sequence")

func (c *clientImpl) GetObject(ctx context.Context, p *GetObjectParams, opts ...CallOption) (*object.Object, error) {
	callOpts := c.defaultCallOptions()

	for i := range opts {
		if opts[i] != nil {
			opts[i](callOpts)
		}
	}

	// create request
	req := new(v2object.GetRequest)

	// initialize request body
	body := new(v2object.GetRequestBody)
	req.SetBody(body)

	// set meta header
	meta := v2MetaHeaderFromOpts(callOpts)

	if err := c.attachV2SessionToken(callOpts, meta, v2SessionReqInfo{
		addr: p.addr.ToV2(),
		verb: v2session.ObjectVerbGet,
	}); err != nil {
		return nil, fmt.Errorf("could not attach session token: %w", err)
	}

	req.SetMetaHeader(meta)

	// fill body fields
	body.SetAddress(p.addr.ToV2())
	body.SetRaw(p.raw)

	// sign the request
	if err := signature.SignServiceMessage(callOpts.key, req); err != nil {
		return nil, fmt.Errorf("signing the request failed: %w", err)
	}

	// open stream
	stream, err := rpcapi.GetObject(c.Raw(), req, client.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("stream opening failed: %w", err)
	}

	var (
		headWas bool
		payload []byte
		obj     = new(v2object.Object)
		resp    = new(v2object.GetResponse)
	)

loop:
	for {
		// receive message from server stream
		err := stream.Read(resp)
		if err != nil {
			if errors.Is(err, io.EOF) {
				if !headWas {
					return nil, io.ErrUnexpectedEOF
				}

				break
			}

			return nil, fmt.Errorf("reading the response failed: %w", err)
		}

		// verify response structure
		if err := signature.VerifyServiceMessage(resp); err != nil {
			return nil, fmt.Errorf("response verification failed: %w", err)
		}

		switch v := resp.GetBody().GetObjectPart().(type) {
		default:
			return nil, fmt.Errorf("unexpected object part %T", v)
		case *v2object.GetObjectPartInit:
			if headWas {
				return nil, errWrongMessageSeq
			}

			headWas = true

			obj.SetObjectID(v.GetObjectID())
			obj.SetSignature(v.GetSignature())

			hdr := v.GetHeader()
			obj.SetHeader(hdr)

			if p.readerHandler != nil {
				p.readerHandler(&objectPayloadReader{
					stream: stream,
				})

				break loop
			}

			if p.w == nil {
				payload = make([]byte, 0, hdr.GetPayloadLength())
			}
		case *v2object.GetObjectPartChunk:
			if !headWas {
				return nil, errWrongMessageSeq
			}

			if p.w != nil {
				if _, err := p.w.Write(v.GetChunk()); err != nil {
					return nil, fmt.Errorf("could not write payload chunk: %w", err)
				}
			} else {
				payload = append(payload, v.GetChunk()...)
			}
		case *v2object.SplitInfo:
			si := object.NewSplitInfoFromV2(v)
			return nil, object.NewSplitInfoError(si)
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

func (p *ObjectHeaderParams) Address() *object.Address {
	if p != nil {
		return p.addr
	}

	return nil
}

func (p *ObjectHeaderParams) WithAllFields() *ObjectHeaderParams {
	if p != nil {
		p.short = false
	}

	return p
}

// AllFields return true if parameter set to return all header fields, returns
// false if parameter set to return only main fields of header.
func (p *ObjectHeaderParams) AllFields() bool {
	if p != nil {
		return !p.short
	}

	return false
}

func (p *ObjectHeaderParams) WithMainFields() *ObjectHeaderParams {
	if p != nil {
		p.short = true
	}

	return p
}

func (p *ObjectHeaderParams) WithRawFlag(v bool) *ObjectHeaderParams {
	if p != nil {
		p.raw = v
	}

	return p
}

func (p *ObjectHeaderParams) RawFlag() bool {
	if p != nil {
		return p.raw
	}

	return false
}

func (c *clientImpl) GetObjectHeader(ctx context.Context, p *ObjectHeaderParams, opts ...CallOption) (*object.Object, error) {
	callOpts := c.defaultCallOptions()

	for i := range opts {
		if opts[i] != nil {
			opts[i](callOpts)
		}
	}

	// create request
	req := new(v2object.HeadRequest)

	// initialize request body
	body := new(v2object.HeadRequestBody)
	req.SetBody(body)

	// set meta header
	meta := v2MetaHeaderFromOpts(callOpts)

	if err := c.attachV2SessionToken(callOpts, meta, v2SessionReqInfo{
		addr: p.addr.ToV2(),
		verb: v2session.ObjectVerbHead,
	}); err != nil {
		return nil, fmt.Errorf("could not attach session token: %w", err)
	}

	req.SetMetaHeader(meta)

	// fill body fields
	body.SetAddress(p.addr.ToV2())
	body.SetMainOnly(p.short)
	body.SetRaw(p.raw)

	// sign the request
	if err := signature.SignServiceMessage(callOpts.key, req); err != nil {
		return nil, fmt.Errorf("signing the request failed: %w", err)
	}

	// send Head request
	resp, err := rpcapi.HeadObject(c.Raw(), req, client.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("sending the request failed: %w", err)
	}

	// verify response structure
	if err := signature.VerifyServiceMessage(resp); err != nil {
		return nil, fmt.Errorf("response verification failed: %w", err)
	}

	var (
		hdr   *v2object.Header
		idSig *v2refs.Signature
	)

	switch v := resp.GetBody().GetHeaderPart().(type) {
	case nil:
		return nil, fmt.Errorf("unexpected header type %T", v)
	case *v2object.ShortHeader:
		if !p.short {
			return nil, fmt.Errorf("wrong header part type: expected %T, received %T",
				(*v2object.ShortHeader)(nil), (*v2object.HeaderWithSignature)(nil),
			)
		}

		h := v

		hdr = new(v2object.Header)
		hdr.SetPayloadLength(h.GetPayloadLength())
		hdr.SetVersion(h.GetVersion())
		hdr.SetOwnerID(h.GetOwnerID())
		hdr.SetObjectType(h.GetObjectType())
		hdr.SetCreationEpoch(h.GetCreationEpoch())
		hdr.SetPayloadHash(h.GetPayloadHash())
		hdr.SetHomomorphicHash(h.GetHomomorphicHash())
	case *v2object.HeaderWithSignature:
		if p.short {
			return nil, fmt.Errorf("wrong header part type: expected %T, received %T",
				(*v2object.HeaderWithSignature)(nil), (*v2object.ShortHeader)(nil),
			)
		}

		hdrWithSig := v
		if hdrWithSig == nil {
			return nil, errNilObjectPart
		}

		hdr = hdrWithSig.GetHeader()
		idSig = hdrWithSig.GetSignature()

		if err := signer.VerifyDataWithSource(
			signature.StableMarshalerWrapper{
				SM: p.addr.ObjectID().ToV2(),
			},
			func() (key, sig []byte) {
				return idSig.GetKey(), idSig.GetSign()
			},
		); err != nil {
			return nil, fmt.Errorf("incorrect object header signature: %w", err)
		}
	case *v2object.SplitInfo:
		si := object.NewSplitInfoFromV2(v)

		return nil, object.NewSplitInfoError(si)
	}

	obj := new(v2object.Object)
	obj.SetHeader(hdr)
	obj.SetSignature(idSig)

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

func (p *RangeDataParams) Address() *object.Address {
	if p != nil {
		return p.addr
	}

	return nil
}

func (p *RangeDataParams) WithRaw(v bool) *RangeDataParams {
	if p != nil {
		p.raw = v
	}

	return p
}

func (p *RangeDataParams) Raw() bool {
	if p != nil {
		return p.raw
	}

	return false
}

func (p *RangeDataParams) WithRange(v *object.Range) *RangeDataParams {
	if p != nil {
		p.r = v
	}

	return p
}

func (p *RangeDataParams) Range() *object.Range {
	if p != nil {
		return p.r
	}

	return nil
}

func (p *RangeDataParams) WithDataWriter(v io.Writer) *RangeDataParams {
	if p != nil {
		p.w = v
	}

	return p
}

func (p *RangeDataParams) DataWriter() io.Writer {
	if p != nil {
		return p.w
	}

	return nil
}

func (c *clientImpl) ObjectPayloadRangeData(ctx context.Context, p *RangeDataParams, opts ...CallOption) ([]byte, error) {
	callOpts := c.defaultCallOptions()

	for i := range opts {
		if opts[i] != nil {
			opts[i](callOpts)
		}
	}

	// create request
	req := new(v2object.GetRangeRequest)

	// initialize request body
	body := new(v2object.GetRangeRequestBody)
	req.SetBody(body)

	// set meta header
	meta := v2MetaHeaderFromOpts(callOpts)

	if err := c.attachV2SessionToken(callOpts, meta, v2SessionReqInfo{
		addr: p.addr.ToV2(),
		verb: v2session.ObjectVerbRange,
	}); err != nil {
		return nil, fmt.Errorf("could not attach session token: %w", err)
	}

	req.SetMetaHeader(meta)

	// fill body fields
	body.SetAddress(p.addr.ToV2())
	body.SetRange(p.r.ToV2())
	body.SetRaw(p.raw)

	// sign the request
	if err := signature.SignServiceMessage(callOpts.key, req); err != nil {
		return nil, fmt.Errorf("signing the request failed: %w", err)
	}

	// open stream
	stream, err := rpcapi.GetObjectRange(c.Raw(), req, client.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("could not create Get payload range stream: %w", err)
	}

	var payload []byte
	if p.w != nil {
		payload = make([]byte, 0, p.r.GetLength())
	}

	resp := new(v2object.GetRangeResponse)

	for {
		// receive message from server stream
		err := stream.Read(resp)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, fmt.Errorf("reading the response failed: %w", err)
		}

		// verify response structure
		if err := signature.VerifyServiceMessage(resp); err != nil {
			return nil, fmt.Errorf("could not verify %T: %w", resp, err)
		}

		switch v := resp.GetBody().GetRangePart().(type) {
		case nil:
			return nil, fmt.Errorf("unexpected range type %T", v)
		case *v2object.GetRangePartChunk:
			if p.w != nil {
				if _, err = p.w.Write(v.GetChunk()); err != nil {
					return nil, fmt.Errorf("could not write payload chunk: %w", err)
				}
			} else {
				payload = append(payload, v.GetChunk()...)
			}
		case *v2object.SplitInfo:
			si := object.NewSplitInfoFromV2(v)

			return nil, object.NewSplitInfoError(si)
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

func (p *RangeChecksumParams) Address() *object.Address {
	if p != nil {
		return p.addr
	}

	return nil
}

func (p *RangeChecksumParams) WithRangeList(rs ...*object.Range) *RangeChecksumParams {
	if p != nil {
		p.rs = rs
	}

	return p
}

func (p *RangeChecksumParams) RangeList() []*object.Range {
	if p != nil {
		return p.rs
	}

	return nil
}

func (p *RangeChecksumParams) WithSalt(v []byte) *RangeChecksumParams {
	if p != nil {
		p.salt = v
	}

	return p
}

func (p *RangeChecksumParams) Salt() []byte {
	if p != nil {
		return p.salt
	}

	return nil
}

func (p *RangeChecksumParams) withChecksumType(t checksumType) *RangeChecksumParams {
	if p != nil {
		p.typ = t
	}

	return p
}

func (c *clientImpl) ObjectPayloadRangeSHA256(ctx context.Context, p *RangeChecksumParams, opts ...CallOption) ([][sha256.Size]byte, error) {
	res, err := c.objectPayloadRangeHash(ctx, p.withChecksumType(checksumSHA256), opts...)
	if err != nil {
		return nil, err
	}

	return res.([][sha256.Size]byte), nil
}

func (c *clientImpl) ObjectPayloadRangeTZ(ctx context.Context, p *RangeChecksumParams, opts ...CallOption) ([][TZSize]byte, error) {
	res, err := c.objectPayloadRangeHash(ctx, p.withChecksumType(checksumTZ), opts...)
	if err != nil {
		return nil, err
	}

	return res.([][TZSize]byte), nil
}

func (c *clientImpl) objectPayloadRangeHash(ctx context.Context, p *RangeChecksumParams, opts ...CallOption) (interface{}, error) {
	callOpts := c.defaultCallOptions()

	for i := range opts {
		if opts[i] != nil {
			opts[i](callOpts)
		}
	}

	// create request
	req := new(v2object.GetRangeHashRequest)

	// initialize request body
	body := new(v2object.GetRangeHashRequestBody)
	req.SetBody(body)

	// set meta header
	meta := v2MetaHeaderFromOpts(callOpts)

	if err := c.attachV2SessionToken(callOpts, meta, v2SessionReqInfo{
		addr: p.addr.ToV2(),
		verb: v2session.ObjectVerbRangeHash,
	}); err != nil {
		return nil, fmt.Errorf("could not attach session token: %w", err)
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
	if err := signature.SignServiceMessage(callOpts.key, req); err != nil {
		return nil, fmt.Errorf("signing the request failed: %w", err)
	}

	// send request
	resp, err := rpcapi.HashObjectRange(c.Raw(), req, client.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("sending the request failed: %w", err)
	}

	// verify response structure
	if err := signature.VerifyServiceMessage(resp); err != nil {
		return nil, fmt.Errorf("response verification failed: %w", err)
	}

	respBody := resp.GetBody()
	respType := respBody.GetType()
	respHashes := respBody.GetHashList()

	if t := p.typ.toV2(); respType != t {
		return nil, fmt.Errorf("invalid checksum type: expected %v, received %v", t, respType)
	} else if reqLn, respLn := len(rsV2), len(respHashes); reqLn != respLn {
		return nil, fmt.Errorf("wrong checksum number: expected %d, received %d", reqLn, respLn)
	}

	var res interface{}

	switch p.typ {
	case checksumSHA256:
		r := make([][sha256.Size]byte, 0, len(respHashes))

		for i := range respHashes {
			if ln := len(respHashes[i]); ln != sha256.Size {
				return nil, fmt.Errorf("invalid checksum length: expected %d, received %d", sha256.Size, ln)
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
				return nil, fmt.Errorf("invalid checksum length: expected %d, received %d", TZSize, ln)
			}

			cs := [TZSize]byte{}
			copy(cs[:], respHashes[i])

			r = append(r, cs)
		}

		res = r
	}

	return res, nil
}

func (p *SearchObjectParams) WithContainerID(v *cid.ID) *SearchObjectParams {
	if p != nil {
		p.cid = v
	}

	return p
}

func (p *SearchObjectParams) ContainerID() *cid.ID {
	if p != nil {
		return p.cid
	}

	return nil
}

func (p *SearchObjectParams) WithSearchFilters(v object.SearchFilters) *SearchObjectParams {
	if p != nil {
		p.filters = v
	}

	return p
}

func (p *SearchObjectParams) SearchFilters() object.SearchFilters {
	if p != nil {
		return p.filters
	}

	return nil
}

func (c *clientImpl) SearchObject(ctx context.Context, p *SearchObjectParams, opts ...CallOption) ([]*object.ID, error) {
	callOpts := c.defaultCallOptions()

	for i := range opts {
		if opts[i] != nil {
			opts[i](callOpts)
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

	if err := c.attachV2SessionToken(callOpts, meta, v2SessionReqInfo{
		addr: v2Addr,
		verb: v2session.ObjectVerbSearch,
	}); err != nil {
		return nil, fmt.Errorf("could not attach session token: %w", err)
	}

	req.SetMetaHeader(meta)

	// fill body fields
	body.SetContainerID(v2Addr.GetContainerID())
	body.SetVersion(searchQueryVersion)
	body.SetFilters(p.filters.ToV2())

	// sign the request
	if err := signature.SignServiceMessage(callOpts.key, req); err != nil {
		return nil, fmt.Errorf("signing the request failed: %w", err)
	}

	// create search stream
	stream, err := rpcapi.SearchObjects(c.Raw(), req, client.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("stream opening failed: %w", err)
	}

	var (
		searchResult []*object.ID
		resp         = new(v2object.SearchResponse)
	)

	for {
		// receive message from server stream
		err := stream.Read(resp)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, fmt.Errorf("reading the response failed: %w", err)
		}

		// verify response structure
		if err := signature.VerifyServiceMessage(resp); err != nil {
			return nil, fmt.Errorf("could not verify %T: %w", resp, err)
		}

		chunk := resp.GetBody().GetIDList()
		for i := range chunk {
			searchResult = append(searchResult, object.NewIDFromV2(chunk[i]))
		}
	}

	return searchResult, nil
}

func (c *clientImpl) attachV2SessionToken(opts *callOptions, hdr *v2session.RequestMetaHeader, info v2SessionReqInfo) error {
	if opts.session == nil {
		return nil
	}

	// Do not resign already prepared session token
	if opts.session.Signature() != nil {
		hdr.SetSessionToken(opts.session.ToV2())
		return nil
	}

	opCtx := new(v2session.ObjectSessionContext)
	opCtx.SetAddress(info.addr)
	opCtx.SetVerb(info.verb)

	lt := new(v2session.TokenLifetime)
	lt.SetIat(info.iat)
	lt.SetNbf(info.nbf)
	lt.SetExp(info.exp)

	body := new(v2session.SessionTokenBody)
	body.SetID(opts.session.ID())
	body.SetOwnerID(opts.session.OwnerID().ToV2())
	body.SetSessionKey(opts.session.SessionKey())
	body.SetContext(opCtx)
	body.SetLifetime(lt)

	token := new(v2session.SessionToken)
	token.SetBody(body)

	signWrapper := signature.StableMarshalerWrapper{SM: token.GetBody()}

	err := signer.SignDataWithHandler(opts.key, signWrapper, func(key []byte, sig []byte) {
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
