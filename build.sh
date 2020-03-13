#!/bin/bash

RUN_NAME="go_spider"

mkdir output output/${RUN_NAME}_log
cp -r config output/
cp -r data output/
export GO111MODULE=on
go build -a -o output/${RUN_NAME}

# gojieba 不支持交叉编译
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o output/${RUN_NAME}
