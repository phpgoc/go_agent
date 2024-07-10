#!/bin/bash

# Check if podman is installed
if command -v podman &> /dev/null; then
    echo "podman"
# Check if docker is installed
elif command -v docker &> /dev/null; then
    echo "docker"
# Neither podman nor docker is installed
else
    echo ""
fi