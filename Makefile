.PHONY: app

build-wasm:
	cd go-wasm/wasm && GOOS=js GOARCH=wasm go build -o  ../../assets/module.wasm

build-wasm-server:
	cd go-wasm/cmd && go build -o ../../bin/wasm-server

build-load-balancer:
	cd go-load-balancer/cmd && go build -o ../../bin/load-balancer

build-layer8-slave-one:
	cd go-layer8-slaves/layer8-slave-one/cmd && go build -o ../../../bin/layer8-slave-one

build-layer8-slave-two:
	 cd go-layer8-slaves/layer8-slave-two/cmd && go build -o ../../../bin/layer8-slave-two

build-layer8-slave-three:
	 cd go-layer8-slaves/layer8-slave-three/cmd && go build -o ../../../bin/layer8-slave-three

load-balancer: # Port 8000
	make build-load-balancer && ./bin/load-balancer

layer8-slave-one: # Port 8001
	make build-layer8-slave-one && ./bin/layer8-slave-one

layer8-slave-two: # Port 8002
	make build-layer8-slave-two && ./bin/layer8-slave-two

layer8-slave-three: # Port 8003
	make build-layer8-slave-three && ./bin/layer8-slave-three

wasm: # Port 9000
	make build-wasm && make build-wasm-server && ./bin/wasm-server

content-server: # Port 9001
	cd go-http-server/cmd && go run main.go
