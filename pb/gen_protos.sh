#!/bin/bash
set -euxo pipefail

docker build -t banana .
docker run --rm \
  -u $(id -u):$(id -g) \
  -v $PWD:/project/proto \
  banana protoc \
    -I /project/proto \
    /project/proto/banana.proto \
    --plugin="protoc-gen-ts=/project/bin/protoc-gen-ts" \
    --proto_path="/project/proto" \
    --go_out=plugins=grpc:/project/proto \
    --ts_out="/project/proto"
