{
  pkgs,
  config,
  iopn ? (import ../. { inherit pkgs; }),
}:
rec {
  start-iopn = pkgs.writeShellScriptBin "start-iopn" ''
    # rely on environment to provide iopnd
    export PATH=${pkgs.test-env}/bin:$PATH
    ${../scripts/start-iopn} ${config.iopn-config} ${config.dotenv} $@
  '';
  start-geth = pkgs.writeShellScriptBin "start-geth" ''
    export PATH=${pkgs.test-env}/bin:${pkgs.go-ethereum}/bin:$PATH
    source ${config.dotenv}
    ${../scripts/start-geth} ${config.geth-genesis} $@
  '';
  start-scripts = pkgs.symlinkJoin {
    name = "start-scripts";
    paths = [
      start-iopn
      start-geth
    ];
  };
}
