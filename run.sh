#!/bin/bash

set -e
echo Create ./data if not exist
mkdir -p ./data
cd ./clientAPI
echo Compile clientAPI service
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -trimpath ./cmd/clientAPI.go

cd ../portDomain
echo Compile portDomain service
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -trimpath ./cmd/portDomain.go

docker-compose up --build
