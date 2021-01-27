#!/usr/bin/env bash
mkdir -p "gen"

for dir in $(find "protocol" -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq); do
    echo "Generating files in ${dir}..."
    find "${dir}" -name '*.proto'

    protoc --experimental_allow_proto3_optional \
    --proto_path=protocol \
    --plugin=protoc-gen-custom=./hrpc \
    --custom_out=./gen \
    --custom_opt="/templates/hrpc-server-go.htmpl" \
    --go_out=./gen \
    --validate_out="lang=go:gen" \
    $(find "${dir}" -name '*.proto')

    protoc --experimental_allow_proto3_optional \
    --proto_path=protocol \
    --plugin=protoc-gen-custom=./hrpc \
    --custom_out=./gen \
    --custom_opt="/templates/hrpc-client-go.htmpl" \
    --go_out=./gen \
    --validate_out="lang=go:gen" \
    $(find "${dir}" -name '*.proto')
done

rsync -a -v gen/github.com/harmony-development/legato/gen/ ./gen

go fmt ./gen/./...