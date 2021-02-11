#!/usr/bin/env bash

mkdir -p "jsgen"

for dir in $(find "protocol" -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq); do
    echo "Generating files in ${dir}..."
    find "${dir}" -name '*.proto'

    protoc \
    --proto_path=protocol \
		--hrpc_out=./jsgen \
    --hrpc_opt="ts_client" \
    $(find "${dir}" -name '*.proto')
done
