#!/bin/bash

./scripts/make_app.sh

export GOPATH=$(pwd)
go install test-server/server
bin/server
