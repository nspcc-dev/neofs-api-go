module github.com/nspcc-dev/neofs-api-go/v2

go 1.17

require (
	github.com/nspcc-dev/neofs-crypto v0.4.0
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/nspcc-dev/rfc6979 v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	gopkg.in/yaml.v3 v3.0.0 // indirect
)

// This version uses broken NeoFS API with incompatible signature
// definitions. See fix in https://github.com/nspcc-dev/neofs-api/pull/203
retract v2.12.0
