#!/bin/bash
docker compose build --pull "$@"
echo "Don't forget to stop the server and run \`docker compose up -d\` again if there was a new java version."