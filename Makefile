.PHONY: app

build:
	cd go-wasm/cmd/wasm && GOOS=js GOARCH=wasm go build -o  ../../assets/module.wasm

serve:
	cd go-wasm/cmd/server && go run main.go

# Use `make run` directly to save this, building and starting the server
run:
	make build && make serve