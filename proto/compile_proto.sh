#!/usr/bin/env bash

protoc --go_out=../backend/models *.proto
../frontend/node_modules/protobufjs/bin/pbjs -t static-module -w commonjs -o compiled.js *.proto
../frontend/node_modules/protobufjs/bin/pbts -o ../frontend/proto-types.d.ts compiled.js
rm compiled.js
