# zkgroup

Go library for the Signal Private Group System.

See [github.com/signalapp/zkgroup](https://github.com/signalapp/zkgroup).

## Build the Rust library

### Build environment

### With [Nix](https://nixos.org/)

Simply enter the Nix shell:

    $ nix-shell

Alternatively, if you are also using [direnv](https://direnv.net/), you can
allow the `.envrc` to enter automatically the Nix shell whenever you enter the
project:

    $ direnv allow

### Without Nix

In order to build the Rust library for all three supported platform, you’ll
first need [Rustup](https://rustup.rs/).

Then, install the stable Rust toolchain:

    $ rustup install stable

Install the targets for the mobile platforms:

    $ rustup target add aarch64-unknown-linux-gnu
    $ rustup target add armv7-unknown-linux-gnueabihf

Install C compilers for these targets. Instructions depend on your distribution,
but you should have `aarch64-unknown-linux-gnu-gcc` and
`armv7l-unknown-linux-gnueabihf-gcc` in your path.

Install `cbindgen` to generate the C header for the Rust FFI:

    $ cargo install cbindgen

### Build instructions

1. Fetch the submodule, if not already done:

        $ git submodule init
        $ git submodule update

2. Build the library:

        $ make

This will build the Rust library into a static library for `amd64`, `arm64` and
`armhf` on Linux, put them in `lib/` and generate `lib/zkgroup.h`.

3. Clean the Rust articafcts:

        $ make clean

You can also remove the generated files by running:

        $ make clean-all
