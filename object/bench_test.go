package object

import (
	"math/rand"
	"testing"
)

func randString(n int) string {
	x := make([]byte, n)
	for i := range x {
		x[i] = byte('a' + rand.Intn('z'-'a'))
	}
	return string(x)
}

func BenchmarkAttributesMarshal(b *testing.B) {
	attrs := make([]*Attribute, 50)
	for i := range attrs {
		attrs[i] = new(Attribute)
		attrs[i].key = SysAttributePrefix + randString(10)
		attrs[i].val = randString(10)
	}
	raw := AttributesToGRPC(attrs)

	b.Run("marshal", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			res := AttributesToGRPC(attrs)
			if len(res) != len(raw) {
				b.FailNow()
			}
		}
	})
	b.Run("unmarshal", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			res, err := AttributesFromGRPC(raw)
			if err != nil || len(res) != len(raw) {
				b.FailNow()
			}
		}
	})
}
