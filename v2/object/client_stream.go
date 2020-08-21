package object

type (
	GetObjectStreamer interface {
		Recv() (*GetResponse, error)
	}

	PutObjectStreamer interface {
		Send(*PutRequest) error
		CloseAndRecv() (*PutResponse, error)
	}

	SearchObjectStreamer interface {
		Recv() (*SearchResponse, error)
	}

	GetRangeObjectStreamer interface {
		Recv() (*GetRangeResponse, error)
	}
)

type (
	getObjectStream struct {
		recv func() (*GetResponse, error)
	}

	putObjectStream struct {
		send func(*PutRequest) error

		closeAndRecv func() (*PutResponse, error)
	}

	searchObjectStream struct {
		recv func() (*SearchResponse, error)
	}

	getRangeObjectStream struct {
		recv func() (*GetRangeResponse, error)
	}
)

func (s *getObjectStream) Recv() (*GetResponse, error) {
	return s.recv()
}

func (p *putObjectStream) Send(request *PutRequest) error {
	return p.send(request)
}

func (p *putObjectStream) CloseAndRecv() (*PutResponse, error) {
	return p.closeAndRecv()
}

func (s *searchObjectStream) Recv() (*SearchResponse, error) {
	return s.recv()
}

func (r *getRangeObjectStream) Recv() (*GetRangeResponse, error) {
	return r.recv()
}
