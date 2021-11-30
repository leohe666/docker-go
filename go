#!/bin/bash

source ./config.properties
docker run --rm -it \
    -u $UID:$UID \
    -e XDG_CACHE_HOME=/tmp/.cache \
    -e GOPROXY=${GOPROXY} \
    -v $PWD:/srv/app \
    -v $PWD/packages:/go \
    -w /srv/app \
    -p ${LOCALPORT}:8080 \
    golang:alpine go $@