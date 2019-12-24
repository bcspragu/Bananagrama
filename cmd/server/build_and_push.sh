#!/bin/bash
set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

CGO_ENABLED=0 GOOS=linux go build -ldflags '-extldflags "-static" -s -w' -o server $DIR

docker build -t docker.bsprague.com/bananagrams $DIR

docker push docker.bsprague.com/bananagrams
