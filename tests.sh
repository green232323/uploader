#!/bin/bash

set -e
echo Run tests in clientAPI
cd ./clientAPI
go test ./...

echo Run tests in portDomain
cd ../portDomain
go test ./...