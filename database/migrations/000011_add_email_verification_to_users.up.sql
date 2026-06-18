-- Up: add email verification fields to users table
ALTER TABLE users
    ADD COLUMN `email_verified` BOOLEAN DEFAULT FALSE,
    ADD COLUMN `verification_code` VARCHAR(255) DEFAULT NULL,
    ADD COLUMN `verification_expire_at` DATETIME DEFAULT NULL;

CREATE INDEX idx_verification_code ON users(verification_code);
