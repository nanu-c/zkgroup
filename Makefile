RUSTFLAGS := RUSTFLAGS='-C link-arg=-s'
BUILD_CMD := cargo build --release

.PHONY: all clean clean-all

golib: rustlib header
	go build

rustlib: libzkgroup-amd64 libzkgroup-arm64 libzkgroup-armhf

libzkgroup-amd64:
	cd lib/zkgroup && \
	$(RUSTFLAGS) $(BUILD_CMD)
	cp lib/zkgroup/target/release/libzkgroup.a lib/libzkgroup_linux_amd64.a

libzkgroup-arm64:
	cd lib/zkgroup && \
	$(RUSTFLAGS) CARGO_TARGET_AARCH64_UNKNOWN_LINUX_GNU_LINKER=aarch64-unknown-linux-gnu-gcc $(BUILD_CMD) --target=aarch64-unknown-linux-gnu
	cp lib/zkgroup/target/aarch64-unknown-linux-gnu/release/libzkgroup.a lib/libzkgroup_linux_arm64.a

libzkgroup-armhf:
	cd lib/zkgroup && \
	$(RUSTFLAGS) CARGO_TARGET_ARMV7_UNKNOWN_LINUX_GNUEABIHF_LINKER=armv7l-unknown-linux-gnueabihf-gcc $(BUILD_CMD) --target=armv7-unknown-linux-gnueabihf
	cp lib/zkgroup/target/armv7-unknown-linux-gnueabihf/release/libzkgroup.a lib/libzkgroup_linux_armhf.a

header:
	cbindgen --lang c lib/zkgroup/rust -o lib/zkgroup.h

clean:
	go clean
	cd lib/zkgroup && \
	cargo clean

clean-all: clean
	rm lib/*.a
	rm lib/*.h
