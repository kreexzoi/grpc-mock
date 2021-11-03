#!/bin/sh

pushd $(dirname "${BASH_SOURCE[0]}")
protoc -I. \
       -I./third_party/googleapis \
       --go_out=plugins=grpc:. \
       --grpc-gateway_out=logtostderr=true:. *.proto
popd
