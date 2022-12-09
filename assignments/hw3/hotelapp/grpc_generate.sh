#!/bin/bash

cd internal/profile/proto 
protoc --go_out=plugins=grpc:. profile.proto
cd ../../geo/proto 
protoc --go_out=plugins=grpc:. geo.proto
cd ../../rate/proto 
protoc --go_out=plugins=grpc:. rate.proto
cd ../../search/proto 
protoc --go_out=plugins=grpc:. search.proto