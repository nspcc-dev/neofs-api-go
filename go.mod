module github.com/nspcc-dev/neofs-api-go

go 1.14

require (
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.1
	github.com/mr-tron/base58 v1.1.3
	github.com/nspcc-dev/neo-go v0.90.0
	github.com/nspcc-dev/neofs-crypto v0.3.0
	github.com/nspcc-dev/netmap v1.7.0
	github.com/nspcc-dev/tzhash v1.4.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.6.0
	github.com/prometheus/client_model v0.2.0
	github.com/spf13/viper v1.7.0
	github.com/stretchr/testify v1.5.1
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
	google.golang.org/grpc v1.29.1
)

// Used for debug reasons
// replace github.com/nspcc-dev/neofs-crypto => ../neofs-crypto
