#! /bin/bash

container_cmd=$1
if [ -z "$container_cmd" ]; then
    container_cmd="docker"
fi
# podman 不需要判断
if [ "$container_cmd" == "podman" ]; then
  exit 0
fi
if  $container_cmd buildx ls | grep -q "default"; then
  exit 0
else
  exit 1
fi