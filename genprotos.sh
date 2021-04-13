#!/usr/bin/env bash
mkdir -p "gen"

for dir in $(find "protocol" -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq); do
    echo "Generating files in ${dir}..."
    find "${dir}" -name '*.proto'

    protoc \
    --proto_path=protocol \
    --hrpc_out=./gen \
    --hrpc_opt=hrpc-server-echo-go:hrpc-scanner:hrpc-client-go \
    $(find "${dir}" -name '*.proto')
done

rsync -a -v gen/github.com/harmony-development/legato/gen/ ./gen

go fmt ./gen/./...