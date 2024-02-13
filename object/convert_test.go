package object

import (
	"testing"

	object "github.com/nspcc-dev/neofs-api-go/v2/object/grpc"
	"github.com/stretchr/testify/require"
)

func TestMatchTypeToGRPCField(t *testing.T) {
	for _, tc := range []struct {
		m MatchType
		g object.MatchType
	}{
		{MatchUnknown, object.MatchType_MATCH_TYPE_UNSPECIFIED},
		{MatchStringEqual, object.MatchType_STRING_EQUAL},
		{MatchStringNotEqual, object.MatchType_STRING_NOT_EQUAL},
		{MatchNotPresent, object.MatchType_NOT_PRESENT},
		{MatchCommonPrefix, object.MatchType_COMMON_PREFIX},
		{MatchNumGT, object.MatchType_NUM_GT},
		{MatchNumGE, object.MatchType_NUM_GE},
		{MatchNumLT, object.MatchType_NUM_LT},
		{MatchNumLE, object.MatchType_NUM_LE},
	} {
		require.Equal(t, tc.g, MatchTypeToGRPCField(tc.m), tc)
		require.Equal(t, tc.m, MatchTypeFromGRPCField(tc.g), tc)
	}
}
