#!/bin/bash

# Migration Rollback Script
# This script rolls back the last migration

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
echo "Migration Rollback"
echo "====================================="
echo "Host: $DB_HOST:$DB_PORT"
echo "Database: $DB_NAME"
echo "====================================="

# List all rollback migrations
echo "Available rollback migrations:"
ls -1 $MIGRATIONS_DIR/*.down.sql | while read file; do
    echo "  - $(basename $file)"
done

echo ""
read -p "Enter migration filename to rollback (e.g., 000005_seed_admin_user.down.sql): " migration_file

# Check if file exists
if [ ! -f "$MIGRATIONS_DIR/$migration_file" ]; then
    echo "Error: Migration file not found: $migration_file"
    exit 1
fi

# Confirm rollback
read -p "Are you sure you want to rollback $migration_file? (yes/no): " confirm
if [ "$confirm" != "yes" ]; then
    echo "Rollback cancelled"
    exit 0
fi

# Run rollback
echo "Rolling back: $migration_file"
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f "$MIGRATIONS_DIR/$migration_file"

echo "====================================="
echo "Rollback completed successfully!"
echo "====================================="
