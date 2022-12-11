#!/bin/sh
set -e

# Generate the Fyne bundle file with all the necessary assets
sh ./cli/bundle.sh;

# Compile and run the application
go run ./main.go;