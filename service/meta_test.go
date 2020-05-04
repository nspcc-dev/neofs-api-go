package service

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type mockedRequest struct {
	msg     string
	name    string
	code    codes.Code
	handler TTLCondition
	RequestMetaHeader
}

func TestMetaRequest(t *testing.T) {
	tests := []mockedRequest{
		{
			name:              "direct to ir node",
			handler:           IRNonForwarding(InnerRingNode),
			RequestMetaHeader: RequestMetaHeader{TTL: NonForwardingTTL},
		},
		{
			code:              codes.InvalidArgument,
			msg:               ErrInvalidTTL.Error(),
			name:              "direct to storage node",
			handler:           IRNonForwarding(StorageNode),
			RequestMetaHeader: RequestMetaHeader{TTL: NonForwardingTTL},
		},
		{
			msg:               ErrInvalidTTL.Error(),
			code:              codes.InvalidArgument,
			name:              "zero ttl",
			handler:           IRNonForwarding(StorageNode),
			RequestMetaHeader: RequestMetaHeader{TTL: ZeroTTL},
		},
		{
			name:              "default to ir node",
			handler:           IRNonForwarding(InnerRingNode),
			RequestMetaHeader: RequestMetaHeader{TTL: SingleForwardingTTL},
		},
		{
			name:              "default to storage node",
			handler:           IRNonForwarding(StorageNode),
			RequestMetaHeader: RequestMetaHeader{TTL: SingleForwardingTTL},
		},
		{
			msg:               "not found",
			code:              codes.NotFound,
			name:              "custom status error",
			RequestMetaHeader: RequestMetaHeader{TTL: SingleForwardingTTL},
			handler:           func(_ uint32) error { return status.Error(codes.NotFound, "not found") },
		},
		{
			msg:               "not found",
			code:              codes.NotFound,
			name:              "custom wrapped status error",
			RequestMetaHeader: RequestMetaHeader{TTL: SingleForwardingTTL},
			handler: func(_ uint32) error {
				err := status.Error(codes.NotFound, "not found")
				err = errors.Wrap(err, "some error context")
				err = errors.Wrap(err, "another error context")
				return err
			},
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			before := tt.GetTTL()
			err := ProcessRequestTTL(&tt, tt.handler)
			if tt.msg != "" {
				require.Errorf(t, err, tt.msg)

				state, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tt.code, state.Code())
				require.Equal(t, tt.msg, state.Message())
			} else {
				require.NoError(t, err)
				require.NotEqualf(t, before, tt.GetTTL(), "ttl should be changed: %d vs %d", before, tt.GetTTL())
			}
		})
	}
}

func TestRequestMetaHeader_SetEpoch(t *testing.T) {
	m := new(ResponseMetaHeader)
	epoch := uint64(3)
	m.SetEpoch(epoch)
	require.Equal(t, epoch, m.GetEpoch())
}

func TestRequestMetaHeader_SetVersion(t *testing.T) {
	m := new(ResponseMetaHeader)
	version := uint32(3)
	m.SetVersion(version)
	require.Equal(t, version, m.GetVersion())
}

func TestRequestMetaHeader_SetRaw(t *testing.T) {
	m := new(RequestMetaHeader)

	m.SetRaw(true)
	require.True(t, m.GetRaw())

	m.SetRaw(false)
	require.False(t, m.GetRaw())
}
