{
  dockerTools,
  runCommandLocal,
  iopn-matrix,
  benchmark-testcase,
}:
let
  patched-iopnd = iopn-matrix.iopnd.overrideAttrs (oldAttrs: {
    patches = oldAttrs.patches or [ ] ++ [ ./testground-iopnd.patch ];
  });
in
let
  tmpDir = runCommandLocal "tmp" { } ''
    mkdir -p $out/tmp/
  '';
in
dockerTools.buildLayeredImage {
  name = "iopn-testground";
  created = "now";
  contents = [
    benchmark-testcase
    patched-iopnd
    tmpDir
  ];
  config = {
    Expose = [
      9090
      26657
      26656
      1317
      26658
      26660
      26659
      30000
    ];
    Cmd = [ "/bin/stateless-testcase" ];
    Env = [ "PYTHONUNBUFFERED=1" ];
  };
}
