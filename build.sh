#!/usr/bin/env bash
set -euo pipefail

COMPOSE_FILE="${1-docker-compose.yml}"

echo "Using $COMPOSE_FILE for building"

echo "Building frontend"
cd frontend
if [ ! -d "node_modules" ]; then
    npm install
fi
npm run build

echo "Copying frontend to app"
cp -r dist ../app/frontend
cd ..

echo "Initiating docker-compose build"
docker-compose -f "$COMPOSE_FILE" build

echo "Running service"
docker-compose -f "$COMPOSE_FILE" up
