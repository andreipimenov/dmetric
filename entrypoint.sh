#!/usr/bin/env bash

set -e

mkdir -p /data/

PORT=${PORT:-8080}

echo '
{
    "port": '$PORT'
}
' > /data/config.json

date '+%Y/%m/%d %H:%M:%S Configuration file created'

cat /data/config.json

exec "$@"