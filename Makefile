.PHONY: app

build-wasm:
	cd go-wasm/wasm && GOOS=js GOARCH=wasm go build -o  ../../assets/module.wasm

build-layer8:
	cd go-layer8/cmd && go build -o ../../bin/layer8-wasm-server

run-server:
	./bin/layer8-wasm-server

run:
	make build-wasm && make build-layer8 && make run-server

serve:
	cd go-http-server && go run main.go