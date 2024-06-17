if ! docker buildx ls | grep -q "default"; then
  exit 1
else
  exit 0
fi