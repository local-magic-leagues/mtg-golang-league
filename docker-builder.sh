#!/bin/bash

IMAGE_NAME="mtg"

echo "  Building from Dockerfile..."
docker build -t $IMAGE_NAME .
echo "  Done"

echo "  Containerizing mtg image"
docker run -d -p 8000:8000 $IMAGE_NAME mtg-golang-league
echo "****  Done. Listening on port 8000    ****"

