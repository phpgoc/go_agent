#!/bin/bash
pushd $(dirname $0) >/dev/null
cmddir=$(pwd)
popd >/dev/null

# 解析参数 -i -o
input=""
project_root=$(realpath $cmddir/../)
proto_dir=$project_root/protos/


while getopts "i:h" opt; do
  case $opt in
    i)
      input=$OPTARG
      ;;
    h)
        echo "Usage: $0 -i input_file -o project_root_dir"
        #输出换行
        echo "" 
        echo "input file if not set, will use all proto files in $protoc_dir"
        protoc_go_dir=$(realpath $cmddir/../protoc_go/)
        echo "output dir  $project_root/proto_go/"
        echo "use this cmd,you must install protoc and protoc-gen-go protoc-gen-go-grpc"
        echo "in Ubuntu,it is easy to install by \"apt install protobuf-compiler protoc-gen-go protoc-gen-go-grpc\""
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
    #find basename
    find_files=$(find $proto_dir -name "*.proto"  -exec basename {} \; | tr '\n' ' ')

    protoc --go_out=$project_root --go-grpc_out=$project_root --proto_path=$proto_dir $find_files
    echo "generate all go code success"
else
    protoc --go_out=$project_root --go-grpc_out=$project_root --proto_path=$proto_dir $input
    echo "generate go code by $input success"
fi
