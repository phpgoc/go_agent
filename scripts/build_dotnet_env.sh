#!/bin/bash
pushd "$(dirname "$0")" >/dev/null || exit 1
cmd_dir=$(pwd)
popd >/dev/null || exit 1

image_name="ahsy-dotnet-test-env"
container_cmd=$(bash "$cmd_dir"/private_docker_or_podman.sh)
if [ -z "$container_cmd" ]; then
    echo "Neither podman nor docker is installed"
    exit 1
fi

# 镜像还没有安装dotnet依赖库的能力

"$cmd_dir"/private_has_docker_buildx.sh "$container_cmd" && \
$container_cmd buildx build -t $image_name "$cmd_dir"/../docker/dotnet/ || \
$container_cmd build -t $image_name "$cmd_dir"/../docker/dotnet/


