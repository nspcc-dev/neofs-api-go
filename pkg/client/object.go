package client

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"

	cryptoalgo "github.com/nspcc-dev/neofs-api-go/crypto/algo"
	neofsecdsa "github.com/nspcc-dev/neofs-api-go/crypto/ecdsa"
	cid "github.com/nspcc-dev/neofs-api-go/pkg/container/id"
	"github.com/nspcc-dev/neofs-api-go/pkg/object"
	apicrypto "github.com/nspcc-dev/neofs-api-go/v2/crypto"
	v2object "github.com/nspcc-dev/neofs-api-go/v2/object"
	v2refs "github.com/nspcc-dev/neofs-api-go/v2/refs"
	rpcapi "github.com/nspcc-dev/neofs-api-go/v2/rpc"
	v2session "github.com/nspcc-dev/neofs-api-go/v2/session"
	"github.com/nspcc-dev/neofs-api-go/v2/signature"
)

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
	key ecdsa.PrivateKey

	chunkPart *v2object.PutObjectPartChunk

	req v2object.PutRequest

	stream rpcapi.PutObjectStream
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

	if err := signature.SignServiceMessage(neofsecdsa.Signer(w.key), &w.req); err != nil {
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

func (x Client) PutObject(ctx context.Context, p *PutObjectParams, opts ...CallOption) (*object.ID, error) {
	callOpts := defaultCallOptions()

	for i := range opts {
		if opts[i] != nil {
			opts[i](callOpts)
		}
	}

	// create request
	var req v2object.PutRequest

	// initialize request body
	body := new(v2object.PutRequestBody)
	req.SetBody(body)

	v2Addr := new(v2refs.Address)
	v2Addr.SetObjectID(p.obj.ID().ToV2())
	v2Addr.SetContainerID(p.obj.ContainerID().ToV2())

	// set meta header
	meta := v2MetaHeaderFromOpts(callOpts)

	if err := x.attachV2SessionToken(callOpts, meta, v2SessionReqInfo{
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
	if err := signature.SignServiceMessage(neofsecdsa.Signer(callOpts.key), &req); err != nil {
		return nil, fmt.Errorf("signing the request failed: %w", err)
	}

	// open stream
	var prm rpcapi.PutObjectPrm

	var res rpcapi.PutObjectRes

	err := rpcapi.PutObject(ctx, x.c, prm, &res)
	if err != nil {
		return nil, fmt.Errorf("transport error: %w", err)
	}

	stream := res.Stream()

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
	err = stream.CloseSend()
	if err != nil {
		return nil, fmt.Errorf("closing the stream failed: %w", err)
	}

	resp := stream.Response()

	// verify response structure
	if err := signature.VerifyServiceMessage(&resp); err != nil {
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
func (x Client) DeleteObject(ctx context.Context, p *DeleteObjectParams, opts ...CallOption) error {
	callOpts := defaultCallOptions()

	for i := range opts {
		if opts[i] != nil {
			opts[i](callOpts)
		}
	}

	// create request
	var req v2object.DeleteRequest

	// initialize request body
	body := new(v2object.DeleteRequestBody)
	req.SetBody(body)

	// set meta header
	meta := v2MetaHeaderFromOpts(callOpts)

	if err := x.attachV2SessionToken(callOpts, meta, v2SessionReqInfo{
		addr: p.addr.ToV2(),
		verb: v2session.ObjectVerbDelete,
	}); err != nil {
		return fmt.Errorf("could not attach session token: %w", err)
	}

	req.SetMetaHeader(meta)

	// fill body fields
	body.SetAddress(p.addr.ToV2())

	// sign the request
	if err := signature.SignServiceMessage(neofsecdsa.Signer(callOpts.key), &req); err != nil {
		return fmt.Errorf("signing the request failed: %w", err)
	}

	// send request
	var prm rpcapi.DeleteObjectPrm

	prm.SetRequest(req)

	var res rpcapi.DeleteObjectRes

	err := rpcapi.DeleteObject(ctx, x.c, prm, &res)
	if err != nil {
		return fmt.Errorf("transport error: %w", err)
	}

	resp := res.Response()

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

func (x Client) GetObject(ctx context.Context, p *GetObjectParams, opts ...CallOption) (*object.Object, error) {
	callOpts := defaultCallOptions()

	for i := range opts {
		if opts[i] != nil {
			opts[i](callOpts)
		}
	}

	// create request
	var req v2object.GetRequest

	// initialize request body
	body := new(v2object.GetRequestBody)
	req.SetBody(body)

	// set meta header
	meta := v2MetaHeaderFromOpts(callOpts)

	if err := x.attachV2SessionToken(callOpts, meta, v2SessionReqInfo{
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
	if err := signature.SignServiceMessage(neofsecdsa.Signer(callOpts.key), &req); err != nil {
		return nil, fmt.Errorf("signing the request failed: %w", err)
	}

	// open stream
	var prm rpcapi.GetObjectPrm

	prm.SetRequest(req)

	var res rpcapi.GetObjectRes

	err := rpcapi.GetObject(ctx, x.c, prm, &res)
	if err != nil {
		return nil, fmt.Errorf("transport error: %w", err)
	}

	stream := res.Stream()

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
					stream: &stream,
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

func (x Client) GetObjectHeader(ctx context.Context, p *ObjectHeaderParams, opts ...CallOption) (*object.Object, error) {
	callOpts := defaultCallOptions()

	for i := range opts {
		if opts[i] != nil {
			opts[i](callOpts)
		}
	}

	// create request
	var req v2object.HeadRequest

	// initialize request body
	body := new(v2object.HeadRequestBody)
	req.SetBody(body)

	// set meta header
	meta := v2MetaHeaderFromOpts(callOpts)

	if err := x.attachV2SessionToken(callOpts, meta, v2SessionReqInfo{
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
	if err := signature.SignServiceMessage(neofsecdsa.Signer(callOpts.key), &req); err != nil {
		return nil, fmt.Errorf("signing the request failed: %w", err)
	}

	// send Head request
	var prm rpcapi.HeadObjectPrm

	prm.SetRequest(req)

	var res rpcapi.HeadObjectRes

	err := rpcapi.HeadObject(ctx, x.c, prm, &res)
	if err != nil {
		return nil, fmt.Errorf("transport error: %w", err)
	}

	resp := res.Response()

	// verify response structure
	if err := signature.VerifyServiceMessage(&resp); err != nil {
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

		key, err := cryptoalgo.UnmarshalKey(cryptoalgo.ECDSA, idSig.GetKey())
		if err != nil {
			return nil, err
		}

		var vPrm apicrypto.VerifyPrm

		vPrm.SetProtoMarshaler(signature.StableMarshalerCrypto(p.addr.ObjectID().ToV2()))
		vPrm.SetSignature(idSig.GetSign())

		if !apicrypto.Verify(key, vPrm) {
			return nil, errors.New("incorrect object header signature")
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

func (x Client) ObjectPayloadRangeData(ctx context.Context, p *RangeDataParams, opts ...CallOption) ([]byte, error) {
	callOpts := defaultCallOptions()

	for i := range opts {
		if opts[i] != nil {
			opts[i](callOpts)
		}
	}

	// create request
	var req v2object.GetRangeRequest

	// initialize request body
	body := new(v2object.GetRangeRequestBody)
	req.SetBody(body)

	// set meta header
	meta := v2MetaHeaderFromOpts(callOpts)

	if err := x.attachV2SessionToken(callOpts, meta, v2SessionReqInfo{
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
	if err := signature.SignServiceMessage(neofsecdsa.Signer(callOpts.key), &req); err != nil {
		return nil, fmt.Errorf("signing the request failed: %w", err)
	}

	// open stream
	var prm rpcapi.GetObjectRangePrm

	prm.SetRequest(req)

	var res rpcapi.GetObjectRangeRes

	err := rpcapi.GetObjectRange(ctx, x.c, prm, &res)
	if err != nil {
		return nil, fmt.Errorf("transport error: %w", err)
	}

	stream := res.Stream()

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

func (x Client) ObjectPayloadRangeSHA256(ctx context.Context, p *RangeChecksumParams, opts ...CallOption) ([][sha256.Size]byte, error) {
	res, err := x.objectPayloadRangeHash(ctx, p.withChecksumType(checksumSHA256), opts...)
	if err != nil {
		return nil, err
	}

	return res.([][sha256.Size]byte), nil
}

func (x Client) ObjectPayloadRangeTZ(ctx context.Context, p *RangeChecksumParams, opts ...CallOption) ([][TZSize]byte, error) {
	res, err := x.objectPayloadRangeHash(ctx, p.withChecksumType(checksumTZ), opts...)
	if err != nil {
		return nil, err
	}

	return res.([][TZSize]byte), nil
}

func (x Client) objectPayloadRangeHash(ctx context.Context, p *RangeChecksumParams, opts ...CallOption) (interface{}, error) {
	callOpts := defaultCallOptions()

	for i := range opts {
		if opts[i] != nil {
			opts[i](callOpts)
		}
	}

	// create request
	var req v2object.GetRangeHashRequest

	// initialize request body
	body := new(v2object.GetRangeHashRequestBody)
	req.SetBody(body)

	// set meta header
	meta := v2MetaHeaderFromOpts(callOpts)

	if err := x.attachV2SessionToken(callOpts, meta, v2SessionReqInfo{
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
	if err := signature.SignServiceMessage(neofsecdsa.Signer(callOpts.key), &req); err != nil {
		return nil, fmt.Errorf("signing the request failed: %w", err)
	}

	// send request
	var prm rpcapi.HashObjectRangePrm

	prm.SetRequest(req)

	var hres rpcapi.HashObjectRangeRes

	err := rpcapi.HashObjectRange(ctx, x.c, prm, &hres)
	if err != nil {
		return nil, fmt.Errorf("transport error: %w", err)
	}

	resp := hres.Response()

	// verify response structure
	if err := signature.VerifyServiceMessage(&resp); err != nil {
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

func (x Client) SearchObject(ctx context.Context, p *SearchObjectParams, opts ...CallOption) ([]*object.ID, error) {
	callOpts := defaultCallOptions()

	for i := range opts {
		if opts[i] != nil {
			opts[i](callOpts)
		}
	}

	// create request
	var req v2object.SearchRequest

	// initialize request body
	body := new(v2object.SearchRequestBody)
	req.SetBody(body)

	v2Addr := new(v2refs.Address)
	v2Addr.SetContainerID(p.cid.ToV2())

	// set meta header
	meta := v2MetaHeaderFromOpts(callOpts)

	if err := x.attachV2SessionToken(callOpts, meta, v2SessionReqInfo{
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
	if err := signature.SignServiceMessage(neofsecdsa.Signer(callOpts.key), &req); err != nil {
		return nil, fmt.Errorf("signing the request failed: %w", err)
	}

	// create search stream
	var prm rpcapi.SearchObjectsPrm

	prm.SetRequest(req)

	var res rpcapi.SearchObjectsRes

	err := rpcapi.SearchObjects(ctx, x.c, prm, &res)
	if err != nil {
		return nil, fmt.Errorf("transport error: %w", err)
	}

	stream := res.Stream()

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

func (x Client) attachV2SessionToken(opts *callOptions, hdr *v2session.RequestMetaHeader, info v2SessionReqInfo) error {
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

	var (
		p   apicrypto.SignPrm
		sig = new(v2refs.Signature)
	)

	p.SetProtoMarshaler(signature.StableMarshalerCrypto(token.GetBody()))
	p.SetTargetSignature(sig)

	err := apicrypto.Sign(neofsecdsa.Signer(opts.key), p)
	if err != nil {
		return err
	}

	hdr.SetSessionToken(token)

	return nil
}
