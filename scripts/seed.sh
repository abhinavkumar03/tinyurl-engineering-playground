#!/bin/sh

set -a
source .env
set +a

PGPASSWORD="$POSTGRES_PASSWORD" \
psql \
"sslmode=require host=$POSTGRES_HOST port=$POSTGRES_PORT dbname=$POSTGRES_DB user=$POSTGRES_USER" \
-f internal/migrations/000002_seed_sample_urls.up.sql