-- Migration: Create password_resets table
-- Version: 000003
-- Description: Create password_resets table for password reset functionality

-- Create password_resets table
CREATE TABLE IF NOT EXISTS password_resets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    token VARCHAR(500) NOT NULL UNIQUE,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    used_at TIMESTAMP WITH TIME ZONE,
    is_used BOOLEAN DEFAULT FALSE,
    ip_address INET
);

-- Create foreign key to users table
ALTER TABLE password_resets
ADD CONSTRAINT fk_password_resets_user_id
FOREIGN KEY (user_id)
REFERENCES users(id)
ON DELETE CASCADE;

-- Create indexes
CREATE INDEX idx_password_resets_user_id ON password_resets(user_id);
CREATE INDEX idx_password_resets_token ON password_resets(token);
CREATE INDEX idx_password_resets_expires_at ON password_resets(expires_at);
CREATE INDEX idx_password_resets_is_used ON password_resets(is_used);

-- Add comments
COMMENT ON TABLE password_resets IS 'Password reset tokens';
COMMENT ON COLUMN password_resets.id IS 'Unique identifier (UUID)';
COMMENT ON COLUMN password_resets.user_id IS 'Reference to user';
COMMENT ON COLUMN password_resets.token IS 'Password reset token';
COMMENT ON COLUMN password_resets.expires_at IS 'Token expiration time';
COMMENT ON COLUMN password_resets.created_at IS 'Token creation timestamp';
COMMENT ON COLUMN password_resets.used_at IS 'Token usage timestamp';
COMMENT ON COLUMN password_resets.is_used IS 'Token usage status';
COMMENT ON COLUMN password_resets.ip_address IS 'IP address when reset was requested';
