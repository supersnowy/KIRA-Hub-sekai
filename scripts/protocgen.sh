#!/usr/bin/env bash

set -eo pipefail

proto_dirs=$(find ./x -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
echo $proto_dirs

for dir in $proto_dirs; do
  protoc -I "$dir" -I "third_party/proto" --gocosmos_out=plugins=interfacetype+grpc,\
Mgoogle/protobuf/any.proto=github.com/KiraCore/cosmos-sdk/codec/types:. \
  $(find "${dir}" -maxdepth 1 -name '*.proto')
done

# move proto files to the right places
cp -r github.com/KiraCore/cosmos-sdk/* ./
rm -rf github.com