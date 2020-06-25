#!/usr/bin/env bash

PROTOC=$(which protoc)

${PROTOC} -I. -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/ \
--go_out=plugins=grpc,\
paths=source_relative:./proto \
proto/prey.proto
${PROTOC} -I. -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/ \
--grpc-gateway_out=logtostderr=true,\
paths=source_relative:./proto \
proto/prey.proto