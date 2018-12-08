#!/bin/bash

pushd ./cmd/server
go build
popd

./cmd/server/server

