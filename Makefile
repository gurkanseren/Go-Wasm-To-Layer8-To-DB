.PHONY: app

build-wasm:
	cd go-wasm/wasm && GOOS=js GOARCH=wasm go build -o  ../../assets/module.wasm

build-wasm-server:
	cd go-wasm/cmd && go build -o ../../bin/wasm-server

build-layer8-slave-one:
	cd go-layer8-slaves/layer8-slave-one/cmd && go build -o ../../../bin/layer8-slave-one

build-layer8-slave-two:
	 cd go-layer8-slaves/layer8-slave-two/cmd && go build -o ../../../bin/layer8-slave-two

layer8-slave-one: # Port 8001
	make build-layer8-slave-one && ./bin/layer8-slave-one

layer8-slave-two: # Port 8002
	make build-layer8-slave-two && ./bin/layer8-slave-two

load-balancer: # Port 8000
	cd go-load-balancer/cmd && go run main.go

wasm: # Port 9000
	make build-wasm && make build-wasm-server && ./bin/wasm-server

content-server: # Port 9001
	cd go-http-server/cmd && go run main.go
