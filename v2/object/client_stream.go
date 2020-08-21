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
	getObjectGRPCStream struct {
		recv func() (*GetResponse, error)
	}

	putObjectGRPCStream struct {
		send func(*PutRequest) error

		closeAndRecv func() (*PutResponse, error)
	}

	searchObjectGRPCStream struct {
		recv func() (*SearchResponse, error)
	}

	getRangeObjectGRPCStream struct {
		recv func() (*GetRangeResponse, error)
	}
)

func (s *getObjectGRPCStream) Recv() (*GetResponse, error) {
	return s.recv()
}

func (p *putObjectGRPCStream) Send(request *PutRequest) error {
	return p.send(request)
}

func (p *putObjectGRPCStream) CloseAndRecv() (*PutResponse, error) {
	return p.closeAndRecv()
}

func (s *searchObjectGRPCStream) Recv() (*SearchResponse, error) {
	return s.recv()
}

func (r *getRangeObjectGRPCStream) Recv() (*GetRangeResponse, error) {
	return r.recv()
}
