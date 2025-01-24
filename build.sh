#!/bin/bash

rm -rf bin
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/server ./
chmod +x bin/server

docker build --push -t shovan1995/go-sample-app:v5 .
