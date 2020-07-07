package acl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEACLFilterWrapper(t *testing.T) {
	s := WrapFilterInfo(nil)

	mt := stringEqual
	s.SetMatchType(mt)
	require.Equal(t, mt, s.MatchType())

	ht := HdrTypeObjUsr
	s.SetHeaderType(ht)
	require.Equal(t, ht, s.HeaderType())

	n := "name"
	s.SetName(n)
	require.Equal(t, n, s.Name())

	v := "value"
	s.SetValue(v)
	require.Equal(t, v, s.Value())
}

func TestEACLTargetWrapper(t *testing.T) {
	s := WrapEACLTarget(nil)

	target := Target(10)
	s.SetTarget(target)
	require.Equal(t, target, s.Target())

	keys := [][]byte{
		{1, 2, 3},
		{4, 5, 6},
	}
	s.SetKeyList(keys)
	require.Equal(t, keys, s.KeyList())
}

func TestEACLRecordWrapper(t *testing.T) {
	s := WrapEACLRecord(nil)

	action := ActionAllow
	s.SetAction(action)
	require.Equal(t, action, s.Action())

	opType := OperationType(5)
	s.SetOperationType(opType)
	require.Equal(t, opType, s.OperationType())

	f1Name := "name1"
	f1 := WrapFilterInfo(nil)
	f1.SetName(f1Name)

	f2Name := "name2"
	f2 := WrapFilterInfo(nil)
	f2.SetName(f2Name)

	s.SetHeaderFilters([]HeaderFilter{f1, f2})

	filters := s.HeaderFilters()
	require.Len(t, filters, 2)
	require.Equal(t, f1Name, filters[0].Name())
	require.Equal(t, f2Name, filters[1].Name())

	target1 := Target(1)
	t1 := WrapEACLTarget(nil)
	t1.SetTarget(target1)

	target2 := Target(2)
	t2 := WrapEACLTarget(nil)
	t2.SetTarget(target2)

	s.SetTargetList([]ExtendedACLTarget{t1, t2})

	targets := s.TargetList()
	require.Len(t, targets, 2)
	require.Equal(t, target1, targets[0].Target())
	require.Equal(t, target2, targets[1].Target())
}

func TestEACLTableWrapper(t *testing.T) {
	s := WrapEACLTable(nil)

	action1 := ExtendedACLAction(1)
	r1 := WrapEACLRecord(nil)
	r1.SetAction(action1)

	action2 := ExtendedACLAction(2)
	r2 := WrapEACLRecord(nil)
	r2.SetAction(action2)

	s.SetRecords([]ExtendedACLRecord{r1, r2})

	records := s.Records()
	require.Len(t, records, 2)
	require.Equal(t, action1, records[0].Action())
	require.Equal(t, action2, records[1].Action())

	data, err := s.MarshalBinary()
	require.NoError(t, err)

	s2 := WrapEACLTable(nil)
	require.NoError(t, s2.UnmarshalBinary(data))

	records1 := s.Records()
	records2 := s2.Records()
	require.Len(t, records1, len(records2))

	for i := range records1 {
		require.Equal(t, records1[i].Action(), records2[i].Action())
		require.Equal(t, records1[i].OperationType(), records2[i].OperationType())

		targets1 := records1[i].TargetList()
		targets2 := records2[i].TargetList()
		require.Len(t, targets1, len(targets2))

		for j := range targets1 {
			require.Equal(t, targets1[j].Target(), targets2[j].Target())
			require.Equal(t, targets1[j].KeyList(), targets2[j].KeyList())
		}

		filters1 := records1[i].HeaderFilters()
		filters2 := records2[i].HeaderFilters()
		require.Len(t, filters1, len(filters2))

		for j := range filters1 {
			require.Equal(t, filters1[j].MatchType(), filters2[j].MatchType())
			require.Equal(t, filters1[j].HeaderType(), filters2[j].HeaderType())
			require.Equal(t, filters1[j].Name(), filters2[j].Name())
			require.Equal(t, filters1[j].Value(), filters2[j].Value())
			require.Equal(t, filters1[j].Value(), filters2[j].Value())
		}
	}
}
