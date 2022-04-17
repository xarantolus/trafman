#!/usr/bin/env bash
set -euo pipefail

cd frontend
npm run build

cp -r dist ../app/frontend

cd ..
