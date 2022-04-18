#!/usr/bin/env bash
set -euo pipefail

cd frontend

echo "Running frontend in the background, this is where you should connect"
npm run serve &
cd ..

echo "Starting debug database"
docker-compose -f "docker-compose.debug.yml" build && docker-compose -f "docker-compose.debug.yml" up &

echo "Running backend"
cd app
export $(cat ../.env | xargs) && DB_HOST=localhost APP_PORT=8080 APP_DEBUG=true go run . &
cd ..

echo "Press enter to kill all services"
read

kill $(jobs -p)
