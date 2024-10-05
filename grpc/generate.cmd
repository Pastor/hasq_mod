@echo off

set PATH=C:\Users\Pastor\go\bin;%PATH%

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative hashq.proto
