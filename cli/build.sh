#!/bin/sh
set -e

BUILD_DOCKER_IMG=0
if [ "$1" = "build" ];
then
  BUILD_DOCKER_IMG=1
fi

mkdir -p ./bin/linux/;
mkdir -p ./bin/windows/;
mkdir -p ./bin/macos/;

echo "Building app for the current OS";

if [ $BUILD_DOCKER_IMG -eq 1  ];
then
  echo "Build the Docker image";
  docker build -f ./docker/Dockerfile-linux --pull --rm -t go-minesweeper-linux:latest .;
fi

echo "Running Docker container to build app";
docker run --rm -v "${PWD}/bin/":"/usr/src/app/bin/" go-minesweeper-linux:latest /bin/sh -c "go build -o ./bin/linux/ && zip -jm ./bin/linux/go-minesweeper.zip ./bin/linux/go-minesweeper";

echo "Finished building app"
