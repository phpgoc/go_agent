#!/bin/bash
pushd $(dirname $0) >/dev/null
cmddir=$(pwd)
popd >/dev/null

# 解析参数 -i -o
input=""
output=""
proto_dir=$(realpath $cmddir/../protos/)
output=$(realpath $cmddir/../)

while getopts "i:h" opt; do
  case $opt in
    i)
      input=$OPTARG
      ;;
    h)
        echo "Usage: $0 -i input_file -o output_dir"
        echo "input file if not set, will use all proto files in $protoc_dir"
        protoc_go_dir=$(realpath $cmddir/../protoc_go/)
        echo "output dir  $protoc_go_dir/proto_go/"
        echo "use this cmd,you must install protoc and protoc-gen-go"
        echo "in Ubuntu,it is easy to install by \"apt install protobuf-compiler protoc-gen-go\""
        exit 0
        ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      exit 1
      ;;
  esac
done

set -e
if [ -z "$input" ]; then
    echo $proto_dir
    #find basename
    find_files=$(find $proto_dir -name "*.proto"  -exec basename {} \; | tr '\n' ' ')
    echo $find_files
    echo $output
    protoc --go_out=$output --proto_path=$proto_dir $find_files
else
    protoc --go_out=$output --proto_path=$proto_dir $input
fi

echo "generate go code success"