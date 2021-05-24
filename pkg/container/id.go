package container

import (
	cid "github.com/nspcc-dev/neofs-api-go/pkg/container/id"
)

// ID represents v2-compatible container identifier.
//
// Deprecated: use cid.ID instead.
type ID = cid.ID

// NewIDFromV2 wraps v2 ContainerID message to ID.
//
// Deprecated: use cid.NewFromV2 instead.
var NewIDFromV2 = cid.NewFromV2

// NewID creates and initializes blank ID.
//
// Works similar to NewIDFromV2(new(ContainerID)).
//
// Deprecated: use cid.New instead.
var NewID = cid.New
