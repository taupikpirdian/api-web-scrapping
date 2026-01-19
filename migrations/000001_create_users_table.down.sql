-- Rollback: Drop users table
-- Version: 000001
-- Description: Drop users table

-- Drop indexes
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_users_is_active;

-- Drop table
DROP TABLE IF EXISTS users;
