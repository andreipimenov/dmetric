#!/usr/bin/env bash

set -e

mkdir -p data

METRIC1_MIN=${METRIC1_MIN:-0}
METRIC1_MAX=${METRIC1_MAX:-10}
METRIC2_MIN=${METRIC2_MIN:-0}
METRIC2_MAX=${METRIC2_MAX:-10}
METRIC3_MIN=${METRIC3_MIN:-0}
METRIC3_MAX=${METRIC3_MAX:-10}
METRIC4_MIN=${METRIC4_MIN:-0}
METRIC4_MAX=${METRIC4_MAX:-10}
METRIC5_MIN=${METRIC5_MIN:-0}
METRIC5_MAX=${METRIC5_MAX:-10}

PORT=${PORT:-8080}
DB_URL=${DB_URL:-postgres://postgres:postgres@postgres:5432/monitor?sslmode=disable}
REDIS_HOST=${REDIS_HOST:-redis}
REDIS_PORT=${REDIS_PORT:-6379}
REDIS_DB=${REDIS_DB:-1}
SMTP_HOST=${SMTP_HOST:-smtp.yandex.ru}
SMTP_PORT=${SMTP_PORT:-587}
SMTP_LOGIN=${SMTP_LOGIN:-dmetric@yandex.ru}
SMTP_PASSWORD=${SMTP_PASSWORD:-Dmetric/}
MAIL_TO=${MAIL_TO:-dmetric@yandex.ru}

echo '
{
    "metric_limits": {
        "metric_1": {
            "min": '$METRIC1_MIN',
            "max": '$METRIC1_MAX'
        },
        "metric_2": {
            "min": '$METRIC2_MIN',
            "max": '$METRIC2_MAX'
        },
        "metric_3": {
            "min": '$METRIC3_MIN',
            "max": '$METRIC3_MAX'
        },
        "metric_4": {
            "min": '$METRIC4_MIN',
            "max": '$METRIC4_MAX'
        },
        "metric_5": {
            "min": '$METRIC5_MIN',
            "max": '$METRIC5_MAX'
        }
    },
    "port": '$PORT',
    "db_url": "'$DB_URL'", 
    "redis_host": "'$REDIS_HOST'",
    "redis_port": '$REDIS_PORT',
    "redis_db": '$REDIS_DB',
    "smtp_host": "'$SMTP_HOST'",
    "smtp_port": '$SMTP_PORT',
    "smtp_login": "'$SMTP_LOGIN'",
    "smtp_password": "'$SMTP_PASSWORD'",
    "mail_to": "'$MAIL_TO'"
}
' > data/config.json

date '+%Y/%m/%d %H:%M:%S Configuration file created'

cat data/config.json

exec "$@"