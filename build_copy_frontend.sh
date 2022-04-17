#!/usr/bin/env bash
set -euo pipefail

cd frontend
npm run build

rm -rf ../app/frontend
cp -r dist ../app/frontend

cd ..
