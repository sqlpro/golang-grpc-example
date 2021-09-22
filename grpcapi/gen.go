package grpcapi

//go:generate -command compile_proto protoc -I../protos
//go:generate compile_proto greeter.proto --go_out=plugins=grpc,paths=source_relative:.
