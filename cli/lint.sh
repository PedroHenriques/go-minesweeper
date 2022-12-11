#!/bin/sh
set -e

BUILD_DOCKER_IMG=0
if [ "$1" = "build" ]; then
  BUILD_DOCKER_IMG=1
fi

if [ $BUILD_DOCKER_IMG -eq 1  ]; then
  echo "Build the Docker image";
  docker build -f ./docker/Dockerfile-linter --pull --rm -t go-minesweeper-linter:latest .;
fi

docker run --rm -v "${PWD}/":"/usr/src/app/" go-minesweeper-linter:latest /bin/sh -c "go mod tidy && sh ./cli/bundle.sh && golangci-lint run -v";