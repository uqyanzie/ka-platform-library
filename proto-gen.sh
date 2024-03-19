#!/bin/bash

rm -rf server/pb/*.go
protoc --proto_path=proto --go_out=server/pb --go_opt=paths=source_relative \
  --go-grpc_out=server/pb --go-grpc_opt=paths=source_relative \
  proto/*.proto

protoc --proto_path=proto \
    --plugin=protoc-gen-ts=client/node_modules/.bin/protoc-gen-ts \
    --plugin=protoc-gen-grpc-web=client/node_modules/.bin/protoc-gen-grpc-web \
    --ts_out="import_style=commonjs,binary:./client/src/pb" \
    --js_out="import_style=commonjs,binary:./client/src/pb" \
    --grpc-web_out=import_style=typescript,mode=grpcwebtext:./client/src/pb \
    proto/*.proto


    