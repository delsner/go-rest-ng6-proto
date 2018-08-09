#!/usr/bin/env bash

protoc --go_out=../backend/models *.proto
../frontend/node_modules/protobufjs/bin/pbjs -t static-module -w commonjs -o ../frontend/proto-types.js *.proto
../frontend/node_modules/protobufjs/bin/pbts -o ../frontend/proto-types.d.ts ../frontend/proto-types.js
