module github.com/nspcc-dev/neofs-api-go

go 1.14

require (
	github.com/gogo/protobuf v1.1.1
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.1
	github.com/mr-tron/base58 v1.1.2
	github.com/nspcc-dev/hrw v1.0.9
	github.com/nspcc-dev/neo-go v0.91.0
	github.com/nspcc-dev/neofs-crypto v0.3.0
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.1
	golang.org/x/sys v0.0.0-20200420163511-1957bb5e6d1f // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20191108220845-16a3f7862a1a // indirect
	google.golang.org/grpc v1.29.1
)

// Used for debug reasons
// replace github.com/nspcc-dev/neofs-crypto => ../neofs-crypto
