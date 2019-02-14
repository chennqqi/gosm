#!/bin/bash


if [ "x$1" = "xdev" ]; then
	go build -ldflags '-X main.Version=$(cat VERSION)'
else
  #! /bin/bash
  if command -v packr2 >/dev/null 2>&1; then 
	true;
  else 
	go get -u github.com/gobuffalo/packr/v2/packr2
  fi
	packr2 build -ldflags "-X main.Version=$(cat VERSION)"
fi
