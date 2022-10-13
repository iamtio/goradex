#!/bin/bash

if [ ! $1 ]; then
    echo "Usage $0 <image_name> [tag]"
    exit 1
fi
ARCH=arm
PLATFORM=linux/$ARCH
IMAGE=$1:${2:-latest}

docker buildx create --name rpi --platform $PLATFORM 2>&1 > /dev/null
docker buildx use rpi
docker buildx build -t $IMAGE --platform=$PLATFORM --load .
docker image save -o "./goradex_docker.$ARCH.tar.gz" $IMAGE
