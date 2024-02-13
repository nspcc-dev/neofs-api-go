package acl

import (
	"testing"

	acl "github.com/nspcc-dev/neofs-api-go/v2/acl/grpc"
	"github.com/stretchr/testify/require"
)

func TestMatchTypeToGRPCField(t *testing.T) {
	require.Equal(t, MatchTypeUnknown, MatchTypeFromGRPCField(acl.MatchType_NUM_LE+1))
	require.Equal(t, acl.MatchType_MATCH_TYPE_UNSPECIFIED, MatchTypeToGRPCField(matchTypeLast))

	for _, tc := range []struct {
		m MatchType
		g acl.MatchType
	}{
		{MatchTypeUnknown, acl.MatchType_MATCH_TYPE_UNSPECIFIED},
		{MatchTypeStringEqual, acl.MatchType_STRING_EQUAL},
		{MatchTypeStringNotEqual, acl.MatchType_STRING_NOT_EQUAL},
		{MatchTypeNotPresent, acl.MatchType_NOT_PRESENT},
		{MatchTypeNumGT, acl.MatchType_NUM_GT},
		{MatchTypeNumGE, acl.MatchType_NUM_GE},
		{MatchTypeNumLT, acl.MatchType_NUM_LT},
		{MatchTypeNumLE, acl.MatchType_NUM_LE},
	} {
		require.Equal(t, tc.g, MatchTypeToGRPCField(tc.m), tc)
		require.Equal(t, tc.m, MatchTypeFromGRPCField(tc.g), tc)
	}
}
