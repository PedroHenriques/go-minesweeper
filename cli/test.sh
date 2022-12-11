#!/bin/sh
set -e

BUILD_DOCKER_IMG=0
DIRS="";
WATCH=0;

if [ "$1" != "" ]; then
  if [ "$1" = "build" ]; then
    BUILD_DOCKER_IMG=1
  else
    if [ "$1" = "-w" ]; then
      WATCH=1
    elif [ "$2" = "" ]; then
      DIRS=$1;
    fi
  fi
fi

if [ $BUILD_DOCKER_IMG -eq 1  ]; then
  echo "Build the Docker image";
  docker build -f ./docker/Dockerfile-linux --pull --rm -t go-minesweeper-linux:latest .;
fi

if [ "$2" != "" ]; then
  DIRS=$2;
fi

if [ "$DIRS" = "" ]; then
  for dir in ./internal/*/ ; do
    DIRS="$DIRS $dir";
  done
fi

if [ $WATCH -eq 1 ]; then
  docker run --rm -v "${PWD}/":"/usr/src/app/" go-minesweeper-linux:latest /bin/sh -c "gow -c test -v -cover $DIRS";
else
  docker run --rm -v "${PWD}/":"/usr/src/app/" go-minesweeper-linux:latest /bin/sh -c "go test -v -cover $DIRS";
fi
