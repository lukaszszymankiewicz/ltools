all:
	GOOS=js GOARCH=wasm go build -v -o ltools.wasm

clean:
	rm -f ltools.wasm
