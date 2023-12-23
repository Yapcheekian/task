#!/bin/bash

set -e

export LOCAL_GOPATH=${GOPATH:-~/go}

[ ! -d "$LOCAL_GOPATH/pkg/mod/cache" ] && mkdir -p "$LOCAL_GOPATH/pkg/mod/cache"

case "$1" in
  test)
    docker-compose run --rm test
    ;;
  build)
    docker-compose run --rm build
    ;;
  task)
    docker-compose run -p 8080:8080 --rm task
    ;;
  *)
    echo "\"$1\" is an unknown command"
    ;;
esac
