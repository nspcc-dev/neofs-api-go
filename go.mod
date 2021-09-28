module github.com/nspcc-dev/neofs-api-go

go 1.16

require (
	github.com/google/uuid v1.1.2
	github.com/mr-tron/base58 v1.1.2
	github.com/nspcc-dev/hrw v1.0.9
	github.com/nspcc-dev/neo-go v0.95.3
	github.com/nspcc-dev/neofs-crypto v0.3.0
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.0.0-20210928044308-7d9f5e0b762b // indirect
	golang.org/x/sys v0.0.0-20210927094055-39ccf1dd6fa6 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20210928142010-c7af6a1a74c9 // indirect
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
)

// Used for debug reasons
// replace github.com/nspcc-dev/neofs-crypto => ../neofs-crypto
