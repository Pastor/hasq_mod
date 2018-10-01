@echo off

%CD%/.bin/bin/protoc.exe -I %CD%/simple %CD%/simple/hasq_simple.proto  --go_out=plugins=grpc:simple
go generate .