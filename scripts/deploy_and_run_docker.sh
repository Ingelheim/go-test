#!/bin/bash

docker build -t go-test .
docker run -t -i -p 8080:8080 go-test
