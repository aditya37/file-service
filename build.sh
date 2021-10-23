#!/usr/bin/env bash
export TAG=$1
export CGO_ENABLED=0 
export GO111MODULE=off

docker-compose -f docker-compose.dev.yml up -d --build --always-recreate-deps