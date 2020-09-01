package object

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"

	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	"github.com/nspcc-dev/neofs-api-go/util/signature"
	"github.com/nspcc-dev/neofs-api-go/v2/object"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	signatureV2 "github.com/nspcc-dev/neofs-api-go/v2/signature"
	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/pkg/errors"
)

// Object represents NeoFS object that provides
// a convenient interface for working in isolation
// from the internal structure of an object.
//
// Object allows to work with the object in read-only
// mode as a reflection of the immutability of objects
// in the system.
type Object struct {
	rwObject
}

type rwObject struct {
	fin bool

	id *ID

	key, sig []byte

	cid *container.ID

	ownerID *owner.ID

	payloadChecksum *refs.Checksum

	// TODO: add other fields
}

// Verify checks if object structure is correct.
func (o *Object) Verify() error {
	if o == nil {
		return nil
	}

	hdr := o.v2Header()

	data, err := hdr.StableMarshal(nil)
	if err != nil {
		return errors.Wrap(err, "could not marshal header")
	}

	hdrChecksum := sha256.Sum256(data)

	if !bytes.Equal(hdrChecksum[:], o.id.val) {
		return errors.New("invalid object identifier")
	}

	if err := signature.VerifyDataWithSource(
		signatureV2.StableMarshalerWrapper{
			SM: o.id.ToV2(),
		},
		func() (key, sig []byte) {
			return o.key, o.sig
		},
	); err != nil {
		return errors.Wrap(err, "invalid object ID signature")
	}

	return nil
}

// ToV2 converts object to v2 Object message.
func (o *Object) ToV2() *object.Object {
	obj, _ := o.rwObject.ToV2(nil)

	return obj
}

func (o *rwObject) v2Header() *object.Header {
	hV2 := new(object.Header)
	hV2.SetContainerID(o.cid.ToV2())
	hV2.SetOwnerID(o.ownerID.ToV2())
	hV2.SetPayloadHash(o.payloadChecksum)
	// TODO: set other fields

	return hV2
}

func (o *rwObject) complete(key *ecdsa.PrivateKey) (*object.Header, error) {
	hdr := o.v2Header()

	hdrData, err := hdr.StableMarshal(nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not marshal header")
	}

	o.id = new(ID)
	o.id.SetSHA256(sha256.Sum256(hdrData))

	if err := signature.SignDataWithHandler(
		key,
		signatureV2.StableMarshalerWrapper{
			SM: o.id.ToV2(),
		},
		func(key []byte, sig []byte) {
			o.key, o.sig = key, sig
		},
	); err != nil {
		return nil, errors.Wrap(err, "could sign object identifier")
	}

	o.fin = true

	return hdr, nil
}

// ToV2 calculates object identifier, signs structure and converts
// it to v2 Object message.
func (o *rwObject) ToV2(key *ecdsa.PrivateKey) (*object.Object, error) {
	if o == nil {
		return nil, nil
	}

	var (
		hdr *object.Header
		err error
	)

	if !o.fin {
		if key == nil {
			return nil, errors.Wrap(crypto.ErrEmptyPrivateKey, "could complete the object")
		}

		if hdr, err = o.complete(key); err != nil {
			return nil, errors.Wrapf(err, "could not complete the object")
		}
	} else {
		hdr = o.v2Header()
	}

	obj := new(object.Object)
	obj.SetObjectID(o.id.ToV2())
	obj.SetHeader(hdr)

	sig := new(refs.Signature)
	sig.SetKey(o.key)
	sig.SetSign(o.sig)
	obj.SetSignature(sig)

	return obj, nil
}

// FromV2 converts v2 Object message to Object instance.
//
// Returns any error encountered which prevented the
// recovery of object data from the message.
func FromV2(oV2 *object.Object) (*Object, error) {
	if oV2 == nil {
		return nil, nil
	}

	id, err := IDFromV2(oV2.GetObjectID())
	if err != nil {
		return nil, errors.Wrap(err, "could not convert object ID")
	}

	hdr := oV2.GetHeader()

	ownerID, err := owner.IDFromV2(hdr.GetOwnerID())
	if err != nil {
		return nil, errors.Wrap(err, "could not convert owner ID")
	}

	cid, err := container.IDFromV2(hdr.GetContainerID())
	if err != nil {
		return nil, errors.Wrap(err, "could not convert container ID")
	}

	// TODO: convert other fields

	sig := oV2.GetSignature()

	return &Object{
		rwObject: rwObject{
			fin:             true,
			id:              id,
			key:             sig.GetKey(),
			sig:             sig.GetSign(),
			cid:             cid,
			ownerID:         ownerID,
			payloadChecksum: hdr.GetPayloadHash(),
		},
	}, nil
}
