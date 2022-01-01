.PHONY: deps
deps: deps-go deps-yarn
	go install github.com/bufbuild/buf/cmd/buf
	go install github.com/evanw/esbuild/cmd/esbuild
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/protobuf/cmd/protoc-gen-go

.PHONY: deps-go
deps-go:
	@command -v go > /dev/null || echo 'You need Go version 1.16 or higher. See https://golang.org/doc/install'

.PHONY: deps-yarn
deps-yarn:
	@command -v yarn > /dev/null || echo 'You need yarn. See https://yarnpkg.com/getting-started/install'

.PHONY: test
test:
	go test -v ./...

.PHONY: dev
dev:
	sh -c 'cd client/goloz-web-client; yarn'
	esbuild \
		client/goloz-web-client/src/index.js \
		'--define:process.env.NODE_ENV="development"' \
		--sourcemap \
		--serve=8000 --servedir=docs --outfile=docs/app.js --bundle

.PHONY: lint
lint:
	buf lint

.PHONY: generate
generate:
	buf protoc \
	    --go_out=. \
	    --go_opt=paths=source_relative \
	    --go-grpc_out=. \
	    --go-grpc_opt=paths=source_relative \
	    --grpc-gateway_out=. \
	    --grpc-gateway_opt=paths=source_relative \
	    --grpc-gateway_opt=generate_unbound_methods=true \
	    proto/goloz/v1/goloz.proto

.PHONY: build
build:
	GOOS=js GOARCH=wasm go build -o docs/goloz.wasm github.com/tmc/goloz/cmd/goloz
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js ./docs/
	sh -c  'cd client/goloz-web-client; yarn'
	esbuild \
		client/goloz-web-client/src/index.js \
		'--define:process.env.NODE_ENV="production"' \
		--sourcemap \
		--outfile=docs/app.js --bundle
