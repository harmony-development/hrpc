#!/usr/bin/env bash
mkdir -p "genruby"

for dir in $(find "protocol" -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq); do
    echo "Generating files in ${dir}..."
    find "${dir}" -name '*.proto'

    protoc \
    --proto_path=protocol \
    --ruby_out=./genruby \
    --hrpc_out=./genruby \
    --hrpc_opt=ruby_client \
    $(find "${dir}" -name '*.proto')
done
