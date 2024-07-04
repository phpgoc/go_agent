#!/bin/bash
pushd "$(dirname "$0")" >/dev/null || exit 1
cmd_dir=$(pwd)
popd >/dev/null || exit 1

image_name="ahsy-dotnet-test-env"

docker run --rm -v "$cmd_dir/../":/app -w /app/dotnet_tests/CallEverything $image_name dotnet run