package query

import (
	"strings"

	"github.com/gogo/protobuf/proto"
)

var (
	_ proto.Message = (*Query)(nil)
	_ proto.Message = (*Filter)(nil)
)

// String returns string representation of Filter.
func (m Filter) String() string {
	b := new(strings.Builder)
	b.WriteString("<Filter '$" + m.Name + "' ")
	switch m.Type {
	case Filter_Exact:
		b.WriteString("==")
	case Filter_Regex:
		b.WriteString("~=")
	default:
		b.WriteString("??")
	}
	b.WriteString(" '" + m.Value + "'>")
	return b.String()
}

// String returns string representation of Query.
func (m Query) String() string {
	b := new(strings.Builder)
	b.WriteString("<Query [")
	ln := len(m.Filters)
	for i := 0; i < ln; i++ {
		b.WriteString(m.Filters[i].String())
		if ln-1 != i {
			b.WriteByte(',')
		}
	}
	b.WriteByte(']')
	return b.String()
}
