#!/bin/bash 

prefix=v2

if [ -z "$1" ]; then 
    echo "usage: ./prepare.sh path/to/neofs-api"
    exit 1
fi 

API_GO_PATH=$(pwd)
API_PATH=$1 

# MOVE FILES FROM API REPO
cd $API_PATH
ARGS=$(find ./ -name '*.proto' -not -path './vendor/*')
for file in $ARGS; do
    dir=$(dirname $file)
    cp -r $dir $API_GO_PATH
done
cd $API_GO_PATH

# MODIFY FILES
for file in $ARGS; do
    TYPES=$(grep '^import' $file | sed 's/import\ \"\(.*\)\/.*/\1/' | sort | uniq)
    PKG=$(grep '^package' $file | sed 's/package\ \(.*\);/\1/')

    TYPES=( "${TYPES[@]}" "${PKG[@]}") # merge two arrays
    TYPES=$(printf "%s\n" "${TYPES[@]}" | sort | uniq) # left only uniq elemetns

    for t in $TYPES; do
        sed -i "s/$t\./$t\.$prefix\./" $file
        sed -i "s/$t\//$t\/$prefix\//" $file
    done 

    sed -i "s/^package\(.*\);/package\1.$prefix;/" $file
    sed -i "s/go_package\(.*\)\";$/go_package\1\/$prefix\";/" $file

    dir=$(dirname $file)
    mkdir $dir/v2 2>/dev/null
    mv $file $dir/v2
done

# COMPILE
make protoc

# REMOVE PROTO FILES
# TO BE DONE AS NEOFS-API WILL BE STABLE
