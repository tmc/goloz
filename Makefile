.PHONY: deps
deps:
	go get github.com/bufbuild/buf/cmd/buf
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get google.golang.org/protobuf/cmd/protoc-gen-go

.PHONY: lint
lint:
	buf lint

.PHONY: generate
generate:
	buf protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	proto/goloz/v1/goloz.proto

.PHONY: build
build:
	GOOS=js GOARCH=wasm go build -o build/goloz.wasm github.com/tmc/goloz/cmd/goloz
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js ./build/
