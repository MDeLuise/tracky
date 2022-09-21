#!/bin/bash

if [ $# -eq 0 ]
  then
    echo "No release version supplied";
    return 1;
fi

RELEASE_VERSION=$1;

cd "$(dirname "$0")";

# Backend
cd ../backend;
mvn package;
docker build -t msdeluise/tracky-backend:$RELEASE_VERSION -f deployment/Dockerfile .;

# Frontend
cd ../frontend;
docker build -t msdeluise/tracky-frontend:$RELEASE_VERSION -f deployment/Dockerfile .;
