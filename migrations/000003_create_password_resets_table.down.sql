-- Rollback: Drop password_resets table
-- Version: 000003
-- Description: Drop password_resets table

-- Drop indexes
DROP INDEX IF EXISTS idx_password_resets_user_id;
DROP INDEX IF EXISTS idx_password_resets_token;
DROP INDEX IF EXISTS idx_password_resets_expires_at;
DROP INDEX IF EXISTS idx_password_resets_is_used;

-- Drop foreign key
ALTER TABLE password_resets
DROP CONSTRAINT IF EXISTS fk_password_resets_user_id;

-- Drop table
DROP TABLE IF EXISTS password_resets;
