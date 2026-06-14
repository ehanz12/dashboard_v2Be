-- Up: replace date column with days JSON array
ALTER TABLE habits
    DROP COLUMN IF EXISTS `date`,
    ADD COLUMN `days` JSON DEFAULT NULL;