.PHONY: app

build-wasm:
	cd go-wasm/wasm && GOOS=js GOARCH=wasm go build -o  ../../assets/module.wasm

build-wasm-server:
	cd go-wasm/cmd && go build -o ../../bin/wasm-server

build-layer8:
	cd go-layer8/cmd && go build -o ../../bin/layer8-server

run-load-balancer: # Port 8000
	cd go-load-balancer/cmd && go run main.go

run-layer8: # Port 8001
	make build-layer8 && ./bin/layer8-server

run-wasm: # Port 8002
	make build-wasm && make build-wasm-server && ./bin/wasm-server

run-content-server: # Port 9000
	cd go-http-server/cmd && go run main.go
