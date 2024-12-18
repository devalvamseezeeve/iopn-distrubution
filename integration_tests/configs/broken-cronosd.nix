{
  pkgs ? import ../../nix { },
}:
let
  iopnd = (pkgs.callPackage ../../. { });
in
iopnd.overrideAttrs (oldAttrs: {
  patches = oldAttrs.patches or [ ] ++ [ ./broken-iopnd.patch ];
})
