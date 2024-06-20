#!/bin/bash
pushd "$(dirname "$0")" >/dev/null
cmd_dir=$(pwd)
popd >/dev/null

# 解析参数 -i -o
input=""
project_root=$(realpath "$cmd_dir"/../)
simple_proto_path=protos/
proto_dir=$project_root/$simple_proto_path


while getopts "i:h" opt; do
  case $opt in
    i)
      input=$OPTARG
      ;;
    h)
        echo "Usage: $0 -i input_file -o project_root_dir"
        #输出换行
        echo "" 
        # shellcheck disable=SC2154
        echo "input file if not set, will use all proto files in $protoc_dir"
        echo "output dir $(realpath "$cmd_dir"/../protoc_go/)/proto_go/"
        echo "use this cmd,you must install protoc and protoc-gen-go protoc-gen-go-grpc"
        echo "in Ubuntu,it is easy to install by
          sudo apt install protobuf-compiler protoc-gen-go protoc-gen-go-grpc"
        echo "or use go:
          go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
          go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
          export PATH=\"\$PATH:\$(go env GOPATH)/bin\""
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
    find_files=$(find "$proto_dir" -name "*.proto"  -exec basename {} \; | tr '\n' ' ')
#    find_files=$(find "$proto_dir" -name "*.proto" -exec basename {} \; | awk -v P=$simple_proto_path '{print P$1}' | tr '\n' ' ')``
#    echo $find_files
#echo $proto_dir
#echo $project_root
    protoc --go_out="$project_root" --go-grpc_out="$project_root"  --proto_path="$proto_dir" $find_files
    echo "generate all go code success"
else
    protoc --go_out="$project_root" --go-grpc_out="$project_root" --proto_path="$proto_dir" $input
    echo "generate go code by $input success"
fi
