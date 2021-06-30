#!/bin/bash -e

EDGEX_STARTUP_DURATION=$(snapctl get startup-duration)

if [ -n "$EDGEX_STARTUP_DURATION" ]; then
  export EDGEX_STARTUP_DURATION
fi

EDGEX_STARTUP_INTERVAL=$(snapctl get startup-interval)

if [ -n "$EDGEX_STARTUP_INTERVAL" ]; then
  export EDGEX_STARTUP_INTERVAL
fi

SERVICE_ENV="$SNAP_DATA/config/res/app-service-configurable.env"
logger "edgex service override: : SERVICE_ENV=$SERVICE_ENV"

if [ -f "$SERVICE_ENV" ]; then
    logger "edgex service override: : sourcing $SERVICE_ENV"
    source "$SERVICE_ENV"
fi

exec "$@"
