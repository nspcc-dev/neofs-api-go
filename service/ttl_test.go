package service

import (
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

type mockedRequest struct {
	msg  string
	ttl  uint32
	name string
	role NodeRole
	code codes.Code
}

func (m *mockedRequest) SetTTL(v uint32) { m.ttl = v }
func (m mockedRequest) GetTTL() uint32   { return m.ttl }

func TestCheckTTLRequest(t *testing.T) {
	tests := []mockedRequest{
		{
			ttl:  NonForwardingTTL,
			role: InnerRingNode,
			name: "direct to ir node",
		},
		{
			ttl:  NonForwardingTTL,
			role: StorageNode,
			code: codes.InvalidArgument,
			msg:  ErrIncorrectTTL.Error(),
			name: "direct to storage node",
		},
		{
			ttl:  ZeroTTL,
			role: StorageNode,
			msg:  ErrZeroTTL.Error(),
			code: codes.InvalidArgument,
			name: "zero ttl",
		},
		{
			ttl:  SingleForwardingTTL,
			role: InnerRingNode,
			name: "default to ir node",
		},
		{
			ttl:  SingleForwardingTTL,
			role: StorageNode,
			name: "default to storage node",
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			before := tt.ttl
			err := CheckTTLRequest(&tt, tt.role)
			if tt.msg != "" {
				require.Errorf(t, err, tt.msg)

				state, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, state.Code(), tt.code)
				require.Equal(t, state.Message(), tt.msg)
			} else {
				require.NoError(t, err)
				require.NotEqualf(t, before, tt.ttl, "ttl should be changed: %d vs %d", before, tt.ttl)
			}
		})
	}
}
