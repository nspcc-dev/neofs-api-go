module github.com/nspcc-dev/neofs-api-go/v2

go 1.16

require (
	github.com/nspcc-dev/neofs-crypto v0.3.0
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
)

// This version uses broken NeoFS API with incompatible signature
// definitions. See fix in https://github.com/nspcc-dev/neofs-api/pull/203
retract v2.12.0
