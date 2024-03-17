#!/bin/bash

docker build -t envoy:1.0 .

docker run -d -p 8080:8080 -p 9901:9901 --network=host envoy:1.0

# docker build -t my-grpc-web-client:1.0 -f client/Dockerfile client

# docker run -d -p 3000:3000 my-grpc-web-client:1.0

docker build -t go-grpc-server:1.0 -f server/Dockerfile server

docker run -d  -p 9090:9090 --env-file .env go-grpc-server:1.0 