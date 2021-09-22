# Makefile

init:
	@go install github.com/smartystreets/goconvey
	@go install github.com/golang/protobuf/protoc-gen-go
	@go mod download

generate:
	@go generate ./...

run: generate
	@go run cmd/main.go

test: generate
	go test ./... -v
