#!/bin/sh
set -e

BUILD_DOCKER_IMG=0

while [ "$#" -gt 0 ]; do
  case "$1" in
    -b|--build) BUILD_DOCKER_IMG=1; shift 1;;

    -*) echo "unknown option: $1" >&2; exit 1;;
    *) shift 1;;
  esac
done

if [ $BUILD_DOCKER_IMG -eq 1  ]; then
  echo "Build the Docker image";
  docker build -f ./docker/Dockerfile-linter --pull --rm -t go-minesweeper-linter:latest .;
fi

docker run --rm -v "${PWD}/":"/usr/src/app/" go-minesweeper-linter:latest /bin/sh -c "go mod tidy && golangci-lint run -v";