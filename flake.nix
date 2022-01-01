{
  description = "Flake for hrpc";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flakeUtils.url = "github:numtide/flake-utils";
    flakeCompat = {
      url = "github:edolstra/flake-compat";
      flake = false;
    };
  };

  outputs = inputs:
    with inputs.flakeUtils.lib; eachSystem [ "x86_64-linux" ] (system:
      let
        pkgs = import inputs.nixpkgs {
          inherit system;
        };
        packages = rec {
          protoc-gen-hrpc = pkgs.callPackage ./build-generic.nix {
            src = inputs.self;
            name = "protoc-gen-hrpc";
            vendorSha256 = "sha256-VoOBEpdkGIonW0AHO7fYsgmyAtDcSU8V/hTR0KHB7i0=";
          };
          protoc-gen-go-hrpc = pkgs.callPackage ./build-generic.nix {
            src = inputs.self;
            name = "protoc-gen-go-hrpc";
            vendorSha256 = "sha256-VoOBEpdkGIonW0AHO7fYsgmyAtDcSU8V/hTR0KHB7i0=";
          };
        };
      in
      {
        inherit packages;
      }
    );
}
