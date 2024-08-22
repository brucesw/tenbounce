#!/bin/bash

# docker build -t tenbounce-image .

# Get the current Git commit SHA
COMMIT_SHA=$(git rev-parse --short HEAD)

if git diff-index --quiet HEAD --; then
  CLEAN="true"
else
  CLEAN="false"
fi

docker tag tenbounce-image us-central1-docker.pkg.dev/tenbounce-prod/tenbounce/tenbounce:$COMMIT_SHA-$CLEAN

docker push us-central1-docker.pkg.dev/tenbounce-prod/tenbounce/tenbounce