package acl

type testExtendedACLTable struct {
	records []ExtendedACLRecord
}

type testRequestInfo struct {
	headers []TypedHeader
	key     []byte
	opType  OperationType
	target  Target
}

type testEACLRecord struct {
	opType  OperationType
	filters []HeaderFilter
	targets []ExtendedACLTarget
	action  ExtendedACLAction
}

type testEACLTarget struct {
	target Target
	keys   [][]byte
}

func (s testEACLTarget) Target() Target {
	return s.target
}

func (s testEACLTarget) KeyList() [][]byte {
	return s.keys
}

func (s testEACLRecord) OperationType() OperationType {
	return s.opType
}

func (s testEACLRecord) HeaderFilters() []HeaderFilter {
	return s.filters
}

func (s testEACLRecord) TargetList() []ExtendedACLTarget {
	return s.targets
}

func (s testEACLRecord) Action() ExtendedACLAction {
	return s.action
}

func (s testRequestInfo) HeadersOfType(typ HeaderType) ([]Header, bool) {
	res := make([]Header, 0, len(s.headers))

	for i := range s.headers {
		if s.headers[i].HeaderType() == typ {
			res = append(res, s.headers[i])
		}
	}

	return res, true
}

func (s testRequestInfo) Key() []byte {
	return s.key
}

func (s testRequestInfo) TypeOf(t OperationType) bool {
	return s.opType == t
}

func (s testRequestInfo) TargetOf(t Target) bool {
	return s.target == t
}

func (s testExtendedACLTable) Records() []ExtendedACLRecord {
	return s.records
}
