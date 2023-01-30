{
  description = "Nix flake for hexto256";

  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    flake-compat = {
      url = "github:edolstra/flake-compat";
      flake = false;
    };
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      rec {
        packages.hexto256 = with pkgs; buildGoModule {
          pname = "hexto256";
          version = "latest";
          src = ./.;

          vendorSha256 = "sha256-pVYrgy5616bY+bH0N+lpMYYK0VDu9xxa9mF/tCquShU=";
        };
        packages.default = packages.hexto256;
      }
    );
}
