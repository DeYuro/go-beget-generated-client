#!/bin/bash

prg=$0
if [ ! -e "$prg" ]; then
  case $prg in
    (*/*) exit 1;;
    (*) prg=$(command -v -- "$prg") || exit;;
  esac
fi
dir=$(
  cd -P -- "$(dirname -- "$prg")" && pwd -P
) || exit
prg=$dir/$(basename -- "$prg") || exit

dr=$(dirname "$prg")

rm -rf $dr/../internal/apiclient/vps
rm -rf $dr/../internal/apiclient/auth
mkdir -p $dr/../internal/apiclient/vps
mkdir -p $dr/../internal/apiclient/auth

docker run \
 --rm \
  -v $dr/../api/vps:/source \
  -v $dr/../internal/apiclient/vps:/local \
  --user $(id -u):$(id -g) \
  openapitools/openapi-generator-cli generate -g go -i /source/openapi.yaml \
  -o /local \

docker run \
 --rm \
  -v $dr/../api/auth:/source \
  -v $dr/../internal/apiclient/auth:/local \
  --user $(id -u):$(id -g) \
  openapitools/openapi-generator-cli generate -g go -i /source/openapi.yaml \
  -o /local \

rm $dr/../internal/apiclient/vps/go.mod
rm $dr/../internal/apiclient/auth/go.mod