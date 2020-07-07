package acl

const (
	_ MatchType = iota
	StringEqual
	StringNotEqual
)

// Maps MatchType to corresponding function.
// 1st argument of function - header value, 2nd - header filter.
var mMatchFns = map[MatchType]func(Header, Header) bool{
	StringEqual: stringEqual,

	StringNotEqual: stringNotEqual,
}

const (
	mResUndefined = iota
	mResMatch
	mResMismatch
)

func stringEqual(header, filter Header) bool {
	return header.Value() == filter.Value()
}

func stringNotEqual(header, filter Header) bool {
	return header.Value() != filter.Value()
}
