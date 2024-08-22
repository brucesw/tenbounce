#!/bin/bash

# Get the current Git commit SHA
COMMIT_SHA=$(git rev-parse --short HEAD)

# Save the SHA to a file
echo -n $COMMIT_SHA > commit_sha.txt

if git diff-index --quiet HEAD --; then
  echo -n "true" > clean.txt
else
  echo -n "false" > clean.txt
fi
