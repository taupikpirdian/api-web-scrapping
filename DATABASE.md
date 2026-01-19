# Database Setup Guide

## Prerequisites

- Docker and Docker Compose installed
- PostgreSQL client (psql) installed (optional)

## Quick Start

### 1. Setup Environment

```bash
# Copy database environment file
cp .env.db.example .env.db

# Edit if needed (defaults are fine for development)
nano .env.db
```

### 2. Start Database

```bash
# Start PostgreSQL container
make db-up

# Check logs
make db-logs
```

### 3. Run Migrations

```bash
# Run all migrations
make db-migrate
```

### 4. Connect to Database

```bash
# Using psql inside container
make db-shell

# Or using psql from host
psql -h localhost -p 5432 -U admin -d api_web_scrapping
```

## Available Commands

### Using Makefile

```bash
# Start database
make db-up

# Stop database
make db-down

# View logs
make db-logs

# Run migrations
make db-migrate

# Rollback migration
make db-rollback

# Connect to database shell
make db-shell

# Backup database
make db-backup

# Restore database
make db-restore

# Reset database (WARNING: Deletes all data)
make db-reset
```

### Using Docker Compose

```bash
# Start database
docker-compose -f migrations/docker-compose-db.yml up -d

# Stop database
docker-compose -f migrations/docker-compose-db.yml down

# View logs
docker-compose -f migrations/docker-compose-db.yml logs -f

# Connect to database
docker exec -it api-web-scrapping-db psql -U admin -d api_web_scrapping
```

## Database Schema

### Tables

1. **users** - User accounts
   - id (UUID, primary key)
   - email (VARCHAR, unique)
   - password (VARCHAR, bcrypt hash)
   - full_name (VARCHAR)
   - created_at (TIMESTAMP)
   - updated_at (TIMESTAMP)
   - is_active (BOOLEAN)
   - is_verified (BOOLEAN)

2. **refresh_tokens** - JWT refresh tokens
   - id (UUID, primary key)
   - user_id (UUID, foreign key)
   - token (VARCHAR, unique)
   - expires_at (TIMESTAMP)
   - created_at (TIMESTAMP)
   - revoked_at (TIMESTAMP)
   - is_revoked (BOOLEAN)
   - device_info (VARCHAR)
   - ip_address (INET)

3. **password_resets** - Password reset tokens
   - id (UUID, primary key)
   - user_id (UUID, foreign key)
   - token (VARCHAR, unique)
   - expires_at (TIMESTAMP)
   - created_at (TIMESTAMP)
   - used_at (TIMESTAMP)
   - is_used (BOOLEAN)
   - ip_address (INET)

4. **audit_logs** - Activity audit trail
   - id (UUID, primary key)
   - user_id (UUID, foreign key, nullable)
   - action (VARCHAR)
   - entity_type (VARCHAR)
   - entity_id (UUID)
   - old_values (JSONB)
   - new_values (JSONB)
   - ip_address (INET)
   - user_agent (VARCHAR)
   - created_at (TIMESTAMP)

## Default Admin User

After running migrations, a default admin user is created:

- **Email**: admin@example.com
- **Password**: (Hash set in migration)
- **Status**: Active and Verified

### Generate Admin Password

```bash
# Using Go
go run scripts/generate-password.go

# Using Bash script
./scripts/generate-password.sh

# Using Python
python3 -c "import bcrypt; print(bcrypt.hashpw(b'admin123', bcrypt.gensalt()).decode())"
```

Copy the generated hash to `migrations/000005_seed_admin_user.up.sql`:

```sql
INSERT INTO users (id, email, password, full_name, is_active, is_verified)
VALUES (
    gen_random_uuid(),
    'admin@example.com',
    '$2a$10$generated_hash_here', -- Paste hash here
    'System Administrator',
    TRUE,
    TRUE
);
```

## PgAdmin Access

PgAdmin is available at: http://localhost:5050

- **Email**: admin@example.com
- **Password**: admin

### Setup PgAdmin

1. Open http://localhost:5050
2. Login with credentials
3. Add new server:
   - Name: API Web Scrapping DB
   - Host: postgres
   - Port: 5432
   - Database: api_web_scrapping
   - Username: admin
   - Password: secret

## Connection Strings

### From Application

```bash
# Docker
DATABASE_URL=postgres://admin:secret@postgres:5432/api_web_scrapping?sslmode=disable

# Local
DATABASE_URL=postgres://admin:secret@localhost:5432/api_web_scrapping?sslmode=disable

# With SSL
DATABASE_URL=postgres://admin:secret@localhost:5432/api_web_scrapping?sslmode=require
```

### From psql

```bash
# Local
psql -h localhost -p 5432 -U admin -d api_web_scrapping

# Docker container
docker exec -it api-web-scrapping-db psql -U admin -d api_web_scrapping

# With password prompt
PGPASSWORD=secret psql -h localhost -p 5432 -U admin -d api_web_scrapping
```

## Backup & Restore

### Backup

```bash
# Using Makefile
make db-backup

# Manual
docker exec api-web-scrapping-db pg_dump -U admin api_web_scrapping > backup.sql

# With timestamp
docker exec api-web-scrapping-db pg_dump -U admin api_web_scrapping > backup_$(date +%Y%m%d_%H%M%S).sql

# Compressed
docker exec api-web-scrapping-db pg_dump -U admin api_web_scrapping | gzip > backup.sql.gz
```

### Restore

