#!/bin/bash

if [ -z "$1" ]; then
    echo "usage: ./prepare.sh path/to/neofs-api"
    exit 1
fi

API_GO_PATH=$(pwd)
API_PATH=$1

# MOVE FILES FROM API REPO
cd "$API_PATH" || exit 1
ARGS=$(find ./ -name '*.proto' -not -path './vendor/*')
for file in $ARGS; do
	dir=$(dirname "$file")
	mkdir -p "$API_GO_PATH/$dir/grpc"
	cp -r "$dir"/* "$API_GO_PATH/$dir/grpc"
done

# MODIFY FILES
cd "$API_GO_PATH" || exit 1
ARGS2=$(find ./ -name '*.proto')
for file in $ARGS2; do
	echo "$file"
	sed -i "s/import\ \"\(.*\)\/\(.*\)\.proto\";/import\ \"\1\/grpc\/\2\.proto\";/" $file
done

cd "$API_GO_PATH" || exit 1
# COMPILE
make protoc

# REMOVE PROTO DEFINITIONS
ARGS=$(find ./$prefix -name '*.proto' -not -path './vendor/*' -not -path './util/*')
for file in $ARGS; do
	rm "$file"
done
