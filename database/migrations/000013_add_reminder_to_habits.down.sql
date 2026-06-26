DROP INDEX idx_reminder_time ON habits;
ALTER TABLE habits
    DROP COLUMN reminder_time,
    DROP COLUMN reminder_enabled;
