package client

import (
	"io"
	"testing"

	neofsecdsatest "github.com/nspcc-dev/neofs-api-go/crypto/ecdsa/test"
	"github.com/nspcc-dev/neofs-api-go/v2/object"
	"github.com/nspcc-dev/neofs-api-go/v2/signature"
	"github.com/stretchr/testify/require"
)

type singleResponseStream struct {
	called bool
	resp   object.GetResponse
}

func (x *singleResponseStream) Read(r *object.GetResponse) error {
	if x.called {
		return io.EOF
	}

	x.called = true

	*r = x.resp

	return nil
}

func chunkResponse(c []byte) (r object.GetResponse) {
	chunkPart := new(object.GetObjectPartChunk)
	chunkPart.SetChunk(c)

	body := new(object.GetResponseBody)
	body.SetObjectPart(chunkPart)

	r.SetBody(body)

	if err := signature.SignServiceMessage(neofsecdsatest.Signer(), &r); err != nil {
		panic(err)
	}

	return
}

func data(sz int) []byte {
	data := make([]byte, sz)

	for i := range data {
		data[i] = byte(i) % ^byte(0)
	}

	return data
}

func checkFullRead(t *testing.T, r io.Reader, buf, payload []byte) {
	var (
		restored []byte
		read     int
	)

	for {
		n, err := r.Read(buf)

		read += n
		restored = append(restored, buf[:n]...)

		if err != nil {
			require.Equal(t, err, io.EOF)
			break

		}
	}

	require.Equal(t, payload, restored)
	require.EqualValues(t, len(payload), read)
}

func TestObjectPayloadReader_Read(t *testing.T) {
	t.Run("read with tail", func(t *testing.T) {
		payload := data(10)

		buf := make([]byte, len(payload)-1)

		var r io.Reader = &objectPayloadReader{
			stream: &singleResponseStream{
				resp: chunkResponse(payload),
			},
		}

		checkFullRead(t, r, buf, payload)
	})
}
