#!/bin/bash
set -e

# Wait for the database to be ready
while ! mysqladmin ping -h"db" --silent; do
    echo "Waiting for database connection..."
    sleep 2
done

# Run the SQL script to initialize the database
mysql -h db -uroot -p$MYSQL_ROOT_PASSWORD < /docker-entrypoint-initdb.d/create_database.sql

# Execute the original command
exec "$@"

