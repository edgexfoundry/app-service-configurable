#!/bin/bash -e

PROFILE_OPT=""

profile=$(snapctl get "profile")
if [ ! -z "$profile" ]; then
    if [ "$profile" != "default" ]; then
        PROFILE_OPT="-profile $profile"
        export SECRETSTORE_TOKENFILE="$SNAP_DATA/app-$profile/secrets-token.json"
    fi
fi

EDGEX_STARTUP_DURATION=$(snapctl get startup-duration)
if [ -n "$EDGEX_STARTUP_DURATION" ]; then
  export EDGEX_STARTUP_DURATION
fi

EDGEX_STARTUP_INTERVAL=$(snapctl get startup-interval)
if [ -n "$EDGEX_STARTUP_INTERVAL" ]; then
  export EDGEX_STARTUP_INTERVAL
fi

$SNAP/bin/app-service-configurable -confdir $SNAP_DATA/config/res $PROFILE_OPT -cp -r

