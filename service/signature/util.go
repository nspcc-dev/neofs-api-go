package signature

import (
	"sync"

	"github.com/pkg/errors"
)

var bytesPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 5<<20)
	},
}

func dataForSignature(src DataSource) ([]byte, error) {
	if src == nil {
		return nil, errors.New("nil source")
	}

	buf := bytesPool.Get().([]byte)

	if size := src.SignedDataLength(); size < 0 {
		return nil, errors.New("negative length")
	} else if size <= cap(buf) {
		buf = buf[:size]
	} else {
		buf = make([]byte, size)
	}

	return src.ReadSignedData(buf)
}
