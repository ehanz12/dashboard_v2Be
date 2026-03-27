CREATE TABLE habit_logs (
    id CHAR(36) PRIMARY KEY,
    habit_id CHAR(36) NOT NULL,
    log_date DATE NOT NULL,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    UNIQUE KEY unique_log (habit_id, log_date),

    FOREIGN KEY (habit_id) REFERENCES habits(id) ON DELETE CASCADE
);