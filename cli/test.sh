#!/bin/sh
set -e

BUILD_DOCKER_IMG=0
WATCH=0;
DIRS="";

while [ "$#" -gt 0 ]; do
  case "$1" in
    -b|--build) BUILD_DOCKER_IMG=1; shift 1;;
    -w|--watch) WATCH=1; shift 1;;

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

CMD="go test -v -cover $DIRS";
DOCKER_FLAGS="";
if [ $WATCH -eq 1 ]; then
  CMD="gow -c test -v -cover $DIRS";
  DOCKER_FLAGS="-it"; # Can not be always added since these docker flags are not supported in Github actions
fi

docker run --rm $DOCKER_FLAGS -v "${PWD}/":"/usr/src/app/" go-minesweeper-linux:latest /bin/sh -c "$CMD";
