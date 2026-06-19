-- Up: add reset password fields to users table
ALTER TABLE users
    ADD COLUMN `reset_password_code` VARCHAR(255) DEFAULT NULL,
    ADD COLUMN `reset_password_expire_at` DATETIME DEFAULT NULL;

CREATE INDEX idx_reset_password_code ON users(reset_password_code);
