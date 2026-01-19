-- Rollback: Remove admin user
-- Version: 000005
-- Description: Remove default admin user

-- Delete admin user
DELETE FROM users WHERE email = 'admin@example.com';
