#!/bin/sh

# exist immediately if the command returns non zero
set -e

echo "Running DB migrations..."

/app/migrate -path /app/migrations -database "$DB_TABLE://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_DB?sslmode=$SSL_MODE" -verbose up

echo "Starting the app..."

# take all the params passed to this script and run it
# runs ./app
exec "$@"