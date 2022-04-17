#!/usr/bin/env bash
set -euo pipefail

cd frontend
if [ ! -d "node_modules" ]; then
    npm install
fi
npm run build

rm -rf ../app/frontend
cp -r dist ../app/frontend

cd ..
