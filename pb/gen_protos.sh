#!/bin/bash
set -euxo pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

docker build -t banana $DIR
docker run --rm \
  -u $(id -u):$(id -g) \
  -v $DIR:/project/proto \
  banana protoc \
    --go_out="plugins=grpc:/project/proto" \
    --plugin="protoc-gen-ts=/project/bin/protoc-gen-ts" \
    --proto_path="/project/proto" \
    --ts_out="service=true:/project/proto" \
    --js_out="import_style=commonjs,binary:/project/proto" \
    /project/proto/banana.proto
mv $DIR/banana_pb.d.ts \
  $DIR/banana_pb_service.d.ts \
  $DIR/banana_pb.js \
  $DIR/banana_pb_service.js \
  $DIR/../frontend/src/proto
