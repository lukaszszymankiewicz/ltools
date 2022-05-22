.PHONY: wasm
wasm:
	rm -f ./html/*.js ./html/*.wasm
	cp $$(go env GOROOT)/misc/wasm/wasm_exec.js ./html/wasm_exec.js
	GOOS=js GOARCH=wasm go build -o ./html/main.wasm .

.PHONY: native
native:
	go build -o ./ltools .

.PHONY: server
server:
	go run server/server.go

.PHONY: clean
clean:
	# JavaScript and HTML files are deleted
	rm -f ./html/*.js ./html/*.wasm
	# compiled binaries are deleted
	rm -f ltools
	# exported levels are deleted
	rm -f *.llv
	# exported levels are deleted
	rm -f sample_name.png
	# exported levels are deleted
	rm -f *.zip
	# old logs file
	rm -f ./logs/*

.PHONY: test
test:
	# run tests in all directories and sub-directories
	go test ./...
