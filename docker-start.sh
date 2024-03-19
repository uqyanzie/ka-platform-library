#!/bin/bash

docker run -d -v ${PWD}/envoy.yaml:/etc/envoy/envoy.yaml:ro -p 8080:8080 -p 9901:9901 envoyproxy/envoy:v1.22.0

# docker build -t react-grpc-web-client:1.0 -f client/Dockerfile client

# docker run -d -p 3000:3000 react-grpc-web-client:1.0

docker build -t go-grpc-server:1.0 -f server/Dockerfile server

docker run -d  -p 9090:9090 --env-file server/.env go-grpc-server:1.0 
