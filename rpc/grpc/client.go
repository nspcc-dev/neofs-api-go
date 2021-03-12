package grpc

// Client represents client for exchanging messages
// with a remote server using gRPC protocol.
type Client struct {
	*cfg
}

// Option is a Client's constructor option.
type Option func(*cfg)

// New creates, configures via options and returns new Client instance.
func New(opts ...Option) *Client {
	c := defaultCfg()

	for _, opt := range opts {
		opt(c)
	}

	return &Client{
		cfg: c,
	}
}
