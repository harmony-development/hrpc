#!/usr/bin/env bash
mkdir -p "genswift"

for dir in $(find "protocol" -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq); do
    echo "Generating files in ${dir}..."
    find "${dir}" -name '*.proto'

    protoc \
    --proto_path=protocol \
    --swift_out=./genswift \
    --hrpc_opt=swift_server \
    --hrpc_out=./genswift \
    $(find "${dir}" -name '*.proto')
done
