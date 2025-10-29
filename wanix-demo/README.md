# Wanix Demo Skeleton

This folder holds a minimal layout for a Wanix-style demo. Extend it by wiring the WASM module to the server and serving richer UI assets.

## Layout

- `go.mod` – module definition for the server and supporting code.
- `Makefile` – handy targets for running the server and compiling the WebAssembly binary.
- `server/` – Go HTTP server that currently serves the static assets.
- `wasm/` – placeholder Go source for building to WASI.
- `static/` – frontend assets served by the server.

## Usage

```bash
make run          # start the HTTP server at http://localhost:8080
make wasm         # build wasm/hello.go to wasm/hello.wasm
make clean        # remove build outputs
```

Feel free to adjust the Makefile targets to match your toolchain (TinyGo, wat2wasm, etc.).
