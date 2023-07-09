.PHONY: app

build-wasm:
	cd go-wasm/cmd/wasm && GOOS=js GOARCH=wasm go build -o  ../../assets/module.wasm

run-wasm:
	cd go-wasm/cmd/server && go run main.go

# Use `make run` directly to save this, building and starting the server
wasm:
	make build-wasm && make run-wasm

build-server:
	cd go-layer8/cmd && go build -o ../../bin/layer8-server

run-server:
	./bin/layer8-server

layer8:
	make build-server && make run-server