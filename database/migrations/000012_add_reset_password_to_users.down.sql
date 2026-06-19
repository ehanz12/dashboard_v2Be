-- Down: remove reset password fields from users table
ALTER TABLE users
    DROP INDEX idx_reset_password_code,
    DROP COLUMN `reset_password_code`,
    DROP COLUMN `reset_password_expire_at`;
