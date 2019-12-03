package storagegroup

import (
	"context"

	"github.com/nspcc-dev/neofs-proto/refs"
)

type (
	// Store is a interface for storing users storage group
	Store interface {
		Lister
		Receiver
		Receptacle
	}

	// Lister defines list function that returns all storage groups
	// created for the passed container
	Lister interface {
		List(ctx context.Context, cid refs.CID) ([]refs.SGID, error)
	}

	// Receiver defines get function that returns asked storage group
	Receiver interface {
		Get(ctx context.Context, cid refs.CID, sgid refs.SGID) (Provider, error)
	}

	// Receptacle defines put function that places storage group in the
	// store.
	Receptacle interface {
		Put(ctx context.Context, sg Provider) error
	}

	// InfoReceiver defines GetSGInfo function that returns storage group
	// that contains passed object ids.
	InfoReceiver interface {
		GetSGInfo(ctx context.Context, cid refs.CID, group []refs.ObjectID) (*StorageGroup, error)
	}
)
