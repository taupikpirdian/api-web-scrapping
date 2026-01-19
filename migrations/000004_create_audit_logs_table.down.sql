-- Rollback: Drop audit_logs table
-- Version: 000004
-- Description: Drop audit_logs table

-- Drop indexes
DROP INDEX IF EXISTS idx_audit_logs_user_id;
DROP INDEX IF EXISTS idx_audit_logs_action;
DROP INDEX IF EXISTS idx_audit_logs_entity;
DROP INDEX IF EXISTS idx_audit_logs_created_at;

-- Drop foreign key
ALTER TABLE audit_logs
DROP CONSTRAINT IF EXISTS fk_audit_logs_user_id;

-- Drop table
DROP TABLE IF EXISTS audit_logs;
