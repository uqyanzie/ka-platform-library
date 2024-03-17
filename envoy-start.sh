#!/bin/bash

docker build -t my-envoy:1.0 .

docker run -d -p 8080:8080 -p 9901:9901 --network=host my-envoy:1.0