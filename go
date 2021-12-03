#!/bin/bash

docker run --rm -it \
    -u $UID:$UID \
    -e TZ="Asia/Shanghai" \
    -e XDG_CACHE_HOME=/tmp/.cache \
    -e GOPROXY=https://goproxy.cn \
    -v $PWD:/srv/app \
    -v $PWD/packages:/go \
    -w /srv/app \
    `# -p 8080:8080` \
    golang:alpine go $@