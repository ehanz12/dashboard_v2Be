ALTER TABLE habits
    ADD COLUMN `reminder_time` TIME DEFAULT NULL,
    ADD COLUMN `reminder_enabled` BOOLEAN DEFAULT FALSE;
CREATE INDEX idx_reminder_time ON habits(reminder_time);