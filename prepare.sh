#!/bin/bash

prefix=v2

if [ -z "$1" ]; then 
    echo "usage: ./prepare.sh path/to/neofs-api"
    exit 1
fi 

API_GO_PATH=$(pwd)
API_PATH=$1 
mkdir $API_GO_PATH/$prefix 2>/dev/null

# MOVE FILES FROM API REPO
cd $API_PATH
ARGS=$(find ./ -name '*.proto' -not -path './vendor/*')
for file in $ARGS; do
    dir=$(dirname $file)
    cp -r $dir $API_GO_PATH/$prefix
done
cd $API_GO_PATH/$prefix

# MODIFY FILES
for file in $ARGS; do
	sed -i "s/import\ \"\(.*\)\";/import\ \"$prefix\/\1\";/" $file
done

# COMPILE
make protoc

# REMOVE PROTO FILES
# TO BE DONE AS NEOFS-API WILL BE STABLE
