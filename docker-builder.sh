#!/bin/bash
set -e

IMAGE_NAME="golang"

# download and create a golang docker container and gets the mtg-golang-league repo
echo "  Creating a golang container..."
docker run $IMAGE_NAME go get -v github.com/local-magic-leagues/mtg-golang-league

# creates a new docker image called mtg-golang-league and commits it
echo "  Committing mtg-golang-container..."
docker commit $(docker ps -lq) mtg-golang-league

# runs the mtg-golang-league image that listens on port 8000
echo "  Running mtg-golang-container..."
docker run -d -p 8000:8000 mtg-golang-league mtg-golang-league
echo "mtg-golang-league container is running and listening on port 8000..."
