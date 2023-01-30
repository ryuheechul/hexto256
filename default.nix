# allows legacy way like followings powered by https://github.com/edolstra/flake-compat
# ```nix
# hexto256 = (import (fetchTarball https://github.com/ryuheechul/hexto256/archive/main.tar.gz)).default
# # or locally
# hexto256 = (import path/to/this/repo).default;
# ```
#
# also can test this file locally via `nix-build`
(import
  (
    let lock = builtins.fromJSON (builtins.readFile ./flake.lock); in
    fetchTarball {
      url = "https://github.com/edolstra/flake-compat/archive/${lock.nodes.flake-compat.locked.rev}.tar.gz";
      sha256 = lock.nodes.flake-compat.locked.narHash;
    }
  )
  { src = ./.; }
).defaultNix
