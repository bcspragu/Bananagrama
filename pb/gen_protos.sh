#!/bin/bash
set -euxo pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

docker build -t banana $DIR
docker run --rm \
  -u $(id -u):$(id -g) \
  -v $DIR:/project/proto \
  banana protoc \
    --go_out="plugins=grpc:/project/proto" \
    --proto_path="/project/proto" \
    --grpc-web_out="import_style=commonjs+dts,mode=grpcwebtext:/project/proto" \
    --js_out="import_style=commonjs:/project/proto" \
    /project/proto/banana.proto

mv $DIR/banana_grpc_web_pb.d.ts $DIR/../frontend/src/proto/banana_pb_service.d.ts
mv $DIR/banana_grpc_web_pb.js $DIR/../frontend/src/proto/banana_pb_service.js
mv $DIR/banana_pb.d.ts \
  $DIR/banana_pb.js \
  $DIR/../frontend/src/proto
