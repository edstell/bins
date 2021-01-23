#!/usr/bin/env bash
set -e

# Usage:
#   ./bin/generate_protobufs service.foo

_realpath() {
	[[ $1 = /* ]] && echo "$1" || echo "$PWD/${1#./}"
}

ROOT=$(_realpath "${1:-$(pwd)}")

if ! [[ -d $ROOT ]]
then
	>&2 echo "Provide a valid directory name"
	exit 1
fi

files=$(find $ROOT -maxdepth 3 -type f -name "*.proto")
for f in $files; do
	echo "Compiling $f";
	protoc -I=$(dirname $f) --proto_path=$GOPATH/src --go_out=$GOPATH/src --service_out=$(dirname $f) --govalidators_out=$GOPATH/src $f;
done
