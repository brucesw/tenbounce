#!/bin/bash

docker build -t tenbounce-image .

docker tag tenbounce-image us-central1-docker.pkg.dev/tenbounce-prod/tenbounce/tenbounce:$COMMIT_SHA

docker push us-central1-docker.pkg.dev/tenbounce-prod/tenbounce/tenbounce