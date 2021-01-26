#!/usr/bin/env bash
mkdir -p "gen"

for dir in $(find "protos" -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq); do
    echo "Generating files in ${dir}..."
    find "${dir}" -name '*.proto'

    protoc --experimental_allow_proto3_optional \
    --proto_path=protos \
    --plugin=protoc-gen-custom=../hrpc \
    --custom_out=./gen \
    --custom_opt="../hrpc-server-go.htmpl" \
    --go_out=./gen \
    --validate_out="lang=go:gen" \
    $(find "${dir}" -name '*.proto')

    protoc --experimental_allow_proto3_optional \
    --proto_path=protos \
    --plugin=protoc-gen-custom=../hrpc \
    --custom_out=./gen \
    --custom_opt="../hrpc-client-go.htmpl" \
    --go_out=./gen \
    --validate_out="lang=go:gen" \
    $(find "${dir}" -name '*.proto')
done

sed -i 's%import "hrpc/jiti"%import "hrpc/gen/hrpc/jiti"%g' $(find gen -name '*.go')
go fmt ./gen/./...
