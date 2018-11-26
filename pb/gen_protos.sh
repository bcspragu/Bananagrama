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

# The output banana_pb_service.js file uses CommonJS exports.* style exports,
# which don't play nice with our default Vue + Webpack configuration. I'm too
# dumb to figure out how to convince Webpack to treat the proto/ directory as a
# module and parse it correctly, but I do know how to use sed to blindly
# replace strings in the generated file so that it works in our environment in
# an extremely fragile way.
sed -i '/^exports\./d' $DIR/banana_pb_service.js
sed -i 's/^var BananaService/export var BananaService/g' $DIR/banana_pb_service.js
sed -i 's/^function BananaServiceClient/export function BananaServiceClient/g' $DIR/banana_pb_service.js

mv $DIR/banana_pb.d.ts \
  $DIR/banana_pb_service.d.ts \
  $DIR/banana_pb.js \
  $DIR/banana_pb_service.js \
  $DIR/../frontend/src/proto
