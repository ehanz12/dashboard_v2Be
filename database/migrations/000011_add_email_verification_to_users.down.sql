-- Down: remove email verification fields from users table
ALTER TABLE users
    DROP INDEX idx_verification_code,
    DROP COLUMN `email_verified`,
    DROP COLUMN `verification_code`,
    DROP COLUMN `verification_expire_at`;
