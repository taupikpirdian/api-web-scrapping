# Database Migrations

This directory contains SQL migration files for the database schema.

## Migration Files

### Migration Naming Convention
```
{version}_{description}.up.sql   - Apply migration
{version}_{description}.down.sql - Rollback migration
```

### Available Migrations

1. **000001_create_users_table**
   - Creates users table with authentication fields
   - Includes email, password, full_name
   - Adds indexes for performance

2. **000002_create_refresh_tokens_table**
   - Creates refresh_tokens table for JWT refresh tokens
   - Tracks token expiration and revocation
   - Stores device info and IP address

3. **000003_create_password_resets_table**
   - Creates password_resets table for password reset functionality
   - Tracks reset token usage and expiration

4. **000004_create_audit_logs_table**
   - Creates audit_logs table for activity tracking
   - Records user actions with IP and user agent
   - Stores old/new values as JSON

5. **000005_seed_admin_user**
   - Creates default admin user
   - **IMPORTANT**: Change password after first login!

6. **000006_create_emitens_table**
   - Creates emitens table for stock/securities data
   - Includes symbol, name, sector
   - Auto-update timestamp trigger

7. **000007_seed_emitens**
   - Seeds Indonesian stock market data (IHSG)
   - Includes banking, tech, consumer goods, energy, etc.
   - 40+ sample emitens

8. **000008_create_stock_price_ticks_table**
   - Creates stock_price_ticks table for historical price data
   - Unique constraint on (emiten_id, price_time)
   - Foreign key to emitens table
   - Multiple indexes for performance

9. **000009_create_stock_price_summary_table**
   - Creates stock_price_summary table for daily OHLC data
   - Open, High, Low, Close prices per day
   - Unique constraint on (emiten_id, date)
   - CHECK constraint for price validation
   - Auto-update timestamp trigger

## Running Migrations

### Using psql (PostgreSQL)

```bash
# Apply all migrations
for migration in migrations/*.up.sql; do
    psql -U username -d database_name -f "$migration"
done

# Rollback specific migration
psql -U username -d database_name -f migrations/000001_create_users_table.down.sql
```

### Using Docker

```bash
# Run migrations with docker exec
docker exec -i postgres-container psql -U username -d database_name < migrations/000001_create_users_table.up.sql

# Or run all migrations
docker exec -i postgres-container psql -U username -d database_name <<EOF
\i /migrations/000001_create_users_table.up.sql
\i /migrations/000002_create_refresh_tokens_table.up.sql
\i /migrations/000003_create_password_resets_table.up.sql
\i /migrations/000004_create_audit_logs_table.up.sql
\i /migrations/000005_seed_admin_user.up.sql
EOF
```

### Using Go Migration Tools

If you're using a Go migration library like `golang-migrate`:

```bash
# Install golang-migrate
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Create migration
migrate create -ext sql -dir migrations -seq create_users_table

# Apply migrations
migrate -path migrations -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" up

# Rollback
migrate -path migrations -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" down 1
```

## Using Complete Schema

For fresh installations, you can use the complete schema:

```bash
psql -U username -d database_name -f migrations/schema.sql
```

## Database Requirements

- PostgreSQL 12 or higher
- UUID extension (uuid-ossp)
- JSONB support for audit logs

## Default Admin User

After running migration 000005, a default admin user is created:

- **Email**: admin@example.com
- **Password**: admin123 (or the bcrypt hash you set)
- **Action Required**: Change password immediately after first login!

### Generate Bcrypt Password Hash

```go
// In Go
import "golang.org/x/crypto/bcrypt"

password := "admin123"
hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
fmt.Println(string(hash))
```

## Best Practices

1. **Version Control**: Always create both up and down migrations
2. **Test Migrations**: Test rollback procedures before production
3. **Backup First**: Always backup database before running migrations
4. **Transaction Safety**: Use transactions for complex migrations
5. **Index Carefully**: Create indexes after bulk data insertions
6. **Review Changes**: Review generated SQL before committing

## Troubleshooting

### Migration Failed

```sql
-- Check current migration version
SELECT * FROM schema_migrations;

-- Rollback to previous version
-- Run corresponding .down.sql file
```

### Duplicate Migration

```sql
-- Check if table exists
SELECT table_name FROM information_schema.tables
WHERE table_schema = 'public';

-- Drop table if needed
DROP TABLE IF EXISTS users CASCADE;
```

### Permission Issues

```sql
-- Grant necessary permissions
GRANT ALL PRIVILEGES ON DATABASE your_database TO your_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO your_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO your_user;
```

## Maintenance

### Clean Old Refresh Tokens

```sql
-- Delete expired refresh tokens older than 30 days
DELETE FROM refresh_tokens
WHERE expires_at < CURRENT_TIMESTAMP - INTERVAL '30 days';
```

### Clean Old Password Reset Tokens

```sql
-- Delete used or expired password reset tokens
DELETE FROM password_resets
WHERE is_used = TRUE
   OR expires_at < CURRENT_TIMESTAMP;
```

### Archive Old Audit Logs

```sql
-- Archive audit logs older than 1 year to a separate table
CREATE TABLE audit_logs_archive AS
SELECT * FROM audit_logs
WHERE created_at < CURRENT_TIMESTAMP - INTERVAL '1 year';

-- Delete archived logs
DELETE FROM audit_logs
WHERE created_at < CURRENT_TIMESTAMP - INTERVAL '1 year';
```
