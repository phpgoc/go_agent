#!/bin/bash
pushd "$(dirname "$0")" >/dev/null
cmd_dir=$(pwd)
popd >/dev/null

# 解析参数 -i -o
input=""
simple_proto_path="protos"
project_root=$(realpath "$cmd_dir"/../)
proto_dir=$project_root/$simple_proto_path
image_name="ahsy-go-agent-env"

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
else
   find_files=$(find "$proto_dir" -name "*.proto"  -exec basename {} \; | grep "$input" |  tr '\n' ' ')
#    protoc --go_out="$project_root" --go-grpc_out="$project_root" --proto_path="$proto_dir" $input

fi
echo "gen these: ""$find_files"
docker run --rm -v "$project_root":/go/src -w /go/src $image_name \
  protoc --go_out=/go/src --go-grpc_out=/go/src --proto_path=/go/src/$simple_proto_path $find_files
echo "generate all go code success"