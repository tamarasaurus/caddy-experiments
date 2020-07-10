#!/bin/bash

docker build --no-cache -f Dockerfile . --tag caddy-experiments && docker run -p 8080:8080 caddy-experiments