#!/bin/sh

set -a
source .env
set +a

for file in internal/migrations/*.up.sql
do
    PGPASSWORD="$POSTGRES_PASSWORD" \
    psql \
    "sslmode=require host=$POSTGRES_HOST port=$POSTGRES_PORT dbname=$POSTGRES_DB user=$POSTGRES_USER" \
    -f "$file"
done