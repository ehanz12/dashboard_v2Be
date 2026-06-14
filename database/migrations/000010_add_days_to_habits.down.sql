-- Down: restore date column and remove days
ALTER TABLE habits
    ADD COLUMN `date` DATE DEFAULT NULL,
    DROP COLUMN IF EXISTS `days`;