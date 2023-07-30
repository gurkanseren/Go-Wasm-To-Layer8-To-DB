.PHONY: app

build-wasm:
	cd go-wasm/wasm && GOOS=js GOARCH=wasm go build -o  ../../assets/module.wasm

build-layer8:
	cd go-layer8/cmd && go build -o ../../bin/layer8-wasm-server

run-server:
	./bin/layer8-wasm-server

run-content-server:
	cd go-http-server/cmd && go run main.go

run-layer8-and-wasm:
	make build-wasm && make build-layer8 && make run-server

run-load-balancer:
	cd go-load-balancer/cmd && go run main.go
