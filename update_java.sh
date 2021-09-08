#!/bin/bash
if ! command -v docker-compose &> /dev/null
then
    COMPOSE="docker compose"
else
    COMPOSE="docker-compose"
fi

$COMPOSE build --pull $@
echo "Don't forget to stop the server and run \`docker compose up -d\` again if there was a new java version."