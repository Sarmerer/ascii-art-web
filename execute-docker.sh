#!/bin/bash
docker build -f Dockerfile -t web .
docker run --name web -p 4241:4241 -d web