PROTO_VERSION=v0.7.0
PROTO_URL=https://github.com/nspcc-dev/neofs-api/archive/$(PROTO_VERSION).tar.gz

B=\033[0;1m
G=\033[0;92m
R=\033[0m

.PHONY: deps format docgen protoc

# Dependencies
deps:
	@echo "${B}${G}=> Golang modules ${R}"
	@go mod tidy -v
	@go mod vendor

	@echo "${B}${G}=> Cleanup old files ${R}"
	@find . -type f -name '*.proto' -not -path './vendor/*' -not -name '*_test.proto' -exec rm {} \;

	@echo "${B}${G}=> NeoFS Proto files ${R}"
	@mkdir -p ./vendor/proto
	@curl -sL -o ./vendor/proto.tar.gz $(PROTO_URL)
	@tar -xzf ./vendor/proto.tar.gz --strip-components 1 -C ./vendor/proto
	@for f in `find ./vendor/proto -type f -name '*.proto' -exec dirname {} \; | sort -u `; do \
		cp $$f/*.proto ./$$(basename $$f); \
	done

	@echo "${B}${G}=> Cleanup ${R}"
	@rm -rf ./vendor/proto
	@rm -rf ./vendor/proto.tar.gz

# Reformat code
format:
	@[ ! -z `which goimports` ] || (echo "install goimports" && exit 2)
	@for f in `find . -type f -name '*.go' -not -path './vendor/*' -not -name '*.pb.go' -prune`; do \
		echo "${B}${G}⇒ Processing $$f ${R}"; \
		goimports -w $$f; \
	done

# Regenerate documentation for protot files:
docgen: deps
	@for f in `find . -type f -name '*.proto' -not -path './vendor/*' -exec dirname {} \; | sort -u `; do \
		echo "${B}${G}⇒ Documentation for $$(basename $$f) ${R}"; \
		protoc \
			--doc_opt=.github/markdown.tmpl,$${f}.md \
			--proto_path=.:./vendor:/usr/local/include \
			--doc_out=docs/ $${f}/*.proto; \
	done

# Regenerate proto files:
protoc: deps
	@echo "${B}${G}=> Cleanup old files ${R}"
	@find . -type f -name '*.pb.go' -not -path './vendor/*' -exec rm {} \;

	@echo "${B}${G}=> Install specific version for gogo-proto ${R}"
	@go list -f '{{.Path}}/...@{{.Version}}' -m github.com/gogo/protobuf | xargs go get -v
	@echo "${B}${G}=> Install specific version for protobuf lib ${R}"
	@go list -f '{{.Path}}/...@{{.Version}}' -m  github.com/golang/protobuf | xargs go get -v
	@echo "${B}${G}=> Protoc generate ${R}"
	@for f in `find . -type f -name '*.proto' -not -path './vendor/*'`; do \
		echo "${B}${G}⇒ Processing $$f ${R}"; \
		protoc \
			--proto_path=.:./vendor:/usr/local/include \
            --gofast_out=plugins=grpc,paths=source_relative:. $$f; \
	done

update: docgen protoc
