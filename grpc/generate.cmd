@echo off

set PATH=E:/go/bin;E:/Programs/protobuf/bin;%PATH%

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative hashq.proto
