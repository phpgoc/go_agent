#!/bin/bash
pushd "$(dirname "$0")" >/dev/null || exit 1
cmd_dir=$(pwd)
popd >/dev/null || exit 1

build_image_name="ahsy-go-agent-env"
p=""

container_cmd=$(bash "$cmd_dir"/private_docker_or_podman.sh)
if [ -z "$container_cmd" ]; then
    echo "Neither podman nor docker is installed"
    exit 1
fi

while getopts "p:h" opt; do
  case $opt in
    p)
      p="-p $OPTARG:$OPTARG"
      ;;
    h)
        echo "Usage: $0 [-p port]"
        #输出换行
        echo ""
        echo " if not set port, will use not use port option"
        echo " notice the port must not be used by other process"
        exit 0
        ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      exit 1
      ;;
  esac
done


$container_cmd run --rm -ti -v "$cmd_dir"/../:/go/src/  $p "$build_image_name" bash

