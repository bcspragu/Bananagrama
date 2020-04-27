#!/bin/bash
docker run --rm -it \
  -w /project \
  --volume $PWD:/project \
  --user $(id -u):$(id -g) \
  node:alpine /bin/sh
