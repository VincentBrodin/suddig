{
  description = "Suddig Devshell";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs, ... }:
    let
      # This helper allows the shell to work on any system (Intel/ARM Linux/Mac)
      supportedSystems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];
      forEachSystem = f: nixpkgs.lib.genAttrs supportedSystems (system: f {
        pkgs = import nixpkgs { inherit system; };
      });
    in
    {
      devShells = forEachSystem ({ pkgs }: {
        default = pkgs.mkShell {
          # Tools go here
          nativeBuildInputs = [
            pkgs.go
            pkgs.gopls
            pkgs.pkg-config
            pkgs.nodePackages.vscode-json-languageserver
          ];

          # Libraries your project links to go here
          buildInputs = [
            pkgs.openssl
          ];

          # Nix automatically handles PKG_CONFIG_PATH when openssl is in buildInputs
          # and pkg-config is in nativeBuildInputs.
          shellHook = ''
          '';
        };
      });
    };
}
