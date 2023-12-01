{
  description = "Flake for setting up my dev environment for Advent of code";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/release-23.11";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      goVersion = 21;
      overlays = [(self: super: {go = super."go_1_${toString goVersion}";})];
      pkgs = import nixpkgs {inherit overlays system;};

      aoc-load = pkgs.writeShellScriptBin "aoc-load" ''
        if [ $1 ]
        then
            curl --cookie "session=$AOC_COOKIE" https://adventofcode.com/$1/day/$2/input > in.txt
        else
            curl --cookie "session=$AOC_COOKIE" `date +https://adventofcode.com/%Y/day/%d/input` > in.txt
        fi
      '';
    in {
      formatter = pkgs.alejandra;
      devShells.default = pkgs.mkShellNoCC {
        packages = with pkgs; [
          go
          gotools
          golangci-lint
          aoc-load
        ];

        shellHook = ''
          ${pkgs.go}/bin/go version
          source ./secrets
        '';
      };
    });
}
