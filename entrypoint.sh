#!/usr/bin/env bash

set -e

mkdir -p /data/

PORT=${PORT:-8080}
DB_URL=${DB_URL:-postgres://postgres:postgres@localhost:5432/monitor?sslmode=disable}
REDIS_HOST=${REDIS_HOST:-redis}
REDIS_PORT=${REDIS_PORT:-6379}
REDIS_DB=${REDIS_DB:-1}

echo '
{
    "port": '$PORT',
    "db_url": "'$DB_URL'", 
    "redis_host": "'$REDIS_HOST'",
    "redis_port": '$REDIS_PORT',
    "redis_db": '$REDIS_DB'
}
' > /data/config.json

date '+%Y/%m/%d %H:%M:%S Configuration file created'

cat /data/config.json

exec "$@"