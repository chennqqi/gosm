#!/bin/bash

./build.sh
sudo docker build -t "sort/gosm:$(cat VERSION)" -f Dockerfile.local .

