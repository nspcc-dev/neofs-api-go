module github.com/nspcc-dev/neofs-api-go/v2

go 1.17

require (
	github.com/nspcc-dev/neofs-crypto v0.4.0
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.28.0
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/nspcc-dev/rfc6979 v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/text v0.3.8 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

// This version uses broken NeoFS API with incompatible signature
// definitions. See fix in https://github.com/nspcc-dev/neofs-api/pull/203
retract v2.12.0
