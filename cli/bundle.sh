#!/bin/sh
set -e

cd ./assets/images/;

APPEND=0;

for file in * ; do
  if [ $APPEND -eq 1 ]; then
    fyne bundle -a -o ../../internal/gui/fyneBundle.go $file;
  else
    fyne bundle --pkg gui -o ../../internal/gui/fyneBundle.go $file;

    APPEND=1;
  fi
done