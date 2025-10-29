.PHONY: run build wasm clean kill-port

PORT ?= 8080

run: kill-port build
	cd server && go run .

build:
	mkdir -p bin
	cd server && go build -o ../bin/wanix-demo

wasm:
	GOOS=wasip1 GOARCH=wasm go build -o wasm/hello.wasm ./wasm

clean:
	rm -rf bin wasm/hello.wasm

kill-port:
	-@pids=$$(lsof -ti tcp:$(PORT)); \
	if [ -n "$$pids" ]; then \
		echo "Killing processes on port $(PORT): $$pids"; \
		kill -9 $$pids; \
	fi
