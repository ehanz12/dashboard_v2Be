-- Up: replace date column with days JSON array
ALTER TABLE habits
    DROP COLUMN `date`,
    ADD COLUMN `days` JSON DEFAULT NULL;