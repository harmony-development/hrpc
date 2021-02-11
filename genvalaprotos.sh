#!/usr/bin/env bash
mkdir -p "genvala"

for dir in $(find "protocol" -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq); do
    echo "Generating files in ${dir}..."
    find "${dir}" -name '*.proto'

    protoc \
    --proto_path=protocol \
    --hrpc_out=./genvala \
    --hrpc_opt=vala_client \
    $(find "${dir}" -name '*.proto')
done
