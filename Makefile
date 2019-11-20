B=\033[0;1m
G=\033[0;92m
R=\033[0m

# Reformat code
format:
	@[ ! -z `which goimports` ] || (echo "install goimports" && exit 2)
	@for f in `find . -type f -name '*.go' -not -path './vendor/*' -not -name '*.pb.go' -prune`; do \
		echo "${B}${G}⇒ Processing $$f ${R}"; \
		goimports -w $$f; \
	done

# Regenerate documentation for protot files:
docgen:
	@for f in `find . -type f -name '*.proto' -not -path './vendor/*' -exec dirname {} \; | sort -u `; do \
		echo "${B}${G}⇒ Documentation for $$(basename $$f) ${R}"; \
		protoc \
			--doc_opt=.github/markdown.tmpl,$${f}.md \
			--proto_path=.:./vendor:/usr/local/include \
			--doc_out=docs/ $${f}/*.proto; \
	done

# Regenerate proto files:
protoc:
	@go mod tidy -v
	@go mod vendor
	# Install specific version for gogo-proto
	@go list -f '{{.Path}}/...@{{.Version}}' -m github.com/gogo/protobuf | xargs go get -v
	# Install specific version for protobuf lib
	@go list -f '{{.Path}}/...@{{.Version}}' -m  github.com/golang/protobuf | xargs go get -v
	# Protoc generate
	@for f in `find . -type f -name '*.proto' -not -path './vendor/*'`; do \
		echo "${B}${G}⇒ Processing $$f ${R}"; \
		protoc \
			--proto_path=.:./vendor:/usr/local/include \
			--gofast_out=plugins=grpc,paths=source_relative:. $$f; \
	done
