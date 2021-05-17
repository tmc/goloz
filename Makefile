.PHONY: deps
deps:
	go get github.com/bufbuild/buf/cmd/buf
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go get github.com/evanw/esbuild/cmd/esbuild

.PHONY: dev
dev:
	esbuild client/goloz-web-client/src/app.js --serve=8000 --servedir=docs --outfile=docs/app.js --bundle

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
	GOOS=js GOARCH=wasm go build -o docs/goloz.wasm github.com/tmc/goloz/cmd/goloz
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js ./docs/
	esbuild client/goloz-web-client/src/app.js --outfile=docs/app.js --bundle
