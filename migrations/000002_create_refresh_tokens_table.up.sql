-- Migration: Create refresh_tokens table
-- Version: 000002
-- Description: Create refresh tokens table for JWT refresh token management

-- Create refresh_tokens table
CREATE TABLE IF NOT EXISTS refresh_tokens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    token VARCHAR(500) NOT NULL UNIQUE,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    revoked_at TIMESTAMP WITH TIME ZONE,
    is_revoked BOOLEAN DEFAULT FALSE,
    device_info VARCHAR(255),
    ip_address INET
);

-- Create foreign key to users table
ALTER TABLE refresh_tokens
ADD CONSTRAINT fk_refresh_tokens_user_id
FOREIGN KEY (user_id)
REFERENCES users(id)
ON DELETE CASCADE;

-- Create indexes
CREATE INDEX idx_refresh_tokens_user_id ON refresh_tokens(user_id);
CREATE INDEX idx_refresh_tokens_token ON refresh_tokens(token);
CREATE INDEX idx_refresh_tokens_expires_at ON refresh_tokens(expires_at);
CREATE INDEX idx_refresh_tokens_is_revoked ON refresh_tokens(is_revoked);

-- Add comments
COMMENT ON TABLE refresh_tokens IS 'Refresh tokens for JWT authentication';
COMMENT ON COLUMN refresh_tokens.id IS 'Unique identifier (UUID)';
COMMENT ON COLUMN refresh_tokens.user_id IS 'Reference to user';
COMMENT ON COLUMN refresh_tokens.token IS 'Refresh token string';
COMMENT ON COLUMN refresh_tokens.expires_at IS 'Token expiration time';
COMMENT ON COLUMN refresh_tokens.created_at IS 'Token creation timestamp';
COMMENT ON COLUMN refresh_tokens.revoked_at IS 'Token revocation timestamp';
COMMENT ON COLUMN refresh_tokens.is_revoked IS 'Token revocation status';
COMMENT ON COLUMN refresh_tokens.device_info IS 'Device information (user agent)';
COMMENT ON COLUMN refresh_tokens.ip_address IS 'IP address when token was created';
