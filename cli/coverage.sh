#!/bin/sh
set -e

BUILD_DOCKER_IMG=0
DIRS="";

while [ "$#" -gt 0 ]; do
  case "$1" in
    -b|--build) BUILD_DOCKER_IMG=1; shift 1;;

    -*) echo "unknown option: $1" >&2; exit 1;;
    *) DIRS="$DIRS $1"; shift 1;;
  esac
done

if [ "$DIRS" = "" ]; then
  for dir in ./internal/*/ ; do
    DIRS="$DIRS $dir";
  done
fi

if [ $BUILD_DOCKER_IMG -eq 1  ]; then
  echo "Build the Docker image";
  docker build -f ./docker/Dockerfile-linux --pull --rm -t go-minesweeper-linux:latest .;
fi

mkdir -p ./coverage/;

docker run --rm -v "${PWD}/":"/usr/src/app/" go-minesweeper-linux:latest /bin/sh -c "go test -coverprofile coverage/coverage.out $DIRS && go tool cover -html coverage/coverage.out -o coverage/coverage.html && gcov2lcov -infile=coverage/coverage.out -outfile=coverage/coverage.lcov";