```bash
# Using Makefile
make db-restore

# Manual
docker exec -i api-web-scrapping-db psql -U admin api_web_scrapping < backup.sql

# From compressed
gunzip -c backup.sql.gz | docker exec -i api-web-scrapping-db psql -U admin api_web_scrapping
```

## Maintenance

### Clean Old Tokens

```sql
-- Delete expired refresh tokens older than 30 days
DELETE FROM refresh_tokens
WHERE expires_at < CURRENT_TIMESTAMP - INTERVAL '30 days';

-- Delete used or expired password reset tokens
DELETE FROM password_resets
WHERE is_used = TRUE OR expires_at < CURRENT_TIMESTAMP;
```

### Archive Audit Logs

```sql
-- Create archive table
CREATE TABLE audit_logs_archive AS
SELECT * FROM audit_logs
WHERE created_at < CURRENT_TIMESTAMP - INTERVAL '1 year';

-- Delete archived logs
DELETE FROM audit_logs
WHERE created_at < CURRENT_TIMESTAMP - INTERVAL '1 year';
```

### Database Statistics

```sql
-- Table sizes
SELECT
    schemaname,
    tablename,
    pg_size_pretty(pg_total_relation_size(schemaname||'.'||tablename)) AS size
FROM pg_tables
WHERE schemaname = 'public'
ORDER BY pg_total_relation_size(schemaname||'.'||tablename) DESC;

-- User count
SELECT COUNT(*) FROM users;

-- Active users
SELECT COUNT(*) FROM users WHERE is_active = TRUE;

-- Active sessions
SELECT COUNT(*) FROM refresh_tokens
WHERE is_revoked = FALSE AND expires_at > CURRENT_TIMESTAMP;
```

## Performance Optimization

### Create Indexes

```sql
-- Composite index for common queries
CREATE INDEX idx_users_email_active ON users(email, is_active);

-- Partial index for active refresh tokens
CREATE INDEX idx_refresh_tokens_active
ON refresh_tokens(user_id, expires_at)
WHERE is_revoked = FALSE;
```

### Vacuum & Analyze

```bash
# Run vacuum and analyze
docker exec api-web-scrapping-db psql -U admin api_web_scrapping -c "VACUUM ANALYZE;"
```

### Update Statistics

```sql
-- Update table statistics
ANALYZE users;
ANALYZE refresh_tokens;
ANALYZE audit_logs;
```

## Troubleshooting

### Database Won't Start

```bash
# Check logs
make db-logs

# Check port conflicts
lsof -i :5432

# Remove old volume (WARNING: Deletes data)
docker-compose -f migrations/docker-compose-db.yml down -v
docker-compose -f migrations/docker-compose-db.yml up -d
```

### Connection Refused

```bash
# Check if database is running
docker ps | grep api-web-scrapping-db

# Check network
docker network inspect api-network

# Test connection
docker exec api-web-scrapping-db pg_isready -U admin
```

### Migration Failed

```bash
# Check current schema
psql -h localhost -U admin -d api_web_scrapping -c "\dt"

# Manually run migration
psql -h localhost -U admin -d api_web_scrapping -f migrations/000001_create_users_table.up.sql

# Rollback
psql -h localhost -U admin -d api_web_scrapping -f migrations/000001_create_users_table.down.sql
```

### Permission Issues

```sql
-- Grant permissions
GRANT ALL PRIVILEGES ON DATABASE api_web_scrapping TO admin;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO admin;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO admin;
```

## Security Best Practices

1. **Change Default Passwords**: Update admin password immediately
2. **Use Strong Passwords**: At least 16 characters for database password
3. **Enable SSL**: Use `sslmode=require` in production
4. **Limit Access**: Use firewall rules to restrict database access
5. **Regular Backups**: Automate daily backups
6. **Monitor Logs**: Check database logs regularly
7. **Update Regularly**: Keep PostgreSQL updated

## Production Deployment

### Environment Variables

```bash
POSTGRES_USER=prod_user
POSTGRES_PASSWORD=very_secure_password_here
POSTGRES_DB=api_web_scrapping_prod
DATABASE_URL=postgres://prod_user:very_secure_password_here@postgres:5432/api_web_scrapping_prod?sslmode=require
```

### Persistent Storage

```yaml
volumes:
  postgres_data:
    driver: local
    driver_opts:
      type: none
      device: /data/postgresql
      o: bind
```

### Resource Limits

```yaml
deploy:
  resources:
    limits:
      cpus: '2'
      memory: 2G
    reservations:
      cpus: '1'
      memory: 1G
```

## Useful Queries

### Find User by Email

```sql
SELECT * FROM users WHERE email = 'user@example.com';
```

### Get User Sessions

```sql
SELECT * FROM user_sessions WHERE user_id = 'user-uuid';
```

### Recent Audit Logs

```sql
SELECT
    al.*,
    u.email
FROM audit_logs al
LEFT JOIN users u ON al.user_id = u.id
ORDER BY al.created_at DESC
LIMIT 100;
```

### Expired Tokens

```sql
SELECT * FROM refresh_tokens
WHERE expires_at < CURRENT_TIMESTAMP
AND is_revoked = FALSE;
```

## Next Steps

1. ✅ Start database: `make db-up`
2. ✅ Run migrations: `make db-migrate`
3. ✅ Generate admin password
4. ✅ Update migration with password hash
5. ✅ Connect and verify: `make db-shell`
6. ✅ Update application config with database connection
7. ✅ Test application with database
