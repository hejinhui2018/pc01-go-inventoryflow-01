#!/usr/bin/env bash
set -euo pipefail
IMAGE="${IMAGE:-pc01-inventoryflow:local}"
PLATFORM="${PLATFORM:-linux/amd64}"
docker build --platform "${PLATFORM}" -f benzhi.Dockerfile -t "${IMAGE}" .
