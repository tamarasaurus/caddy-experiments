#!/bin/bash

docker build -f Dockerfile . --tag caddy-experiments && docker run -p 8080:8080 caddy-experiments