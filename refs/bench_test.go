package refs

import (
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkObjectIDSlice(b *testing.B) {
	for _, size := range []int{0, 1, 50} {
		b.Run(strconv.Itoa(size)+" elements", func(b *testing.B) {
			benchmarkObjectIDSlice(b, size)
		})
	}
}

func benchmarkObjectIDSlice(b *testing.B, size int) {
	ids := make([]ObjectID, size)
	for i := range ids {
		ids[i].val = make([]byte, 32)
		rand.Read(ids[i].val)
	}
	raw := ObjectIDListToGRPCMessage(ids)

	b.Run("to grpc message", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			raw := ObjectIDListToGRPCMessage(ids)
			if len(raw) != len(ids) {
				b.FailNow()
			}
		}
	})
	b.Run("from grpc message", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ids, err := ObjectIDListFromGRPCMessage(raw)
			if err != nil || len(raw) != len(ids) {
				b.FailNow()
			}
		}
	})
	b.Run("marshal", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf := make([]byte, ObjectIDNestedListSize(1, ids))
			n := ObjectIDNestedListMarshal(1, buf, ids)
			if n != len(buf) {
				b.FailNow()
			}
		}
	})
}
