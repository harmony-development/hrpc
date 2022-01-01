{ stdenv
, fetchFromGitHub
, buildGoModule
, lib
, src
, vendorSha256
, name
,
}: buildGoModule {
  inherit name src vendorSha256;
  
  version = "0.0.0";
  subPackages = [ name ];
}
