#!/bin/bash
pushd "$(dirname "$0")" >/dev/null || exit 1
cmd_dir=$(pwd)
popd >/dev/null || exit 1

build_image_name="ahsy-go-agent-env"


docker run --rm -ti -v "$cmd_dir"/../:/go/src/ "$build_image_name" bash

