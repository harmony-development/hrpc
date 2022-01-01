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
          protoc-gen-hrpc = pkgs.callPackage ./build-hrpc-gen.nix {
            src = inputs.self;
          };
        };
      in
      {
        inherit packages;
      }
    );
}
