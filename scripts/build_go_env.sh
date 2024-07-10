#!/bin/bash
pushd "$(dirname "$0")" >/dev/null || exit 1
cmd_dir=$(pwd)
popd >/dev/null || exit 1

container_cmd=$(bash "$cmd_dir"/private_docker_or_podman.sh)
if [ -z "$container_cmd" ]; then
    echo "Neither podman nor docker is installed"
    exit 1
fi

cp "$cmd_dir"/../go.mod  "$cmd_dir"/../docker/go/

image_name="ahsy-go-agent-env"

"$cmd_dir"/private_has_docker_buildx.sh "$container_cmd" && \
$container_cmd buildx build -t $image_name "$cmd_dir"/../docker/go/ || \
$container_cmd build -t $image_name "$cmd_dir"/../docker/go/

rm -rf "$cmd_dir"/../docker/go/go.mod "$cmd_dir"/../docker/go/go.sum

