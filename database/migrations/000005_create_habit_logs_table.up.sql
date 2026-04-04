CREATE TABLE habit_logs (
    id CHAR(36) PRIMARY KEY,
    habit_id CHAR(36) NOT NULL,
    log_date DATE NOT NULL,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    UNIQUE KEY unique_log (habit_id, log_date),

    INDEX idx_habit_logs_habit (habit_id),
    INDEX idx_habit_logs_date (log_date),
    INDEX idx_habit_logs_habit_date (habit_id, log_date),

    FOREIGN KEY (habit_id) REFERENCES habits(id) ON DELETE CASCADE
);