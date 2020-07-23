package eacl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEACLFilterWrapper(t *testing.T) {
	s := WrapFilterInfo(nil)

	mt := StringEqual
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
	s := WrapTarget(nil)

	group := Group(3)
	s.SetGroup(group)
	require.Equal(t, group, s.Group())

	keys := [][]byte{
		{1, 2, 3},
		{4, 5, 6},
	}
	s.SetKeyList(keys)
	require.Equal(t, keys, s.KeyList())
}

func TestEACLRecordWrapper(t *testing.T) {
	s := WrapRecord(nil)

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

	group1 := Group(1)
	t1 := WrapTarget(nil)
	t1.SetGroup(group1)

	group2 := Group(2)
	t2 := WrapTarget(nil)
	t2.SetGroup(group2)

	s.SetTargetList([]Target{t1, t2})

	targets := s.TargetList()
	require.Len(t, targets, 2)
	require.Equal(t, group1, targets[0].Group())
	require.Equal(t, group2, targets[1].Group())
}

func TestEACLTableWrapper(t *testing.T) {
	s := WrapTable(nil)

	action1 := Action(1)
	r1 := WrapRecord(nil)
	r1.SetAction(action1)

	action2 := Action(2)
	r2 := WrapRecord(nil)
	r2.SetAction(action2)

	s.SetRecords([]Record{r1, r2})

	records := s.Records()
	require.Len(t, records, 2)
	require.Equal(t, action1, records[0].Action())
	require.Equal(t, action2, records[1].Action())

	s2, err := UnmarshalTable(MarshalTable(s))
	require.NoError(t, err)

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
			require.Equal(t, targets1[j].Group(), targets2[j].Group())
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
