#!/bin/bash
set -euxo pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

docker build -t banana $DIR
docker run --rm \
  -u $(id -u):$(id -g) \
  -v $DIR:/project/proto \
  banana protoc \
    -I /project/proto \
    /project/proto/banana.proto \
    --plugin="protoc-gen-ts=/project/bin/protoc-gen-ts" \
    --proto_path="/project/proto" \
    --go_out=plugins=grpc:/project/proto \
    --ts_out=service=true:"/project/proto"
mv $DIR/banana_pb.d.ts $DIR/banana_pb_service.d.ts $DIR/../frontend/src
rm banana_pb_service.js
