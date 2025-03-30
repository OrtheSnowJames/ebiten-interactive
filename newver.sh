#!/bin/bash

# Check if the correct number of arguments is provided
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <commit-message> <tag>"
    exit 1
fi

COMMIT_MESSAGE=$2
TAG=$1

# Stage all changes
git add .

# Commit with the provided message
git commit -m "$COMMIT_MESSAGE"

# Create a new tag
git tag "$TAG"

# Push the commit and the tag
git push
git push origin "$TAG"