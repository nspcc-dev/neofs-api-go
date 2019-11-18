protoc:
	@go mod tidy -v
	@go mod vendor
	# Install specific version for gogo-proto
	@go list -f '{{.Path}}/...@{{.Version}}' -m github.com/gogo/protobuf | xargs go get -v
	# Install specific version for protobuf lib
	@go list -f '{{.Path}}/...@{{.Version}}' -m  github.com/golang/protobuf | xargs go get -v
	# Protoc generate
	@find . -type f -name '*.proto' -not -path './vendor/*' \
		-exec protoc \
		--proto_path=.:./vendor \
		--gofast_out=plugins=grpc,paths=source_relative:. '{}' \;
