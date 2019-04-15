#!/usr/bin/env bash

PROJECT="go-echo-server"
TAG="${PROJECT}:latest"

PORT="18080"

echo "removing old container"
docker rm $(docker stop $(docker ps -a -q --filter ancestor=${TAG} --format="{{.ID}}"))
echo

echo "building new image"
docker build --rm=true -t ${TAG} .
echo

echo "running new container with:"
echo -e "\tdocker run -p ${PORT}:${PORT} --name ${PROJECT} -d ${TAG}"
echo
