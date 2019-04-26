#!/bin/sh
output=$(pwd)/proto/
input=$(pwd)/proto/*.proto
protoc -I/usr/local/include -I. -I$(pwd)/proto  $input --go_out=plugins=grpc:$output
protoc -I/usr/local/include -I. -I$(pwd)/proto  --grpc-gateway_out=logtostderr=true:$output $input
