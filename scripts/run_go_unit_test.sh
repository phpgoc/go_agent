#!/bin/bash
pushd "$(dirname "$0")" >/dev/null || exit 1
cmd_dir=$(pwd)
popd >/dev/null || exit 1

image_name="ahsy-go-agent-env"

container_cmd=$(bash "$cmd_dir"/private_docker_or_podman.sh)
if [ -z "$container_cmd" ]; then
    echo "Neither podman nor docker is installed"
    exit 1
fi

# 手动 输入需要测试的包
$container_cmd run --rm -v "$cmd_dir"/../:/go/src/ "$image_name" \
                  go test go-agent/utils go-agent/services/apache go-agent/services/nginx