[![Coverage Status](https://coveralls.io/repos/github/PedroHenriques/go-minesweeper/badge.svg?branch=main)](https://coveralls.io/github/PedroHenriques/go-minesweeper?branch=main)
![ci workflow](https://github.com/PedroHenriques/go-minesweeper/actions/workflows/ci.yml/badge.svg?branch=main)
![cd workflow](https://github.com/PedroHenriques/go-minesweeper/actions/workflows/cd.yml/badge.svg)

# Go Minesweeper

## How to play

Your goal is to reveal all the tiles that don't have a mine.
Revealed tiles that are adjacent to mines will have a number. This number indicates how many mines are in adjacent tiles.

- left mouse button: reveals a tile
- right mouse button: adds/removes a flag from a tile, preventing it from being revealed
- left + right mouse buttons (on a revealed tile): reveales all adjacent tiles, only if enough flags are placed. This function prevents acidental mine hits.

## Binaries

You can download the binaries [here](http://pedrojhenriques.com/games/go-minesweeper/)

## Building the binaries

### Prerequisites

- **Docker:** Install documentation [here](https://docs.docker.com/get-docker/)

### Building the binaries

On a terminal, from the root of the repo, run
```sh
sh cli/build.sh [OPTIONS]

Options:
-b, --build: will build the docker image used to compile the code
--version: version of the build. Will be used to name the binary file.
```

The binaries will be available on the directory `bin/`

## Run the game from source

### Prerequisites

1. Make sure you have the latest version of golang installed

2. Install Fyne's dependencies for your OS. Consult them [here](https://developer.fyne.io/started/#prerequisites)

### Running the game

On a terminal, from the root of the repo, run
```sh
go run ./main.go
```

## Development tools

### Updating the asset bundle file

On a terminal, from the root of the repo, run
```sh
sh cli/bundle.sh
```

The script will catch all files inside `assets/images/`

### Running the linters

On a terminal, from the root of the repo, run
```sh
sh cli/lint.sh [OPTIONS]

Options:
-b, --build: will build the docker image used to run the linters
```

### Running the tests

On a terminal, from the root of the repo, run
```sh
sh cli/test.sh [OPTIONS] [dir1 dir2 ...]

Options:
-b, --build: will build the docker image used to run the tests
-w, --watch: run the tests in watch mode

dir1 dir2 ...: the directories to look for test files. Defaults to all subdirectories of internal/
```

### Running the test coverage

On a terminal, from the root of the repo, run
```sh
sh cli/coverage.sh [OPTIONS]

Options:
-b, --build: will build the docker image used to run the tests
```

The output will be available on the directory `coverage/`