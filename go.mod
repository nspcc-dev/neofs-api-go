module github.com/nspcc-dev/neofs-api-go

go 1.16

require (
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.1
	github.com/mr-tron/base58 v1.1.2
	github.com/nspcc-dev/hrw v1.0.9
	github.com/nspcc-dev/neo-go v0.95.3
	github.com/nspcc-dev/rfc6979 v0.2.0
	github.com/stretchr/testify v1.6.1
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.23.0
)

// Used for debug reasons
// replace github.com/nspcc-dev/neofs-crypto => ../neofs-crypto
