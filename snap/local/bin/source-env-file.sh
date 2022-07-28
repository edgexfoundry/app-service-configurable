#!/bin/bash -e

SERVICE="app-service-configurable"
ENV_FILE="$SNAP_DATA/config/res/$SERVICE.env"
TAG="edgex-$SERVICE."$(basename "$0")

if [ -f "$ENV_FILE" ]; then
    logger --tag="$TAG" "sourcing $ENV_FILE"
    set -o allexport
    source "$ENV_FILE" set
    set +o allexport
fi

exec "$@"

