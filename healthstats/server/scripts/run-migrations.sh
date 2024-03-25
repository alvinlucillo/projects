#!/bin/sh

# Wait for the database to be ready
dockerize -wait tcp://${DB_HOST}:${DB_PORT} -timeout 60s

# Run migrations
echo "y" | /go/bin/migrate -source file://files -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable ${MIGRATE_COMMAND}