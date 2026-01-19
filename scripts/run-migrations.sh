#!/bin/bash

# Migration Runner Script
# This script runs all pending migrations

set -e

# Database configuration
DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}
DB_NAME=${DB_NAME:-api_web_scrapping}
DB_USER=${DB_USER:-admin}
DB_PASSWORD=${DB_PASSWORD:-secret}

# Migration directory
MIGRATIONS_DIR="./migrations"

echo "====================================="
echo "Database Migration Runner"
echo "====================================="
echo "Host: $DB_HOST:$DB_PORT"
echo "Database: $DB_NAME"
echo "User: $DB_USER"
echo "====================================="

# Check if psql is installed
if ! command -v psql &> /dev/null; then
    echo "Error: psql is not installed"
    exit 1
fi

# Create database if it doesn't exist
echo "Creating database if not exists..."
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d postgres <<EOF
SELECT 'CREATE DATABASE $DB_NAME' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '$DB_NAME')\\gexec
EOF

# Run migrations in order
echo "Running migrations..."
for migration in $MIGRATIONS_DIR/*.up.sql; do
    if [ -f "$migration" ]; then
        filename=$(basename "$migration")
        echo "Running migration: $filename"
        PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f "$migration"
        echo "✓ Migration completed: $filename"
    fi
done

echo "====================================="
echo "All migrations completed successfully!"
echo "====================================="
