#!/bin/bash
pushd "$(dirname "$0")" >/dev/null || exit 1
cmd_dir=$(pwd)
popd >/dev/null || exit 1


cp "$cmd_dir"/../go.mod  "$cmd_dir"/../docker/go/

image_name="ahsy-go-agent-env"

"$cmd_dir"/private_has_docker_buildx.sh && \
docker buildx build -t $image_name "$cmd_dir"/../docker/go/ || \
docker build -t $image_name "$cmd_dir"/../docker/go/

rm -rf "$cmd_dir"/../docker/go/go.mod "$cmd_dir"/../docker/go/go.sum

