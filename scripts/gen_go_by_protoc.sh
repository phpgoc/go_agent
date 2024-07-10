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

container_cmd=$(bash "$cmd_dir"/private_docker_or_podman.sh)
if [ -z "$container_cmd" ]; then
    echo "Neither podman nor docker is installed"
    exit 1
fi

while getopts "i:h" opt; do
  case $opt in
    i)
      input=$OPTARG
      ;;
    h)
        echo "Usage: $0 [-i input_file]"
        #输出换行
        echo "" 
        # shellcheck disable=SC2154
        echo " if not set input file, will use all proto files in ""$proto_dir"
        echo "input use grep, hello will match helloworld.proto"
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
fi
echo "gen these: ""$find_files"
$container_cmd run --rm -v "$project_root":/go/src -w /go/src $image_name \
  protoc --go_out=/go/src --go-grpc_out=/go/src --proto_path=/go/src/$simple_proto_path $find_files
echo "generate all go code success"