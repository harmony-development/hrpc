#!/usr/bin/env bash
mkdir -p "gendart"

for dir in $(find "protocol" -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq); do
    echo "Generating files in ${dir}..."
    find "${dir}" -name '*.proto'

    protoc \
    --proto_path=protocol \
    --dart_out=./gendart \
    --hrpc_out=./gendart \
    --hrpc_opt=dart_client \
    $(find "${dir}" -name '*.proto')
done
