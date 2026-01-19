-- Migration: Seed admin user
-- Version: 000005
-- Description: Create default admin user
-- Note: Password is 'admin123' hashed with bcrypt
-- IMPORTANT: Change this password immediately after first login!

-- Insert admin user
-- Password: admin123 (hashed with bcrypt, cost 10)
INSERT INTO users (id, email, password, full_name, is_active, is_verified)
VALUES (
    gen_random_uuid(),
    'admin@example.com',
    '$2a$10$YourBcryptHashedPasswordHere', -- Replace with actual bcrypt hash
    'System Administrator',
    TRUE,
    TRUE
)
ON CONFLICT (email) DO NOTHING;

-- Add comment
COMMENT ON TABLE users IS 'Default admin user created. Email: admin@example.com. Please change password immediately.';
