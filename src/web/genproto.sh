#!/bin/bash

# create dir if not exist
dldir="./genproto"
[ ! -d "$dldir" ] && mkdir -p "$dldir"

# make sure path contains the go binarys
PATH=$PATH:$GOPATH/bin
protodir=../../pb

# executing the protoc command
protoc --proto_path=$protodir \
    --go_out=genproto \
    --go_opt=paths=source_relative \
    --go_opt=Mwalnut.proto=walnut/ \
    --go-grpc_out=genproto \
    --go-grpc_opt=paths=source_relative \
    --go-grpc_opt=Mwalnut.proto=walnut/ \
    $protodir/walnut.proto
