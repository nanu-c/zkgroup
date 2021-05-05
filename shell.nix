{ pkgs ? import <nixpkgs> {
  overlays = [ (import (builtins.fetchTarball "https://github.com/oxalica/rust-overlay/archive/master.tar.gz")) ];
} }:

with pkgs;

let
  rust-toolchain = rust-bin.stable.latest.default.override {
    targets = [
      "aarch64-unknown-linux-gnu"
      "armv7-unknown-linux-gnueabihf"
    ];
  };
in

mkShell {
  buildInputs = [
    go
    rust-toolchain
    rust-cbindgen

    # NOTE: This is not the best way to create a cross-compilation environment
    # in Nix, as libraries cannot be handled this way. However, it works
    # currently for our use case, and allow us to have all three build
    # environments in a single Nix shell.
    pkgsCross.aarch64-multiplatform.pkgsBuildHost.gcc
    pkgsCross.armv7l-hf-multiplatform.pkgsBuildHost.gcc
  ];
}
