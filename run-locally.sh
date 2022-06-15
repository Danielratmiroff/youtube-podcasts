#!/bin/bash

# Builds docker image and runs it in localhost:8080
docker build -t myapp .
docker run -it -p 8080:9000 myapp
