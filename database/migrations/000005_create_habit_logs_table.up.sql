CREATE TABLE habit_logs (
    id CHAR(36) NOT NULL,
    habit_id CHAR(36) NOT NULL,
    log_date DATE NOT NULL,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id),

    UNIQUE KEY unique_habit_log (habit_id, log_date),
    INDEX idx_habit_logs_habit (habit_id),

    CONSTRAINT fk_habit_logs_habit
        FOREIGN KEY (habit_id)
        REFERENCES habits(id)
        ON DELETE CASCADE
) ENGINE=InnoDB;
