package acl

type testTypedHeader struct {
	t HeaderType
	k string
	v string
}

type testHeaderSrc struct {
	hs []TypedHeader
}

type testHeaderFilter struct {
	TypedHeader
	t MatchType
}

func (s testHeaderFilter) MatchType() MatchType {
	return s.t
}

func (s testHeaderSrc) HeadersOfType(typ HeaderType) ([]Header, bool) {
	res := make([]Header, 0, len(s.hs))

	for i := range s.hs {
		if s.hs[i].HeaderType() == typ {
			res = append(res, s.hs[i])
		}
	}

	return res, true
}

func (s testTypedHeader) Name() string {
	return s.k
}

func (s testTypedHeader) Value() string {
	return s.v
}

func (s testTypedHeader) HeaderType() HeaderType {
	return s.t
}
