module github.com/nspcc-dev/neofs-api-go

go 1.14

require (
	github.com/alecthomas/participle v0.6.0
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.1
	github.com/mr-tron/base58 v1.1.2
	github.com/nspcc-dev/hrw v1.0.9
	github.com/nspcc-dev/neo-go v0.91.0
	github.com/nspcc-dev/neofs-crypto v0.3.0
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.1
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.23.0
)

// Used for debug reasons
// replace github.com/nspcc-dev/neofs-crypto => ../neofs-crypto
