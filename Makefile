.PHONY: wasm
wasm:
	rm -f ./html/*.js ./html/*.wasm
	cp $$(go env GOROOT)/misc/wasm/wasm_exec.js ./html/wasm_exec.js
	GOOS=js GOARCH=wasm go build -o ./html/main.wasm .

.PHONY: native
native:
	go build -o ./build/pong .

.PHONY: server
server:
	go run server/server.go

.PHONY: clean
clean:
	rm -f ./html/*.js ./html/*.wasm
