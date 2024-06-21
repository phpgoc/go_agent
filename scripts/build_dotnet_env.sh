#!/bin/bash
pushd "$(dirname "$0")" >/dev/null || exit 1
cmd_dir=$(pwd)
popd >/dev/null || exit 1

image_name="ahsy-dotnet-test-env"

# 镜像还没有安装dotnet依赖库的能力

"$cmd_dir"/private_has_docker_buildx.sh && \
docker buildx build -t $image_name "$cmd_dir"/../docker/dotnet/ || \
docker build -t $image_name "$cmd_dir"/../docker/dotnet/


