#!/bin/bash
pushd "$(dirname "$0")" >/dev/null || exit 1
cmd_dir=$(pwd)
popd >/dev/null || exit 1

build_image_name="ahsy-go-agent-env"

container_cmd=$(bash "$cmd_dir"/private_docker_or_podman.sh)
if [ -z "$container_cmd" ]; then
    echo "Neither podman nor docker is installed"
    exit 1
fi

set -e

$container_cmd run --rm -v "$cmd_dir"/../:/go/src/ "$build_image_name" \
  sh -c "CGO_ENABLED=0 go build -ldflags \"-s -w\" -o  bin/go-agent cmd/go-agent/main.go"

echo "Build go-agent success"
echo "Output: $(realpath "$cmd_dir"/../bin)/go-agent"