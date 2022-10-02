{
  description = "Scriptable image destruction & manipulation tool, used to power the next generation of Borik";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils/master";
  };

  outputs = { self, nixpkgs, flake-utils }:
    let
      version = builtins.substring 0 8 self.lastModifiedDate;
    in
      flake-utils.lib.eachDefaultSystem (system:
        let
          pkgs = import nixpkgs { inherit system; };
        in
        {
          devShell = pkgs.mkShell {
            buildInputs = [
              pkgs.go
              pkgs.gotools
              pkgs.imagemagick6
              pkgs.pkg-config
            ];
            shellHook = ''
              export CGO_CFLAGS_ALLOW=-Xpreprocessor
            '';
          };

          packages.default = pkgs.buildGoModule {
            pname = "sorik";
            inherit version;
            src = ./.;

            nativeBuildInputs = [ pkgs.pkg-config ];
            buildInputs = [ pkgs.imagemagick6 ];

            vendorSha256 = "sha256-fwjfnvVQteIfGPaPcxlNSWGFj/dqrbJOQClP1YD1BPg=";

            meta = with pkgs.lib; {
              description = "Scriptable image destruction & manipulation tool, used to power the next generation of Borik";
              homepage = https://github.com/fogo-sh/sorik;
              license = licenses.mit;
              platforms = platforms.linux ++ platforms.darwin;
            };
          };

          apps.default = {
            type = "app";
            program = "${self.packages.${system}.default}/bin/sorik";
          };
        });
}
