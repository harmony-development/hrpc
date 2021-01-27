#!/usr/bin/env bash
mkdir -p "cppgen"

for dir in $(find "protocol" -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq); do
    echo "Generating files in ${dir}..."
    find "${dir}" -name '*.proto'

    protoc --experimental_allow_proto3_optional \
    --proto_path=protocol \
    --plugin=protoc-gen-custom=./hrpc \
    --custom_out=./cppgen \
    --custom_opt="qt_cpp_client" \
    --cpp_out=./cppgen \
    $(find "${dir}" -name '*.proto')
done
