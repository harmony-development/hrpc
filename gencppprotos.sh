#!/usr/bin/env bash

mkdir -p "cppgen"

for dir in $(find "protocol" -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq); do
    echo "Generating files in ${dir}..."
    find "${dir}" -name '*.proto'

    protoc \
    --proto_path=protocol \
    --hrpc_out=./cppgen \
    --hrpc_opt="qt_cpp_client:qt_cpp_proto" \
    $(find "${dir}" -name '*.proto')
done
