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

$container_cmd run --rm -v "$cmd_dir/../":/app -w /app/dotnet_tests/CallEverything $image_name dotnet run