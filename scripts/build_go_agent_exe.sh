#!/bin/bash
pushd "$(dirname "$0")" >/dev/null || exit 1
cmd_dir=$(pwd)
popd >/dev/null || exit 1

build_image_name="ahsy-go-agent-env"

set -e

docker run --rm -v "$cmd_dir"/../:/go/src/ "$build_image_name" \
  sh -c "CGO_ENABLED=0 go build -o  bin/go-agent cmd/go-agent/main.go"

echo "Build go-agent success"
echo "Output: $(realpath "$cmd_dir"/../bin)/go-agent"