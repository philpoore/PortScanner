HOST=127.0.0.1

go:
	cd go && go build

go-run: go
	./go/portscanner $(HOST)

rust:
	cd rust && cargo build --release

rust-run: rust
	./rust/target/release/portscanner

.PHONY: go rust