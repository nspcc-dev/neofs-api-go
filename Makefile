#!/usr/bin/make -f
SHELL = bash

VERSION ?= $(shell git describe --tags --match "v*" --abbrev=8 --dirty --always)

.PHONY: dep fmts fmt imports protoc test lint version help

# Pull go dependencies
dep:
	@printf "⇒ Tidy requirements : "
	CGO_ENABLED=0 \
	GO111MODULE=on \
	go mod tidy -v && echo OK
	@printf "⇒ Download requirements: "
	CGO_ENABLED=0 \
	GO111MODULE=on \
	go mod download && echo OK
	@printf "⇒ Install test requirements: "
	CGO_ENABLED=0 \
	GO111MODULE=on \
	go test ./... && echo OK

# Run all code formatters
fmts: fmt imports

# Reformat code
fmt:
	@echo "⇒ Processing gofmt check"
	@for f in `find . -type f -name '*.go' -not -path './vendor/*' -not -name '*.pb.go' -prune`; do \
		GO111MODULE=on gofmt -s -w $$f; \
	done

# Reformat imports
imports:
	@echo "⇒ Processing goimports check"
	@for f in `find . -type f -name '*.go' -not -path './vendor/*' -not -name '*.pb.go' -prune`; do \
		GO111MODULE=on goimports -w $$f; \
	done

# Regenerate code for proto files
protoc:
	@GOPRIVATE=github.com/nspcc-dev go mod vendor
	# Install specific version for protobuf lib
	@go list -f '{{.Path}}/...@{{.Version}}' -m  google.golang.org/protobuf | xargs go install -v
	# Protoc generate
	@for f in `find . -type f -name '*.proto' -not -path './vendor/*'`; do \
		echo "⇒ Processing $$f "; \
		protoc \
			--proto_path=.:./vendor:/usr/local/include \
			--go_out=. --go_opt=paths=source_relative \
			--go-grpc_opt=require_unimplemented_servers=false \
			--go-grpc_out=. --go-grpc_opt=paths=source_relative $$f; \
	done
	rm -rf vendor

# Run Unit Test with go test
test:
	@echo "⇒ Running go test"
	@GO111MODULE=on go test ./...

# Run linters
lint:
	@golangci-lint run

# Print version
version:
	@echo $(VERSION)

# Show this help prompt
help:
	@echo '  Usage:'
	@echo ''
	@echo '    make <target>'
	@echo ''
	@echo '  Targets:'
	@echo ''
	@awk '/^#/{ comment = substr($$0,3) } comment && /^[a-zA-Z][a-zA-Z0-9_-]+ ?:/{ print "   ", $$1, comment }' $(MAKEFILE_LIST) | column -t -s ':' | grep -v 'IGNORE' | sort -u
