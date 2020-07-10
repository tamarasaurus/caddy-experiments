#!/bin/bash

docker build --no-cache -f Dockerfile . --tag caddy-experiments-1 && docker run -p 8080:8080 caddy-experiments-1