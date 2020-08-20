package client

import (
	"github.com/pkg/errors"
)

type Protocol uint32

const (
	_ Protocol = iota
	ProtoGRPC
)

var ErrProtoUnsupported = errors.New("unsupported protocol")

func (p Protocol) String() string {
	switch p {
	case ProtoGRPC:
		return "GRPC"
	default:
		return "UNKNOWN"
	}
}
