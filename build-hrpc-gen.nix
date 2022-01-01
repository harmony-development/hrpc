{ stdenv
, fetchFromGitHub
, buildGoModule
, lib
, src
,
}: buildGoModule {
  name = "protoc-gen-hrpc";
  version = "main";

  src = src;

  subPackages = [ "protoc-gen-hrpc" ];

  vendorSha256 = "sha256-VoOBEpdkGIonW0AHO7fYsgmyAtDcSU8V/hTR0KHB7i0=";
}
