#!/bin/bash
pushd "$(dirname "$0")" >/dev/null || exit 1
cmd_dir=$(pwd)
popd >/dev/null || exit 1

image_name="ahsy-go-agent-env"

# 手动 输入需要测试的包
docker run --rm -v "$cmd_dir"/../:/go/src/ "$image_name" \
                  go test go-agent/utils