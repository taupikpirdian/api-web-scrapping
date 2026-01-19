-- Migration: Create audit_logs table
-- Version: 000004
-- Description: Create audit_logs table for tracking user activities

-- Create audit_logs table
CREATE TABLE IF NOT EXISTS audit_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID,
    action VARCHAR(100) NOT NULL,
    entity_type VARCHAR(100),
    entity_id UUID,
    old_values JSONB,
    new_values JSONB,
    ip_address INET,
    user_agent VARCHAR(500),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create foreign key to users table (optional, can be NULL for system actions)
ALTER TABLE audit_logs
ADD CONSTRAINT fk_audit_logs_user_id
FOREIGN KEY (user_id)
REFERENCES users(id)
ON DELETE SET NULL;

-- Create indexes
CREATE INDEX idx_audit_logs_user_id ON audit_logs(user_id);
CREATE INDEX idx_audit_logs_action ON audit_logs(action);
CREATE INDEX idx_audit_logs_entity ON audit_logs(entity_type, entity_id);
CREATE INDEX idx_audit_logs_created_at ON audit_logs(created_at);

-- Add comments
COMMENT ON TABLE audit_logs IS 'Audit trail for user activities';
COMMENT ON COLUMN audit_logs.id IS 'Unique identifier (UUID)';
COMMENT ON COLUMN audit_logs.user_id IS 'Reference to user (NULL for system actions)';
COMMENT ON COLUMN audit_logs.action IS 'Action performed (login, logout, create, update, delete)';
COMMENT ON COLUMN audit_logs.entity_type IS 'Type of entity affected';
COMMENT ON COLUMN audit_logs.entity_id IS 'ID of entity affected';
COMMENT ON COLUMN audit_logs.old_values IS 'Previous values (JSON)';
COMMENT ON COLUMN audit_logs.new_values IS 'New values (JSON)';
COMMENT ON COLUMN audit_logs.ip_address IS 'IP address of the request';
COMMENT ON COLUMN audit_logs.user_agent IS 'User agent string';
COMMENT ON COLUMN audit_logs.created_at IS 'Timestamp of the action';
