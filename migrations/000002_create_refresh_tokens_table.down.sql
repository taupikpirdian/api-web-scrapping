-- Rollback: Drop refresh_tokens table
-- Version: 000002
-- Description: Drop refresh_tokens table

-- Drop indexes
DROP INDEX IF EXISTS idx_refresh_tokens_user_id;
DROP INDEX IF EXISTS idx_refresh_tokens_token;
DROP INDEX IF EXISTS idx_refresh_tokens_expires_at;
DROP INDEX IF EXISTS idx_refresh_tokens_is_revoked;

-- Drop foreign key
ALTER TABLE refresh_tokens
DROP CONSTRAINT IF EXISTS fk_refresh_tokens_user_id;

-- Drop table
DROP TABLE IF EXISTS refresh_tokens;
