{ lib, config, ... }:
let
  cfg = config.facter;
in
{

  options.facter = with lib; {
    report = mkOption {
      type = types.path;
      description = mdDoc "Path to a report.json that was generated by nixos-facter.";
    };
  };

  config.facter = {

  };
}
